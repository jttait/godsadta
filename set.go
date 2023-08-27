package godsa

type Set[T comparable] struct {
	m map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	s := Set[T]{}
	s.m = make(map[T]struct{})
	return &s
}

func (s *Set[T]) Size() int {
	return len(s.m)
}

func (s *Set[T]) Add(i T) bool {
	_, ok := s.m[i]
	if ok {
		return false
	}
	s.m[i] = struct{}{}
	return true
}

func (s *Set[T]) Remove(i T) bool {
	_, ok := s.m[i]
	if !ok {
		return false
	}
	delete(s.m, i)
	return true
}

func (s *Set[T]) Contains(i T) bool {
	_, ok := s.m[i]
	return ok
}

func (s *Set[T]) Equals(t *Set[T]) bool {
	for key, _ := range s.m {
		_, ok := t.m[key]
		if !ok {
			return false
		}
	}
	return true
}
