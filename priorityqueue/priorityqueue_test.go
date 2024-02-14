package priorityqueue

import (
	"testing"

	"github.com/jttait/godsadta/assert"
)

func getPriorityQueueImplementations() []PriorityQueue[int] {
	return []PriorityQueue[int]{
		NewMaxHeapPriorityQueue[int](),
	}
}

func TestShouldBeSizeZeroForNewlyInstantiatedPriorityQueue(t *testing.T) {
	for _, q := range getPriorityQueueImplementations() {
		result := q.Size()
		assert.AssertEqual(result, 0, t)
	}
}

func TestShouldBeSizeOneWhenItemAddedToNewlyInstantiatedPriorityQueue(t *testing.T) {
	for _, q := range getPriorityQueueImplementations() {
		q.Insert(5)
		result := q.Size()
		assert.AssertEqual(result, 1, t)
	}
}

func TestShouldBeSizeTwoWhenAddingTwoItems(t *testing.T) {
	for _, q := range getPriorityQueueImplementations() {
		q.Insert(5)
		q.Insert(6)
		result := q.Size()
		assert.AssertEqual(result, 2, t)
	}
}

func TestShouldBeSizeTwoWhenAddingTwoIdenticalItems(t *testing.T) {
	for _, q := range getPriorityQueueImplementations() {
		q.Insert(5)
		q.Insert(5)
		result := q.Size()
		assert.AssertEqual(result, 2, t)
	}
}

func TestShouldBeHighestPriorityItemWhenPollingPriorityQueue(t *testing.T) {
	for _, q := range getPriorityQueueImplementations() {
		q.Insert(5)
		q.Insert(3)
		result, _ := q.Extract()
		assert.AssertEqual(result, 5, t)
	}
}

func TestShouldBeTrueWhenPollingNonEmptyPriorityQueue(t *testing.T) {
	for _, q := range getPriorityQueueImplementations() {
		q.Insert(5)
		_, ok := q.Extract()
		assert.AssertTrue(ok, t)
	}
}

func TestShouldBeFalseWhenPollingEmptyPriorityQueue(t *testing.T) {
	for _, q := range getPriorityQueueImplementations() {
		_, ok := q.Extract()
		assert.AssertFalse(ok, t)
	}
}

func TestShouldBeHighestPriorityItemWhenPeekingPriorityQueue(t *testing.T) {
	for _, q := range getPriorityQueueImplementations() {
		q.Insert(5)
		q.Insert(3)
		result, _ := q.Peek()
		assert.AssertEqual(result, 5, t)
	}
}

func TestShouldBeTrueWhenPeekingNonEmptyPriorityQueue(t *testing.T) {
	for _, q := range getPriorityQueueImplementations() {
		q.Insert(5)
		_, ok := q.Peek()
		assert.AssertTrue(ok, t)
	}
}

func TestShouldBeFalseWhenPeekingEmptyPriorityQueue(t *testing.T) {
	for _, q := range getPriorityQueueImplementations() {
		_, ok := q.Peek()
		assert.AssertFalse(ok, t)
	}
}

func TestShouldBeSameSizeAfterPeekingNonEmptyPriorityQueue(t *testing.T) {
	for _, q := range getPriorityQueueImplementations() {
		q.Insert(5)
		_, _ = q.Peek()
		result := q.Size()
		assert.AssertEqual(result, 1, t)
	}
}

func TestShouldBeSizeZeroAfterPollingPriorityQueueOfSizeOne(t *testing.T) {
	for _, q := range getPriorityQueueImplementations() {
		q.Insert(5)
		_, _ = q.Extract()
		result := q.Size()
		assert.AssertEqual(result, 0, t)
	}
}
