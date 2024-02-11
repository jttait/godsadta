package lrucache

import (
	"testing"

	"github.com/jttait/godsa/assert"
)

func TestShouldPutAndGet(t *testing.T) {
	c := NewLLMapLRUCache[int, int](5)
	c.Put(1, 2)
	got, _ := c.Get(1)
	want := 2
	assert.AssertEqual(got, want, t)
}

func TestShouldPutTwiceAndGet(t *testing.T) {
	c := NewLLMapLRUCache[int, int](5)
	c.Put(1, 2)
	c.Put(3, 4)
	got, _ := c.Get(3)
	want := 4
	assert.AssertEqual(got, want, t)
}

func TestShouldReturnFalseForNonExistentKey(t *testing.T) {
	c := NewLLMapLRUCache[int, int](5)
	_, got := c.Get(1)
	assert.AssertFalse(got, t)
}

func TestShouldReturnTrueForExistentKey(t *testing.T) {
	c := NewLLMapLRUCache[int, int](5)
	c.Put(1, 2)
	_, got := c.Get(1)
	assert.AssertTrue(got, t)
}

func TestShouldEvictLeastRecentlyUsedIfSizeExceeded(t *testing.T) {
	c := NewLLMapLRUCache[int, int](1)
	c.Put(1, 2)
	c.Put(3, 4)
	_, got := c.Get(1)
	assert.AssertFalse(got, t)
}

func TestShouldNotEvictLeastRecentlyUsedIfSizeNotExceeded(t *testing.T) {
	c := NewLLMapLRUCache[int, int](2)
	c.Put(1, 2)
	c.Put(3, 4)
	_, got := c.Get(1)
	assert.AssertTrue(got, t)
}

func TestShouldReturnFalseWhenPuttingIntoZeroSizeLRUCache(t *testing.T) {
	c := NewLLMapLRUCache[int, int](0)
	c.Put(1, 2)
	_, got := c.Get(1)
	assert.AssertFalse(got, t)
}
