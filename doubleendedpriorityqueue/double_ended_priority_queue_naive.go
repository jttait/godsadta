package doubleendedpriorityqueue

import (
	"cmp"
	"sort"
)

type DoubleEndedPriorityQueueNaive[T cmp.Ordered] struct {
	s []T
}

func NewDoubleEndedPriorityQueueNaive[T cmp.Ordered]() *DoubleEndedPriorityQueueNaive[T] {
	q := DoubleEndedPriorityQueueNaive[T]{}
	return &q
}

func (q *DoubleEndedPriorityQueueNaive[T]) Insert(i T) {
	q.s = append(q.s, i)
}

func (q *DoubleEndedPriorityQueueNaive[T]) Size() int {
	return len(q.s)
}

func (q *DoubleEndedPriorityQueueNaive[T]) ExtractMax() (T, bool) {
	if len(q.s) == 0 {
		var zeroValue T
		return zeroValue, false
	}
	sort.Slice(q.s, func(i, j int) bool { return cmp.Less(q.s[j], q.s[i]) })
	result := q.s[0]
	q.s = q.s[1:]
	return result, true
}

func (q *DoubleEndedPriorityQueueNaive[T]) ExtractMin() (T, bool) {
	if len(q.s) == 0 {
		var zeroValue T
		return zeroValue, false
	}
	sort.Slice(q.s, func(i, j int) bool { return cmp.Less(q.s[j], q.s[i]) })
	result := q.s[len(q.s)-1]
	q.s = q.s[:len(q.s)-1]
	return result, true
}

func (q *DoubleEndedPriorityQueueNaive[T]) PeekMax() (T, bool) {
	if len(q.s) == 0 {
		var zeroValue T
		return zeroValue, false
	}
	sort.Slice(q.s, func(i, j int) bool { return cmp.Less(q.s[j], q.s[i]) })
	result := q.s[0]
	return result, true
}

func (q *DoubleEndedPriorityQueueNaive[T]) PeekMin() (T, bool) {
	if len(q.s) == 0 {
		var zeroValue T
		return zeroValue, false
	}
	sort.Slice(q.s, func(i, j int) bool { return cmp.Less(q.s[j], q.s[i]) })
	result := q.s[len(q.s)-1]
	return result, true
}
