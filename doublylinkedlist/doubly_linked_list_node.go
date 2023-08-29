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
