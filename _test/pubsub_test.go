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
