package singlelinkedlist

import "fmt"

type node struct {
	value interface{}
	next *node
}


type linkedList struct {
	size int
	head *node
}


func (ll *linkedList) empty() bool {
	if ll.head == nil {
		return true
	} else {
		return false
	}
}


func (ll *linkedList) pushFront(newValue interface{}) {
	newNode := new(node)
	newNode.value = newValue
	if ll.head == nil {
		ll.head = newNode
		ll.size += 1
	} else {
		newNode.next = ll.head
		ll.head = newNode
	}
}


func (ll *linkedList) pushBack(newValue interface{}) {
	newNode := new(node)
	newNode.value = newValue
	tempHead := ll.head
	if tempHead == nil {
		ll.head = newNode
	} else {
		for tempHead.next != nil {
			tempHead = tempHead.next
		}
		tempHead.next = newNode
	}
	ll.size += 1
}

func (ll *linkedList) printList() {
	if ll.empty() {
		fmt.Println(nil)
	} else {
		for cur := ll.head; cur != nil; cur = cur.next {
			fmt.Printf("%+v\n", cur)
		}
	}
}


func (ll *linkedList) popFront() *node {
	if ll.empty() {
		return nil

	} else {
		returnNode := ll.head
		ll.head = ll.head.next
		return returnNode
	}
}


func (ll *linkedList) popBack() *node {
	tempCur := ll.head
	tempPrev := tempCur
	for tempCur.next != nil {
		tempPrev = tempCur
		tempCur = tempCur.next
	}
	tempPrev.next = nil
	return tempCur
}

func (ll *linkedList) find(value interface{}) *node {
	for cur := ll.head; cur != nil; cur = cur.next {
		if cur.value == value{
			return cur
		}
	}
	return nil
}


func (ll *linkedList) findKth(kth int) *node {
	tempKth := 0
	for cur := ll.head; cur != nil; cur = cur.next {
		if tempKth == kth{
			return cur
		}
		tempKth += 1
	}
	return nil
}


func (ll *linkedList) clear() {

	for ll.popFront() != nil{
	}

}


func (ll *linkedList) peakBack() {
	tempHead := ll.head
	if tempHead == nil {
		fmt.Println(nil)
	} else {
		for tempHead.next != nil {
			tempHead = tempHead.next
		}
		fmt.Printf("%+v\n", tempHead)
	}
}

func (ll *linkedList) peakFront() {
	if ll.empty(){
		fmt.Println(nil)
	} else {
		fmt.Printf("%+v\n", ll.head)
	}
}