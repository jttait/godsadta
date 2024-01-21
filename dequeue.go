package godsa

type DEQueue[T any] interface {
	Size() int
	InsertFront(i T)
	InsertLast(i T)
	RemoveFront() (T, bool)
	RemoveLast() (T, bool)
}
