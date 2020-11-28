package singlelinkedlist

import (
	"reflect"
	"testing"
)

func TestLinkedList(t *testing.T) {
	list := new(linkedList)
	list.pushFront(1)
	list.pushFront("test")
	list.pushFront(3.0)

	t.Run("Test pushFront()", func(t *testing.T) {
		want := []interface{}{3.0, "tests", 1}
		got := []interface{}{}
		current := list.head
		got = append(got, current.value)
		for current.next != nil {
			current = current.next
			got = append(got, current.value)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}