package dequeue

import (
	"github.com/jttait/godsa/linkedlist"
)

// LLDEQueue is a double-ended queue implemented using a doubly-linked list.
type LLDEQueue[T any] struct {
	list *linkedlist.DLL[T]
}

// NewLLDEQueue instantiates a new queue and returns a pointer to it.
func NewLLDEQueue[T any](values ...T) *LLDEQueue[T] {
	q := LLDEQueue[T]{}
	q.list = linkedlist.NewDLL[T]()
	for _, v := range values {
		q.list.InsertLast(v)
	}
	return &q
}

// Size returns the number of items in the queue.
func (q *LLDEQueue[T]) Size() int {
	return q.list.Size()
}

// InsertFront inserts a new item at the front of the queue.
func (q *LLDEQueue[T]) InsertFront(i T) {
	q.list.InsertFront(i)
}

// InsertLast inserts a new item at the end of the queue.
func (q *LLDEQueue[T]) InsertLast(i T) {
	q.list.InsertLast(i)
}

// RemoveFirst removes and returns the item at the front of the queue.
func (q *LLDEQueue[T]) RemoveFront() (T, bool) {
	return q.list.RemoveFront()
}

// RemoveLast removes and returns the item at the end of the queue.
func (q *LLDEQueue[T]) RemoveLast() (T, bool) {
	return q.list.RemoveLast()
}
