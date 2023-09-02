package list

type ArrayList[T any] struct {
	array []T
}

func NewArrayList[T any]() *ArrayList[T] {
	l := ArrayList[T]{}
	return &l
}

func (l *ArrayList[T]) Size() int {
	return len(l.array)
}

func (l *ArrayList[T]) Append(i T) {
	l.array = append(l.array, i)
}

func (l *ArrayList[T]) Prepend(i T) {
	l.array = append([]T{i}, l.array...)
}

func (l *ArrayList[T]) Get(index int) (T, bool) {
	if index >= len(l.array) {
		var zeroValue T
		return zeroValue, false
	}
	return l.array[index], true
}

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

func (l *ArrayList[T]) Map(f func(T) T) {
	for i := 0; i < len(l.array); i++ {
		l.array[i] = f(l.array[i])
	}
}

func (l *ArrayList[T]) Filter(f func(T) bool) *ArrayList[T] {
	result := NewArrayList[T]()
	for i := 0; i < len(l.array); i++ {
		if f(l.array[i]) {
			result.Append(l.array[i])
		}
	}
	return result
}
