package chess_state

import (
	"context"
	"fmt"

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
	stateSender.StateSend(fmt.Sprintf("Waiting for move, %s turn.", game.Position().Turn()))
	ctx, cancel := context.WithCancel(context.Background())

	resultChan := make(chan inputResult)
	for _, input := range p.Inputs {
		input := input
		go func() {
			move, err := input.MakeMove(ctx, stateSender, game)
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
