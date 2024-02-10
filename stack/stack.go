// Package stack provides an interface and implementations for the stack abstract data type
package stack

type Stack[T any] interface {
	Size() int
	Push(i T)
	Pop() (T, bool)
	Peek() (T, bool)
}
