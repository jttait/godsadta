package godsa

type Stack[T any] struct {
	array []T
}

func NewStack[T any]() *Stack[T] {
	s := Stack[T]{}
	return &s
}

func (s *Stack[T]) Size() int {
	return len(s.array)
}

func (s *Stack[T]) Push(i T) {
	s.array = append(s.array, i)
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.Size() == 0 {
		var zeroValue T
		return zeroValue, false
	}
	result := s.array[len(s.array)-1]
	s.array = s.array[:len(s.array)-1]
	return result, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.Size() == 0 {
		var zeroValue T
		return zeroValue, false
	}
	result := s.array[len(s.array)-1]
	return result, true
}
