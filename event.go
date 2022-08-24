package pubsub

import (
	"github.com/emillis/cacheMachine"
	"github.com/emillis/idGen"
	"sync"
)

//===========[STATIC]====================================================================================================

var jenny = idGen.NewGenerator(&idGen.Requirements{Length: 24})

//===========[STRUCTS]====================================================================================================

//Event definition
type Event[TVal any] struct {
	subscribers cacheMachine.Cache[string, func(TVal)]
	tmpIdList   cacheMachine.Cache[string, struct{}]
	mx          sync.RWMutex
}

//Publish sends the value supplied to all the subscribers to the Event
func (e *Event[TVal]) Publish(v TVal) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	subs := e.subscribers.GetAll()

	wg.Add(len(subs))

	go func() {
		for _, sub := range subs {
			sub(v)
			wg.Done()
		}
	}()

	return wg
}

//Subscribe adds a new subscription to the list to be called once a value gets published to the event
func (e *Event[TVal]) Subscribe(callback func(v TVal)) Subscription[TVal] {
	var tmpId string

	e.mx.Lock()
	for {
		tmpId = jenny.Random()

		if e.subscribers.Exist(tmpId) || e.tmpIdList.Exist(tmpId) {
			continue
		}

		e.tmpIdList.Add(tmpId, struct{}{})

		break
	}
	e.mx.Unlock()

	//Removing the ID from the temporary store once the function exits
	defer e.tmpIdList.Remove(tmpId)

	e.subscribers.Add(tmpId, callback)

	return Subscription[TVal]{
		id: tmpId,
		e:  e,
	}
}

//Unsubscribe removes callback with the supplied ID from the call list
func (e *Event[TVal]) Unsubscribe(id string) {
	e.subscribers.Remove(id)
}

//===========[FUNCTIONALITY]====================================================================================================

//New initiates and returns a new Event which you can Publish and Subscribe to
func New[TVal any]() Event[TVal] {
	return Event[TVal]{
		subscribers: cacheMachine.New[string, func(TVal)](nil),
		tmpIdList:   cacheMachine.New[string, struct{}](nil),
		mx:          sync.RWMutex{},
	}
}
