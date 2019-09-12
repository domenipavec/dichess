package chess_state

import (
	"github.com/freeeve/uci"
	"github.com/notnil/chess"
)

type UciPlayer struct {
	engine *uci.Engine
}

func NewUciPlayer() (*UciPlayer, error) {
	engine, err := uci.NewEngine("/home/pi/stockfish")
	if err != nil {
		return nil, err
	}
	return &UciPlayer{engine}, nil
}

func (p *UciPlayer) MakeMove(game *chess.Game) (*Move, error) {
	if err := p.engine.SetFEN(game.FEN()); err != nil {
		return nil, err
	}
	result, err := p.engine.Go(0, "", 1000)
	if err != nil {
		return nil, err
	}
	if err := game.MoveStr(result.BestMove); err != nil {
		return nil, err
	}
	return &Move{
		ShouldMove: true,
		ShouldSay:  true,
	}, nil
}

func (p *UciPlayer) Close() error {
	p.engine.Close()
	return nil
}
