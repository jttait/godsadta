package linkedlist

import "testing"

func TestShouldBeSizeZeroForNewlyInstantiatedSinglyLinkedList(t *testing.T) {
	d := NewSinglyLinkedList[int]()
	result := d.Size()
	want := 0
	if result != want {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldBeSizeOneAfterInsertingIntoSinglyLinkedList(t *testing.T) {
	d := NewSinglyLinkedList[int]()
	d.InsertFront(5)
	result := d.Size()
	want := 1
	if result != want {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldBeSizeZeroAfterInsertingAndRemovingFromSinglyLinkedList(t *testing.T) {
	d := NewSinglyLinkedList[int]()
	d.InsertFront(5)
	_, _ = d.RemoveFront()
	result := d.Size()
	want := 0
	if result != want {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldRemoveMultipleItemsOfSinglyLinkedList(t *testing.T) {
	d := NewSinglyLinkedList[int]()
	d.InsertFront(3)
	d.InsertFront(2)
	d.InsertFront(1)
	result, _ := d.RemoveFront()
	want := 1
	if result != want {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
	result, _ = d.RemoveFront()
	want = 2
	if result != want {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
	result, _ = d.RemoveFront()
	want = 3
	if result != want {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldPeekFromNonEmptySinglyLinkedList(t *testing.T) {
	d := NewSinglyLinkedList[int]()
	d.InsertFront(5)
	result, _ := d.PeekFront()
	want := 5
	if result != want {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldBeTrueWhenPeekingFromNonEmptySinglyLinkedList(t *testing.T) {
	d := NewSinglyLinkedList[int]()
	d.InsertFront(5)
	_, ok := d.PeekFront()
	if !ok {
		t.Fatalf("Want true. Got false.")
	}
}

func TestShouldBeFalseWhenPeekingFromNonEmptySinglyLinkedList(t *testing.T) {
	d := NewSinglyLinkedList[int]()
	_, ok := d.PeekFront()
	if ok {
		t.Fatalf("Want false. Got true.")
	}
}

func TestShouldBeTrueWhenRemovingFromNonEmptySinglyLinkedList(t *testing.T) {
	d := NewSinglyLinkedList[int]()
	d.InsertFront(1)
	_, ok := d.RemoveFront()
	if !ok {
		t.Fatalf("Got false. Want true.")
	}
}

func TestShouldBeFalseWhenRemovingFromEmptySinglyLinkedList(t *testing.T) {
	d := NewSinglyLinkedList[int]()
	_, ok := d.RemoveFront()
	if ok {
		t.Fatalf("Got true. Want false.")
	}
}

func TestShouldBeFalseWhenGettingFromEmptySinglyLinkedList(t *testing.T) {
	d := NewSinglyLinkedList[int]()
	_, ok := d.Get(0)
	if ok {
		t.Fatalf("Got true. Want false.")
	}
}

func TestShouldBeFalseWhenGettingFromIndexOutsideSinglyLinkedList(t *testing.T) {
	d := NewSinglyLinkedList[int]()
	d.InsertFront(5)
	_, ok := d.Get(1)
	if ok {
		t.Fatalf("Got true. Want false.")
	}
}

func TestShouldBeTrueWhenGettingFromIndexInsideSinglyLinkedList(t *testing.T) {
	d := NewSinglyLinkedList[int]()
	d.InsertFront(5)
	_, ok := d.Get(0)
	if !ok {
		t.Fatalf("Got false. Want true.")
	}
}

func TestShouldGetFromIndexInSinglyLinkedList(t *testing.T) {
	d := NewSinglyLinkedList[int]()
	d.InsertFront(5)
	d.InsertFront(4)
	result, _ := d.Get(0)
	want := 4
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
	result, _ = d.Get(1)
	want = 5
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldApplyMapToSinglyLinkedList(t *testing.T) {
	d := NewSinglyLinkedList[int](1, 2)
	d.Map(func(i int) int { return 2 * i })
	result, _ := d.Get(0)
	want := 2
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}
