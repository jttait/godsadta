package queue

import "testing"

func TestShouldBeSizeZeroForNewlyInstantiatedQueueLinkedList(t *testing.T) {
	q := NewQueueLinkedList[int]()
	result := q.Size()
	assertEqual(result, 0, t)
}

func TestShouldBeSizeOneWhenItemAddedToNewlyInstantiatedQueueLinkedList(t *testing.T) {
	q := NewQueueLinkedList[int]()
	q.Insert(5)
	result := q.Size()
	assertEqual(result, 1, t)
}

func TestShouldRemoveItemFromQueueLinkedListOfSizeOne(t *testing.T) {
	q := NewQueueLinkedList[int]()
	q.Insert(5)
	result, _ := q.Remove()
	assertEqual(result, 5, t)
}

func TestShouldBeFalseWhenRemovingFromEmptyQueueLinkedList(t *testing.T) {
	q := NewQueueLinkedList[int]()
	_, ok := q.Remove()
	assertFalse(ok, t)
}

func TestShouldBeTrueWhenRemovingFromNonEmptyQueueLinkedList(t *testing.T) {
	q := NewQueueLinkedList[int]()
	q.Insert(5)
	_, ok := q.Remove()
	assertTrue(ok, t)
}

func TestShouldRemoveIntegerThatWasAddedEarliestFromQueueLinkedList(t *testing.T) {
	q := NewQueueLinkedList[int]()
	q.Insert(5)
	q.Insert(6)
	result, _ := q.Remove()
	assertEqual(result, 5, t)
}

func TestShouldRemoveStringThatWasAddedEarliestFromQueueLinkedList(t *testing.T) {
	q := NewQueueLinkedList[string]()
	q.Insert("five")
	q.Insert("six")
	result, _ := q.Remove()
	assertEqual(result, "five", t)
}

func TestShouldRemoveMultipleStringsFromQueueLinkedList(t *testing.T) {
	q := NewQueueLinkedList[string]()
	q.Insert("five")
	q.Insert("six")
	q.Insert("seven")
	result, _ := q.Remove()
	assertEqual(result, "five", t)
	result, _ = q.Remove()
	assertEqual(result, "six", t)
	result, _ = q.Remove()
	assertEqual(result, "seven", t)
	_, ok := q.Remove()
	assertFalse(ok, t)
}

func assertTrue(i bool, t *testing.T) {
	if !i {
		t.Fatalf("Got false. Want true.")
	}
}

func assertFalse(i bool, t *testing.T) {
	if i {
		t.Fatalf("Got true. Want false.")
	}
}

func assertEqual[T comparable](got T, want T, t *testing.T) {
	if got != want {
		t.Fatalf("Got %v. Want %v.\n", got, want)
	}
}
