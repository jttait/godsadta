// Package set provides a set abstract data type and associated methods
package godsa

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
