// Package maxheap provides the max heap data structure and associated methods
package maxheap

import "cmp"

type MaxHeap[T cmp.Ordered] interface {
	Insert(i T)
	Extract() (T, bool)
	Size() int
	Peek() (T, bool)
}
