package list

import "github.com/jttait/godsa/doublylinkedlist"

type LinkedListList[T any] struct {
	list *doublylinkedlist.DoublyLinkedList[T]
}

func NewLinkedListList[T any]() *LinkedListList[T] {
	l := LinkedListList[T]{}
	l.list = doublylinkedlist.NewDoublyLinkedList[T]()
	return &l
}

func (l *LinkedListList[T]) Size() int {
	return l.list.Size()
}

func (l *LinkedListList[T]) Append(i T) {
	l.list.InsertLast(i)
}

func (l *LinkedListList[T]) Get(index int) T {

}
