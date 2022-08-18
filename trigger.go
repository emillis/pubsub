package pubsub

import (
	"sync"
)

//===========[INTERFACES]====================================================================================================

type IDer interface {
	Id() string
}

//===========[STRUCTS]====================================================================================================

//Trigger stores the value to be sent to subscribers
type Trigger[TVal any] struct {
	val          TVal
	eventHistory map[string]struct{}
	mx           sync.RWMutex
}

//Value returns the value passed along with the trigger
func (t *Trigger[TVal]) Value() TVal {
	t.mx.RLock()
	defer t.mx.RUnlock()
	return t.val
}

//SetValue replaces the value in the trigger
func (t *Trigger[TVal]) SetValue(v TVal) {
	t.mx.Lock()
	defer t.mx.Unlock()
	t.val = v
}

//AddEventId adds a new event ID to internal cache. This is used later to prevent infinite loops from forming when
//subscribers call other events that call the initial event and in turn starting the infinite loop
func (t *Trigger[TVal]) AddEventId(id string) {
	t.mx.Lock()
	defer t.mx.Unlock()
	t.eventHistory[id] = struct{}{}
}

//CreateEventException allows you to pass this trigger back to the event for the second time
//func (t *Trigger[TVal]) CreateEventException(id string) {
//	t.mx.Lock()
//	defer t.mx.Unlock()
//	t.eventHistory[id] = struct{}{}
//}

//===========[FUNCTIONALITY]====================================================================================================

//NewTrigger initiates and returns a new Trigger
func NewTrigger[TVal any](v TVal) Trigger[TVal] {
	return Trigger[TVal]{
		val:          v,
		eventHistory: make(map[string]struct{}),
		mx:           sync.RWMutex{},
	}
}
