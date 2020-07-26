package main

import (
	"fmt"
)

type Item struct {
	Value    interface{}
	nextItem *Item
}

type Stack struct {
	Top  *Item
	Size uint64
}

//Size
func (s *Stack) size() uint64 {
	return s.Size
}

//isEmptyS
func (s *Stack) isEmpty() bool {
	if s.Size == 0 {
		return true
	} else {
		return false
	}
}

//push
func (s *Stack) push(newItem *Item) {
	if s.Top == nil {
		s.Top = newItem
	} else {
		tempItem := s.Top
		s.Top = newItem
		s.Top.nextItem = tempItem
	}
	s.Size++
}

//pop
func (s *Stack) pop() Item {

	returnItem := Item{Value: s.Top.Value, nextItem: s.Top.nextItem}
	s.Top = s.Top.nextItem
	s.Size--
	return returnItem

}

//peak
func (s *Stack) peak() {
	fmt.Printf("%v %p\n", s.Top.Value, s.Top.nextItem)
}

//print stack
func (s *Stack) print() {
	tempItem := s.Top
	for tempItem != nil {
		fmt.Printf("%v %p\n", tempItem.Value, tempItem.nextItem)
		tempItem = tempItem.nextItem
	}
}

func main() {
}
