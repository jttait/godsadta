package priorityqueue

import (
	"cmp"

	"github.com/jttait/godsa/maxheap"
)

// PriorityQueue is a list of elements which each have a priority. When polling the queue, the
// highest priority item is returned.
type MaxHeapPriorityQueue[T cmp.Ordered] struct {
	maxheap maxheap.MaxHeap[T]
}

// NewPriorityQueue instantiates a priority queue and returns a pointer to it.
func NewMaxHeapPriorityQueue[T cmp.Ordered]() *MaxHeapPriorityQueue[T] {
	q := MaxHeapPriorityQueue[T]{}
	q.maxheap = maxheap.NewMaxHeapArray[T]()
	return &q
}

// Size returns the number of items in the priority queue.
func (q *MaxHeapPriorityQueue[T]) Size() int {
	return q.maxheap.Size()
}

// Add adds an item the priority queue.
func (q *MaxHeapPriorityQueue[T]) Insert(i T) {
	q.maxheap.Insert(i)
}

// Poll removes the highest-priority item and returns it. It returns a Boolean that is false if the
// priority queue was empty. In cases where the priority queue was empty then the zero value of the
// type will be returned but this is meaningless.
func (q *MaxHeapPriorityQueue[T]) Extract() (T, bool) {
	return q.maxheap.Extract()
}

// Peek returns the highest-priority item in the priority queue but, unlike poll, does not remove it.
// It also returns a Boolean that is false if the priority queue was empty. In
// cases where the priority queue was empty, the zero value for the type will be returned but this is
// meaningless.
func (q *MaxHeapPriorityQueue[T]) Peek() (T, bool) {
	return q.maxheap.Peek()
}
