package chess_state

import (
	"log"

	"github.com/notnil/chess"
)

type Player interface {
	MakeMove(StateSender, *chess.Game) (*Move, error)
	Close() error
}

type Move struct {
	ShouldMove bool
	ShouldSay  bool
}

type Game struct {
	Observers    *Observers
	StateSenders *StateSenders
	Game         *chess.Game
	Players      []Player
}

func NewGame(player1, player2 Player, Observers *Observers, StateSenders *StateSenders) *Game {
	return &Game{
		Game:         chess.NewGame(chess.UseNotation(chess.LongAlgebraicNotation{})),
		Observers:    Observers,
		StateSenders: StateSenders,
		Players:      []Player{player1, player2},
	}
}

func (g *Game) Play() error {
	defer func() {
		for _, player := range g.Players {
			if err := player.Close(); err != nil {
				log.Printf("Could not close player: %v", err)
			}
		}
	}()

	var move *Move
	for g.Game.Outcome() == chess.NoOutcome {
		g.Observers.Update(g.StateSenders, g, move)

		var player Player
		if g.Game.Position().Turn() == chess.White {
			player = g.Players[0]
		} else {
			player = g.Players[1]
		}
		newMove, err := player.MakeMove(g.StateSenders, g.Game)
		if err != nil {
			return err
		}
		move = newMove
	}

	// handle outcome

	return nil
}

func Square(x, y int) chess.Square {
	return chess.Square(x + 8*y)
}
