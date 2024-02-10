package depriorityqueue

import "cmp"

type DEPriorityQueue[T cmp.Ordered] interface {
	Size() int
	Insert(T)
	ExtractMax() (T, bool)
	ExtractMin() (T, bool)
	PeekMin() (T, bool)
	PeekMax() (T, bool)
}
