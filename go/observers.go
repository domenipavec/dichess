package main

import (
	"log"
	"sync"

	"github.com/notnil/chess"
)

type Observer interface {
	Update(*chess.Game) error
}

type Observers struct {
	observers []Observer
	mutex     sync.Mutex
}

func (o *Observers) Add(observer Observer) {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	o.observers = append(o.observers, observer)
}

// Update calls update on all observers, removing observers that return error.
func (o *Observers) Update(game *chess.Game) {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	i := 0
	for _, observer := range o.observers {
		if err := observer.Update(game); err != nil {
			log.Println(err)
		}
		o.observers[i] = observer
	}
	o.observers = o.observers[:i]
}
