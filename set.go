package godsa

// Set is a data structure that stores unique values without ordering.
type Set[T comparable] struct {
	m map[T]struct{}
}

// NewSet instantiates a set and returns a pointer to it.
func NewSet[T comparable]() *Set[T] {
	s := Set[T]{}
	s.m = make(map[T]struct{})
	return &s
}

// Size returns the number of items in the set.
func (s *Set[T]) Size() int {
	return len(s.m)
}

// Add adds an item to the set. It returns false if the item was already in the set.
func (s *Set[T]) Add(i T) bool {
	_, ok := s.m[i]
	if ok {
		return false
	}
	s.m[i] = struct{}{}
	return true
}

// Remove removes an item from the set. It returns false if the item was not in the set.
func (s *Set[T]) Remove(i T) bool {
	_, ok := s.m[i]
	if !ok {
		return false
	}
	delete(s.m, i)
	return true
}

// Contains returns true if the item is in the set.
func (s *Set[T]) Contains(i T) bool {
	_, ok := s.m[i]
	return ok
}

// Equals compares two sets and returns true if they contain exactly the same items.
func (s *Set[T]) Equals(t *Set[T]) bool {
	for key, _ := range s.m {
		_, ok := t.m[key]
		if !ok {
			return false
		}
	}
	return true
}
