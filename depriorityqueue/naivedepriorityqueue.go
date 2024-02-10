package depriorityqueue

import (
	"cmp"
	"sort"
)

// NaiveDEPriorityQueue is a naive implementation of a double-ended priority queue.
type NaiveDEPriorityQueue[T cmp.Ordered] struct {
	s []T
}

// NewNaiveDEPriorityQueue instantiates a NaiveDEPriorityQueue and returns a pointer to it.
func NewNaiveDEPriorityQueue[T cmp.Ordered]() *NaiveDEPriorityQueue[T] {
	q := NaiveDEPriorityQueue[T]{}
	return &q
}

// Insert inserts an element into the queue.
func (q *NaiveDEPriorityQueue[T]) Insert(i T) {
	q.s = append(q.s, i)
}

// Size returns the number of items in the queue.
func (q *NaiveDEPriorityQueue[T]) Size() int {
	return len(q.s)
}

// ExtractMax removes the highest-priority item and returns it. A Boolean is also returned that is
// false if the queue was empty.
func (q *NaiveDEPriorityQueue[T]) ExtractMax() (T, bool) {
	if len(q.s) == 0 {
		var zeroValue T
		return zeroValue, false
	}
	sort.Slice(q.s, func(i, j int) bool { return cmp.Less(q.s[j], q.s[i]) })
	result := q.s[0]
	q.s = q.s[1:]
	return result, true
}

// ExtractMin removes the lowest-priority item and returns it. A Boolean is also returned that is
// false if the queue was empty.
func (q *NaiveDEPriorityQueue[T]) ExtractMin() (T, bool) {
	if len(q.s) == 0 {
		var zeroValue T
		return zeroValue, false
	}
	sort.Slice(q.s, func(i, j int) bool { return cmp.Less(q.s[j], q.s[i]) })
	result := q.s[len(q.s)-1]
	q.s = q.s[:len(q.s)-1]
	return result, true
}

// PeekMax returns the highest-priority item but, unlike ExtractMax, does not remove it. A Boolean
// is also returned that is false if the queue was empty.
func (q *NaiveDEPriorityQueue[T]) PeekMax() (T, bool) {
	if len(q.s) == 0 {
		var zeroValue T
		return zeroValue, false
	}
	sort.Slice(q.s, func(i, j int) bool { return cmp.Less(q.s[j], q.s[i]) })
	result := q.s[0]
	return result, true
}

// PeekMin returns the lowest-prioity item but, unlike ExtractMin, does not remove it. A Boolean is
// also returned that is false if the queue was empty.
func (q *NaiveDEPriorityQueue[T]) PeekMin() (T, bool) {
	if len(q.s) == 0 {
		var zeroValue T
		return zeroValue, false
	}
	sort.Slice(q.s, func(i, j int) bool { return cmp.Less(q.s[j], q.s[i]) })
	result := q.s[len(q.s)-1]
	return result, true
}
