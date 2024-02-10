// Package minheap provides an interface and implementations for the min heap data structure
package minheap

import "cmp"

type MinHeap[T cmp.Ordered] interface {
	Insert(i T)
	Extract() T
	Size() int
	Peek() T
}
