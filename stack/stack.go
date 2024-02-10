package stack

type Stack[T any] interface {
	Size() int
	Push(i T)
	Pop() (T, bool)
	Peek() (T, bool)
}
