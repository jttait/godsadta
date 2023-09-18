package queue

import (
	"testing"

	"github.com/jttait/godsa/assert"
)

func TestShouldBeSizeZeroForNewlyInstantiatedQueueLinkedList(t *testing.T) {
	q := NewQueueLinkedList[int]()
	result := q.Size()
	assert.AssertEqual(result, 0, t)
}

func TestShouldBeSizeOneWhenItemAddedToNewlyInstantiatedQueueLinkedList(t *testing.T) {
	q := NewQueueLinkedList[int]()
	q.Insert(5)
	result := q.Size()
	assert.AssertEqual(result, 1, t)
}

func TestShouldRemoveItemFromQueueLinkedListOfSizeOne(t *testing.T) {
	q := NewQueueLinkedList[int]()
	q.Insert(5)
	result, _ := q.Remove()
	assert.AssertEqual(result, 5, t)
}

func TestShouldBeFalseWhenRemovingFromEmptyQueueLinkedList(t *testing.T) {
	q := NewQueueLinkedList[int]()
	_, ok := q.Remove()
	assert.AssertFalse(ok, t)
}

func TestShouldBeTrueWhenRemovingFromNonEmptyQueueLinkedList(t *testing.T) {
	q := NewQueueLinkedList[int]()
	q.Insert(5)
	_, ok := q.Remove()
	assert.AssertTrue(ok, t)
}

func TestShouldRemoveIntegerThatWasAddedEarliestFromQueueLinkedList(t *testing.T) {
	q := NewQueueLinkedList[int]()
	q.Insert(5)
	q.Insert(6)
	result, _ := q.Remove()
	assert.AssertEqual(result, 5, t)
}

func TestShouldRemoveStringThatWasAddedEarliestFromQueueLinkedList(t *testing.T) {
	q := NewQueueLinkedList[string]()
	q.Insert("five")
	q.Insert("six")
	result, _ := q.Remove()
	assert.AssertEqual(result, "five", t)
}

func TestShouldRemoveMultipleStringsFromQueueLinkedList(t *testing.T) {
	q := NewQueueLinkedList[string]()
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
