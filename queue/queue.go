// Package queue provides the queue abstract data type and associated methods
package queue

import (
	"github.com/jttait/godsa/doublylinkedlist"
)

// Queue is a first-in, first-out data structure.
type Queue[T any] struct {
	list *doublylinkedlist.DoublyLinkedList[T]
}

// NewQueue instantiates a new queue and returns a pointer to it.
func NewQueue[T any]() *Queue[T] {
	q := Queue[T]{}
	q.list = doublylinkedlist.NewDoublyLinkedList[T]()
	return &q
}

// Size returns the number of items in the queue.
func (q *Queue[T]) Size() int {
	return q.list.Size()
}

// Add inserts a new item at the end of the queue.
func (q *Queue[T]) Insert(i T) {
	q.list.InsertLast(i)
}

// Remove removes and returns the item at the head of the queue.
func (q *Queue[T]) Remove() (T, bool) {
	return q.list.RemoveFront()
}
