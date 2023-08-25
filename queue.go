package godsa

type Queue struct {
	array []int
}

func NewQueue() *Queue {
	q := Queue{}
	return &q
}

func (q *Queue) Size() int {
	return len(q.array)
}

func (q *Queue) Add(i int) {
	q.array = append(q.array, i)
}

func (q *Queue) Remove() (int, bool) {
	if q.Size() == 0 {
		return 0, false
	}
	result := q.array[0]
	q.array = q.array[1:]
	return result, true
}
