package stack

import (
	"github.com/jttait/godsadta/linkedlist"
)

// Stack is a last-in, first-out data structure.
type LLStack[T any] struct {
	list linkedlist.LinkedList[T]
}

// NewStack instantiates a new stack and returns a pointer to it.
func NewLLStack[T any]() *LLStack[T] {
	s := LLStack[T]{}
	s.list = linkedlist.NewSLL[T]()
	return &s
}

// Size returns the number of items in the stack.
func (s *LLStack[T]) Size() int {
	return s.list.Size()
}

// Push adds a new item to the top of the stack.
func (s *LLStack[T]) Push(i T) {
	s.list.InsertFront(i)
}

// Pop removes the item at the top of the stack and returns it. It also returns a Boolean which is
// false if the stack is empty (and so it was not possible to pop an item). In
// the case where the stack is empty, the zero value of the type in the stack will be returned but
// this is meaningless.
func (s *LLStack[T]) Pop() (T, bool) {
	return s.list.RemoveFront()
}

// Peek returns the item at the top of the stack but, unlike pop, does not remove it. It also
// returns a Boolean which is false if the stack is empty (and so it was not possible to peek an
// item). In the case where the stack is empty, the zero value of the type in the
// stack will be returned but this is meaningless.
func (s *LLStack[T]) Peek() (T, bool) {
	return s.list.PeekFront()
}
