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

//Add add node
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

//Insert insert node
func (dll *DoublyLinkList) Insert(index int, val int) *DoublyNode {
	if dll.Head == nil {
		dll.Head = &DoublyNode{val, nil, nil}
		return dll.Head
	}

	count := 0
	current := dll.Head
	for current.Next != nil {
		if index == count {
			newNode := &DoublyNode{val, current.Next, current}
			current.Next = newNode
			newNode.Next.Prev = newNode
			dll.Size++
			return dll.Head
		}
		current = current.Next
		count++
	}

	newNode := &DoublyNode{val, nil, current}
	current.Next = newNode
	dll.Size++

	return dll.Head
}

//
func (dll *DoublyLinkList) Remove(val int) bool {
	if dll.Head == nil {
		return false
	}

	current := dll.Head
	var previous *DoublyNode

	if current.Val == val {
		if current.Next != nil {
			dll.Head = current.Next
			dll.Head.Prev = nil
		} else {
			dll.Head = nil
		}
		return true
	}

	for current != nil {
		if current.Val == val {
			if current.Next != nil {
				previous.Next = current.Next
				previous.Next.Prev = previous
			} else {
				previous.Next = current.Next
			}
			return true
		}
		previous = current
		current = current.Next
	}
	
	return false
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
