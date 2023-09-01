package queue

import "testing"

func TestShouldBeSizeZeroForNewlyInstantiatedQueueLinkedList(t *testing.T) {
	want := 0
	q := NewQueueLinkedList[int]()
	result := q.Size()
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldBeSizeOneWhenItemAddedToNewlyInstantiatedQueueLinkedList(t *testing.T) {
	want := 1
	q := NewQueueLinkedList[int]()
	q.Insert(5)
	result := q.Size()
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldRemoveItemFromQueueLinkedListOfSizeOne(t *testing.T) {
	want := 5
	q := NewQueueLinkedList[int]()
	q.Insert(5)
	result, _ := q.Remove()
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldBeFalseWhenRemovingFromEmptyQueueLinkedList(t *testing.T) {
	want := false
	q := NewQueueLinkedList[int]()
	_, ok := q.Remove()
	if want != ok {
		t.Fatalf("Want %v. Got %v\n", want, ok)
	}
}

func TestShouldBeTrueWhenRemovingFromNonEmptyQueueLinkedList(t *testing.T) {
	want := true
	q := NewQueueLinkedList[int]()
	q.Insert(5)
	_, ok := q.Remove()
	if want != ok {
		t.Fatalf("Want %v. Got %v\n", want, ok)
	}
}

func TestShouldRemoveIntegerThatWasAddedEarliestFromQueueLinkedList(t *testing.T) {
	want := 5
	q := NewQueueLinkedList[int]()
	q.Insert(5)
	q.Insert(6)
	result, _ := q.Remove()
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldRemoveStringThatWasAddedEarliestFromQueueLinkedList(t *testing.T) {
	want := "five"
	q := NewQueueLinkedList[string]()
	q.Insert("five")
	q.Insert("six")
	result, _ := q.Remove()
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldRemoveMultipleStringsFromQueueLinkedList(t *testing.T) {
	q := NewQueueLinkedList[string]()
	q.Insert("five")
	q.Insert("six")
	q.Insert("seven")
	result, _ := q.Remove()
	want := "five"
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
	result, _ = q.Remove()
	want = "six"
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
	result, _ = q.Remove()
	want = "seven"
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
	_, ok := q.Remove()
	if ok {
		t.Fatalf("Want false. Got true.")
	}
}
