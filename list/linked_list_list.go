package list

import "github.com/jttait/godsa/linkedlist"

type LinkedListList[T any] struct {
	list *linkedlist.DoublyLinkedList[T]
}

func NewLinkedListList[T any]() *LinkedListList[T] {
	l := LinkedListList[T]{}
	l.list = linkedlist.NewDoublyLinkedList[T]()
	return &l
}

func (l *LinkedListList[T]) Size() int {
	return l.list.Size()
}

func (l *LinkedListList[T]) Append(i T) {
	l.list.InsertLast(i)
}

func (l *LinkedListList[T]) Get(index int) (T, bool) {
	return l.list.Get(index)
}

func (l *LinkedListList[T]) Map(f func(T) T) {
	l.list.Map(f)
}

func (l *LinkedListList[T]) Filter(f func(T) bool) *LinkedListList[T] {
	l.list = l.list.Filter(f)
	return l
}
