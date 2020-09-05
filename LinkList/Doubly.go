package LinkList

import "fmt"

//DoublyNode doubly node
type DoublyNode struct {
	Val        int
	Next, Prev *DoublyNode
}

//DoublyLinkList doubly linked list
type DoublyLinkList struct {
	Head *DoublyNode
	Size int
}

//NewDoublyLinkList create DoublyLinkList
func NewDoublyLinkList() *DoublyLinkList {
	return &DoublyLinkList{nil, 0}
}

//
func (dll *DoublyLinkList) Add(val int) *DoublyNode {
	if dll.Head == nil {
		dll.Head = &DoublyNode{val, nil, nil}
		dll.Size++
		return dll.Head
	}

	current := dll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &DoublyNode{val, nil, current}
	dll.Size++
	return dll.Head
}

//
func (dll *DoublyLinkList) Insert() {

}

//
func (dll *DoublyLinkList) Remove() {

}

//
func (dll *DoublyLinkList) Reverse() {

}

//
func (dll *DoublyLinkList) Print() {
	if dll.Head == nil {
		fmt.Println("no node")
		return
	}

	current := dll.Head

	for current != nil {
		fmt.Printf("%d->", current.Val)
		current = current.Next
	}

	fmt.Println()
}
