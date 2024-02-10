package mapset

import "github.com/jttait/godsa/set"

// Set is an abstract data type that stores unique values without ordering.
type MapSet[T comparable] struct {
	m map[T]struct{}
}

// NewSet instantiates a set and returns a pointer to it.
func NewMapSet[T comparable](i ...T) *MapSet[T] {
	s := MapSet[T]{}
	s.m = make(map[T]struct{})
	for _, j := range i {
		s.m[j] = struct{}{}
	}
	return &s
}

// Size returns the number of items in the set.
func (s *MapSet[T]) Size() int {
	return len(s.m)
}

// Add adds an item to the set. It returns false if the item was already in the set.
func (s *MapSet[T]) Insert(i T) bool {
	_, ok := s.m[i]
	if ok {
		return false
	}
	s.m[i] = struct{}{}
	return true
}

// Remove removes an item from the set. It returns false if the item was not in the set.
func (s *MapSet[T]) Remove(i T) bool {
	_, ok := s.m[i]
	if !ok {
		return false
	}
	delete(s.m, i)
	return true
}

// Contains returns true if the item is in the set.
func (s *MapSet[T]) Contains(i T) bool {
	_, ok := s.m[i]
	return ok
}

// Equals compares two sets and returns true if they contain exactly the same items.
func (s *MapSet[T]) Equals(t set.Set[T]) bool {
	u := t.(*MapSet[T])
	for key, _ := range s.m {
		_, ok := u.m[key]
		if !ok {
			return false
		}
	}
	return true
}

func (s *MapSet[T]) Iter() []T {
	result := []T{}
	for k, _ := range s.m {
		result = append(result, k)
	}
	return result
}

// Union returns a set containing items that are in either set
func (s *MapSet[T]) Union(t set.Set[T]) set.Set[T] {
	result := NewMapSet[T]()
	for _, value := range s.Iter() {
		_ = result.Insert(value)
	}
	for _, value := range t.Iter() {
		_ = result.Insert(value)
	}
	return result
}

// Intersection returns a set containing only the items that are in both sets.
func (s *MapSet[T]) Intersection(t set.Set[T]) set.Set[T] {
	result := NewMapSet[T]()
	for _, value := range s.Iter() {
		if t.Contains(value) {
			result.Insert(value)
		}
	}
	return result
}

// Difference returns a set containing items that are in this set but are not in the set passed in
// the argument.
func (s *MapSet[T]) Difference(t set.Set[T]) set.Set[T] {
	result := NewMapSet[T]()
	for _, value := range s.Iter() {
		_ = result.Insert(value)
	}
	for _, value := range t.Iter() {
		result.Remove(value)
	}
	return result
}
