// Package singlylinkedlist provides the singly-linked list data structure and associated methods
package singlylinkedlist

// SinglyLinkedListNode is a data structure that contains a value and a pointer to the next node in
// the list.
type SinglyLinkedListNode struct {
	Val  int
	Next *SinglyLinkedListNode
}

// NewSinglyLinkedListNode instantiates a singly-linked list node and returns a pointer to it.
func NewSinglyLinkedListNode(i int) *SinglyLinkedListNode {
	n := SinglyLinkedListNode{}
	n.Val = i
	return &n
}

// InsertAfter inserts a new node after the current one. If there is already a node there then the
// newly-inserted node will point to this after the insert operation is complete.
func (n *SinglyLinkedListNode) InsertAfter(i int) {
	if n.Next == nil {
		n.Next = NewSinglyLinkedListNode(i)
		return
	}
	temp := n.Next
	n.Next = NewSinglyLinkedListNode(i)
	n.Next.Next = temp
}

// RemoveNext will remove the next node from the linked list. The Next pointer for the current node
// will be updated to point to the one after the removed node.
func (n *SinglyLinkedListNode) RemoveNext() {
	n.Next = n.Next.Next
}
