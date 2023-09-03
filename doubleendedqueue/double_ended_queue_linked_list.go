// Package doubleendedqueue provides the double-ended queue abstract data type and associated
// methods
package doubleendedqueue

import "github.com/jttait/godsa/linkedlist"

// Double-ended queue enables inserting and polling from either end
type DoubleEndedQueueLinkedList[T any] struct {
	list *linkedlist.DoublyLinkedList[T]
}

// NewDoubleEndedQueue instantiates a new double-ended queue and returns a pointer to it.
func NewDoubleEndedQueueLinkedList[T any]() *DoubleEndedQueueLinkedList[T] {
	q := DoubleEndedQueueLinkedList[T]{}
	q.list = linkedlist.NewDoublyLinkedList[T]()
	return &q
}

// Size returns the number of items in the double-ended queue.
func (q *DoubleEndedQueueLinkedList[T]) Size() int {
	return q.list.Size()
}

// InsertFront inserts a new item at the front of the double-ended queue.
func (q *DoubleEndedQueueLinkedList[T]) InsertFront(i T) {
	q.list.InsertFront(i)
}

// InsertLast inserts a new item at the end of the double-ended queue.
func (q *DoubleEndedQueueLinkedList[T]) InsertLast(i T) {
	q.list.InsertLast(i)
}

// RemoveFirst removes and returns the item at the front of the double-ended queue.
func (q *DoubleEndedQueueLinkedList[T]) RemoveFront() (T, bool) {
	return q.list.RemoveFront()
}

// RemoveLast removes and returns the item at the end of the double-ended queue.
func (q *DoubleEndedQueueLinkedList[T]) RemoveLast() (T, bool) {
	return q.list.RemoveLast()
}
