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
	want := NewImmutableArrayList[int](5, 6, 7)
	result := l.Equal(want)
	assert.AssertTrue(result, t)
}

func TestShouldPrependMultipleItemsToArrayList(t *testing.T) {
	l := NewImmutableArrayList[int]()
	l = l.Prepend(5)
	l = l.Prepend(6)
	l = l.Prepend(7)
	want := NewImmutableArrayList[int](7, 6, 5)
	result := l.Equal(want)
	assert.AssertTrue(result, t)
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
	want := NewImmutableArrayList[int](1, 3)
	result := l.Equal(want)
	assert.AssertTrue(result, t)
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
	l, _ = l.Insert(1, 3)
	want := NewImmutableArrayList[int](2, 3, 1)
	result := l.Equal(want)
	assert.AssertTrue(result, t)
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
	l := NewImmutableArrayList[int](1, 2, 3)
	l = l.Map(func(i int) int { return i * 2 })
	want := NewImmutableArrayList[int](2, 4, 6)
	result := l.Equal(want)
	assert.AssertTrue(result, t)
}

func TestShouldApplyFilterToArrayList(t *testing.T) {
	l := NewImmutableArrayList[int](1, 2, 3, 4)
	l = l.Filter(func(i int) bool { return i%2 == 0 })
	want := NewImmutableArrayList[int](2, 4)
	result := l.Equal(want)
	assert.AssertTrue(result, t)
}

func TestShouldBeTrueForIdenticalImmutableArrayLists(t *testing.T) {
	a := NewImmutableArrayList[int](1, 2, 3, 4)
	b := NewImmutableArrayList[int](1, 2, 3, 4)
	result := a.Equal(b)
	assert.AssertTrue(result, t)
}

func TestShouldBeFalseForDifferentImmutableArrayLists(t *testing.T) {
	a := NewImmutableArrayList[int](1, 2, 3, 4)
	b := NewImmutableArrayList[int](1, 2, 3, 5)
	result := a.Equal(b)
	assert.AssertFalse(result, t)
}
