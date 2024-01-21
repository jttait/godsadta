package godsa

import (
	"testing"

	"github.com/jttait/godsa/assert"
)

func TestShouldBeSizeZeroForNewlyInstantiatedLLDEQueue(t *testing.T) {
	q := NewLLDEQueue[int]()
	result := q.Size()
	assert.AssertEqual(result, 0, t)
}

func TestShouldBeSizeOneWhenItemAddedToNewlyInstantiatedLLDEQueue(t *testing.T) {
	q := NewLLDEQueue[int](5)
	result := q.Size()
	assert.AssertEqual(result, 1, t)
}

func TestShouldRemoveItemFromLLDEQueueOfSizeOne(t *testing.T) {
	q := NewLLDEQueue[int](5)
	result, _ := q.RemoveFront()
	assert.AssertEqual(result, 5, t)
}

func TestShouldBeFalseWhenRemovingFromEmptyLLDEQueue(t *testing.T) {
	q := NewLLDEQueue[int]()
	_, ok := q.RemoveFront()
	assert.AssertFalse(ok, t)
}

func TestShouldBeTrueWhenRemovingFromNonEmptyLLDEQueue(t *testing.T) {
	q := NewLLDEQueue[int](1, 2, 3, 4, 5)
	_, ok := q.RemoveFront()
	assert.AssertTrue(ok, t)
}

func TestShouldRemoveFrontFromLLDEQueueOfIntergers(t *testing.T) {
	q := NewLLDEQueue[int](1, 2, 3, 4, 5)
	result, _ := q.RemoveFront()
	assert.AssertEqual(result, 1, t)
}

func TestShouldRemoveLastFromLLDEQueueOfStrings(t *testing.T) {
	q := NewLLDEQueue[string]("one", "two", "three", "four", "five")
	result, _ := q.RemoveLast()
	assert.AssertEqual(result, "five", t)
	result, _ = q.RemoveLast()
	assert.AssertEqual(result, "four", t)
}

func TestShouldRemoveFrontFromLLDEQueueOfStrings(t *testing.T) {
	q := NewLLDEQueue[string]("one", "two", "three", "four", "five")
	result, _ := q.RemoveFront()
	assert.AssertEqual(result, "one", t)
	result, _ = q.RemoveFront()
	assert.AssertEqual(result, "two", t)
}
