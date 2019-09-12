package chess_state

import (
	"github.com/notnil/chess"
)

type Player interface {
	MakeMove(*chess.Game) (*Move, error)
	Close() error
}

type Move struct {
	ShouldMove bool
	ShouldSay  bool
}

type Game struct {
	game      *chess.Game
	current   int
	players   []Player
	observers *Observers
}

func NewGame(player1, player2 Player, observers *Observers) *Game {
	return &Game{
		game:      chess.NewGame(chess.UseNotation(chess.LongAlgebraicNotation{})),
		observers: observers,
		players:   []Player{player1, player2},
	}
}

func (g *Game) Play() error {
	var move *Move
	for g.game.Outcome() == chess.NoOutcome {
		g.observers.Update(g.game, move)

		newMove, err := g.players[g.current].MakeMove(g.game)
		if err != nil {
			return err
		}
		move = newMove

		g.current++
		if g.current > 1 {
			g.current = 0
		}
	}

	// handle outcome

	for _, player := range g.players {
		if err := player.Close(); err != nil {
			return err
		}
	}

	return nil
}

func Square(x, y int) chess.Square {
	return chess.Square(x + 8*y)
}
