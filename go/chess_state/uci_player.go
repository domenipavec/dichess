package chess_state

import (
	"context"

	"github.com/freeeve/uci"
	"github.com/matematik7/dichess/go/bluetoothpb"
	"github.com/notnil/chess"
	"github.com/pkg/errors"
)

type UciPlayer struct {
	settings bluetoothpb.SettingsProvider
}

func NewUciPlayer(settings bluetoothpb.SettingsProvider) (*UciPlayer, error) {

	return &UciPlayer{settings}, nil
}

type result struct {
	bestMove string
	err      error
}

func (p *UciPlayer) MakeMove(ctx context.Context, stateSender StateSender, game *chess.Game) (*Move, error) {
	stateSender.StateSend("Computer's turn. Thinking...")

	cs := p.settings.GetSettings().ComputerSettings

	engine, err := uci.NewEngine("/home/pi/stockfish")
	if err != nil {
		return nil, err
	}
	if err := engine.SendOption("Skill Level", cs.SkillLevel); err != nil {
		return nil, errors.Wrap(err, "could not set skill level")
	}
	if err := engine.SendOption("UCI_LimitStrength", cs.LimitStrength); err != nil {
		return nil, errors.Wrap(err, "could not set limit strength")
	}
	if err := engine.SendOption("UCI_Elo", cs.Elo); err != nil {
		return nil, errors.Wrap(err, "could not set elo")
	}

	if err := engine.SetFEN(game.FEN()); err != nil {
		return nil, err
	}

	resultChan := make(chan result)
	go func() {
		defer engine.Close()
		r, err := engine.Go(0, "", int64(cs.TimeLimitMs))
		if err != nil {
			resultChan <- result{err: err}
		}
		resultChan <- result{bestMove: r.BestMove}
	}()

	select {
	case <-ctx.Done():
		return &Move{}, nil
	case r := <-resultChan:
		if r.err != nil {
			return nil, r.err
		}
		move, err := chess.LongAlgebraicNotation{}.Decode(game.Position(), r.bestMove)
		if err != nil {
			return nil, err
		}
		return &Move{
			Move:       move,
			ShouldMove: true,
			ShouldSay:  true,
		}, nil
	}

}

func (p *UciPlayer) Close() error {
	return nil
}
