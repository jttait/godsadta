package dequeue

import (
	"testing"

	"github.com/jttait/godsa/assert"
)

func getDEQueueImplementationsInt(values ...int) []DEQueue[int] {
	return []DEQueue[int]{
		NewLLDEQueue[int](values...),
	}
}

func getDEQueueImplementationsString(values ...string) []DEQueue[string] {
	return []DEQueue[string]{
		NewLLDEQueue[string](values...),
	}
}

func TestShouldBeSizeZeroForNewlyInstantiatedDEQueue(t *testing.T) {
	for _, q := range getDEQueueImplementationsInt() {
		result := q.Size()
		assert.AssertEqual(result, 0, t)
	}
}

func TestShouldBeSizeOneWhenItemAddedToNewlyInstantiatedDEQueue(t *testing.T) {
	for _, q := range getDEQueueImplementationsInt(5) {
		result := q.Size()
		assert.AssertEqual(result, 1, t)
	}
}

func TestShouldRemoveItemFromDEQueueOfSizeOne(t *testing.T) {
	for _, q := range getDEQueueImplementationsInt(5) {
		result, _ := q.RemoveFront()
		assert.AssertEqual(result, 5, t)
	}
}

func TestShouldBeFalseWhenRemovingFromEmptyDEQueue(t *testing.T) {
	for _, q := range getDEQueueImplementationsInt() {
		_, ok := q.RemoveFront()
		assert.AssertFalse(ok, t)
	}
}

func TestShouldBeTrueWhenRemovingFromNonEmptyDEQueue(t *testing.T) {
	for _, q := range getDEQueueImplementationsInt(1, 2, 3, 4, 5) {
		_, ok := q.RemoveFront()
		assert.AssertTrue(ok, t)
	}
}

func TestShouldRemoveFrontFromDEQueueOfIntegers(t *testing.T) {
	for _, q := range getDEQueueImplementationsInt(1, 2, 3, 4, 5) {
		result, _ := q.RemoveFront()
		assert.AssertEqual(result, 1, t)
	}
}

func TestShouldRemoveLastFromDEQueueOfStrings(t *testing.T) {
	for _, q := range getDEQueueImplementationsString("one", "two", "three", "four", "five") {
		result, _ := q.RemoveLast()
		assert.AssertEqual(result, "five", t)
		result, _ = q.RemoveLast()
		assert.AssertEqual(result, "four", t)
	}
}

func TestShouldRemoveFrontFromDEQueueOfStrings(t *testing.T) {
	for _, q := range getDEQueueImplementationsString("one", "two", "three", "four", "five") {
		result, _ := q.RemoveFront()
		assert.AssertEqual(result, "one", t)
		result, _ = q.RemoveFront()
		assert.AssertEqual(result, "two", t)
	}
}
