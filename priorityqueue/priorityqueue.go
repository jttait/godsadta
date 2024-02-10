// Package priorityqueue provides an interface and implementations for the priority queue abstract
// data type
package priorityqueue

import "cmp"

type PriorityQueue[T cmp.Ordered] interface {
	Size() int
	Insert(T)
	Extract() (T, bool)
	Peek() (T, bool)
}
