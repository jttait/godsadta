package doubleendedqueue

import (
	"testing"

	"github.com/jttait/godsa/assert"
)

func TestShouldBeSizeZeroForNewlyInstantiatedDoubleEndedQueueLinkedList(t *testing.T) {
	q := NewDoubleEndedQueueLinkedList[int]()
	result := q.Size()
	assert.AssertEqual(result, 0, t)
}

func TestShouldBeSizeOneWhenItemAddedToNewlyInstantiatedDoubleEndedQueueLinkedList(t *testing.T) {
	q := NewDoubleEndedQueueLinkedList[int](5)
	result := q.Size()
	assert.AssertEqual(result, 1, t)
}

func TestShouldRemoveItemFromDoubleEndedQueueLinkedListOfSizeOne(t *testing.T) {
	q := NewDoubleEndedQueueLinkedList[int](5)
	result, _ := q.RemoveFront()
	assert.AssertEqual(result, 5, t)
}

func TestShouldBeFalseWhenRemovingFromEmptyDoubleEndedQueueLinkedList(t *testing.T) {
	q := NewDoubleEndedQueueLinkedList[int]()
	_, ok := q.RemoveFront()
	assert.AssertFalse(ok, t)
}

func TestShouldBeTrueWhenRemovingFromNonEmptyDoubleEndedQueueLinkedList(t *testing.T) {
	q := NewDoubleEndedQueueLinkedList[int](1, 2, 3, 4, 5)
	_, ok := q.RemoveFront()
	assert.AssertTrue(ok, t)
}

func TestShouldRemoveFrontFromDoubleEndedQueueLinkedListOfIntergers(t *testing.T) {
	q := NewDoubleEndedQueueLinkedList[int](1, 2, 3, 4, 5)
	result, _ := q.RemoveFront()
	assert.AssertEqual(result, 1, t)
}

func TestShouldRemoveLastFromDoubleEndedQueueLinkedListOfStrings(t *testing.T) {
	q := NewDoubleEndedQueueLinkedList[string]("one", "two", "three", "four", "five")
	result, _ := q.RemoveLast()
	assert.AssertEqual(result, "five", t)
	result, _ = q.RemoveLast()
	assert.AssertEqual(result, "four", t)
}

func TestShouldRemoveFrontFromDoubleEndedQueueLinkedListOfStrings(t *testing.T) {
	q := NewDoubleEndedQueueLinkedList[string]("one", "two", "three", "four", "five")
	result, _ := q.RemoveFront()
	assert.AssertEqual(result, "one", t)
	result, _ = q.RemoveFront()
	assert.AssertEqual(result, "two", t)
}
