// Package doubleendedqueue provides the double-ended queue abstract data type and associated
// methods
package doubleendedqueue

import "github.com/jttait/godsa/doublylinkedlist"

// Double-ended queue enables inserting and polling from either end
type DoubleEndedQueue[T any] struct {
	dummyHead *doublylinkedlist.DoublyLinkedListNode[T]
	dummyTail *doublylinkedlist.DoublyLinkedListNode[T]
	size      int
}

// NewDoubleEndedQueue instantiates a new double-ended queue and returns a pointer to it.
func NewDoubleEndedQueue[T any]() *DoubleEndedQueue[T] {
	q := DoubleEndedQueue[T]{}
	var zeroValue T
	q.dummyHead = doublylinkedlist.NewDoublyLinkedListNode[T](zeroValue)
	q.dummyTail = doublylinkedlist.NewDoublyLinkedListNode[T](zeroValue)
	q.dummyHead.Next = q.dummyTail
	q.dummyTail.Prev = q.dummyHead
	q.size = 0
	return &q
}

// Size returns the number of items in the double-ended queue.
func (q *DoubleEndedQueue[T]) Size() int {
	return q.size
}

// Add inserts a new item at the end of the queue.
func (q *DoubleEndedQueue[T]) InsertFront(i T) {
	insertedNode := doublylinkedlist.NewDoublyLinkedListNode[T](i)
	currentHead := q.dummyHead.Next

	insertedNode.Prev = q.dummyHead
	insertedNode.Next = currentHead
	currentHead.Prev = insertedNode
	q.dummyHead.Next = insertedNode

	q.size += 1
}

func (q *DoubleEndedQueue[T]) InsertLast(i T) {
	insertedNode := doublylinkedlist.NewDoublyLinkedListNode[T](i)
	currentTail := q.dummyTail.Prev

	insertedNode.Next = q.dummyTail
	insertedNode.Prev = currentTail
	currentTail.Next = insertedNode
	q.dummyTail.Prev = insertedNode

	q.size += 1
}

// Remove removes and returns the item at the front of the double-ended queue.
func (q *DoubleEndedQueue[T]) RemoveFront() (T, bool) {
	if q.size == 0 {
		var zeroValue T
		return zeroValue, false
	}
	result := q.dummyHead.Next.Val
	head := q.dummyHead.Next
	newHead := head.Next
	q.dummyHead.Next = newHead
	newHead.Prev = q.dummyHead
	q.size -= 1
	return result, true
}

func (q *DoubleEndedQueue[T]) RemoveLast() (T, bool) {
	if q.size == 0 {
		var zeroValue T
		return zeroValue, false
	}
	result := q.dummyTail.Prev.Val
	tail := q.dummyTail.Prev
	newTail := tail.Prev
	q.dummyTail.Prev = newTail
	newTail.Next = q.dummyTail
	q.size -= 1
	return result, true
}
