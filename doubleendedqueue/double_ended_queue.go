package doubleendedqueue

type DoubleEndedQueue[T any] interface {
	Size()
	InsertFront(i T)
	InsertLast(i T)
	RemoveFront() (T, bool)
	RemoveLast() (T, bool)
}
