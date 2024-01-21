package godsa

import (
	"testing"

	"github.com/jttait/godsa/assert"
)

func TestShouldBeSizeZeroForNewlyInstantiatedLLQueue(t *testing.T) {
	q := NewLLQueue[int]()
	result := q.Size()
	assert.AssertEqual(result, 0, t)
}

func TestShouldBeSizeOneWhenItemAddedToNewlyInstantiatedLLQueue(t *testing.T) {
	q := NewLLQueue[int]()
	q.Insert(5)
	result := q.Size()
	assert.AssertEqual(result, 1, t)
}

func TestShouldRemoveItemFromLLQueueOfSizeOne(t *testing.T) {
	q := NewLLQueue[int]()
	q.Insert(5)
	result, _ := q.Remove()
	assert.AssertEqual(result, 5, t)
}

func TestShouldBeFalseWhenRemovingFromEmptyLLQueue(t *testing.T) {
	q := NewLLQueue[int]()
	_, ok := q.Remove()
	assert.AssertFalse(ok, t)
}

func TestShouldBeTrueWhenRemovingFromNonEmptyLLQueue(t *testing.T) {
	q := NewLLQueue[int]()
	q.Insert(5)
	_, ok := q.Remove()
	assert.AssertTrue(ok, t)
}

func TestShouldRemoveIntegerThatWasAddedEarliestFromLLQueue(t *testing.T) {
	q := NewLLQueue[int]()
	q.Insert(5)
	q.Insert(6)
	result, _ := q.Remove()
	assert.AssertEqual(result, 5, t)
}

func TestShouldRemoveStringThatWasAddedEarliestFromLLQueue(t *testing.T) {
	q := NewLLQueue[string]()
	q.Insert("five")
	q.Insert("six")
	result, _ := q.Remove()
	assert.AssertEqual(result, "five", t)
}

func TestShouldRemoveMultipleStringsFromLLQueue(t *testing.T) {
	q := NewLLQueue[string]()
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
