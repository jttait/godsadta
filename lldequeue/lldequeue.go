// Package dequeue provides the double-ended queue abstract data type and associated
// methods
package lldequeue

import (
	"github.com/jttait/godsa/dll"
)

// Double-ended queue enables inserting and polling from either end
type LLDEQueue[T any] struct {
	list *dll.DLL[T]
}

// NewDoubleEndedQueue instantiates a new double-ended queue and returns a pointer to it.
func NewLLDEQueue[T any](values ...T) *LLDEQueue[T] {
	q := LLDEQueue[T]{}
	q.list = dll.NewDLL[T]()
	for _, v := range values {
		q.list.InsertLast(v)
	}
	return &q
}

// Size returns the number of items in the double-ended queue.
func (q *LLDEQueue[T]) Size() int {
	return q.list.Size()
}

// InsertFront inserts a new item at the front of the double-ended queue.
func (q *LLDEQueue[T]) InsertFront(i T) {
	q.list.InsertFront(i)
}

// InsertLast inserts a new item at the end of the double-ended queue.
func (q *LLDEQueue[T]) InsertLast(i T) {
	q.list.InsertLast(i)
}

// RemoveFirst removes and returns the item at the front of the double-ended queue.
func (q *LLDEQueue[T]) RemoveFront() (T, bool) {
	return q.list.RemoveFront()
}

// RemoveLast removes and returns the item at the end of the double-ended queue.
func (q *LLDEQueue[T]) RemoveLast() (T, bool) {
	return q.list.RemoveLast()
}
