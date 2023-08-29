// Package doublylinkedlist provides the doubly-linked list data structure and associated methods
package doublylinkedlist

// DoublyLinkedListNode is a data structure that contains a value and pointers to the previous and
// next nodes in the list.
type DoublyLinkedListNode[T any] struct {
	Val  T
	Prev *DoublyLinkedListNode[T]
	Next *DoublyLinkedListNode[T]
}

// NewDoublyLinkedListNode instantiates a doubly-linked list node and returns a pointer to it.
func NewDoublyLinkedListNode[T any](i T) *DoublyLinkedListNode[T] {
	d := DoublyLinkedListNode[T]{}
	d.Val = i
	return &d
}

// InsertAfter inserts a new node after the current one. If there is already a node there then the
// previous and next pointers of all three nodes will be updated as required.
func (n *DoublyLinkedListNode[T]) InsertAfter(i T) {
	if n.Next == nil {
		insertedNode := NewDoublyLinkedListNode[T](i)
		insertedNode.Prev = n
		n.Next = insertedNode
		return
	}
	temp := n.Next
	insertedNode := NewDoublyLinkedListNode[T](i)
	insertedNode.Prev = n
	insertedNode.Next = temp
	n.Next = insertedNode
}

// Remove removes the current node and updates the pointers of the previous and next nodes as
// required.
func (n *DoublyLinkedListNode[T]) Remove() {
	prev := n.Prev
	next := n.Next
	if prev != nil {
		prev.Next = next
	}
	if next != nil {
		next.Prev = prev
	}
}
