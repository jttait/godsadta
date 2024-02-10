package queue

import (
	"testing"

	"github.com/jttait/godsa/assert"
)

func getQueueImplementationsInt() []Queue[int] {
	return []Queue[int]{
		NewLLQueue[int](),
	}
}

func getQueueImplementationsString() []Queue[string] {
	return []Queue[string]{
		NewLLQueue[string](),
	}
}

func TestShouldBeSizeZeroForNewlyInstantiatedQueue(t *testing.T) {
	for _, q := range getQueueImplementationsInt() {
		result := q.Size()
		assert.AssertEqual(result, 0, t)
	}
}

func TestShouldBeSizeOneWhenItemAddedToNewlyInstantiatedQueue(t *testing.T) {
	for _, q := range getQueueImplementationsInt() {
		q.Insert(5)
		result := q.Size()
		assert.AssertEqual(result, 1, t)
	}
}

func TestShouldRemoveItemFromQueueOfSizeOne(t *testing.T) {
	for _, q := range getQueueImplementationsInt() {
		q.Insert(5)
		result, _ := q.Remove()
		assert.AssertEqual(result, 5, t)
	}
}

func TestShouldBeFalseWhenRemovingFromEmptyQueue(t *testing.T) {
	for _, q := range getQueueImplementationsInt() {
		_, ok := q.Remove()
		assert.AssertFalse(ok, t)
	}
}

func TestShouldBeTrueWhenRemovingFromNonEmptyQueue(t *testing.T) {
	for _, q := range getQueueImplementationsInt() {
		q.Insert(5)
		_, ok := q.Remove()
		assert.AssertTrue(ok, t)
	}
}

func TestShouldRemoveIntegerThatWasAddedEarliestFromQueue(t *testing.T) {
	for _, q := range getQueueImplementationsInt() {
		q.Insert(5)
		q.Insert(6)
		result, _ := q.Remove()
		assert.AssertEqual(result, 5, t)
	}
}

func TestShouldRemoveStringThatWasAddedEarliestFromQueue(t *testing.T) {
	for _, q := range getQueueImplementationsString() {
		q.Insert("five")
		q.Insert("six")
		result, _ := q.Remove()
		assert.AssertEqual(result, "five", t)
	}
}

func TestShouldRemoveMultipleStringsFromQueue(t *testing.T) {
	for _, q := range getQueueImplementationsString() {
		q.Insert("five")
		q.Insert("six")
		q.Insert("seven")
		result, _ := q.Remove()
		assert.AssertEqual(result, "five", t)
		result, _ = q.Remove()
		assert.AssertEqual(result, "six", t)
		result, _ = q.Remove()
		assert.AssertEqual(result, "seven", t)
		_, ok := q.Remove()
		assert.AssertFalse(ok, t)
	}
}
