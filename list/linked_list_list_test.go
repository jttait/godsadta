package list

import "testing"

func TestShouldInstantiateLinkedListListOfSizeZero(t *testing.T) {
	l := NewLinkedListList[int]()
	result := l.Size()
	assertEqual(result, 0, t)
}

func TestShouldAppendMultipleItemsToLinkedListList(t *testing.T) {
	l := NewLinkedListList[int]()
	l.Append(5)
	l.Append(6)
	l.Append(7)
	result, _ := l.Get(0)
	assertEqual(result, 5, t)
	result, _ = l.Get(1)
	assertEqual(result, 6, t)
	result, _ = l.Get(2)
	assertEqual(result, 7, t)
}

func TestShouldIncreaseSizeWhenAppendingToLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	l.Append(5)
	result := l.Size()
	assertEqual(result, 1, t)
	l.Append(6)
	result = l.Size()
	assertEqual(result, 2, t)
	l.Append(7)
	result = l.Size()
	assertEqual(result, 3, t)
}

func TestShouldPrependMultipleItemsToLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	l.Prepend(5)
	l.Prepend(6)
	l.Prepend(7)
	result, _ := l.Get(0)
	assertEqual(result, 7, t)
	result, _ = l.Get(1)
	assertEqual(result, 6, t)
	result, _ = l.Get(2)
	assertEqual(result, 5, t)
}

func TestShouldIncreaseSizeWhenPrependingToLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	l.Prepend(5)
	result := l.Size()
	assertEqual(result, 1, t)
	l.Prepend(6)
	result = l.Size()
	assertEqual(result, 2, t)
	l.Prepend(7)
	result = l.Size()
	assertEqual(result, 3, t)
}

func TestShouldBeFalseIfGetOnEmptyLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	_, ok := l.Get(0)
	assertFalse(ok, t)
}

func TestShouldBeTrueIfGetIndexIsWithinLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	l.Append(5)
	_, ok := l.Get(0)
	assertTrue(ok, t)
}

func TestShouldBeFalseIfGetIndexIsOutsideLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	l.Append(5)
	_, ok := l.Get(1)
	assertFalse(ok, t)
}

func TestShouldRemoveItemAtIndexOfLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	_ = l.Remove(1)
	result, _ := l.Get(1)
	assertEqual(result, 3, t)
	_ = l.Remove(1)
	_, ok := l.Get(1)
	assertFalse(ok, t)
}

func TestShouldBeFalseWhenRemovingFromEmptyLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	ok := l.Remove(0)
	assertFalse(ok, t)
}

func TestShouldBeFalseWhenRemovingFromIndexOutsideLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	l.Append(5)
	ok := l.Remove(1)
	assertFalse(ok, t)
}

func TestShouldBeTrueWhenRemovingFromIndexInsideLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	l.Append(5)
	ok := l.Remove(0)
	assertTrue(ok, t)
}

func TestShouldInsertItemAtIndexOfLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	l.Append(1)
	_ = l.Insert(0, 2)
	result, _ := l.Get(0)
	assertEqual(result, 2, t)
}

func TestShouldShiftTheItemCurrentlyAtTheIndexToRightWithLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	l.Append(1)
	_ = l.Insert(0, 2)
	result, _ := l.Get(1)
	assertEqual(result, 1, t)
}

func TestShouldBeFalseWhenInsertingIntoEmptyLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	ok := l.Insert(0, 1)
	assertFalse(ok, t)
}

func TestShouldBeTrueWhenInsertingIntoLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	l.Append(1)
	ok := l.Insert(0, 2)
	assertTrue(ok, t)
}

func TestShouldBeFalseWhenInsertingIntoIndexOutsideLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	l.Append(1)
	ok := l.Insert(1, 2)
	assertFalse(ok, t)
}

func TestShouldApplyMapToLinkedListList(t *testing.T) {
	l := NewLinkedListList[int]()
	l.Append(1)
	l.Map(func(i int) int { return i * 2 })
	result, _ := l.Get(0)
	assertEqual(result, 2, t)
}

func TestShouldApplyFilterToLinkedListList(t *testing.T) {
	l := NewLinkedListList[int]()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Filter(func(i int) bool { return i%2 == 0 })
	result, _ := l.Get(0)
	assertEqual(result, 2, t)
	result, _ = l.Get(1)
	assertEqual(result, 4, t)
}
