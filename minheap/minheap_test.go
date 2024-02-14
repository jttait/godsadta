package minheap

import (
	"testing"

	"github.com/jttait/godsadta/assert"
)

func getMinHeapImplementations(values ...int) []MinHeap[int] {
	return []MinHeap[int]{
		NewArrayMinHeap[int](values...),
	}
}

func TestShouldInstantiateMinHeap(t *testing.T) {
	for _, m := range getMinHeapImplementations() {
		if m == nil {
			t.Fatalf("Should not be nil.")
		}
	}
}

func TestShouldInsertItemToMinHeap(t *testing.T) {
	for _, m := range getMinHeapImplementations() {
		m.Insert(5)
	}
}

func TestShouldExtractFromMinHeapWithOneItem(t *testing.T) {
	for _, m := range getMinHeapImplementations(5) {
		result := m.Extract()
		assert.AssertEqual(result, 5, t)
	}
}

func TestShouldExtractFromMinHeapWithTwoItems(t *testing.T) {
	for _, m := range getMinHeapImplementations(5, 15, 10) {
		result := m.Extract()
		assert.AssertEqual(result, 5, t)
	}
}

func TestShouldExtractMultipleTimesFromMinHeap(t *testing.T) {
	for _, m := range getMinHeapImplementations(78, 39, 24, 98, 72, 77, 71, 68, 53, 41, 60, 54, 92, 36, 75, 47, 33, 80, 9, 61) {
		result := m.Extract()
		assert.AssertEqual(result, 9, t)
		result = m.Extract()
		assert.AssertEqual(result, 24, t)
		result = m.Extract()
		assert.AssertEqual(result, 33, t)
	}
}

func TestShouldPeekTopItemFromMinHeap(t *testing.T) {
	for _, m := range getMinHeapImplementations(5, 15, 10) {
		result := m.Peek()
		assert.AssertEqual(result, 5, t)
	}
}

func TestShouldBeSameSizeAfterPeekingMinHeap(t *testing.T) {
	for _, m := range getMinHeapImplementations(5, 15, 10) {
		_ = m.Peek()
		result := m.Size()
		assert.AssertEqual(result, 3, t)
	}
}

func TestShouldBeSizeZeroForNewlyInstantiatedMinHeap(t *testing.T) {
	for _, m := range getMinHeapImplementations() {
		result := m.Size()
		assert.AssertEqual(result, 0, t)
	}
}

func TestShouldBeSizeTwoForMinHeapWithTwoItems(t *testing.T) {
	for _, m := range getMinHeapImplementations(1, 2) {
		result := m.Size()
		assert.AssertEqual(result, 2, t)
	}
}
