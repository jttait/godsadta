package linkedlist_test

import (
	"testing"

	"github.com/jttait/godsa/assert"
	"github.com/jttait/godsa/dll"
	"github.com/jttait/godsa/linkedlist"
	"github.com/jttait/godsa/sll"
)

func getImplementations() []linkedlist.LinkedList[int] {
	return []linkedlist.LinkedList[int]{
		sll.NewSLL[int](),
		dll.NewDLL[int](),
	}
}

func TestShouldBeSizeZeroForNewlyInstantiatedLinkedList(t *testing.T) {
	for _, ll := range getImplementations() {
		result := ll.Size()
		assert.AssertEqual(result, 0, t)
	}
}

func TestShouldBeSizeOneAfterInsertingIntoLinkedList(t *testing.T) {
	for _, ll := range getImplementations() {
		ll.InsertFront(5)
		result := ll.Size()
		assert.AssertEqual(result, 1, t)
	}
}

func TestShouldBeSizeZeroAfterInsertingAndRemovingFromLinkedList(t *testing.T) {
	for _, ll := range getImplementations() {
		ll.InsertFront(5)
		_, _ = ll.RemoveFront()
		result := ll.Size()
		assert.AssertEqual(result, 0, t)
	}
}

func TestShouldRemoveMultipleItemsOfLinkedList(t *testing.T) {
	for _, ll := range getImplementations() {
		ll.InsertLast(1)
		ll.InsertLast(2)
		ll.InsertLast(3)
		result, _ := ll.RemoveFront()
		assert.AssertEqual(result, 1, t)
		result, _ = ll.RemoveFront()
		assert.AssertEqual(result, 2, t)
		result, _ = ll.RemoveFront()
		assert.AssertEqual(result, 3, t)
	}
}

func TestShouldPeekFromNonEmptyLinkedList(t *testing.T) {
	for _, ll := range getImplementations() {
		ll.InsertLast(5)
		result, _ := ll.PeekFront()
		assert.AssertEqual(result, 5, t)
	}
}

func TestShouldPeekLastFromNonEmptyLinkedList(t *testing.T) {
	for _, ll := range getImplementations() {
		ll.InsertLast(5)
		result, _ := ll.PeekLast()
		assert.AssertEqual(result, 5, t)
	}
}

func TestShouldBeTrueWhenPeekingLastFromNonEmptyLinkedList(t *testing.T) {
	for _, ll := range getImplementations() {
		ll.InsertLast(5)
		_, ok := ll.PeekFront()
		assert.AssertTrue(ok, t)
	}
}

func TestShouldBeFalseWhenPeekingLastFromEmptyLinkedList(t *testing.T) {
	for _, ll := range getImplementations() {
		_, ok := ll.PeekLast()
		assert.AssertFalse(ok, t)
	}
}

func TestShouldBeTrueWhenPeekingFromNonEmptyLinkedList(t *testing.T) {
	for _, ll := range getImplementations() {
		ll.InsertLast(5)
		_, ok := ll.PeekFront()
		assert.AssertTrue(ok, t)
	}
}

func TestShouldBeFalseWhenPeekingFromNonEmptyLinkedList(t *testing.T) {
	for _, ll := range getImplementations() {
		_, ok := ll.PeekFront()
		assert.AssertFalse(ok, t)
	}
}

func TestShouldBeTrueWhenRemovingFrontFromNonEmptyLinkedList(t *testing.T) {
	for _, ll := range getImplementations() {
		ll.InsertLast(5)
		_, ok := ll.RemoveFront()
		assert.AssertTrue(ok, t)
	}
}

func TestShouldBeFalseWhenRemovingFrontFromEmptyLinkedList(t *testing.T) {
	for _, ll := range getImplementations() {
		_, ok := ll.RemoveFront()
		assert.AssertFalse(ok, t)
	}
}

func TestShouldBeFalseWhenGettingFromEmptyLinkedList(t *testing.T) {
	for _, ll := range getImplementations() {
		_, ok := ll.Get(0)
		assert.AssertFalse(ok, t)
	}
}

func TestShouldBeFalseWhenGettingFromIndexOutsideLinkedList(t *testing.T) {
	for _, ll := range getImplementations() {
		ll.InsertLast(5)
		_, ok := ll.Get(1)
		assert.AssertFalse(ok, t)
	}
}

func TestShouldBeTrueWhenGettingFromIndexInsideLinkedList(t *testing.T) {
	for _, ll := range getImplementations() {
		ll.InsertLast(5)
		_, ok := ll.Get(0)
		assert.AssertTrue(ok, t)
	}
}

func TestShouldGetFromIndexInLinkedList(t *testing.T) {
	for _, ll := range getImplementations() {
		ll.InsertLast(4)
		ll.InsertLast(5)
		result, _ := ll.Get(0)
		assert.AssertEqual(result, 4, t)
		result, _ = ll.Get(1)
		assert.AssertEqual(result, 5, t)
	}
}

func TestShouldInsertIntoLinkedList(t *testing.T) {
	for _, ll := range getImplementations() {
		ll.InsertFront(5)
		_ = ll.Insert(0, 6)
		result, _ := ll.Get(0)
		assert.AssertEqual(result, 6, t)
	}
}

func TestShouldRemoveFromLinkedList(t *testing.T) {
	for _, ll := range getImplementations() {
		ll.InsertLast(5)
		ll.InsertLast(6)
		_ = ll.Remove(0)
		result, _ := ll.Get(0)
		assert.AssertEqual(result, 6, t)
	}
}

func TestShouldBeTrueWhenRemovingFromLinkedList(t *testing.T) {
	for _, ll := range getImplementations() {
		ll.InsertFront(5)
		ok := ll.Remove(0)
		assert.AssertTrue(ok, t)
	}
}

func TestShouldBeFalseWhenRemovingFromEmptyLinkedList(t *testing.T) {
	for _, ll := range getImplementations() {
		ok := ll.Remove(0)
		assert.AssertFalse(ok, t)
	}
}

func TestShouldBeTrueWhenInsertingIntoLinkedList(t *testing.T) {
	for _, ll := range getImplementations() {
		ll.InsertFront(5)
		ok := ll.Insert(0, 6)
		assert.AssertTrue(ok, t)
	}
}

func TestShouldBeFalseWhenInsertingIntoEmptyLinkedList(t *testing.T) {
	for _, ll := range getImplementations() {
		ok := ll.Insert(0, 6)
		assert.AssertFalse(ok, t)
	}
}

func TestShouldApplyMapToLinkedList(t *testing.T) {
	for _, ll := range getImplementations() {
		ll.InsertLast(1)
		ll.InsertLast(2)
		ll = ll.Map(func(i int) int { return 2 * i })
		result, _ := ll.Get(0)
		assert.AssertEqual(result, 2, t)
	}
}
