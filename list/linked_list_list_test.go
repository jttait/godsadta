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
	want := NewLinkedListList[int](5, 6, 7)
	assert.AssertTrue(l.Equal(want), t)
}

func TestShouldPrependMultipleItemsToLinkedListList(t *testing.T) {
	l := NewLinkedListList[int]()
	l.Prepend(5)
	l.Prepend(6)
	l.Prepend(7)
	want := NewLinkedListList[int](7, 6, 5)
	assert.AssertTrue(l.Equal(want), t)
}

func TestShouldBeFalseIfGetOnEmptyLinkedListList(t *testing.T) {
	l := NewLinkedListList[int]()
	_, ok := l.Get(0)
	assert.AssertFalse(ok, t)
}

func TestShouldBeTrueIfGetIndexIsWithinLinkedListList(t *testing.T) {
	l := NewLinkedListList[int](5)
	_, ok := l.Get(0)
	assert.AssertTrue(ok, t)
}

func TestShouldBeFalseIfGetIndexIsOutsideLinkedListList(t *testing.T) {
	l := NewLinkedListList[int](5)
	_, ok := l.Get(1)
	assert.AssertFalse(ok, t)
}

func TestShouldRemoveItemAtIndexOfLinkedListList(t *testing.T) {
	l := NewLinkedListList[int](1, 2, 3)
	_ = l.Remove(1)
	result, _ := l.Get(1)
	assert.AssertEqual(result, 3, t)
}

func TestShouldBeFalseWhenRemovingFromEmptyLinkedListList(t *testing.T) {
	l := NewLinkedListList[int]()
	ok := l.Remove(0)
	assert.AssertFalse(ok, t)
}

func TestShouldBeFalseWhenRemovingFromIndexOutsideLinkedListList(t *testing.T) {
	l := NewLinkedListList[int](5)
	ok := l.Remove(1)
	assert.AssertFalse(ok, t)
}

func TestShouldBeTrueWhenRemovingFromIndexInsideLinkedListList(t *testing.T) {
	l := NewLinkedListList[int](5)
	ok := l.Remove(0)
	assert.AssertTrue(ok, t)
}

func TestShouldInsertItemAtIndexOfLinkedListList(t *testing.T) {
	l := NewLinkedListList[int](1)
	_ = l.Insert(0, 2)
	result, _ := l.Get(0)
	assert.AssertEqual(result, 2, t)
}

func TestShouldShiftTheItemCurrentlyAtTheIndexToRightWithLinkedListList(t *testing.T) {
	l := NewLinkedListList[int](1)
	_ = l.Insert(0, 2)
	result, _ := l.Get(1)
	assert.AssertEqual(result, 1, t)
}

func TestShouldBeFalseWhenInsertingIntoEmptyLinkedListList(t *testing.T) {
	l := NewLinkedListList[int]()
	ok := l.Insert(0, 1)
	assert.AssertFalse(ok, t)
}

func TestShouldBeTrueWhenInsertingIntoLinkedListList(t *testing.T) {
	l := NewLinkedListList[int](1)
	ok := l.Insert(0, 2)
	assert.AssertTrue(ok, t)
}

func TestShouldBeFalseWhenInsertingIntoIndexOutsideLinkedListList(t *testing.T) {
	l := NewLinkedListList[int](1)
	ok := l.Insert(1, 2)
	assert.AssertFalse(ok, t)
}

func TestShouldApplyMapToLinkedListList(t *testing.T) {
	l := NewLinkedListList[int](1, 2, 3)
	l.Map(func(i int) int { return i * 2 })
	want := NewLinkedListList[int](2, 4, 6)
	assert.AssertTrue(l.Equal(want), t)
}

func TestShouldApplyFilterToLinkedListList(t *testing.T) {
	l := NewLinkedListList[int](1, 2, 3, 4)
	l.Filter(func(i int) bool { return i%2 == 0 })
	want := NewLinkedListList[int](2, 4)
	assert.AssertTrue(l.Equal(want), t)
}

func TestShouldBeTrueForEqualOnTwoIdenticalLists(t *testing.T) {
	a := NewLinkedListList[int](1, 2, 3)
	b := NewLinkedListList[int](1, 2, 3)
	result := a.Equal(b)
	assert.AssertTrue(result, t)
}

func TestShouldBeFalseForEqualOnTwoDifferentLists(t *testing.T) {
	a := NewLinkedListList[int](1, 2, 3)
	b := NewLinkedListList[int](1, 2, 4)
	result := a.Equal(b)
	assert.AssertFalse(result, t)
}
