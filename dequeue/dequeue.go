package dequeue

type DEQueue[T any] interface {
	Size()
	InsertFront(i T)
	InsertLast(i T)
	RemoveFront() (T, bool)
	RemoveLast() (T, bool)
}
