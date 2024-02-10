// Package depriorityqueue provides interface and implementations for double-ended priority queue
// abstract data type
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
