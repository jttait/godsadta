package godsa

type Stack struct {
	array []int
}

func NewStack() *Stack {
	s := Stack{}
	return &s
}

func (s *Stack) Size() int {
	return len(s.array)
}

func (s *Stack) Push(i int) {
	s.array = append(s.array, i)
}

func (s *Stack) Pop() (int, bool) {
	if s.Size() == 0 {
		return 0, false
	}
	result := s.array[len(s.array)-1]
	s.array = s.array[:len(s.array)-1]
	return result, true
}

func (s *Stack) Peek() (int, bool) {
	if s.Size() == 0 {
		return 0, false
	}
	result := s.array[len(s.array)-1]
	return result, true
}
