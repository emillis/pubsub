package test

import (
	"pubsub"
	"reflect"
	"testing"
)

func Test_New(t *testing.T) {
	event := pubsub.New[string]()

	expected := "pubsub.Event[string]"
	got := reflect.TypeOf(event).String()

	if expected != got {
		t.Errorf("Got wrong type. Expected %s, got %s", expected, got)
	}
}

func TestEvent_Publish(t *testing.T) {
	e := pubsub.New[string]()
	expecting := "hello_world"
	got := ""

	callback := func(s string) {
		got = expecting
	}

	e.Subscribe(callback)

	wg := e.Publish(expecting)

	wg.Wait()

	if got != expecting {
		t.Errorf("In TestEvent_Publish expeted to have %s, got %s", expecting, got)
	}
}

func TestEvent_Subscribe(t *testing.T) {
	e := pubsub.New[string]()
	expecting := "hello_world"
	got := ""

	callback := func(s string) {
		got = expecting
	}

	e.Subscribe(callback)

	wg := e.Publish(expecting)

	wg.Wait()

	if got != expecting {
		t.Errorf("In Event_Subscribe expeted to have %s, got %s", expecting, got)
	}
}

func TestSubscription_Cancel(t *testing.T) {
	e := pubsub.New[string]()
	expecting := "hello_world"
	got := ""

	callback := func(s string) {
		got = expecting
	}

	subscription := e.Subscribe(callback)

	wg := e.Publish(expecting)

	wg.Wait()

	if got != expecting {
		t.Errorf("Expected to have `got` and `expecting` variables to be the same, got %s and %s", expecting, got)
	}

	subscription.Cancel()

	wg = e.Publish("new_value")

	wg.Wait()

	if got != expecting {
		t.Errorf("After subscription cancellation, expected to have same values in variables `got` and `expected`, got %s and %s", expecting, got)
	}
}

func TestEvent_Unsubscribe(t *testing.T) {
	e := pubsub.New[string]()
	expecting := "hello_world"
	got := ""

	callback := func(s string) {
		got = expecting
	}

	subscription := e.Subscribe(callback)

	wg := e.Publish(expecting)

	wg.Wait()

	if got != expecting {
		t.Errorf("Expected to have `got` and `expecting` variables to be the same, got %s and %s", expecting, got)
	}

	e.Unsubscribe(subscription.Id())

	wg = e.Publish("new_value")

	wg.Wait()

	if got != expecting {
		t.Errorf("After unsubscription, expected to have same values in variables `got` and `expected`, got %s and %s", expecting, got)
	}
}
