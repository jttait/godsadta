// Package doubleendedqueue provides the double-ended queue abstract data type and associated
// methods
package doubleendedqueue

import "github.com/jttait/godsa/doublylinkedlist"

// Double-ended queue enables inserting and polling from either end
type DoubleEndedQueue[T any] struct {
	list *doublylinkedlist.DoublyLinkedList[T]
}

// NewDoubleEndedQueue instantiates a new double-ended queue and returns a pointer to it.
func NewDoubleEndedQueue[T any]() *DoubleEndedQueue[T] {
	q := DoubleEndedQueue[T]{}
	q.list = doublylinkedlist.NewDoublyLinkedList[T]()
	return &q
}

// Size returns the number of items in the double-ended queue.
func (q *DoubleEndedQueue[T]) Size() int {
	return q.list.Size()
}

// InsertFront inserts a new item at the front of the double-ended queue.
func (q *DoubleEndedQueue[T]) InsertFront(i T) {
	q.list.InsertFront(i)
}

// InsertLast inserts a new item at the end of the double-ended queue.
func (q *DoubleEndedQueue[T]) InsertLast(i T) {
	q.list.InsertLast(i)
}

// RemoveFirst removes and returns the item at the front of the double-ended queue.
func (q *DoubleEndedQueue[T]) RemoveFront() (T, bool) {
	return q.list.RemoveFront()
}

// RemoveLast removes and returns the item at the end of the double-ended queue.
func (q *DoubleEndedQueue[T]) RemoveLast() (T, bool) {
	return q.list.RemoveLast()
}
