package hardware

import (
	"context"
	"fmt"
	"log"
	"math"

	"github.com/hashicorp/go-multierror"
	"github.com/matematik7/dichess/go/chess_state"
	"github.com/notnil/chess"
	"github.com/pkg/errors"
)

func (h *Hardware) Do(fs ...func() error) error {
	errs := make(chan error)
	for _, f := range fs {
		f := f
		go func() {
			errs <- f()
		}()
	}

	var result error
	for range fs {
		if err := <-errs; err != nil {
			result = multierror.Append(result, err)
		}
	}
	return result
}

func (h *Hardware) undoAllMovedPieces() error {
	for _, p := range h.movedPieces {
		if err := h.move(p.x2, p.y2, p.x1, p.y1, p.color, true); err != nil {
			return err
		}
	}
	h.movedPieces = nil
	return nil
}

func (h *Hardware) checkAndMove(position *chess.Position, x1, y1, x2, y2 float64) error {
	piece := position.Board().Piece(chess_state.Square(int(x1), int(y1)))
	if piece == chess.NoPiece {
		return nil
	}

	if err := h.move(x1, y1, x2, y2, piece.Color(), true); err != nil {
		return errors.Wrap(err, "couldn't check and move")
	}

	h.movedPieces = append(h.movedPieces, movedPiece{
		x1:    x1,
		y1:    y1,
		x2:    x2,
		y2:    y2,
		color: piece.Color(),
	})

	return nil
}

func (h *Hardware) move(x1, y1, x2, y2 float64, color chess.Color, rotate bool) error {
	if err := h.Do(
		func() error { return h.xAxis.GoTo(x1, 40) },
		func() error { return h.yAxis.GoTo(y1, 40) },
	); err != nil {
		return err
	}

	if err := h.coil.On(); err != nil {
		return err
	}

	log.Printf("move %v piece (%f, %f) -> (%f, %f)", color, x1, y1, x2, y2)
	if rotate {
		angle := 0.0
		if color == chess.White {
			angle = math.Atan2(x2-x1, y2-y1) / math.Pi * 180
		} else {
			angle = math.Atan2(x1-x2, y1-y2) / math.Pi * 180
		}
		log.Printf("angle: %f", angle)
		if err := h.coil.Rotate(int(angle)); err != nil {
			return err
		}
	}

	var vx, vy int
	sx := math.Abs(x2 - x1)
	sy := math.Abs(y2 - y1)
	if sx > sy {
		vx = 80
		vy = int(sx * float64(vx) / sy)
		if vy > 255 {
			vy = 255
		}
	} else {
		vy = 80
		vx = int(sy * float64(vy) / sx)
		if vx > 255 {
			vx = 255
		}
	}
	if err := h.Do(
		func() error { return h.xAxis.GoTo(x2, uint8(vx)) },
		func() error { return h.yAxis.GoTo(y2, uint8(vy)) },
	); err != nil {
		return err
	}

	if rotate {
		if err := h.coil.Rotate(0); err != nil {
			return err
		}
	}

	if err := h.coil.Off(); err != nil {
		return err
	}

	return nil
}

func (h *Hardware) Update(ctx context.Context, stateSender chess_state.StateSender, game *chess_state.Game, move *chess_state.Move) error {
	if !h.Settings.GetSettings().AutoMove {
		return nil
	}
	stateSender.StateSend("Moving pieces.")
	if h.fake {
		return nil
	}
	if move == nil || !move.ShouldMove || move.Undo {
		return nil
	}
	if move.Undo {
		// wait for undo observing ctx
		return nil
	}
	moves := game.Game.Moves()
	if len(moves) < 1 {
		return nil
	}
	gameMove := moves[len(moves)-1]
	if gameMove.HasTag(chess.KingSideCastle) || gameMove.HasTag(chess.QueenSideCastle) {
		stateSender.StateSend("Waiting for castling to complete.")
		return h.WaitFor(ctx, int(gameMove.S2().File()), int(gameMove.S2().Rank()), true)
	}
	gamePosition := game.Game.Positions()[len(game.Game.Positions())-2]
	piece := gamePosition.Board().Piece(gameMove.S1())
	if gameMove.HasTag(chess.Capture) {
		stateSender.StateSend(fmt.Sprintf("Waiting for captured piece on %v to be removed.", gameMove.S2().String()))
		if err := h.WaitFor(ctx, int(gameMove.S2().File()), int(gameMove.S2().Rank()), false); err != nil {
			return err
		}
	}

	if err := h.xAxis.SetCurrent(77); err != nil {
		return errors.Wrap(err, "couldn't set working current")
	}
	if err := h.yAxis.SetCurrent(77); err != nil {
		return errors.Wrap(err, "couldn't set working current")
	}
	defer func() {
		if err := h.xAxis.SetCurrent(10); err != nil {
			log.Printf("Couldn't set idle current: %v", err)
		}
		if err := h.yAxis.SetCurrent(10); err != nil {
			log.Printf("Couldn't set idle current: %v", err)
		}
	}()

	x1 := float64(gameMove.S1().File())
	y1 := float64(gameMove.S1().Rank())
	x2 := float64(gameMove.S2().File())
	y2 := float64(gameMove.S2().Rank())

	isDiagonal := math.Abs(x1-x2) > 0 && math.Abs(y1-y2) > 0
	if piece == chess.WhiteKnight || piece == chess.BlackKnight || (isDiagonal && (piece == chess.WhiteQueen || piece == chess.BlackQueen)) {
		if math.Abs(y2-y1) == 2 {
			if err := h.checkAndMove(gamePosition, x1, y1+(y2-y1)*0.5, x1-(x2-x1)*0.5, y1+(y2-y1)*0.5); err != nil {
				return err
			}
			if err := h.checkAndMove(gamePosition, x2, y1+(y2-y1)*0.5, x2+(x2-x1)*0.5, y1+(y2-y1)*0.5); err != nil {
				return err
			}
		} else {
			if err := h.checkAndMove(gamePosition, x1+(x2-x1)*0.5, y1, x1+(x2-x1)*0.5, y1-(y2-y1)*0.5); err != nil {
				return err
			}
			if err := h.checkAndMove(gamePosition, x1+(x2-x1)*0.5, y2, x1+(x2-x1)*0.5, y2+(y2-y1)*0.5); err != nil {
				return err
			}
		}
	}
	if piece == chess.WhiteBishop || piece == chess.BlackBishop {
		x := x1
		y := y1
		dx := (x2 - x1) / math.Abs(x2-x1)
		dy := (y2 - y1) / math.Abs(y2-y1)
		for {
			if int(x) == int(x2) && int(y) == int(y2) {
				break
			}
			if err := h.checkAndMove(gamePosition, x+dx, y, x+1.25*dx, y-0.25*dy); err != nil {
				return err
			}
			if err := h.checkAndMove(gamePosition, x, y+dy, x-0.25*dx, y+1.25*dx); err != nil {
				return err
			}

			x += dx
			y += dy
		}
	}

	if err := h.move(
		x1, y1, x2, y2,
		gamePosition.Turn(),
		true, /* rotate */
	); err != nil {
		return err
	}

	if err := h.undoAllMovedPieces(); err != nil {
		return err
	}

	return nil
}
