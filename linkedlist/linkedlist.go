// Package linkedlist provides an interface and implementations for the linked list data structure
package linkedlist

type LinkedList[T any] interface {
	Size() int
	InsertFront(i T)
	InsertLast(i T)
	Insert(int, T) bool
	RemoveFront() (T, bool)
	RemoveLast() (T, bool)
	PeekFront() (T, bool)
	PeekLast() (T, bool)
	Get(int) (T, bool)
	Remove(int) bool
	Equal(LinkedList[T]) bool
	Map(func(T) T) LinkedList[T]
	Filter(func(T) bool) LinkedList[T]
}
