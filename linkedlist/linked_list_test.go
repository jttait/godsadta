package linkedlist

import (
	"testing"
)

func getImplementations() []LinkedList[int] {
	return []LinkedList[int]{
		NewSinglyLinkedList[int](),
		NewDoublyLinkedList[int](),
	}
}

func TestShouldBeSizeZeroForNewlyInstantiatedLinkedList(t *testing.T) {
	for _, ll := range getImplementations() {
		result := ll.Size()
		want := 0
		if result != want {
			t.Fatalf("Got %v. Want %v.\n", result, want)
		}
	}
}

func TestShouldBeSizeOneAfterInsertingIntoLinkedList(t *testing.T) {
	for _, ll := range getImplementations() {
		ll.InsertFront(5)
		result := ll.Size()
		want := 1
		if result != want {
			t.Fatalf("Got %v. Want %v.\n", result, want)
		}
	}
}

func TestShouldBeSizeZeroAfterInsertingAndRemovingFromLinkedList(t *testing.T) {
	for _, ll := range getImplementations() {
		ll.InsertFront(5)
		_, _ = ll.RemoveFront()
		result := ll.Size()
		want := 0
		if result != want {
			t.Fatalf("Got %v. Want %v.\n", result, want)
		}
	}
}

func TestShouldRemoveMultipleItemsOfLinkedList(t *testing.T) {
	for _, ll := range getImplementations() {
		ll.InsertLast(1)
		ll.InsertLast(2)
		ll.InsertLast(3)
		result, _ := ll.RemoveFront()
		want := 1
		if result != want {
			t.Fatalf("Got %v. Want %v.\n", result, want)
		}
		result, _ = ll.RemoveFront()
		want = 2
		if result != want {
			t.Fatalf("Got %v. Want %v.\n", result, want)
		}
		result, _ = ll.RemoveFront()
		want = 3
		if result != want {
			t.Fatalf("Got %v. Want %v.\n", result, want)
		}
	}
}

func TestShouldPeekFromNonEmptyLinkedList(t *testing.T) {
	for _, ll := range getImplementations() {
		ll.InsertLast(5)
		result, _ := ll.PeekFront()
		want := 5
		if result != want {
			t.Fatalf("Got %v. Want %v.\n", result, want)
		}
	}
}

func TestShouldPeekLastFromNonEmptyLinkedList(t *testing.T) {
	for _, ll := range getImplementations() {
		ll.InsertLast(5)
		result, _ := ll.PeekLast()
		want := 5
		if result != want {
			t.Fatalf("Got %v. Want %v.\n", result, want)
		}
	}
}

func TestShouldBeTrueWhenPeekingLastFromNonEmptyLinkedList(t *testing.T) {
	for _, ll := range getImplementations() {
		ll.InsertLast(5)
		_, ok := ll.PeekFront()
		if !ok {
			t.Fatalf("Want true. Got false.")
		}
	}
}

func TestShouldBeFalseWhenPeekingLastFromEmptyLinkedList(t *testing.T) {
	for _, ll := range getImplementations() {
		_, ok := ll.PeekLast()
		if ok {
			t.Fatalf("Want false. Got true. LL type: %T", ll)
		}
	}
}

func TestShouldBeTrueWhenPeekingFromNonEmptyLinkedList(t *testing.T) {
	for _, ll := range getImplementations() {
		ll.InsertLast(5)
		_, ok := ll.PeekFront()
		if !ok {
			t.Fatalf("Want true. Got false.")
		}
	}
}

func TestShouldBeFalseWhenPeekingFromNonEmptyLinkedList(t *testing.T) {
	for _, ll := range getImplementations() {
		_, ok := ll.PeekFront()
		if ok {
			t.Fatalf("Want false. Got true. LL type: %T", ll)
		}
	}
}

func TestShouldBeTrueWhenRemovingFromNonEmptyLinkedList(t *testing.T) {
	for _, ll := range getImplementations() {
		ll.InsertLast(5)
		_, ok := ll.RemoveFront()
		if !ok {
			t.Fatalf("Got false. Want true.")
		}
	}
}

func TestShouldBeFalseWhenRemovingFromEmptyLinkedList(t *testing.T) {
	for _, ll := range getImplementations() {
		_, ok := ll.RemoveFront()
		if ok {
			t.Fatalf("Got true. Want false.")
		}
	}
}

func TestShouldBeFalseWhenGettingFromEmptyLinkedList(t *testing.T) {
	for _, ll := range getImplementations() {
		_, ok := ll.Get(0)
		if ok {
			t.Fatalf("Got true. Want false.")
		}
	}
}

func TestShouldBeFalseWhenGettingFromIndexOutsideLinkedList(t *testing.T) {
	for _, ll := range getImplementations() {
		ll.InsertLast(5)
		_, ok := ll.Get(1)
		if ok {
			t.Fatalf("Got true. Want false.")
		}
	}
}

func TestShouldBeTrueWhenGettingFromIndexInsideLinkedList(t *testing.T) {
	for _, ll := range getImplementations() {
		ll.InsertLast(5)
		_, ok := ll.Get(0)
		if !ok {
			t.Fatalf("Got false. Want true.")
		}
	}
}

func TestShouldGetFromIndexInLinkedList(t *testing.T) {
	for _, ll := range getImplementations() {
		ll.InsertLast(4)
		ll.InsertLast(5)
		result, _ := ll.Get(0)
		want := 4
		if want != result {
			t.Fatalf("Got %v. Want %v.\n", result, want)
		}
		result, _ = ll.Get(1)
		want = 5
		if want != result {
			t.Fatalf("Got %v. Want %v.\n", result, want)
		}
	}
}

func TestShouldApplyMapToLinkedList(t *testing.T) {
	for _, ll := range getImplementations() {
		ll.InsertLast(1)
		ll.InsertLast(2)
		ll = ll.Map(func(i int) int { return 2 * i })
		result, _ := ll.Get(0)
		want := 2
		if want != result {
			t.Fatalf("Got %v. Want %v.\n", result, want)
		}
	}
}
