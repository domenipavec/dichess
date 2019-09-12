package chess_state

import (
	"context"

	"github.com/notnil/chess"
)

type HumanPlayer struct {
	Inputs []HumanInput
}

type HumanInput interface {
	MakeMove(context.Context, *chess.Game) (*Move, error)
}

type inputResult struct {
	err  error
	move *Move
}

func (p *HumanPlayer) MakeMove(game *chess.Game) (*Move, error) {
	ctx, cancel := context.WithCancel(context.Background())

	resultChan := make(chan inputResult)
	for _, input := range p.Inputs {
		input := input
		go func() {
			move, err := input.MakeMove(ctx, game)
			resultChan <- inputResult{err: err, move: move}
		}()
	}

	result := <-resultChan
	cancel()
	return result.move, result.err
}

func (p *HumanPlayer) Close() error {
	return nil
}
