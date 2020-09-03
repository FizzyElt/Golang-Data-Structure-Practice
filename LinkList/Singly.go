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

//Add add node
func (sll *SinglyLinkList) Add(val int) *SinglyNode {
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

//Insert insert node
func (sll SinglyLinkList) Insert(index, val int) *SinglyNode {
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

// Remove remove node
func (sll *SinglyLinkList) Remove(val int) bool {
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

// Reverse reverse Linklist
func (sll *SinglyLinkList) Reverse() *SinglyNode {
	if sll.Head == nil || sll.Head.Next == nil {
		return sll.Head
	}

	current := sll.Head
	preceding := sll.Head.Next
	var previous *SinglyNode

	for preceding != nil {
		current.Next = previous
		previous = current
		current = preceding
		preceding = preceding.Next
	}

	current.Next=previous
	sll.Head = current

	return sll.Head
}

// Print 1->2->3
func (sll *SinglyLinkList) Print() {
	if sll.Head == nil {
		fmt.Println("no node")
		return
	}

	current := sll.Head

	for current != nil {
		fmt.Printf("%d->", current.Val)
		current = current.Next
	}
	fmt.Println()

}
