package depriorityqueue

import (
	"testing"

	"github.com/jttait/godsa/assert"
)

func TestShouldInsertIntoNaiveQueue(t *testing.T) {
	q := NewNaiveDEPriorityQueue[int]()
	q.Insert(5)
	result := q.Size()
	assert.AssertEqual(result, 1, t)
}

func TestShouldExtractMaxFromNaiveQueue(t *testing.T) {
	q := NewNaiveDEPriorityQueue[int]()
	q.Insert(5)
	q.Insert(7)
	q.Insert(3)
	result, _ := q.ExtractMax()
	assert.AssertEqual(result, 7, t)
	result, _ = q.ExtractMax()
	assert.AssertEqual(result, 5, t)
	result, _ = q.ExtractMax()
	assert.AssertEqual(result, 3, t)
	_, ok := q.ExtractMax()
	assert.AssertFalse(ok, t)
}

func TestShouldExtractMinFromNaiveQueue(t *testing.T) {
	q := NewNaiveDEPriorityQueue[int]()
	q.Insert(5)
	q.Insert(7)
	q.Insert(3)
	result, _ := q.ExtractMin()
	assert.AssertEqual(result, 3, t)
	result, _ = q.ExtractMin()
	assert.AssertEqual(result, 5, t)
	result, _ = q.ExtractMin()
	assert.AssertEqual(result, 7, t)
	_, ok := q.ExtractMax()
	assert.AssertFalse(ok, t)
}

func TestShouldPeekMaxFromNaiveQueue(t *testing.T) {
	q := NewNaiveDEPriorityQueue[int]()
	q.Insert(5)
	q.Insert(7)
	q.Insert(3)
	result, _ := q.PeekMax()
	assert.AssertEqual(result, 7, t)
	result, _ = q.PeekMax()
	assert.AssertEqual(result, 7, t)
}

func TestShouldPeekMinFromNaiveQueue(t *testing.T) {
	q := NewNaiveDEPriorityQueue[int]()
	q.Insert(5)
	q.Insert(7)
	q.Insert(3)
	result, _ := q.PeekMin()
	assert.AssertEqual(result, 3, t)
	result, _ = q.PeekMin()
	assert.AssertEqual(result, 3, t)
}
