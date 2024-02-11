// Package lrucache provides an interface and implementations for LRU cache abstract data type
package lrucache

type LRUCache[S comparable, T comparable] interface {
	Get(S) T
	Put(S, T)
}
