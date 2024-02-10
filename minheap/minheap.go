package minheap

import "cmp"

type MinHeap[T cmp.Ordered] interface {
	Insert(i T)
	Extract() T
	Size() int
	Peek() T
}
