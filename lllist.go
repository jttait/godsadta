package godsa

import (
	"github.com/jttait/godsa/dll"
	"github.com/jttait/godsa/linkedlist"
)

type LLList[T any] struct {
	list linkedlist.LinkedList[T]
}

func NewLLList[T any](values ...T) *LLList[T] {
	l := LLList[T]{}
	l.list = dll.NewDLL[T]()
	for _, v := range values {
		l.Append(v)
	}
	return &l
}

// Size returns the number of items in the list.
func (l *LLList[T]) Size() int {
	return l.list.Size()
}

// Append adds the given item at the end of the list.
func (l *LLList[T]) Append(i T) {
	l.list.InsertLast(i)
}

// Prepend adds the given item at the start of the list.
func (l *LLList[T]) Prepend(i T) {
	l.list.InsertFront(i)
}

// Remove removes the item at the given index of the list. It also returns a Boolean that is true
// if the index is within the bounds of the list.
func (l *LLList[T]) Remove(index int) bool {
	return l.list.Remove(index)
}

func (l *LLList[T]) Insert(index int, i T) bool {
	return l.list.Insert(index, i)
}

// Get returns the item at the given index of the list. It also returns a Boolean that is true if
// the index is within the bounds of the list.
func (l *LLList[T]) Get(index int) (T, bool) {
	return l.list.Get(index)
}

// Map applies the given function to each item in the list.
func (l *LLList[T]) Map(f func(T) T) {
	l.list = l.list.Map(f)
}

// Filter applies the given predicate function to each item in the list and returns a list
// containing all items for which the predicate was true.
func (l *LLList[T]) Filter(f func(T) bool) {
	l.list = l.list.Filter(f)
}
