package godsa

type PriorityQueue struct {
	array []int
}

func NewPriorityQueue() *PriorityQueue {
	q := PriorityQueue{}
	return &q
}

func (q *PriorityQueue) Size() int {
	return len(q.array)
}

func (q *PriorityQueue) Add(i int) {
	q.array = append(q.array, i)
}

func (q *PriorityQueue) Remove(j int) bool {
	for i, v := range q.array {
		if v == j {
			q.array[i] = q.array[len(q.array)-1]
			q.array = q.array[:len(q.array)-1]
			return true
		}
	}
	return false
}

func (q *PriorityQueue) Poll() (int, bool) {
	if q.Size() == 0 {
		return 0, false
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

func (q *PriorityQueue) Peek() (int, bool) {
	if q.Size() == 0 {
		return 0, false
	}
	max := q.array[0]
	for i := 1; i < q.Size(); i++ {
		if q.array[i] > max {
			max = q.array[i]
		}
	}
	return max, true

}
