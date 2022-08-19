package pubsub

import (
	"github.com/emillis/cacheMachine"
)

type Event[TVal any] struct {
	subscribers cacheMachine.Cache[string, func(v Trigger[TVal])]
}

//TODO: Add method, unsubscribe(id string)

func (e *Event[TVal]) New(v TVal) {

}

//TODO: return subscription that the client can use to cancel it

func (e *Event[TVal]) Subscribe(callback func(v Trigger[TVal])) {
	e.subscribers.Add("1", callback)
}
