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
}

func (cll *CircularLinkList) Add() {

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
	}

	fmt.Println()
}
