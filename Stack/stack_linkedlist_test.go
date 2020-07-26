package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test Push
func TestStack_Push(t *testing.T) {
	stack := Stack{Top: nil, Size: 0}
	stack.push(&(Item{Value: "test", nextItem: nil}))
	assert.Equal(t, "test", stack.Top.Value, "Value to top item and expected item is not equal")
	assert.NotEmpty(t, 122313, stack.Top.Value, "Value to top item and expected item is not equal")
}

// Test isEmpty
func TestStack_isEmpty(t *testing.T) {
	stack := Stack{Top: nil, Size: 0}
	stack.push(&(Item{Value: "test", nextItem: nil}))
	assert.Equal(t, stack.isEmpty(), false)

	stack.pop()
	assert.Equal(t, stack.isEmpty(), true)

}

func TestStack_Pop(t *testing.T) {
	stack := Stack{Top: nil, Size: 0}
	stack.push(&(Item{Value: "test", nextItem: nil}))
	stack.push(&(Item{Value: 1234, nextItem: nil}))

	// test first pop
	item1 := stack.pop()
	assert.Equal(t, item1.Value, 1234)

	// test second pop
	item2 := stack.pop()
	assert.Equal(t, item2.Value, "test")

}
