package godsa

import (
	"testing"

	"github.com/jttait/godsa/assert"
)

func getListImplementations(values ...int) []List[int] {
	return []List[int]{
		NewLLList[int](values...),
		NewSliceList[int](values...),
	}
}

func TestShouldInstantiateListOfSizeZero(t *testing.T) {
	for _, l := range getListImplementations() {
		result := l.Size()
		assert.AssertEqual(result, 0, t)
	}
}

func TestShouldAppendMultipleItemsToList(t *testing.T) {
	for _, l := range getListImplementations() {
		l.Append(5)
		l.Append(6)
		l.Append(7)
		firstElement, _ := l.Get(0)
		secondElement, _ := l.Get(1)
		thirdElement, _ := l.Get(2)
		assert.AssertEqual(firstElement, 5, t)
		assert.AssertEqual(secondElement, 6, t)
		assert.AssertEqual(thirdElement, 7, t)
	}
}

func TestShouldPrependMultipleItemsToList(t *testing.T) {
	for _, l := range getListImplementations() {
		l.Prepend(5)
		l.Prepend(6)
		l.Prepend(7)
		firstElement, _ := l.Get(0)
		secondElement, _ := l.Get(1)
		thirdElement, _ := l.Get(2)
		assert.AssertEqual(firstElement, 7, t)
		assert.AssertEqual(secondElement, 6, t)
		assert.AssertEqual(thirdElement, 5, t)
	}
}

func TestShouldBeFalseIfGetOnEmptyList(t *testing.T) {
	for _, l := range getListImplementations() {
		_, ok := l.Get(0)
		assert.AssertFalse(ok, t)
	}
}

func TestShouldBeTrueIfGetIndexIsWithinList(t *testing.T) {
	for _, l := range getListImplementations(5) {
		_, ok := l.Get(0)
		assert.AssertTrue(ok, t)
	}
}

func TestShouldBeFalseIfGetIndexIsOutsideList(t *testing.T) {
	for _, l := range getListImplementations(5) {
		_, ok := l.Get(1)
		assert.AssertFalse(ok, t)
	}
}

func TestShouldRemoveItemAtIndexOfList(t *testing.T) {
	for _, l := range getListImplementations(1, 2, 3) {
		_ = l.Remove(1)
		result, _ := l.Get(1)
		assert.AssertEqual(result, 3, t)
	}
}

func TestShouldBeFalseWhenRemovingFromEmptyList(t *testing.T) {
	for _, l := range getListImplementations() {
		ok := l.Remove(0)
		assert.AssertFalse(ok, t)
	}
}

func TestShouldBeFalseWhenRemovingFromIndexOutsideList(t *testing.T) {
	for _, l := range getListImplementations(5) {
		ok := l.Remove(1)
		assert.AssertFalse(ok, t)
	}
}

func TestShouldBeTrueWhenRemovingFromIndexInsideList(t *testing.T) {
	for _, l := range getListImplementations(5) {
		ok := l.Remove(0)
		assert.AssertTrue(ok, t)
	}
}

func TestShouldInsertItemAtIndexOfList(t *testing.T) {
	for _, l := range getListImplementations(1) {
		_ = l.Insert(0, 2)
		result, _ := l.Get(0)
		assert.AssertEqual(result, 2, t)
	}
}

func TestShouldShiftTheItemCurrentlyAtTheIndexToRightWithList(t *testing.T) {
	for _, l := range getListImplementations(1) {
		_ = l.Insert(0, 2)
		result, _ := l.Get(1)
		assert.AssertEqual(result, 1, t)
	}
}

func TestShouldBeFalseWhenInsertingIntoEmptyList(t *testing.T) {
	for _, l := range getListImplementations() {
		ok := l.Insert(0, 1)
		assert.AssertFalse(ok, t)
	}
}

func TestShouldBeTrueWhenInsertingIntoList(t *testing.T) {
	for _, l := range getListImplementations(1) {
		ok := l.Insert(0, 2)
		assert.AssertTrue(ok, t)
	}
}

func TestShouldBeFalseWhenInsertingIntoIndexOutsideList(t *testing.T) {
	for _, l := range getListImplementations(1) {
		ok := l.Insert(1, 2)
		assert.AssertFalse(ok, t)
	}
}

func TestShouldApplyMapToList(t *testing.T) {
	for _, l := range getListImplementations(1, 2, 3) {
		l.Map(func(i int) int { return i * 2 })
		firstElement, _ := l.Get(0)
		secondElement, _ := l.Get(1)
		thirdElement, _ := l.Get(2)
		assert.AssertEqual(firstElement, 2, t)
		assert.AssertEqual(secondElement, 4, t)
		assert.AssertEqual(thirdElement, 6, t)
	}
}

func TestShouldApplyFilterToList(t *testing.T) {
	for _, l := range getListImplementations(1, 2, 3, 4) {
		l.Filter(func(i int) bool { return i%2 == 0 })
		firstElement, _ := l.Get(0)
		secondElement, _ := l.Get(1)
		assert.AssertEqual(firstElement, 2, t)
		assert.AssertEqual(secondElement, 4, t)
	}
}
