// Package stack provides the stack abstract data type and associated methods
package stack

import "github.com/jttait/godsa/singlylinkedlist"

// Stack is a last-in, first-out data structure.
type Stack[T any] struct {
	list *singlylinkedlist.SinglyLinkedList[T]
}

// NewStack instantiates a new stack and returns a pointer to it.
func NewStack[T any]() *Stack[T] {
	s := Stack[T]{}
	s.list = singlylinkedlist.NewSinglyLinkedList[T]()
	return &s
}

// Size returns the number of items in the stack.
func (s *Stack[T]) Size() int {
	return s.list.Size()
}

// Push adds a new item to the top of the stack. Time complexity is O(1)
func (s *Stack[T]) Push(i T) {
	s.list.Insert(i)
}

// Pop removes the item at the top of the stack and returns it. It also returns a Boolean which is
// false if the stack is empty (and so it was not possible to pop an item). In
// the case where the stack is empty, the zero value of the type in the stack will be returned but
// this is meaningless. Time complexity is O(1).
func (s *Stack[T]) Pop() (T, bool) {
	return s.list.Remove()
}

// Peek returns the item at the top of the stack but, unlike pop, does not remove it. It also
// returns a Boolean which is false if the stack is empty (and so it was not possible to peek an
// item). In the case where the stack is empty, the zero value of the type in the
// stack will be returned but this is meaningless. Time complexity is O(1).
func (s *Stack[T]) Peek() (T, bool) {
	return s.list.Peek()
}
