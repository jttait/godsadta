// Package queue provides the queue abstract data type and associated methods
package queue

import "github.com/jttait/godsa/linkedlist"

// Queue is a first-in, first-out data structure.
type QueueLinkedList[T any] struct {
	list *linkedlist.DoublyLinkedList[T]
}

// NewQueue instantiates a new queue and returns a pointer to it.
func NewQueueLinkedList[T any]() *QueueLinkedList[T] {
	q := QueueLinkedList[T]{}
	q.list = linkedlist.NewDoublyLinkedList[T]()
	return &q
}

// Size returns the number of items in the queue.
func (q *QueueLinkedList[T]) Size() int {
	return q.list.Size()
}

// Add inserts a new item at the end of the queue.
func (q *QueueLinkedList[T]) Insert(i T) {
	q.list.InsertLast(i)
}

// Remove removes and returns the item at the head of the queue.
func (q *QueueLinkedList[T]) Remove() (T, bool) {
	return q.list.RemoveFront()
}
