package main

import (
	"errors"
	"fmt"
)

type Item struct {
	value    interface{}
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
func (s *Stack) pop() (interface{}, error) {

	if s.isEmpty() {
		return nil, errors.New("stack is empty")
	} else {
		returnItem := Item{value: s.Top.value, nextItem: s.Top.nextItem}
		s.Top = s.Top.nextItem
		s.Size--
		return returnItem, nil
	}
}

//peak
func (s *Stack) peak() {
	fmt.Printf("%v %p\n", s.Top.value, s.Top.nextItem)
}

//print stack
func (s *Stack) print() {
	tempItem := s.Top
	for tempItem != nil {
		fmt.Printf("%v %p\n", tempItem.value, tempItem.nextItem)
		tempItem = tempItem.nextItem
	}
}

func main() {
}
