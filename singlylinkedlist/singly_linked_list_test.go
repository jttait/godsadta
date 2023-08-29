package singlylinkedlist

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
	d.Insert(5)
	result := d.Size()
	want := 1
	if result != want {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldBeSizeZeroAfterInsertingAndRemovingFromSinglyLinkedList(t *testing.T) {
	d := NewSinglyLinkedList[int]()
	d.Insert(5)
	_, _ = d.Remove()
	result := d.Size()
	want := 0
	if result != want {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldRemoveMultipleItemsOfSinglyLinkedList(t *testing.T) {
	d := NewSinglyLinkedList[int]()
	d.Insert(3)
	d.Insert(2)
	d.Insert(1)
	result, _ := d.Remove()
	want := 1
	if result != want {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
	result, _ = d.Remove()
	want = 2
	if result != want {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
	result, _ = d.Remove()
	want = 3
	if result != want {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldPeekFromNonEmptySinglyLinkedList(t *testing.T) {
	d := NewSinglyLinkedList[int]()
	d.Insert(5)
	result, _ := d.Peek()
	want := 5
	if result != want {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldBeTrueWhenPeekingFromNonEmptySinglyLinkedList(t *testing.T) {
	d := NewSinglyLinkedList[int]()
	d.Insert(5)
	_, ok := d.Peek()
	if !ok {
		t.Fatalf("Want true. Got false.")
	}
}

func TestShouldBeFalseWhenPeekingFromNonEmptySinglyLinkedList(t *testing.T) {
	d := NewSinglyLinkedList[int]()
	_, ok := d.Peek()
	if ok {
		t.Fatalf("Want false. Got true.")
	}
}

func TestShouldBeTrueWhenRemovingFromNonEmptySinglyLinkedList(t *testing.T) {
	d := NewSinglyLinkedList[int]()
	d.Insert(1)
	_, ok := d.Remove()
	if !ok {
		t.Fatalf("Got false. Want true.")
	}
}

func TestShouldBeFalseWhenRemovingFromEmptySinglyLinkedList(t *testing.T) {
	d := NewSinglyLinkedList[int]()
	_, ok := d.Remove()
	if ok {
		t.Fatalf("Got true. Want false.")
	}
}
