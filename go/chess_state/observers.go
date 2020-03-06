package chess_state

import (
	"context"
	"log"
	"sync"
)

// Observer should use game.Game as read only
type Observer interface {
	Update(context.Context, StateSender, *Game, *Move) error
}

type Observers struct {
	observers []Observer
	mutex     sync.Mutex
}

func (o *Observers) Add(observer Observer) int {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	o.observers = append(o.observers, observer)
	return len(o.observers) - 1
}

func (o *Observers) Remove(id int) {
	o.observers[id] = nil
}

// Update calls update on all observers
func (o *Observers) Update(ctx context.Context, stateSender StateSender, game *Game, move *Move) {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	errChan := make(chan error)
	for _, observer := range o.observers {
		if observer == nil {
			continue
		}
		observer := observer
		go func() {
			errChan <- observer.Update(ctx, stateSender, game, move)
		}()
	}

	for _, observer := range o.observers {
		if observer == nil {
			continue
		}

		if err := <-errChan; err != nil {
			log.Println(err)
			continue
		}
	}
}

type LoggingObserver struct{}

func (o *LoggingObserver) Update(_ context.Context, _ StateSender, game *Game, move *Move) error {
	log.Println(game.Game.Position().Board().Draw())
	return nil
}
