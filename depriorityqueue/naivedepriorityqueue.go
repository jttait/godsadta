package depriorityqueue

import (
	"cmp"
	"sort"
)

type NaiveDEPriorityQueue[T cmp.Ordered] struct {
	s []T
}

func NewNaiveDEPriorityQueue[T cmp.Ordered]() *NaiveDEPriorityQueue[T] {
	q := NaiveDEPriorityQueue[T]{}
	return &q
}

func (q *NaiveDEPriorityQueue[T]) Insert(i T) {
	q.s = append(q.s, i)
}

func (q *NaiveDEPriorityQueue[T]) Size() int {
	return len(q.s)
}

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

func (q *NaiveDEPriorityQueue[T]) PeekMax() (T, bool) {
	if len(q.s) == 0 {
		var zeroValue T
		return zeroValue, false
	}
	sort.Slice(q.s, func(i, j int) bool { return cmp.Less(q.s[j], q.s[i]) })
	result := q.s[0]
	return result, true
}

func (q *NaiveDEPriorityQueue[T]) PeekMin() (T, bool) {
	if len(q.s) == 0 {
		var zeroValue T
		return zeroValue, false
	}
	sort.Slice(q.s, func(i, j int) bool { return cmp.Less(q.s[j], q.s[i]) })
	result := q.s[len(q.s)-1]
	return result, true
}
