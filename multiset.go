package godsa

type Multiset[T any] interface {
	Count(T) int
	Add(T)
	Size() int
	Remove(T) bool
}
