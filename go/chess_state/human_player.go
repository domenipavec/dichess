package chess_state

import (
	"context"
	"fmt"
	"strings"

	"github.com/notnil/chess"
	"github.com/pkg/errors"
)

type HumanPlayer struct {
	Inputs []HumanInput
}

type HumanInput interface {
	MakeMove(context.Context, StateSender, *chess.Game) (*Move, error)
}

type inputResult struct {
	err  error
	move *Move
}

func (p *HumanPlayer) MakeMove(ctx context.Context, stateSender StateSender, game *chess.Game) (*Move, error) {
	state := fmt.Sprintf("Waiting for move, %s turn.", strings.ToLower(game.Position().Turn().Name()))
	moves := game.Moves()
	if len(moves) > 0 && moves[len(moves)-1].HasTag(chess.Check) {
		state = "Check. " + state
	}
	stateSender.StateSend(state)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	resultChan := make(chan inputResult)
	for _, input := range p.Inputs {
		input := input
		go func() {
			move, err := input.MakeMove(ctx, stateSender, game)
			select {
			case resultChan <- inputResult{err: err, move: move}:
			default:
			}
		}()
	}

	errored := 0
	for {
		select {
		case result := <-resultChan:
			if result.err != nil {
				errored++
				if errored >= len(p.Inputs) {
					return nil, errors.Wrapf(result.err, "all human inputs errored, last error")
				}
			} else {
				return result.move, nil
			}
		case <-ctx.Done():
			return &Move{}, nil
		}
	}
}

func (p *HumanPlayer) Close() error {
	return nil
}
