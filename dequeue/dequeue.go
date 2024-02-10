// Package dequeue defines an interface for double-ended queues
package dequeue

type DEQueue[T any] interface {
	Size() int
	InsertFront(i T)
	InsertLast(i T)
	RemoveFront() (T, bool)
	RemoveLast() (T, bool)
}
