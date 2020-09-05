package LinkList

type CircularNode struct {
	Val  int
	Next *CircularNode
}

type CircularLinkList struct {
	Head *CircularNode
}
