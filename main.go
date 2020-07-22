package main

import "fmt"

type Item struct {
	value    interface{}
	nextItem *Item
}

type Stack struct {
	top   *Item
	count uint64
}

//push
func (s *Stack) push(newItem *Item) {
	if s.top == nil {
		s.top = newItem
	} else {
		tempItem := s.top
		s.top = newItem
		s.top.nextItem = tempItem
	}
	s.count++
}

//pop
func (s *Stack) pop() Item {
	returnItem := Item{value: s.top.value, nextItem: s.top.nextItem}
	s.top = s.top.nextItem
	s.count--
	return returnItem
}

//peak
func (s *Stack) peak() {
	fmt.Printf("%v %p\n", s.top.value, s.top.nextItem)
}

//print stack
func (s *Stack) print() {

	tempItem := s.top
	for tempItem != nil {
		fmt.Printf("%v %p\n", tempItem.value, tempItem.nextItem)
		tempItem = tempItem.nextItem
	}
}

func main() {
	Item1 := Item{value: "TEST Item", nextItem: nil}
	stack := Stack{top: nil, count: 0}
	fmt.Println(Item1)

	stack.push(&Item1)
	stack.print()

	newItems := []string{"watermelon"}

	for _, item := range newItems {
		newItem := Item{value: item, nextItem: nil}
		stack.push(&newItem)
	}

	stack.print()
	fmt.Println(stack.count)

	// Pop test
	ItemPop := stack.pop()
	fmt.Println(ItemPop)
	stack.print()

	//Peak test
	Item2 := Item{value: "KENNETH Item", nextItem: nil}
	Item5 := Item{value: "KENNETH Item5", nextItem: nil}
	stack.push(&Item2)

	stack.push(&Item5)
	Item6 := Item{value: 2, nextItem: nil}
	stack.push(&Item6)
	stack.peak()
	fmt.Println(stack.count)

	stack.print()
}
