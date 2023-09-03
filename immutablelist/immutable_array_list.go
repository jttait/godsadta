package immutablelist

// ImmutableArrayList implements the ImmutableList interface using a Go array.
type ImmutableArrayList[T any] struct {
	array []T
}

// NewImmutableArrayList instantiates an ImmutableArrayList and returns a pointer to it.
func NewImmutableArrayList[T any](values ...T) *ImmutableArrayList[T] {
	l := ImmutableArrayList[T]{values}
	return &l
}

// Size returns the number of items in the ImmutableArrayList.
func (l *ImmutableArrayList[T]) Size() int {
	return len(l.array)
}

// Append adds the given item to the end of the ImmutableArrayList.
func (l *ImmutableArrayList[T]) Append(i T) *ImmutableArrayList[T] {
	r := NewImmutableArrayList[T]()
	for _, v := range l.array {
		r.array = append(r.array, v)
	}
	r.array = append(r.array, i)
	return r
}

// Prepend adds the given item to the start of the ImmutableArrayList.
func (l *ImmutableArrayList[T]) Prepend(i T) *ImmutableArrayList[T] {
	r := NewImmutableArrayList[T]()
	r.array = append(r.array, i)
	for _, v := range l.array {
		r.array = append(r.array, v)
	}
	return r
}

// Get returns the item at the given index of the ImmutableArrayList. It also returns a Boolean that
// is true if the index is within the bounds of the ImmutableArrayList.
func (l *ImmutableArrayList[T]) Get(index int) (T, bool) {
	if index >= len(l.array) {
		var zeroValue T
		return zeroValue, false
	}
	return l.array[index], true
}

// Remove removes the item at the given index of the ImmutableArrayList. It returns a Boolean that
// is true if the index is within the bounds of the ImmutableArrayList.
func (l *ImmutableArrayList[T]) Remove(index int) (*ImmutableArrayList[T], bool) {
	if index >= len(l.array) {
		return l, false
	}
	r := NewImmutableArrayList[T]()
	for i, v := range l.array {
		if i != index {
			r.array = append(r.array, v)
		}
	}
	return r, true
}

// Insert adds the given item at the given index of the ImmutableArrayList. All following items are
// shifted right. If the given index is not within the bounds of the ImmutableArrayList then no
// changes will be made. A Boolean is returned that is true if the index is within the bounds of the
// ImmutableArrayList and the item was inserted.
func (l *ImmutableArrayList[T]) Insert(index int, value T) (*ImmutableArrayList[T], bool) {
	if index < 0 || index >= l.Size() {
		return l, false
	}
	r := NewImmutableArrayList[T]()
	for i, v := range l.array {
		if i == index {
			r.array = append(r.array, value)
		}
		r.array = append(r.array, v)
	}
	return r, true
}

// Map applies the given function to all items in the ImmutableArrayList.
func (l *ImmutableArrayList[T]) Map(f func(T) T) *ImmutableArrayList[T] {
	r := NewImmutableArrayList[T]()
	for _, v := range l.array {
		r.array = append(r.array, f(v))
	}
	return r
}

// Filter applies the given predicate function to all items in the ImmutableArrayList and returns a
// new ImmutableArrayList that only contains items for which the predicate was true.
func (l *ImmutableArrayList[T]) Filter(f func(T) bool) *ImmutableArrayList[T] {
	r := NewImmutableArrayList[T]()
	for _, v := range l.array {
		if f(v) {
			r.array = append(r.array, v)
		}
	}
	return r
}
