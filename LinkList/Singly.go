package LinkList

import "fmt"

//SinglyNode singly node
type SinglyNode struct {
	Val  int
	Next *SinglyNode
}

//SinglyLinkList singly linklist
type SinglyLinkList struct {
	Head *SinglyNode
}

//NewSinglyLinkList create new singly linklist
func NewSinglyLinkList() *SinglyLinkList {
	return &SinglyLinkList{nil}
}

//AddNode add node
func (sll *SinglyLinkList) AddNode(val int) *SinglyNode {
	if sll.Head == nil {
		sll.Head = &SinglyNode{val, nil}
		return sll.Head
	}

	current := sll.Head

	for current.Next != nil {
		current = current.Next
	}

	current.Next = &SinglyNode{val, nil}

	return sll.Head
}

//InsertNode insert node
func (sll SinglyLinkList) InsertNode(index, val int) *SinglyNode {
	if sll.Head == nil {
		sll.Head = &SinglyNode{val, nil}
		return sll.Head
	}

	var count int
	current := sll.Head

	for current.Next != nil {
		if count == index {
			node := &SinglyNode{val, nil}
			node.Next = current.Next
			current.Next = node
			return sll.Head
		}
		current = current.Next
		count++
	}

	current.Next = &SinglyNode{val, nil}

	return sll.Head
}

// RemoveNode remove node
func (sll *SinglyLinkList) RemoveNode(val int) bool {
	if sll.Head == nil {
		return false
	}

	current := sll.Head
	var previous *SinglyNode

	for current != nil {
		if current.Val == val {
			if previous == nil {
				sll.Head = current.Next
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

// PrintNodes 1->2->3
func (sll *SinglyLinkList) PrintNodes() {
	if sll.Head == nil {
		fmt.Println("no node")
		return
	}

	current := sll.Head

	for current.Next != nil {
		fmt.Printf("%d->", current.Val)
		current = current.Next
	}
	fmt.Println()

}
