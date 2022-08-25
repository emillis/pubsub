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
