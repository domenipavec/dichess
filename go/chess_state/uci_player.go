package chess_state

import (
	"github.com/freeeve/uci"
	"github.com/matematik7/dichess/go/bluetoothpb"
	"github.com/notnil/chess"
	"github.com/pkg/errors"
)

type UciPlayer struct {
	engine *uci.Engine

	timeLimit int64
}

func NewUciPlayer(cs *bluetoothpb.Settings_ComputerSettings) (*UciPlayer, error) {
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

	return &UciPlayer{engine, int64(cs.TimeLimitMs)}, nil
}

func (p *UciPlayer) MakeMove(stateSender StateSender, game *chess.Game) (*Move, error) {
	stateSender.StateSend("Computer's turn. Thinking...")
	if err := p.engine.SetFEN(game.FEN()); err != nil {
		return nil, err
	}
	result, err := p.engine.Go(0, "", p.timeLimit)
	if err != nil {
		return nil, err
	}
	move, err := chess.LongAlgebraicNotation{}.Decode(game.Position(), result.BestMove)
	if err != nil {
		return nil, err
	}
	return &Move{
		Move:       move,
		ShouldMove: true,
		ShouldSay:  true,
	}, nil
}

func (p *UciPlayer) Close() error {
	p.engine.Close()
	return nil
}
