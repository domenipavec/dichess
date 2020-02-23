package hardware

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/matematik7/dichess/go/chess_state"
	"github.com/notnil/chess"
	"github.com/pkg/errors"
)

func (h *Hardware) MakeMove(ctx context.Context, stateSender chess_state.StateSender, game *chess.Game) (*chess_state.Move, error) {
	time.Sleep(time.Millisecond * 100)
	initialData, err := h.matrix.Read()
	if err != nil {
		return nil, err
	}
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
			stateSender.StateSend(fmt.Sprintf("Invalid move from %v to %v.", from, to))
			continue
		}
		if len(validMoves) > 1 {
			stateSender.StateSend(fmt.Sprintf("Ambiguous move from %v to %v: %v", from, to, validMoves))
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

func (h *Hardware) StartGame(stateSender chess_state.StateSender) error {
	for {
		missing := "Waiting for all pieces to start the game. Missing: "
		commaNeeded := false
		addMissing := func(i, j int) {
			if !commaNeeded {
				commaNeeded = true
			} else {
				missing += ", "
			}
			missing += string([]byte{byte('A' + i), byte('1' + j)})
		}

		time.Sleep(time.Second)
		data, err := h.ReadMatrix()
		if err != nil {
			log.Fatal(err)
		}

		done := true
		for i := 0; i < 8; i++ {
			for j := 0; j < 2; j++ {
				if !data[i][j] {
					addMissing(i, j)
					done = false
				}
			}
		}
		for i := 0; i < 8; i++ {
			for j := 6; j < 8; j++ {
				if i == 0 && j == 6 {
					continue
				}
				if i == 2 && j == 6 {
					continue
				}
				if !data[i][j] {
					addMissing(i, j)
					done = false
				}
			}
		}
		if done {
			break
		}

		missing += "."
		stateSender.StateSend(missing)
	}
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
