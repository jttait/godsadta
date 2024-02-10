// Package immutable list provides an interface and implementations for the immutable list abstract
// data type
package immutablelist

type ImmutableList[T any] interface {
	Size() int
	Prepend(i T) ImmutableList[T]
	Append(i T) ImmutableList[T]
	Remove(index int) (ImmutableList[T], bool)
	Insert(index int, i T) (ImmutableList[T], bool)
	Get(index int) T
	Map(func(T) T) ImmutableList[T]
	Filter(func(T) bool) ImmutableList[T]
	Equal(ImmutableList[T]) bool
}
