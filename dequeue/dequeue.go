// Package dequeue provides an interface and implementations for the double-ended queue abstract
// data type
package dequeue

type DEQueue[T any] interface {
	Size() int
	InsertFront(i T)
	InsertLast(i T)
	RemoveFront() (T, bool)
	RemoveLast() (T, bool)
}
