package godsa

import (
	"testing"

	"github.com/jttait/godsa/assert"
)

func getImplementations() []List[int] {
	return []List[int]{
		NewLLList[int](),
		NewSliceList[int](),
	}
}

func TestShouldInstantiateListOfSizeZero(t *testing.T) {
	for _, l := range getImplementations() {
		result := l.Size()
		assert.AssertEqual(result, 0, t)
	}
}

func TestShouldAppendMultipleItemsToList(t *testing.T) {
	l := NewLLList[int]()
	l.Append(5)
	l.Append(6)
	l.Append(7)
	want := NewLLList[int](5, 6, 7)
	assert.AssertTrue(l.Equal(want), t)
}

func TestShouldPrependMultipleItemsToList(t *testing.T) {
	l := NewLLList[int]()
	l.Prepend(5)
	l.Prepend(6)
	l.Prepend(7)
	want := NewLLList[int](7, 6, 5)
	assert.AssertTrue(l.Equal(want), t)
}

func TestShouldBeFalseIfGetOnEmptyList(t *testing.T) {
	l := NewLLList[int]()
	_, ok := l.Get(0)
	assert.AssertFalse(ok, t)
}

func TestShouldBeTrueIfGetIndexIsWithinList(t *testing.T) {
	l := NewLLList[int](5)
	_, ok := l.Get(0)
	assert.AssertTrue(ok, t)
}

func TestShouldBeFalseIfGetIndexIsOutsideList(t *testing.T) {
	l := NewLLList[int](5)
	_, ok := l.Get(1)
	assert.AssertFalse(ok, t)
}

func TestShouldRemoveItemAtIndexOfList(t *testing.T) {
	l := NewLLList[int](1, 2, 3)
	_ = l.Remove(1)
	result, _ := l.Get(1)
	assert.AssertEqual(result, 3, t)
}

func TestShouldBeFalseWhenRemovingFromEmptyList(t *testing.T) {
	l := NewLLList[int]()
	ok := l.Remove(0)
	assert.AssertFalse(ok, t)
}

func TestShouldBeFalseWhenRemovingFromIndexOutsideList(t *testing.T) {
	l := NewLLList[int](5)
	ok := l.Remove(1)
	assert.AssertFalse(ok, t)
}

func TestShouldBeTrueWhenRemovingFromIndexInsideList(t *testing.T) {
	l := NewLLList[int](5)
	ok := l.Remove(0)
	assert.AssertTrue(ok, t)
}

func TestShouldInsertItemAtIndexOfList(t *testing.T) {
	l := NewLLList[int](1)
	_ = l.Insert(0, 2)
	result, _ := l.Get(0)
	assert.AssertEqual(result, 2, t)
}

func TestShouldShiftTheItemCurrentlyAtTheIndexToRightWithList(t *testing.T) {
	l := NewLLList[int](1)
	_ = l.Insert(0, 2)
	result, _ := l.Get(1)
	assert.AssertEqual(result, 1, t)
}

func TestShouldBeFalseWhenInsertingIntoEmptyList(t *testing.T) {
	l := NewLLList[int]()
	ok := l.Insert(0, 1)
	assert.AssertFalse(ok, t)
}

func TestShouldBeTrueWhenInsertingIntoList(t *testing.T) {
	l := NewLLList[int](1)
	ok := l.Insert(0, 2)
	assert.AssertTrue(ok, t)
}

func TestShouldBeFalseWhenInsertingIntoIndexOutsideList(t *testing.T) {
	l := NewLLList[int](1)
	ok := l.Insert(1, 2)
	assert.AssertFalse(ok, t)
}

func TestShouldApplyMapToList(t *testing.T) {
	l := NewLLList[int](1, 2, 3)
	l.Map(func(i int) int { return i * 2 })
	want := NewLLList[int](2, 4, 6)
	assert.AssertTrue(l.Equal(want), t)
}

func TestShouldApplyFilterToList(t *testing.T) {
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
