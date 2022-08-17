package pubsub

import (
	"sync"
)

type Trigger[TVal any] struct {
	val TVal
	mx  sync.RWMutex
}

func (t *Trigger[TVal]) Value() TVal {
	t.mx.RLock()
	defer t.mx.RUnlock()
	return t.val
}

//===========[FUNCTIONALITY]====================================================================================================

func NewTrigger[TVal any](v TVal) Trigger[TVal] {
	return Trigger[TVal]{
		val: v,
		mx:  sync.RWMutex{},
	}
}
