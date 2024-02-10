// Package maxheap provides an interface and implementations for the max heap data structure
package maxheap

import "cmp"

type MaxHeap[T cmp.Ordered] interface {
	Insert(i T)
	Extract() (T, bool)
	Size() int
	Peek() (T, bool)
}
