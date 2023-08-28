// Package set provides a set abstract data type and associated methods
package set

// Set is an abstract data type that stores unique values without ordering.
type Set[T comparable] struct {
	m map[T]struct{}
}

// NewSet instantiates a set and returns a pointer to it.
func NewSet[T comparable](i ...T) *Set[T] {
	s := Set[T]{}
	s.m = make(map[T]struct{})
	for _, j := range i {
		s.m[j] = struct{}{}
	}
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

func (s *Set[T]) Iter() []T {
	result := []T{}
	for k, _ := range s.m {
		result = append(result, k)
	}
	return result
}

// Union returns a set containing items that are in either set
func (s *Set[T]) Union(t *Set[T]) *Set[T] {
	for _, value := range t.Iter() {
		_ = s.Add(value)
	}
	return s
}

// Intersection returns a set containing only the items that are in both sets.
func (s *Set[T]) Intersection(t *Set[T]) *Set[T] {
	result := NewSet[T]()
	for _, value := range s.Iter() {
		if t.Contains(value) {
			result.Add(value)
		}
	}
	return result
}

// Difference returns a set containing items that are in this set but are not in the set passed in
// the argument.
func (s *Set[T]) Difference(t *Set[T]) *Set[T] {
	for _, value := range t.Iter() {
		s.Remove(value)
	}
	return s
}
