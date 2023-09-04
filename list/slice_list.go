package list

import "reflect"

// SliceList implements the List interface using a Go array.
type SliceList[T any] struct {
	slice []T
}

// NewSliceList instantiates an SliceList and returns a pointer to it.
func NewSliceList[T any](values ...T) *SliceList[T] {
	l := SliceList[T]{values}
	return &l
}

// Size returns the number of items in the SliceList.
func (l *SliceList[T]) Size() int {
	return len(l.slice)
}

// Append adds the given item to the end of the SliceList.
func (l *SliceList[T]) Append(i T) {
	l.slice = append(l.slice, i)
}

// Prepend adds the given item to the start of the SliceList.
func (l *SliceList[T]) Prepend(i T) {
	l.slice = append([]T{i}, l.slice...)
}

// Get returns the item at the given index of the SliceList. It also returns a Boolean that is true
// if the index is within the bounds of the SliceList.
func (l *SliceList[T]) Get(index int) (T, bool) {
	if index >= len(l.slice) {
		var zeroValue T
		return zeroValue, false
	}
	return l.slice[index], true
}

// Equal returns true if both SliceLists are equal
func (l *SliceList[T]) Equal(m *SliceList[T]) bool {
	if l.Size() != m.Size() {
		return false
	}
	for i := 0; i < l.Size(); i++ {
		lValue, _ := l.Get(i)
		mValue, _ := m.Get(i)
		if !reflect.DeepEqual(lValue, mValue) {
			return false
		}
	}
	return true
}

// Remove removes the item at the given index of the SliceList. It returns a Boolean that is true
// if the index is within the bounds of the SliceList.
func (l *SliceList[T]) Remove(index int) bool {
	if index >= len(l.slice) {
		return false
	}
	newLength := 0
	for i, v := range l.slice {
		if i != index {
			l.slice[newLength] = v
			newLength += 1
		}
	}
	l.slice = l.slice[:newLength]
	return true
}

// Insert adds the given item at the given index of the SliceList. All following items are shifted
// right. If the given index is not within the bounds of the SliceList then no changes will be made.
// A Boolean is returned that is true if the index is within the bounds of the SliceList and the
// item was inserted.
func (l *SliceList[T]) Insert(index int, i T) bool {
	if index < 0 || index >= l.Size() {
		return false
	}
	newSize := l.Size() + 1
	temp := l.slice[index]
	l.slice[index] = i
	for j := index + 1; j < newSize-1; j++ {
		l.slice[index], temp = temp, l.slice[index]
	}
	l.slice = append(l.slice, temp)
	return true
}

// Map applies the given function to all items in the SliceList.
func (l *SliceList[T]) Map(f func(T) T) {
	for i, v := range l.slice {
		l.slice[i] = f(v)
	}
}

// Filter applies the given predicate function to all items in the SliceList and returns a new
// ArrayList that only contains items for which the predicate was true.
func (l *SliceList[T]) Filter(f func(T) bool) {
	originalIndex := 0
	newIndex := 0
	originalLength := len(l.slice)
	newLength := 0
	for originalIndex < originalLength {
		if f(l.slice[originalIndex]) {
			l.slice[newIndex] = l.slice[originalIndex]
			newIndex += 1
			newLength += 1
		}
		originalIndex += 1
	}
	l.slice = l.slice[:newLength]
}
