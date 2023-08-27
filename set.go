package godsa

type Set struct {
	m map[int]struct{}
}

func NewSet() *Set {
	s := Set{}
	s.m = make(map[int]struct{})
	return &s
}

func (s *Set) Size() int {
	return len(s.m)
}

func (s *Set) Add(i int) bool {
	_, ok := s.m[i]
	if ok {
		return false
	}
	s.m[i] = struct{}{}
	return true
}

func (s *Set) Remove(i int) bool {
	_, ok := s.m[i]
	if !ok {
		return false
	}
	delete(s.m, i)
	return true
}

func (s *Set) Contains(i int) bool {
	_, ok := s.m[i]
	return ok
}

func (s *Set) Equals(t *Set) bool {
	for key, _ := range s.m {
		_, ok := t.m[key]
		if !ok {
			return false
		}
	}
	return true
}
