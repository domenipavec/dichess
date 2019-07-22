package main

import (
	"github.com/freeeve/uci"
	"github.com/notnil/chess"
)

type Player interface {
	MakeMove(*chess.Game) error
	Close() error
}

type ChessState struct {
	game      *chess.Game
	current   int
	players   []Player
	observers *Observers
}

func NewChessState(player1, player2 Player, observers *Observers) *ChessState {
	return &ChessState{
		game:      chess.NewGame(chess.UseNotation(chess.LongAlgebraicNotation{})),
		observers: observers,
		players:   []Player{player1, player2},
	}
}

func (cs *ChessState) Play() error {
	for cs.game.Outcome() == chess.NoOutcome {
		cs.observers.Update(cs.game)

		err := cs.players[cs.current].MakeMove(cs.game)
		if err != nil {
			return err
		}

		cs.current++
		if cs.current > 1 {
			cs.current = 0
		}
	}

	// handle outcome

	for _, player := range cs.players {
		if err := player.Close(); err != nil {
			return err
		}
	}

	return nil
}

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

func (p *UciPlayer) MakeMove(game *chess.Game) error {
	if err := p.engine.SetFEN(game.FEN()); err != nil {
		return err
	}
	result, err := p.engine.Go(0, "", 1000)
	if err != nil {
		return err
	}
	if err := game.MoveStr(result.BestMove); err != nil {
		return err
	}
	return nil
}

func (p *UciPlayer) Close() error {
	p.engine.Close()
	return nil
}
