package godsa

import "cmp"

type PriorityQueue[T cmp.Ordered] struct {
	array []T
}

func NewPriorityQueue[T cmp.Ordered]() *PriorityQueue[T] {
	q := PriorityQueue[T]{}
	return &q
}

func (q *PriorityQueue[T]) Size() int {
	return len(q.array)
}

func (q *PriorityQueue[T]) Add(i T) {
	q.array = append(q.array, i)
}

func (q *PriorityQueue[T]) Remove(j T) bool {
	for i, v := range q.array {
		if v == j {
			q.array[i] = q.array[len(q.array)-1]
			q.array = q.array[:len(q.array)-1]
			return true
		}
	}
	return false
}

func (q *PriorityQueue[T]) Poll() (T, bool) {
	if q.Size() == 0 {
		var zeroValue T
		return zeroValue, false
	}
	max := q.array[0]
	maxIndex := 0
	for i := 1; i < q.Size(); i++ {
		if q.array[i] > max {
			max = q.array[i]
			maxIndex = i
		}
	}
	q.array[maxIndex] = q.array[len(q.array)-1]
	q.array = q.array[:len(q.array)-1]
	return max, true
}

func (q *PriorityQueue[T]) Peek() (T, bool) {
	if q.Size() == 0 {
		var zeroValue T
		return zeroValue, false
	}
	max := q.array[0]
	for i := 1; i < q.Size(); i++ {
		if q.array[i] > max {
			max = q.array[i]
		}
	}
	return max, true

}
