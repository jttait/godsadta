package list

import (
	"testing"

	"github.com/jttait/godsa/assert"
)

func TestShouldInstantiateArrayListOfSizeZero(t *testing.T) {
	l := NewSliceList[int]()
	result := l.Size()
	assert.AssertEqual(result, 0, t)
}

func TestShouldAppendMultipleItemsToSliceList(t *testing.T) {
	l := NewSliceList[int]()
	l.Append(5)
	l.Append(6)
	l.Append(7)
	want := NewSliceList[int](5, 6, 7)
	assert.AssertTrue(l.Equal(want), t)
}

func TestShouldPrependMultipleItemsToSliceList(t *testing.T) {
	l := NewSliceList[int]()
	l.Prepend(5)
	l.Prepend(6)
	l.Prepend(7)
	want := NewSliceList[int](7, 6, 5)
	assert.AssertTrue(l.Equal(want), t)
}

func TestShouldBeFalseIfGetOnEmptySliceList(t *testing.T) {
	l := NewSliceList[int]()
	_, ok := l.Get(0)
	assert.AssertFalse(ok, t)
}

func TestShouldBeTrueIfGetIndexIsWithinSliceList(t *testing.T) {
	l := NewSliceList[int](5)
	_, ok := l.Get(0)
	assert.AssertTrue(ok, t)
}

func TestShouldBeFalseIfGetIndexIsOutsideSliceList(t *testing.T) {
	l := NewSliceList[int](5)
	_, ok := l.Get(1)
	assert.AssertFalse(ok, t)
}

func TestShouldRemoveItemAtIndexOfSliceList(t *testing.T) {
	l := NewSliceList[int](1, 2, 3)
	_ = l.Remove(1)
	want := NewSliceList[int](1, 3)
	assert.AssertTrue(l.Equal(want), t)
}

func TestShouldBeFalseWhenRemovingFromEmptySliceList(t *testing.T) {
	l := NewSliceList[int]()
	ok := l.Remove(0)
	assert.AssertFalse(ok, t)
}

func TestShouldBeFalseWhenRemovingFromIndexOutsideSliceList(t *testing.T) {
	l := NewSliceList[int](5)
	ok := l.Remove(1)
	assert.AssertFalse(ok, t)
}

func TestShouldBeTrueWhenRemovingFromIndexInsideSliceList(t *testing.T) {
	l := NewSliceList[int](5)
	ok := l.Remove(0)
	assert.AssertTrue(ok, t)
}

func TestShouldInsertItemAtIndexOfSliceList(t *testing.T) {
	l := NewSliceList[int](1)
	_ = l.Insert(0, 2)
	want := NewSliceList[int](2, 1)
	assert.AssertTrue(l.Equal(want), t)
}

func TestShouldBeFalseWhenInsertingIntoEmptySliceList(t *testing.T) {
	l := NewSliceList[int]()
	ok := l.Insert(0, 1)
	assert.AssertFalse(ok, t)
}

func TestShouldBeTrueWhenInsertingIntoSliceList(t *testing.T) {
	l := NewSliceList[int](1)
	ok := l.Insert(0, 2)
	assert.AssertTrue(ok, t)
}

func TestShouldBeFalseWhenInsertingIntoIndexOutsideSliceList(t *testing.T) {
	l := NewSliceList[int](1)
	ok := l.Insert(1, 2)
	assert.AssertFalse(ok, t)
}

func TestShouldApplyMapToSliceList(t *testing.T) {
	l := NewSliceList[int](1)
	l.Map(func(i int) int { return i * 2 })
	want := NewSliceList[int](2)
	assert.AssertTrue(l.Equal(want), t)
}

func TestShouldApplyFilterToSliceList(t *testing.T) {
	l := NewSliceList[int](1, 2, 3, 4)
	l.Filter(func(i int) bool { return i%2 == 0 })
	want := NewSliceList[int](2, 4)
	assert.AssertTrue(l.Equal(want), t)
}

func TestShouldBeTrueForEqualOnTwoIdenticalSliceLists(t *testing.T) {
	a := NewSliceList[int](1, 2, 3)
	b := NewSliceList[int](1, 2, 3)
	result := a.Equal(b)
	assert.AssertTrue(result, t)
}

func TestShouldBeFalseForEqualOnTwoDifferentSliceLists(t *testing.T) {
	a := NewSliceList[int](1, 2, 3)
	b := NewSliceList[int](1, 2, 4)
	result := a.Equal(b)
	assert.AssertFalse(result, t)
}
