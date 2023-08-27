package godsa

// Queue is a first-in, first-out data structure.
type Queue[T any] struct {
	array []T
}

// NewQueue instantiates a new queue and returns a pointer to it.
func NewQueue[T any]() *Queue[T] {
	q := Queue[T]{}
	return &q
}

// Size returns the number of items in the queue.
func (q *Queue[T]) Size() int {
	return len(q.array)
}

// Add inserts a new item at the end of the queue.
func (q *Queue[T]) Add(i T) {
	q.array = append(q.array, i)
}

// Remove removes and returns the item at the head of the queue.
func (q *Queue[T]) Remove() (T, bool) {
	if q.Size() == 0 {
		var zeroValue T
		return zeroValue, false
	}
	result := q.array[0]
	q.array = q.array[1:]
	return result, true
}
