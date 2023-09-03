package list

import (
	"testing"

	"github.com/jttait/godsa/assert"
)

func TestShouldInstantiateLinkedListListOfSizeZero(t *testing.T) {
	l := NewLinkedListList[int]()
	result := l.Size()
	assert.AssertEqual(result, 0, t)
}

func TestShouldAppendMultipleItemsToLinkedListList(t *testing.T) {
	l := NewLinkedListList[int]()
	l.Append(5)
	l.Append(6)
	l.Append(7)
	result, _ := l.Get(0)
	assert.AssertEqual(result, 5, t)
	result, _ = l.Get(1)
	assert.AssertEqual(result, 6, t)
	result, _ = l.Get(2)
	assert.AssertEqual(result, 7, t)
}

func TestShouldIncreaseSizeWhenAppendingToLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	l.Append(5)
	result := l.Size()
	assert.AssertEqual(result, 1, t)
	l.Append(6)
	result = l.Size()
	assert.AssertEqual(result, 2, t)
	l.Append(7)
	result = l.Size()
	assert.AssertEqual(result, 3, t)
}

func TestShouldPrependMultipleItemsToLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	l.Prepend(5)
	l.Prepend(6)
	l.Prepend(7)
	result, _ := l.Get(0)
	assert.AssertEqual(result, 7, t)
	result, _ = l.Get(1)
	assert.AssertEqual(result, 6, t)
	result, _ = l.Get(2)
	assert.AssertEqual(result, 5, t)
}

func TestShouldIncreaseSizeWhenPrependingToLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	l.Prepend(5)
	result := l.Size()
	assert.AssertEqual(result, 1, t)
	l.Prepend(6)
	result = l.Size()
	assert.AssertEqual(result, 2, t)
	l.Prepend(7)
	result = l.Size()
	assert.AssertEqual(result, 3, t)
}

func TestShouldBeFalseIfGetOnEmptyLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	_, ok := l.Get(0)
	assert.AssertFalse(ok, t)
}

func TestShouldBeTrueIfGetIndexIsWithinLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	l.Append(5)
	_, ok := l.Get(0)
	assert.AssertTrue(ok, t)
}

func TestShouldBeFalseIfGetIndexIsOutsideLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	l.Append(5)
	_, ok := l.Get(1)
	assert.AssertFalse(ok, t)
}

func TestShouldRemoveItemAtIndexOfLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	_ = l.Remove(1)
	result, _ := l.Get(1)
	assert.AssertEqual(result, 3, t)
	_ = l.Remove(1)
	_, ok := l.Get(1)
	assert.AssertFalse(ok, t)
}

func TestShouldBeFalseWhenRemovingFromEmptyLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	ok := l.Remove(0)
	assert.AssertFalse(ok, t)
}

func TestShouldBeFalseWhenRemovingFromIndexOutsideLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	l.Append(5)
	ok := l.Remove(1)
	assert.AssertFalse(ok, t)
}

func TestShouldBeTrueWhenRemovingFromIndexInsideLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	l.Append(5)
	ok := l.Remove(0)
	assert.AssertTrue(ok, t)
}

func TestShouldInsertItemAtIndexOfLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	l.Append(1)
	_ = l.Insert(0, 2)
	result, _ := l.Get(0)
	assert.AssertEqual(result, 2, t)
}

func TestShouldShiftTheItemCurrentlyAtTheIndexToRightWithLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	l.Append(1)
	_ = l.Insert(0, 2)
	result, _ := l.Get(1)
	assert.AssertEqual(result, 1, t)
}

func TestShouldBeFalseWhenInsertingIntoEmptyLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	ok := l.Insert(0, 1)
	assert.AssertFalse(ok, t)
}

func TestShouldBeTrueWhenInsertingIntoLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	l.Append(1)
	ok := l.Insert(0, 2)
	assert.AssertTrue(ok, t)
}

func TestShouldBeFalseWhenInsertingIntoIndexOutsideLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	l.Append(1)
	ok := l.Insert(1, 2)
	assert.AssertFalse(ok, t)
}

func TestShouldApplyMapToLinkedListList(t *testing.T) {
	l := NewLinkedListList[int]()
	l.Append(1)
	l.Map(func(i int) int { return i * 2 })
	result, _ := l.Get(0)
	assert.AssertEqual(result, 2, t)
}

func TestShouldApplyFilterToLinkedListList(t *testing.T) {
	l := NewLinkedListList[int]()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Filter(func(i int) bool { return i%2 == 0 })
	result, _ := l.Get(0)
	assert.AssertEqual(result, 2, t)
	result, _ = l.Get(1)
	assert.AssertEqual(result, 4, t)
}
