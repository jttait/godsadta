package lrucache

import (
	"github.com/jttait/godsa/linkedlist"
)

type CacheElement[S comparable, T comparable] struct {
	key   S
	value T
}

type LLMapLRUCache[S comparable, T comparable] struct {
	ll   *linkedlist.DLL[CacheElement[S, T]]
	m    map[S]linkedlist.DLLNode[CacheElement[S, T]]
	size int
}

func NewLLMapLRUCache[S comparable, T comparable](size int) *LLMapLRUCache[S, T] {
	c := LLMapLRUCache[S, T]{
		linkedlist.NewDLL[CacheElement[S, T]](),
		map[S]linkedlist.DLLNode[CacheElement[S, T]]{},
		size,
	}
	return &c
}

func (c *LLMapLRUCache[S, T]) Put(key S, value T) {
	if c.size == 0 {
		return
	}
	ptr, ok := c.m[key]
	if ok {
		ptr.Prev.Next = ptr.Next
	} else if len(c.m) >= c.size {
		lastVal, _ := c.ll.PeekLast()
		delete(c.m, lastVal.key)
		c.ll.RemoveLast()
	}
	newNode := linkedlist.NewDLLNode[CacheElement[S, T]](CacheElement[S, T]{key, value})
	c.ll.InsertNodeFront(newNode)
	c.m[key] = *newNode
}

func (c *LLMapLRUCache[S, T]) Get(key S) (T, bool) {
	ptr, ok := c.m[key]
	if !ok {
		var zeroValue T
		return zeroValue, false
	}
	return ptr.Val.value, true
}
