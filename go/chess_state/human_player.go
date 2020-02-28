package chess_state

import (
	"context"
	"fmt"
	"strings"

	"github.com/notnil/chess"
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

func (p *HumanPlayer) MakeMove(stateSender StateSender, game *chess.Game) (*Move, error) {
	stateSender.StateSend(fmt.Sprintf("Waiting for move, %s turn.", strings.ToLower(game.Position().Turn().Name())))
	ctx, cancel := context.WithCancel(context.Background())

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

	result := <-resultChan
	cancel()
	return result.move, result.err
}

func (p *HumanPlayer) Close() error {
	return nil
}
