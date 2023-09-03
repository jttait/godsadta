package list

// ArrayList implements the List interface using a Go array.
type ArrayList[T any] struct {
	array []T
}

// NewArrayList instantiates an ArrayList and returns a pointer to it.
func NewArrayList[T any]() *ArrayList[T] {
	l := ArrayList[T]{}
	return &l
}

// Size returns the number of items in the ArrayList.
func (l *ArrayList[T]) Size() int {
	return len(l.array)
}

// Append adds the given item to the end of the ArrayList.
func (l *ArrayList[T]) Append(i T) {
	l.array = append(l.array, i)
}

// Prepend adds the given item to the start of the ArrayList.
func (l *ArrayList[T]) Prepend(i T) {
	l.array = append([]T{i}, l.array...)
}

// Get returns the item at the given index of the ArrayList. It also returns a Boolean that is true
// if the index is within the bounds of the ArrayList.
func (l *ArrayList[T]) Get(index int) (T, bool) {
	if index >= len(l.array) {
		var zeroValue T
		return zeroValue, false
	}
	return l.array[index], true
}

// Remove removes the item at the given index of the ArrayList. It returns a Boolean that is true
// if the index is within the bounds of the ArrayList.
func (l *ArrayList[T]) Remove(index int) bool {
	if index >= len(l.array) {
		return false
	}
	newLength := 0
	for i, v := range l.array {
		if i != index {
			l.array[newLength] = v
			newLength += 1
		}
	}
	l.array = l.array[:newLength]
	return true
}

// Insert adds the given item at the given index of the ArrayList. All following items are shifted
// right. If the given index is not within the bounds of the ArrayList then no changes will be made.
// A Boolean is returned that is true if the index is within the bounds of the ArrayList and the
// item was inserted.
func (l *ArrayList[T]) Insert(index int, i T) bool {
	if index < 0 || index >= l.Size() {
		return false
	}
	newSize := l.Size() + 1
	temp := l.array[index]
	l.array[index] = i
	for j := index + 1; j < newSize-1; j++ {
		l.array[index], temp = temp, l.array[index]
	}
	l.array = append(l.array, temp)
	return true
}

// Map applies the given function to all items in the ArrayList.
func (l *ArrayList[T]) Map(f func(T) T) {
	for i, v := range l.array {
		l.array[i] = f(v)
	}
}

// Filter applies the given predicate function to all items in the ArrayList and returns a new
// ArrayList that only contains items for which the predicate was true.
func (l *ArrayList[T]) Filter(f func(T) bool) {
	originalIndex := 0
	newIndex := 0
	originalLength := len(l.array)
	newLength := 0
	for originalIndex < originalLength {
		if f(l.array[originalIndex]) {
			l.array[newIndex] = l.array[originalIndex]
			newIndex += 1
			newLength += 1
		}
		originalIndex += 1
	}
	l.array = l.array[:newLength]
}
