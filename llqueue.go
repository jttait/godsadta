package godsa

type LLQueue[T any] struct {
	list *DLL[T]
}

// NewLLQueue instantiates a new linked list queue and returns a pointer to it.
func NewLLQueue[T any]() *LLQueue[T] {
	q := LLQueue[T]{}
	q.list = NewDLL[T]()
	return &q
}

// Size returns the number of items in the queue.
func (q *LLQueue[T]) Size() int {
	return q.list.Size()
}

// Add inserts a new item at the end of the queue.
func (q *LLQueue[T]) Insert(i T) {
	q.list.InsertLast(i)
}

// Remove removes and returns the item at the head of the queue.
func (q *LLQueue[T]) Remove() (T, bool) {
	return q.list.RemoveFront()
}
