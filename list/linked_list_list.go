package list

import "github.com/jttait/godsa/linkedlist"

type LinkedListList[T any] struct {
	list linkedlist.LinkedList[T]
}

func NewLinkedListList[T any](values ...T) *LinkedListList[T] {
	l := LinkedListList[T]{}
	l.list = linkedlist.NewDoublyLinkedList[T]()
	for _, v := range values {
		l.Append(v)
	}
	return &l
}

// Size returns the number of items in the list.
func (l *LinkedListList[T]) Size() int {
	return l.list.Size()
}

// Append adds the given item at the end of the list.
func (l *LinkedListList[T]) Append(i T) {
	l.list.InsertLast(i)
}

// Prepend adds the given item at the start of the list.
func (l *LinkedListList[T]) Prepend(i T) {
	l.list.InsertFront(i)
}

// Remove removes the item at the given index of the list. It also returns a Boolean that is true
// if the index is within the bounds of the list.
func (l *LinkedListList[T]) Remove(index int) bool {
	return l.list.Remove(index)
}

func (l *LinkedListList[T]) Insert(index int, i T) bool {
	return l.list.Insert(index, i)
}

// Get returns the item at the given index of the list. It also returns a Boolean that is true if
// the index is within the bounds of the list.
func (l *LinkedListList[T]) Get(index int) (T, bool) {
	return l.list.Get(index)
}

// Map applies the given function to each item in the list.
func (l *LinkedListList[T]) Map(f func(T) T) {
	l.list = l.list.Map(f)
}

// Filter applies the given predicate function to each item in the list and returns a list
// containing all items for which the predicate was true.
func (l *LinkedListList[T]) Filter(f func(T) bool) {
	l.list = l.list.Filter(f)
}
