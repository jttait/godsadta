package linkedlist

import "testing"

func TestShouldBeSizeZeroForNewlyInstantiatedDoublyLinkedList(t *testing.T) {
	d := NewDoublyLinkedList[int]()
	result := d.Size()
	want := 0
	if result != want {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldBeSizeOneAfterInsertingFrontOfDoublyLinkedList(t *testing.T) {
	d := NewDoublyLinkedList[int]()
	d.InsertFront(5)
	result := d.Size()
	want := 1
	if result != want {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldBeSizeOneAfterInsertingLastOfDoublyLinkedList(t *testing.T) {
	d := NewDoublyLinkedList[int]()
	d.InsertLast(5)
	result := d.Size()
	want := 1
	if result != want {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldBeSizeZeroAfterInsertingFrontAndRemovingFrontOfDoublyLinkedList(t *testing.T) {
	d := NewDoublyLinkedList[int]()
	d.InsertFront(5)
	_, _ = d.RemoveFront()
	result := d.Size()
	want := 0
	if result != want {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldBeSizeZeroAfterInsertingLastAndRemovingLastOfDoublyLinkedList(t *testing.T) {
	d := NewDoublyLinkedList[int]()
	d.InsertLast(5)
	_, _ = d.RemoveLast()
	result := d.Size()
	want := 0
	if result != want {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldBeSizeZeroAfterInsertingFrontAndRemovingLastOfDoublyLinkedList(t *testing.T) {
	d := NewDoublyLinkedList[int]()
	d.InsertFront(5)
	_, _ = d.RemoveLast()
	result := d.Size()
	want := 0
	if result != want {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldBeSizeZeroAfterInsertingLastAndRemovingFrontOfDoublyLinkedList(t *testing.T) {
	d := NewDoublyLinkedList[int]()
	d.InsertLast(5)
	_, _ = d.RemoveFront()
	result := d.Size()
	want := 0
	if result != want {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldRemoveFrontForMultipleItemsOfDoublyLinkedList(t *testing.T) {
	d := NewDoublyLinkedList[int]()
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

func TestShouldRemoveLastForMultipleItemsOfDoublyLinkedList(t *testing.T) {
	d := NewDoublyLinkedList[int]()
	d.InsertFront(3)
	d.InsertFront(2)
	d.InsertFront(1)
	result, _ := d.RemoveLast()
	want := 3
	if result != want {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
	result, _ = d.RemoveLast()
	want = 2
	if result != want {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
	result, _ = d.RemoveLast()
	want = 1
	if result != want {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldBeTrueWhenRemoveFirstForNonEmptyDoublyLinkedList(t *testing.T) {
	d := NewDoublyLinkedList[int]()
	d.InsertFront(1)
	_, ok := d.RemoveFront()
	if !ok {
		t.Fatalf("Got false. Want true.")
	}
}

func TestShouldBeFalseWhenRemoveFirstForEmptyDoublyLinkedList(t *testing.T) {
	d := NewDoublyLinkedList[int]()
	_, ok := d.RemoveFront()
	if ok {
		t.Fatalf("Got true. Want false.")
	}
}

func TestShouldBeTrueWhenRemoveLastForNonEmptyDoublyLinkedList(t *testing.T) {
	d := NewDoublyLinkedList[int]()
	d.InsertFront(1)
	_, ok := d.RemoveLast()
	if !ok {
		t.Fatalf("Got false. Want true.")
	}
}

func TestShouldBeFalseWhenRemoveLastForEmptyDoublyLinkedList(t *testing.T) {
	d := NewDoublyLinkedList[int]()
	_, ok := d.RemoveLast()
	if ok {
		t.Fatalf("Got true. Want false.")
	}
}

func TestShouldBeFalseWhenGettingFromEmptyDoublyLinkedList(t *testing.T) {
	d := NewDoublyLinkedList[int]()
	_, ok := d.Get(0)
	if ok {
		t.Fatalf("Got true. Want false.")
	}
}

func TestShouldBeFalseWhenGettingFromIndexOutsideDoublyLinkedList(t *testing.T) {
	d := NewDoublyLinkedList[int]()
	d.InsertFront(5)
	_, ok := d.Get(1)
	if ok {
		t.Fatalf("Got true. Want false.")
	}
}

func TestShouldBeTrueWhenGettingFromIndexInsideDoublyLinkedList(t *testing.T) {
	d := NewDoublyLinkedList[int]()
	d.InsertFront(5)
	_, ok := d.Get(0)
	if !ok {
		t.Fatalf("Got false. Want true.")
	}
}

func TestShouldGetFromIndexInDoublyLinkedList(t *testing.T) {
	d := NewDoublyLinkedList[int]()
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

func TestShouldApplyMapToDoublyLinkedList(t *testing.T) {
	d := NewDoublyLinkedList[int](1, 2)
	d = d.Map(func(i int) int { return 2 * i })
	result, _ := d.Get(0)
	want := 2
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}
