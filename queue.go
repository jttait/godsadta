package godsa

type Queue[T any] struct {
	array []T
}

func NewQueue[T any]() *Queue[T] {
	q := Queue[T]{}
	return &q
}

func (q *Queue[T]) Size() int {
	return len(q.array)
}

func (q *Queue[T]) Add(i T) {
	q.array = append(q.array, i)
}

func (q *Queue[T]) Remove() (T, bool) {
	if q.Size() == 0 {
		var zeroValue T
		return zeroValue, false
	}
	result := q.array[0]
	q.array = q.array[1:]
	return result, true
}
