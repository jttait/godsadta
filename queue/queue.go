package queue

type Queue[T any] interface {
	Size() int
	Insert(i T)
	Remove() (T, bool)
}
