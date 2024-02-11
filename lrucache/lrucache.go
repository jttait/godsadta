package lrucache

type LRUCache[S comparable, T comparable] interface {
	Get(S) T
	Put(S, T)
}
