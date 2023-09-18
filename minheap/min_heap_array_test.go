package minheap

import (
	"testing"

	"github.com/jttait/godsa/assert"
)

func TestShouldInstantiateMinHeapArray(t *testing.T) {
	m := NewMinHeapArray[int]()
	if m == nil {
		t.Fatalf("Should not be nil.")
	}
}

func TestShouldInsertItemToMinHeapArray(t *testing.T) {
	m := NewMinHeapArray[int]()
	m.Insert(5)
}

func TestShouldExtractFromMinHeapArrayWithOneItem(t *testing.T) {
	m := NewMinHeapArray[int](5)
	result := m.Extract()
	assert.AssertEqual(result, 5, t)
}

func TestShouldExtractFromMinHeapArrayWithTwoItems(t *testing.T) {
	m := NewMinHeapArray[int](5, 15, 10)
	result := m.Extract()
	assert.AssertEqual(result, 5, t)
}

func TestShouldExtractMultipleTimesFromMinHeapArray(t *testing.T) {
	m := NewMinHeapArray[int](78, 39, 24, 98, 72, 77, 71, 68, 53, 41, 60, 54, 92, 36, 75, 47, 33, 80, 9, 61)
	result := m.Extract()
	assert.AssertEqual(result, 9, t)
	result = m.Extract()
	assert.AssertEqual(result, 24, t)
	result = m.Extract()
	assert.AssertEqual(result, 33, t)
}

func TestShouldPeekTopItemFromMinHeapArray(t *testing.T) {
	m := NewMinHeapArray[int](5, 15, 10)
	result := m.Peek()
	assert.AssertEqual(result, 5, t)
}

func TestShouldBeSameSizeAfterPeekingMinHeapArray(t *testing.T) {
	m := NewMinHeapArray[int](5, 15, 10)
	_ = m.Peek()
	result := m.Size()
	assert.AssertEqual(result, 3, t)
}

func TestShouldBeSizeZeroForNewlyInstantiatedMinHeapArray(t *testing.T) {
	m := NewMinHeapArray[int]()
	result := m.Size()
	assert.AssertEqual(result, 0, t)
}

func TestShouldBeSizeTwoForMinHeapArrayWithTwoItems(t *testing.T) {
	m := NewMinHeapArray[int](1, 2)
	result := m.Size()
	assert.AssertEqual(result, 2, t)
}
