package linkedlist

type LinkedList[T any] interface {
	Size() int
	InsertFront(i T)
	InsertLast(i T)
	RemoveFront() (T, bool)
	RemoveLast() (T, bool)
	PeekFront() (T, bool)
	PeekLast() (T, bool)
	Get(int) (T, bool)
}
