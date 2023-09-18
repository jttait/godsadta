package priorityqueue

import "cmp"

type PriorityQueue[T cmp.Ordered] interface {
	Size() int
	Insert(T)
	Extract() (T, bool)
	Peek() (T, bool)
}
