package hardware

import (
	"context"
	"log"
	"math"
	"time"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/matematik7/dichess/go/chess_state"
	"github.com/notnil/chess"
	"github.com/pkg/errors"
	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/conn/i2c"
	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/host"
)

type Hardware struct {
	initialized bool

	matrix *ReedMatrix
	xAxis  *Axis
	yAxis  *Axis
	coil   *Coil

	movedPieces []movedPiece
}

type movedPiece struct {
	x1, y1 float64
	x2, y2 float64
	color  chess.Color
}

func New() *Hardware {
	return &Hardware{
		matrix: &ReedMatrix{},
		yAxis: &Axis{
			MotorDriver: &MotorDriver{
				Dev: &i2c.Dev{Addr: 21},
			},
			MinusOffset: 220,
			FirstOffset: 0,
			EveryOffset: 247,
			LastOffset:  0,
		},
		xAxis: &Axis{
			MotorDriver: &MotorDriver{
				Dev: &i2c.Dev{Addr: 22},
			},
			MinusOffset: 250,
			FirstOffset: 0,
			EveryOffset: 247,
			LastOffset:  0,
		},
		coil: &Coil{
			MotorDriver: &MotorDriver{
				Dev: &i2c.Dev{Addr: 23},
			},
		},
	}
}

func (h *Hardware) Initialize() error {
	if _, err := host.Init(); err != nil {
		return errors.Wrap(err, "couldn't initialize periph")
	}
	bus, err := i2creg.Open("1")
	if err != nil {
		return errors.Wrap(err, "couldn't open i2c bus 1")
	}
	h.xAxis.MotorDriver.Dev.Bus = bus
	h.yAxis.MotorDriver.Dev.Bus = bus
	h.coil.MotorDriver.Dev.Bus = bus
	h.matrix.Rows = []gpio.PinIn{
		gpioreg.ByName("GPIO4"),
		gpioreg.ByName("GPIO17"),
		gpioreg.ByName("GPIO27"),
		gpioreg.ByName("GPIO22"),
		gpioreg.ByName("GPIO5"),
		gpioreg.ByName("GPIO6"),
		gpioreg.ByName("GPIO13"),
		gpioreg.ByName("GPIO19"),
	}
	h.matrix.Columns = []gpio.PinIO{
		gpioreg.ByName("GPIO21"),
		gpioreg.ByName("GPIO20"),
		gpioreg.ByName("GPIO16"),
		gpioreg.ByName("GPIO12"),
		gpioreg.ByName("GPIO25"),
		gpioreg.ByName("GPIO24"),
		gpioreg.ByName("GPIO23"),
		gpioreg.ByName("GPIO18"),
	}

	// coil should be initialized first so it doesn't heat for nothing
	if err := h.coil.Initialize(); err != nil {
		return errors.Wrap(err, "couldn't initialize coil")
	}

	if err := h.matrix.Initialize(); err != nil {
		return errors.Wrap(err, "couldn't initialize reed matrix")
	}

	// Must home axis with addr 21 before addr 22
	if err := h.yAxis.Initialize(); err != nil {
		return errors.Wrap(err, "couldn't initialize y-axis")
	}
	if err := h.xAxis.Initialize(); err != nil {
		return errors.Wrap(err, "couldn't initialize x-axis")
	}

	h.initialized = true

	return nil
}

func (h *Hardware) ReadMatrix() ([][]bool, error) {
	if !h.initialized {
		data := make([][]bool, 8)
		for i := 0; i < 8; i++ {
			data[i] = make([]bool, 8)
		}
		return data, nil
	}
	return h.matrix.Read()
}

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

	log.Printf("move %v piece (%v, %v) -> (%v, %v)", color, x1, y1, x2, y2)
	if rotate {
		angle := 0.0
		if color == chess.White {
			angle = math.Atan2(x2-x1, y2-y1) / math.Pi * 180
		} else {
			angle = math.Atan2(x1-x2, y1-y2) / math.Pi * 180
		}
		log.Printf("angle: %v", angle)
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

func (h *Hardware) Update(game *chess.Game, move *chess_state.Move) error {
	if move == nil || !move.ShouldMove {
		return nil
	}
	moves := game.Moves()
	if len(moves) < 1 {
		return nil
	}
	gameMove := moves[len(moves)-1]
	gamePosition := game.Positions()[len(game.Positions())-2]
	piece := gamePosition.Board().Piece(gameMove.S1())
	if (piece != chess.WhitePawn && piece != chess.BlackPawn) || gameMove.HasTag(chess.Capture) {
		for {
			data, err := h.matrix.Read()
			if err != nil {
				return err
			}
			if data[gameMove.S2().File()][gameMove.S2().Rank()] {
				break
			}
		}
		return nil
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

	// Implement castle and capture
	// Implement promotion
	// Implement taking pieces
	// Implement knight moves
	// Implement laufer moves
	if piece == chess.WhiteKnight || piece == chess.BlackKnight {
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

func (h *Hardware) MakeMove(ctx context.Context, game *chess.Game) (*chess_state.Move, error) {
	time.Sleep(time.Millisecond * 100)
	initialData, err := h.matrix.Read()
	if err != nil {
		return nil, err
	}
	log.Println("Got initial")
	prev := initialData
	for {
		select {
		case <-ctx.Done():
			return nil, nil
		case <-time.NewTimer(time.Millisecond * 1000).C:
		}

		data, err := h.matrix.Read()
		if err != nil {
			return nil, err
		}

		sameAsPrev := true
		for i := range data {
			for j := range data[i] {
				if data[i][j] != prev[i][j] {
					sameAsPrev = false
				}
			}
		}
		// wait for same reading 2 times in a row
		prev = data
		if !sameAsPrev {
			continue
		}

		var from []chess.Square
		var to []chess.Square
		for i := 0; i < 8; i++ {
			for j := 0; j < 8; j++ {
				if initialData[i][j] && !data[i][j] {
					from = append(from, chess_state.Square(i, j))
				}
				if !initialData[i][j] && data[i][j] {
					to = append(to, chess_state.Square(i, j))
				}
			}
		}
		log.Printf("from %v, to %v", from, to)
		if len(from) == 0 {
			continue
		}

		var validMoves []*chess.Move
		for _, move := range game.ValidMoves() {
			for _, s1 := range from {
				if move.S1() == s1 && move.HasTag(chess.Capture) && len(to) == 0 {
					validMoves = append(validMoves, move)
					break
				}
				for _, s2 := range to {
					if move.S1() == s1 && move.S2() == s2 {
						validMoves = append(validMoves, move)
						break
					}
				}
			}
		}

		if len(validMoves) < 1 {
			log.Printf("Invalid move from %v to %v", from, to)
			continue
		}
		if len(validMoves) > 1 {
			log.Printf("Ambiguous move from %v to %v: %v", from, to, validMoves)
			continue
		}

		if err := game.Move(validMoves[0]); err != nil {
			return nil, errors.Wrap(err, "couldn't make a valid move")
		}

		return &chess_state.Move{
			ShouldMove: false,
			ShouldSay:  true,
		}, nil
	}
}
