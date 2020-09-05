package LinkList

import "fmt"

// CircularNode circular node
type CircularNode struct {
	Val  int
	Next *CircularNode
}

//CircularLinkList circular linked list
type CircularLinkList struct {
	Head *CircularNode
	Size int
}

func NewCircularLinkList() *CircularLinkList {
	return &CircularLinkList{nil, 0}
}

func (cll *CircularLinkList) Add(val int) *CircularNode {
	if cll.Head == nil {
		cll.Head = &CircularNode{val, nil}
		cll.Head.Next = cll.Head
		cll.Size++
		return cll.Head
	}
	current := cll.Head
	for current.Next != cll.Head {
		current = current.Next
	}
	current.Next = &CircularNode{val, cll.Head}
	return cll.Head
}

func (cll *CircularLinkList) Insert() {

}

func (cll *CircularLinkList) Remove() {

}

func (cll *CircularLinkList) Reverse() {

}

func (cll *CircularLinkList) Print() {
	if cll.Head == nil {
		fmt.Println("no node")
		return
	}
	current := cll.Head
	for current.Next != cll.Head {
		fmt.Printf("%d->", current.Val)
		current = current.Next
	}

	fmt.Println()
}
