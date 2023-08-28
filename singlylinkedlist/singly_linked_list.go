// Package singlylinkedlist provides the singly-linked list data structure and associated methods
package singlylinkedlist

// SinglyLinkedListNode is a data structure that contains a value and a pointer to the next node in
// the list.
type SinglyLinkedListNode[T any] struct {
	Val  T
	Next *SinglyLinkedListNode[T]
}

// NewSinglyLinkedListNode instantiates a singly-linked list node and returns a pointer to it.
func NewSinglyLinkedListNode[T any](i T) *SinglyLinkedListNode[T] {
	n := SinglyLinkedListNode[T]{}
	n.Val = i
	return &n
}

// InsertAfter inserts a new node after the current one. If there is already a node there then the
// newly-inserted node will point to this after the insert operation is complete.
func (n *SinglyLinkedListNode[T]) InsertAfter(i T) {
	if n.Next == nil {
		n.Next = NewSinglyLinkedListNode[T](i)
		return
	}
	temp := n.Next
	n.Next = NewSinglyLinkedListNode[T](i)
	n.Next.Next = temp
}

// RemoveNext will remove the next node from the linked list. The Next pointer for the current node
// will be updated to point to the one after the removed node.
func (n *SinglyLinkedListNode[T]) RemoveNext() {
	n.Next = n.Next.Next
}
