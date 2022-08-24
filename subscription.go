package pubsub

//===========[TYPES]====================================================================================================

type Subscription[TVal any] struct {
	id string
	e  *Event[TVal]
}

//Cancel cancels the subscription
func (s Subscription[TVal]) Cancel() {
	if s.e == nil {
		return
	}

	s.e.Unsubscribe(s.id)
}
