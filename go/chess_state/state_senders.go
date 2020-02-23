package chess_state

import (
	"log"
	"sync"
)

// StateSender should use game.Game as read only
type StateSender interface {
	StateSend(string)
}

type StateSenders struct {
	lastState    string
	stateSenders []StateSender
	mutex        sync.Mutex
}

func (o *StateSenders) Add(stateSender StateSender) int {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	o.stateSenders = append(o.stateSenders, stateSender)

	return len(o.stateSenders) - 1
}

func (o *StateSenders) Remove(id int) {
	o.stateSenders[id] = nil
}

func (o *StateSenders) GetLastState() string {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	return o.lastState
}

// StateSend calls StateSend on all stateSenders
func (o *StateSenders) StateSend(state string) {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	if o.lastState == state {
		return
	}
	o.lastState = state

	for _, stateSender := range o.stateSenders {
		if stateSender == nil {
			continue
		}
		stateSender := stateSender
		go func() {
			stateSender.StateSend(state)
		}()
	}
}

type LoggingStateSender struct{}

func (o *LoggingStateSender) StateSend(state string) {
	log.Printf("State: %s", state)
}
