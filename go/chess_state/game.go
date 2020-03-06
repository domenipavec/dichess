package chess_state

import (
	"context"
	"log"

	"github.com/notnil/chess"
)

type Player interface {
	MakeMove(context.Context, StateSender, *chess.Game) (*Move, error)
	Close() error
}

type Move struct {
	*chess.Move
	ShouldMove bool
	ShouldSay  bool
	Undo       bool
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

func (g *Game) Play(ctx context.Context) error {
	defer func() {
		for _, player := range g.Players {
			if err := player.Close(); err != nil {
				log.Printf("Could not close player: %v", err)
			}
		}
	}()

	var move *Move
	for g.Game.Outcome() == chess.NoOutcome {
		g.Observers.Update(ctx, g.StateSenders, g, move)

		select {
		case <-ctx.Done():
			return nil
		default:
		}

		var player Player
		if g.Game.Position().Turn() == chess.White {
			player = g.Players[0]
		} else {
			player = g.Players[1]
		}
		newMove, err := player.MakeMove(ctx, g.StateSenders, g.Game)
		if err != nil {
			return err
		}

		select {
		case <-ctx.Done():
			return nil
		default:
		}

		if newMove.Undo {
			if err := g.undoMove(); err != nil {
				return err
			}
		} else {
			if err := g.Game.Move(newMove.Move); err != nil {
				return err
			}
		}
		move = newMove
	}

	g.Observers.Update(ctx, g.StateSenders, g, move)
	g.StateSenders.StateSend(g.Game.Outcome().String())

	// handle outcome

	return nil
}

func (g *Game) undoMove() error {
	moves := g.Game.Moves()

	var otherPlayer Player
	if g.Game.Position().Turn() == chess.White {
		otherPlayer = g.Players[1]
	} else {
		otherPlayer = g.Players[0]
	}
	if _, ok := otherPlayer.(*UciPlayer); ok {
		if len(moves) < 2 {
			return nil
		}
		moves = moves[:len(moves)-2]
	} else {
		if len(moves) < 1 {
			return nil
		}
		moves = moves[:len(moves)-1]
	}

	newGame := chess.NewGame(chess.UseNotation(chess.LongAlgebraicNotation{}))
	for _, move := range moves {
		if err := newGame.Move(move); err != nil {
			return err
		}
	}
	g.Game = newGame

	return nil
}

func Square(x, y int) chess.Square {
	return chess.Square(x + 8*y)
}
