// Package queue provides the queue abstract data type and associated methods
package queue

import (
	"github.com/jttait/godsa/singlylinkedlist"
)

// Queue is a first-in, first-out data structure.
type Queue[T any] struct {
	dummyHead *singlylinkedlist.SinglyLinkedListNode[T]
	dummyTail *singlylinkedlist.SinglyLinkedListNode[T]
	size      int
}

// NewQueue instantiates a new queue and returns a pointer to it.
func NewQueue[T any]() *Queue[T] {
	q := Queue[T]{}
	var zeroValue T
	q.dummyHead = singlylinkedlist.NewSinglyLinkedListNode[T](zeroValue)
	q.dummyTail = singlylinkedlist.NewSinglyLinkedListNode[T](zeroValue)
	q.dummyHead.Next = q.dummyTail
	return &q
}

// Size returns the number of items in the queue.
func (q *Queue[T]) Size() int {
	return q.size
}

// Add inserts a new item at the end of the queue.
func (q *Queue[T]) Add(i T) {
	q.dummyTail.Val = i
	var zeroValue T
	q.dummyTail.Next = singlylinkedlist.NewSinglyLinkedListNode[T](zeroValue)
	q.dummyTail = q.dummyTail.Next
	q.size += 1
}

// Remove removes and returns the item at the head of the queue.
func (q *Queue[T]) Remove() (T, bool) {
	if q.Size() == 0 {
		var zeroValue T
		return zeroValue, false
	}
	result := q.dummyHead.Next.Val
	q.dummyHead = q.dummyHead.Next
	q.size -= 1
	return result, true
}
