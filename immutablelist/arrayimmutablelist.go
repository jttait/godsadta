package immutablelist

import "reflect"

// ImmutableArrayList implements the ImmutableList interface using a Go array.
type ArrayImmutableList[T any] struct {
	array []T
}

// NewImmutableArrayList instantiates an ImmutableArrayList and returns a pointer to it.
func NewArrayImmutableList[T any](values ...T) *ArrayImmutableList[T] {
	l := ArrayImmutableList[T]{values}
	return &l
}

// Size returns the number of items in the ImmutableArrayList.
func (l *ArrayImmutableList[T]) Size() int {
	return len(l.array)
}

// Append adds the given item to the end of the ImmutableArrayList.
func (l *ArrayImmutableList[T]) Append(val T) *ArrayImmutableList[T] {
	slice := []T{}
	for _, v := range l.array {
		slice = append(slice, v)
	}
	slice = append(slice, val)
	r := NewArrayImmutableList[T](slice...)
	return r
}

// Prepend adds the given item to the start of the ImmutableArrayList.
func (l *ArrayImmutableList[T]) Prepend(val T) *ArrayImmutableList[T] {
	slice := []T{}
	slice = append(slice, val)
	for _, v := range l.array {
		slice = append(slice, v)
	}
	r := NewArrayImmutableList[T](slice...)
	return r
}

// Get returns the item at the given index of the ImmutableArrayList. It also returns a Boolean that
// is true if the index is within the bounds of the ImmutableArrayList.
func (l *ArrayImmutableList[T]) Get(index int) (T, bool) {
	if index >= len(l.array) {
		var zeroValue T
		return zeroValue, false
	}
	return l.array[index], true
}

// Remove removes the item at the given index of the ImmutableArrayList. It returns a Boolean that
// is true if the index is within the bounds of the ImmutableArrayList.
func (l *ArrayImmutableList[T]) Remove(index int) (*ArrayImmutableList[T], bool) {
	if index >= len(l.array) {
		return l, false
	}
	slice := []T{}
	for i, v := range l.array {
		if i != index {
			slice = append(slice, v)
		}
	}
	r := NewArrayImmutableList[T](slice...)
	return r, true
}

// Insert adds the given item at the given index of the ImmutableArrayList. All following items are
// shifted right. If the given index is not within the bounds of the ImmutableArrayList then no
// changes will be made. A Boolean is returned that is true if the index is within the bounds of the
// ImmutableArrayList and the item was inserted.
func (l *ArrayImmutableList[T]) Insert(index int, value T) (*ArrayImmutableList[T], bool) {
	if index < 0 || index >= l.Size() {
		return l, false
	}
	slice := []T{}
	for i, v := range l.array {
		if i == index {
			slice = append(slice, value)
		}
		slice = append(slice, v)
	}
	r := NewArrayImmutableList[T](slice...)
	return r, true
}

// Map applies the given function to all items in the ImmutableArrayList.
func (l *ArrayImmutableList[T]) Map(f func(T) T) *ArrayImmutableList[T] {
	slice := []T{}
	for _, v := range l.array {
		slice = append(slice, f(v))
	}
	r := NewArrayImmutableList[T](slice...)
	return r
}

// Filter applies the given predicate function to all items in the ImmutableArrayList and returns a
// new ImmutableArrayList that only contains items for which the predicate was true.
func (l *ArrayImmutableList[T]) Filter(f func(T) bool) *ArrayImmutableList[T] {
	slice := []T{}
	for _, v := range l.array {
		if f(v) {
			slice = append(slice, v)
		}
	}
	r := NewArrayImmutableList[T](slice...)
	return r
}

func (l *ArrayImmutableList[T]) Equal(b *ArrayImmutableList[T]) bool {
	if l.Size() != b.Size() {
		return false
	}
	for i := range l.array {
		if !reflect.DeepEqual(l.array[i], b.array[i]) {
			return false
		}
	}
	return true
}
