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

//NewCircularLinkList create CircularLinkList
func NewCircularLinkList() *CircularLinkList {
	return &CircularLinkList{nil, 0}
}

//Add add node
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
	cll.Size++

	return cll.Head
}

//Insert insert node
func (cll *CircularLinkList) Insert(index int, val int) *CircularNode {
	if cll.Head == nil {
		cll.Head = &CircularNode{val, nil}
		cll.Head.Next = cll.Head
		cll.Size++
		return cll.Head
	}

	current := cll.Head
	count := 0

	for current.Next != cll.Head {
		if index == count {
			node := &CircularNode{val, current.Next}
			current.Next = node
			cll.Size++
			return cll.Head
		}
		current = current.Next
		count++
	}

	node := &CircularNode{val, current.Next}
	current.Next = node
	cll.Size++

	return cll.Head
}

//Remove remove node
func (cll *CircularLinkList) Remove(val int) bool {
	if cll.Head == nil {
		return false
	}

	current := cll.Head
	var previous *CircularNode

	// head 是移除對象
	if current.Val == val {
		newHead := current.Next
		for current.Next != cll.Head {
			current = current.Next
		}
		current.Next = newHead
		cll.Head = newHead
		cll.Size--
		return true
	}

	previous = current
	current = current.Next

	for current.Next != cll.Head {
		if current.Val == val {
			previous.Next = current.Next
			cll.Size--
			return true
		}
		previous = current
		current = current.Next
	}

	//尾部是刪除對象
	if current.Val == val {
		previous.Next = cll.Head
		cll.Size--
		return true
	}

	return false
}

//Reverse reverse linked list
func (cll *CircularLinkList) Reverse() *CircularNode {
	if cll.Head == nil || cll.Head.Next == cll.Head {
		return cll.Head
	}

	tail := cll.Head
	current := cll.Head
	preceding := cll.Head.Next
	var previous *CircularNode

	for preceding.Next != cll.Head {
		current.Next = previous
		previous = current
		current = preceding
		preceding = preceding.Next
	}

	current.Next = previous
	preceding.Next = current
	tail.Next = preceding //尾部接上
	cll.Head = preceding

	return cll.Head
}

//Print print node
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
	fmt.Printf("%d->", current.Val)
	fmt.Println()
}
