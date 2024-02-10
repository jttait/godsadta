// Package multiset provides an interface and implementations for the multiset abstract data type
package multiset

type Multiset[T any] interface {
	Count(T) int
	Add(T)
	Size() int
	Remove(T) bool
}
