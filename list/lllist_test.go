package list

import (
	"testing"

	"github.com/jttait/godsa/assert"
)

func TestShouldInstantiateLLListOfSizeZero(t *testing.T) {
	l := NewLLList[int]()
	result := l.Size()
	assert.AssertEqual(result, 0, t)
}

func TestShouldAppendMultipleItemsToLLList(t *testing.T) {
	l := NewLinkedListList[int]()
	l.Append(5)
	l.Append(6)
	l.Append(7)
	want := NewLLList[int](5, 6, 7)
	assert.AssertTrue(l.Equal(want), t)
}

func TestShouldPrependMultipleItemsToLLList(t *testing.T) {
	l := NewLLList[int]()
	l.Prepend(5)
	l.Prepend(6)
	l.Prepend(7)
	want := NewLLList[int](7, 6, 5)
	assert.AssertTrue(l.Equal(want), t)
}

func TestShouldBeFalseIfGetOnEmptyLLList(t *testing.T) {
	l := NewLLList[int]()
	_, ok := l.Get(0)
	assert.AssertFalse(ok, t)
}

func TestShouldBeTrueIfGetIndexIsWithinLLList(t *testing.T) {
	l := NewLLList[int](5)
	_, ok := l.Get(0)
	assert.AssertTrue(ok, t)
}

func TestShouldBeFalseIfGetIndexIsOutsideLLList(t *testing.T) {
	l := NewLLList[int](5)
	_, ok := l.Get(1)
	assert.AssertFalse(ok, t)
}

func TestShouldRemoveItemAtIndexOfLLList(t *testing.T) {
	l := NewLLList[int](1, 2, 3)
	_ = l.Remove(1)
	result, _ := l.Get(1)
	assert.AssertEqual(result, 3, t)
}

func TestShouldBeFalseWhenRemovingFromEmptyLLList(t *testing.T) {
	l := NewLLList[int]()
	ok := l.Remove(0)
	assert.AssertFalse(ok, t)
}

func TestShouldBeFalseWhenRemovingFromIndexOutsideLLList(t *testing.T) {
	l := NewLLList[int](5)
	ok := l.Remove(1)
	assert.AssertFalse(ok, t)
}

func TestShouldBeTrueWhenRemovingFromIndexInsideLLList(t *testing.T) {
	l := NewLLList[int](5)
	ok := l.Remove(0)
	assert.AssertTrue(ok, t)
}

func TestShouldInsertItemAtIndexOfLLList(t *testing.T) {
	l := NewLLList[int](1)
	_ = l.Insert(0, 2)
	result, _ := l.Get(0)
	assert.AssertEqual(result, 2, t)
}

func TestShouldShiftTheItemCurrentlyAtTheIndexToRightWithLLList(t *testing.T) {
	l := NewLLList[int](1)
	_ = l.Insert(0, 2)
	result, _ := l.Get(1)
	assert.AssertEqual(result, 1, t)
}

func TestShouldBeFalseWhenInsertingIntoEmptyLLList(t *testing.T) {
	l := NewLLList[int]()
	ok := l.Insert(0, 1)
	assert.AssertFalse(ok, t)
}

func TestShouldBeTrueWhenInsertingIntoLLList(t *testing.T) {
	l := NewLLList[int](1)
	ok := l.Insert(0, 2)
	assert.AssertTrue(ok, t)
}

func TestShouldBeFalseWhenInsertingIntoIndexOutsideLLList(t *testing.T) {
	l := NewLLList[int](1)
	ok := l.Insert(1, 2)
	assert.AssertFalse(ok, t)
}

func TestShouldApplyMapToLLList(t *testing.T) {
	l := NewLLList[int](1, 2, 3)
	l.Map(func(i int) int { return i * 2 })
	want := NewLLList[int](2, 4, 6)
	assert.AssertTrue(l.Equal(want), t)
}

func TestShouldApplyFilterToLLList(t *testing.T) {
	l := NewLLList[int](1, 2, 3, 4)
	l.Filter(func(i int) bool { return i%2 == 0 })
	want := NewLLList[int](2, 4)
	assert.AssertTrue(l.Equal(want), t)
}

func TestShouldBeTrueForEqualOnTwoIdenticalLists(t *testing.T) {
	a := NewLLList[int](1, 2, 3)
	b := NewLLList[int](1, 2, 3)
	result := a.Equal(b)
	assert.AssertTrue(result, t)
}

func TestShouldBeFalseForEqualOnTwoDifferentLists(t *testing.T) {
	a := NewLLList[int](1, 2, 3)
	b := NewLLList[int](1, 2, 4)
	result := a.Equal(b)
	assert.AssertFalse(result, t)
}
