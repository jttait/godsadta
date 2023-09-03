package immutablelist

import (
	"testing"

	"github.com/jttait/godsa/assert"
)

func TestShouldInstantiateArrayListOfSizeZero(t *testing.T) {
	l := NewImmutableArrayList[int]()
	result := l.Size()
	assert.AssertEqual(result, 0, t)
}

func TestShouldAppendMultipleItemsToArrayList(t *testing.T) {
	l := NewImmutableArrayList[int]()
	l = l.Append(5)
	l = l.Append(6)
	l = l.Append(7)
	result, _ := l.Get(0)
	assert.AssertEqual(result, 5, t)
	result, _ = l.Get(1)
	assert.AssertEqual(result, 6, t)
	result, _ = l.Get(2)
	assert.AssertEqual(result, 7, t)
}

func TestShouldIncreaseSizeWhenAppendingToArrayList(t *testing.T) {
	l := NewImmutableArrayList[int]()
	l = l.Append(5)
	result := l.Size()
	assert.AssertEqual(result, 1, t)
	l = l.Append(6)
	result = l.Size()
	assert.AssertEqual(result, 2, t)
	l = l.Append(7)
	result = l.Size()
	assert.AssertEqual(result, 3, t)
}

func TestShouldPrependMultipleItemsToArrayList(t *testing.T) {
	l := NewImmutableArrayList[int]()
	l = l.Prepend(5)
	l = l.Prepend(6)
	l = l.Prepend(7)
	result, _ := l.Get(0)
	assert.AssertEqual(result, 7, t)
	result, _ = l.Get(1)
	assert.AssertEqual(result, 6, t)
	result, _ = l.Get(2)
	assert.AssertEqual(result, 5, t)
}

func TestShouldIncreaseSizeWhenPrependingToArrayList(t *testing.T) {
	l := NewImmutableArrayList[int]()
	l = l.Prepend(5)
	result := l.Size()
	assert.AssertEqual(result, 1, t)
	l = l.Prepend(6)
	result = l.Size()
	assert.AssertEqual(result, 2, t)
	l = l.Prepend(7)
	result = l.Size()
	assert.AssertEqual(result, 3, t)
}

func TestShouldBeFalseIfGetOnEmptyArray(t *testing.T) {
	l := NewImmutableArrayList[int]()
	_, ok := l.Get(0)
	assert.AssertFalse(ok, t)
}

func TestShouldBeTrueIfGetIndexIsWithinArray(t *testing.T) {
	l := NewImmutableArrayList[int](5)
	_, ok := l.Get(0)
	assert.AssertTrue(ok, t)
}

func TestShouldBeFalseIfGetIndexIsOutsideArray(t *testing.T) {
	l := NewImmutableArrayList[int](5)
	_, ok := l.Get(1)
	assert.AssertFalse(ok, t)
}

func TestShouldRemoveItemAtIndex(t *testing.T) {
	l := NewImmutableArrayList[int](1, 2, 3)
	l, _ = l.Remove(1)
	result, _ := l.Get(1)
	assert.AssertEqual(result, 3, t)
	l, _ = l.Remove(1)
	_, ok := l.Get(1)
	assert.AssertFalse(ok, t)
}

func TestShouldBeFalseWhenRemovingFromEmptyArrayList(t *testing.T) {
	l := NewImmutableArrayList[int]()
	_, ok := l.Remove(0)
	assert.AssertFalse(ok, t)
}

func TestShouldBeFalseWhenRemovingFromIndexOutsideArrayList(t *testing.T) {
	l := NewImmutableArrayList[int](5)
	l, ok := l.Remove(1)
	assert.AssertFalse(ok, t)
}

func TestShouldBeTrueWhenRemovingFromIndexInsideArrayList(t *testing.T) {
	l := NewImmutableArrayList[int](5)
	l, ok := l.Remove(0)
	assert.AssertTrue(ok, t)
}

func TestShouldInsertItemAtIndex(t *testing.T) {
	l := NewImmutableArrayList[int](1)
	l, _ = l.Insert(0, 2)
	result, _ := l.Get(0)
	assert.AssertEqual(result, 2, t)
}

func TestShouldShiftTheItemCurrentlyAtTheIndexToRight(t *testing.T) {
	l := NewImmutableArrayList[int](1)
	l, _ = l.Insert(0, 2)
	result, _ := l.Get(1)
	assert.AssertEqual(result, 1, t)
}

func TestShouldBeFalseWhenInsertingIntoEmptyArrayList(t *testing.T) {
	l := NewImmutableArrayList[int]()
	l, ok := l.Insert(0, 1)
	assert.AssertFalse(ok, t)
}

func TestShouldBeTrueWhenInsertingIntoArrayList(t *testing.T) {
	l := NewImmutableArrayList[int](1)
	l, ok := l.Insert(0, 2)
	assert.AssertTrue(ok, t)
}

func TestShouldBeFalseWhenInsertingIntoIndexOutsideArrayList(t *testing.T) {
	l := NewImmutableArrayList[int](1)
	l, ok := l.Insert(1, 2)
	assert.AssertFalse(ok, t)
}

func TestShouldApplyMapToArrayList(t *testing.T) {
	l := NewImmutableArrayList[int](1)
	l = l.Map(func(i int) int { return i * 2 })
	result, _ := l.Get(0)
	assert.AssertEqual(result, 2, t)
}

func TestShouldApplyFilterToArrayList(t *testing.T) {
	l := NewImmutableArrayList[int](1, 2, 3, 4)
	l = l.Filter(func(i int) bool { return i%2 == 0 })
	result, _ := l.Get(0)
	assert.AssertEqual(result, 2, t)
	result, _ = l.Get(1)
	assert.AssertEqual(result, 4, t)
}
