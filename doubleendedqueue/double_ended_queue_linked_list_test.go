package doubleendedqueue

import "testing"

func TestShouldBeSizeZeroForNewlyInstantiatedDoubleEndedQueueLinkedList(t *testing.T) {
	want := 0
	q := NewDoubleEndedQueueLinkedList[int]()
	result := q.Size()
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldBeSizeOneWhenItemAddedToNewlyInstantiatedDoubleEndedQueueLinkedList(t *testing.T) {
	want := 1
	q := NewDoubleEndedQueueLinkedList[int]()
	q.InsertFront(5)
	result := q.Size()
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldRemoveItemFromDoubleEndedQueueLinkedListOfSizeOne(t *testing.T) {
	want := 5
	q := NewDoubleEndedQueueLinkedList[int]()
	q.InsertFront(5)
	result, _ := q.RemoveFront()
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldBeFalseWhenRemovingFromEmptyDoubleEndedQueueLinkedList(t *testing.T) {
	want := false
	q := NewDoubleEndedQueueLinkedList[int]()
	_, ok := q.RemoveFront()
	if want != ok {
		t.Fatalf("Want %v. Got %v\n", want, ok)
	}
}

func TestShouldBeTrueWhenRemovingFromNonEmptyDoubleEndedQueueLinkedList(t *testing.T) {
	q := createDoubleEndedQueueOfIntegers()
	want := true
	_, ok := q.RemoveFront()
	if want != ok {
		t.Fatalf("Want %v. Got %v\n", want, ok)
	}
}

func TestShouldRemoveFrontFromDoubleEndedQueueLinkedListOfIntergers(t *testing.T) {
	q := createDoubleEndedQueueOfIntegers()
	want := 1
	result, _ := q.RemoveFront()
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldRemoveLastFromDoubleEndedQueueLinkedListOfStrings(t *testing.T) {
	q := createDoubleEndedQueueOfStrings()
	result, _ := q.RemoveLast()
	want := "five"
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
	result, _ = q.RemoveLast()
	want = "four"
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}

}

func TestShouldRemoveFrontFromDoubleEndedQueueLinkedListOfStrings(t *testing.T) {
	q := createDoubleEndedQueueOfStrings()
	result, _ := q.RemoveFront()
	want := "one"
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
	result, _ = q.RemoveFront()
	want = "two"
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}

}

func createDoubleEndedQueueOfIntegers() *DoubleEndedQueueLinkedList[int] {
	q := NewDoubleEndedQueueLinkedList[int]()
	q.InsertFront(5)
	q.InsertFront(4)
	q.InsertFront(3)
	q.InsertFront(2)
	q.InsertFront(1)
	return q
}

func createDoubleEndedQueueOfStrings() *DoubleEndedQueueLinkedList[string] {
	q := NewDoubleEndedQueueLinkedList[string]()
	q.InsertFront("five")
	q.InsertFront("four")
	q.InsertFront("three")
	q.InsertFront("two")
	q.InsertFront("one")
	return q
}
