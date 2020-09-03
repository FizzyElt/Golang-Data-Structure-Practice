package LinkList

import "fmt"

type SinglyNode struct {
	Val  int
	Next *SinglyNode
}
type SinglyLinkList struct {
	Head *SinglyNode
}

func NewSinglyLinkList() *SinglyLinkList {
	return &SinglyLinkList{nil}
}

func (sll *SinglyLinkList) AddNode(val int) *SinglyNode {
	if sll.Head == nil {
		sll.Head = &SinglyNode{val, nil}
		return sll.Head
	}

	var current *SinglyNode = sll.Head

	for current.Next != nil {
		current = current.Next
	}

	current.Next = &SinglyNode{val, nil}

	return sll.Head
}

func (sll SinglyLinkList) InsertNode(index, val int) *SinglyNode {
	if sll.Head == nil {
		sll.Head = &SinglyNode{val, nil}
		return sll.Head
	}

	var count int = 0
	var current *SinglyNode = sll.Head

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

func (sll *SinglyLinkList) RemoveNode(val int) bool {
	if sll.Head == nil {
		return false
	}
	var current *SinglyNode = sll.Head
	var previous *SinglyNode = nil
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

func (sll *SinglyLinkList) PrintNodes() {
	if sll.Head == nil {
		fmt.Println("no node")
		return
	}

	var current *SinglyNode = sll.Head

	for current.Next != nil {
		fmt.Printf("%d->", current.Val)
		current = current.Next
	}
	fmt.Println()

}
