// Package set provides an interface and implementations for the set abstract data type
package set

type Set[T comparable] interface {
	Size() int
	Insert(i T) bool
	Remove(i T) bool
	Contains(i T) bool
	Equals(i Set[T]) bool
	Iter() []T
	Union(i Set[T]) Set[T]
	Intersection(i Set[T]) Set[T]
	Difference(i Set[T]) Set[T]
}
