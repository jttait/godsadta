// Package queue provides an interface and implementations for the queue abstract data type
package queue

type Queue[T any] interface {
	Size() int
	Insert(i T)
	Remove() (T, bool)
}
