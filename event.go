package pubsub

import (
	"github.com/emillis/cacheMachine"
)

type Trigger[TVal any] TVal

type Event[TVal any] struct {
	subscribers cacheMachine.Cache[string, func(TVal)]
}

//TODO: Add method, unsubscribe(id string)

func (e *Event[TVal]) Trigger(v TVal) {

}

//TODO: return subscription that the client can use to cancel it

func (e *Event[TVal]) Subscribe(callback func(v TVal)) {
	e.subscribers.Add("1", callback)
}
