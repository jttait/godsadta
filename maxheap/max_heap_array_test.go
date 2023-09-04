package maxheap

import (
	"testing"

	"github.com/jttait/godsa/assert"
)

func TestShouldInstantiateMaxHeap(t *testing.T) {
	m := NewMaxHeapArray[int]()
	if m == nil {
		t.Fatalf("Should not be nil.")
	}
}

func TestShouldInsertItemToMaxHeap(t *testing.T) {
	m := NewMaxHeapArray[int]()
	m.Insert(5)
}

func TestShouldExtractFromMaxHeapWithOneItem(t *testing.T) {
	m := NewMaxHeapArray[int]()
	m.Insert(5)
	result, _ := m.Extract()
	assert.AssertEqual(result, 5, t)
}

func TestShouldExtractFromMaxHeapWithTwoItems(t *testing.T) {
	m := NewMaxHeapArray[int](5, 15, 10)
	result, _ := m.Extract()
	assert.AssertEqual(result, 15, t)
}

func TestShouldExtractMultipleTimesFromMaxHeap(t *testing.T) {
	m := NewMaxHeapArray[int](78, 39, 24, 98, 72, 77, 71, 68, 53, 41, 60, 54, 92, 36, 75, 47, 33, 80, 9, 61)
	result, _ := m.Extract()
	assert.AssertEqual(result, 98, t)
	result, _ = m.Extract()
	assert.AssertEqual(result, 92, t)
	result, _ = m.Extract()
	assert.AssertEqual(result, 80, t)
}

func TestShouldPeekTopItemFromMaxHeap(t *testing.T) {
	m := NewMaxHeapArray[int](5, 15, 10)
	result, _ := m.Peek()
	assert.AssertEqual(result, 15, t)
}

func TestShouldBeSameSizeAfterPeekingMaxHeap(t *testing.T) {
	m := NewMaxHeapArray[int](5, 15, 10)
	_, _ = m.Peek()
	result := m.Size()
	assert.AssertEqual(result, 3, t)
}

func TestShouldBeSizeZeroForNewlyInstantiatedMaxHeap(t *testing.T) {
	m := NewMaxHeapArray[int]()
	result := m.Size()
	assert.AssertEqual(result, 0, t)
}

func TestShouldBeSizeTwoForMaxHeapWithTwoItems(t *testing.T) {
	m := NewMaxHeapArray[int](1, 2)
	result := m.Size()
	assert.AssertEqual(result, 2, t)
}
