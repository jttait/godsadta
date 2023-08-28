// Package priorityqueue provides the priority queue data structure and associated methods
package priorityqueue

import "cmp"

// PriorityQueue is a list of elements which each have a priority. When polling the queue, the
// highest priority item is returned.
type PriorityQueue[T cmp.Ordered] struct {
	array []T
}

// NewPriorityQueue instantiates a priority queue and returns a pointer to it.
func NewPriorityQueue[T cmp.Ordered]() *PriorityQueue[T] {
	q := PriorityQueue[T]{}
	return &q
}

// Size returns the number of items in the priority queue.
func (q *PriorityQueue[T]) Size() int {
	return len(q.array)
}

// Add adds an item the priority queue.
func (q *PriorityQueue[T]) Add(i T) {
	q.array = append(q.array, i)
}

// Remove removes one of the items from the priority queue. If there are multiple items with the
// same value then only one of them will be removed. It returns a Boolean that is true if an item
// was in the priority queue.
func (q *PriorityQueue[T]) Remove(j T) bool {
	for i, v := range q.array {
		if v == j {
			q.array[i] = q.array[len(q.array)-1]
			q.array = q.array[:len(q.array)-1]
			return true
		}
	}
	return false
}

// Poll removes the highest-priority item and returns it. It returns a Boolean that is false if the
// priority queue was empty. In cases where the priority queue was empty then the zero value of the
// type will be returned but this is meaningless.
func (q *PriorityQueue[T]) Poll() (T, bool) {
	if q.Size() == 0 {
		var zeroValue T
		return zeroValue, false
	}
	max := q.array[0]
	maxIndex := 0
	for i := 1; i < q.Size(); i++ {
		if q.array[i] > max {
			max = q.array[i]
			maxIndex = i
		}
	}
	q.array[maxIndex] = q.array[len(q.array)-1]
	q.array = q.array[:len(q.array)-1]
	return max, true
}

// Peek returns the highest-priority item in the priority queue but, unlike poll, does not remove it.
// It also returns a Boolean that is false if the priority queue was empty. In
// cases where the priority queue was empty, the zero value for the type will be returned but this is
// meaningless.
func (q *PriorityQueue[T]) Peek() (T, bool) {
	if q.Size() == 0 {
		var zeroValue T
		return zeroValue, false
	}
	max := q.array[0]
	for i := 1; i < q.Size(); i++ {
		if q.array[i] > max {
			max = q.array[i]
		}
	}
	return max, true
}
