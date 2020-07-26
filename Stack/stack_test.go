package main

import (
	"testing"
)

// Test Push
func testPush(t *testing.T) {
	stack := Stack{top: nil, count: 0}
	stack.push(&(Item{value: "TEST Item", nextItem: nil}))

	if stack.top.value != "TEST Item" {
		t.Error("Expected Item value to be: \"TEST Item\" but got ", stack.top.value)
	}
}
