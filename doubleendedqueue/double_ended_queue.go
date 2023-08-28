// Package doubleendedqueue provides the double-ended queue abstract data type and associated
// methods
package doubleendedqueue

// Double-ended queue enables inserting and polling from either end
type DoubleEndedQueue[T any] struct {
	array []T
}

// NewDoubleEndedQueue instantiates a new double-ended queue and returns a pointer to it.
func NewDoubleEndedQueue[T any]() *DoubleEndedQueue[T] {
	q := DoubleEndedQueue[T]{}
	return &q
}

// Size returns the number of items in the double-ended queue.
func (q *DoubleEndedQueue[T]) Size() int {
	return len(q.array)
}

// Add inserts a new item at the end of the queue.
func (q *DoubleEndedQueue[T]) InsertFront(i T) {
	q.array = append(q.array, i)
}

func (q *DoubleEndedQueue[T]) InsertLast(i T) {
	q
}

// Remove removes and returns the item at the front of the double-ended queue.
func (q *DoubleEndedQueue[T]) RemoveFront() (T, bool) {
	if q.Size() == 0 {
		var zeroValue T
		return zeroValue, false
	}
	result := q.array[0]
	q.array = q.array[1:]
	return result, true
}

func (q *DoubleEndedQueue[T]) RemoveLast() (T, bool) {
	if q.Size() == 0 {
		var zeroValue T
		return zeroValue, false
	}
	result := q.array[len(q.array)-1]
	q.array = q.array[:len(q.array)-1]
	return result, true
}
