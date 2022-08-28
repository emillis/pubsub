package pubsub

//===========[TYPES]====================================================================================================

//Subscription defines subscription type
type Subscription[TVal any] struct {
	id string
	e  *Event[TVal]
}

//Id returns the subscription ID
func (s Subscription[TVal]) Id() string {
	return s.id
}

//Cancel cancels the subscription
func (s Subscription[TVal]) Cancel() {
	if s.e == nil {
		return
	}

	s.e.Unsubscribe(s.id)
}
