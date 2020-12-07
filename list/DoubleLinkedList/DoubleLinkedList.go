package doublelinkedlist

type node struct {
	value interface{}
	next  *node
	prev  *node
}

type linkedList struct {
	size int
	head *node
	tail *node
}

func (ll *linkedList) empty() bool {
	if ll.head == nil {
		return true
	} else {
		return false
	}
}

func (ll *linkedList) pushFront(newValue interface{}) {

}

func (ll *linkedList) pushBack(newValue interface{}) {

}

func (ll *linkedList) printList() {

}

func (ll *linkedList) popFront() *node {

}

func (ll *linkedList) popBack() *node {

}

func (ll *linkedList) find(value interface{}) *node {

}

func (ll *linkedList) findKth(kth int) *node {

}

func (ll *linkedList) clear() {

}

func (ll *linkedList) peakBack() {

}

func (ll *linkedList) peakFront() {

}
