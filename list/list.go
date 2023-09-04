package list

type List[T any] interface {
	Size() int
	Prepend(i T)
	Append(i T)
	Remove(index int) bool
	Insert(index int, i T) bool
	Get(index int) T
	Equal(List[T]) bool
	Map(func(T) T)
	Filter(func(T) bool)
}
