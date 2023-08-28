package doubleendedqueue

import "testing"

func TestShouldBeSizeZeroForNewlyInstantiatedDoubleEndedQueue(t *testing.T) {
	want := 0
	q := NewDoubleEndedQueue[int]()
	result := q.Size()
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldBeSizeOneWhenItemAddedToNewlyInstantiatedDoubleEndedQueue(t *testing.T) {
	want := 1
	q := NewDoubleEndedQueue[int]()
	q.InsertFront(5)
	result := q.Size()
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldRemoveItemFromDoubleEndedQueueOfSizeOne(t *testing.T) {
	want := 5
	q := NewDoubleEndedQueue[int]()
	q.InsertFront(5)
	result, _ := q.RemoveFront()
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldBeFalseWhenRemovingFromEmptyDoubleEndedQueue(t *testing.T) {
	want := false
	q := NewDoubleEndedQueue[int]()
	_, ok := q.RemoveFront()
	if want != ok {
		t.Fatalf("Want %v. Got %v\n", want, ok)
	}
}

func TestShouldBeTrueWhenRemovingFromNonEmptyDoubleEndedQueue(t *testing.T) {
	want := true
	q := NewDoubleEndedQueue[int]()
	q.InsertFront(5)
	_, ok := q.RemoveFront()
	if want != ok {
		t.Fatalf("Want %v. Got %v\n", want, ok)
	}
}

func TestShouldRemoveIntegerThatWasAddedEarliest(t *testing.T) {
	want := 5
	q := NewDoubleEndedQueue[int]()
	q.InsertFront(5)
	q.InsertFront(6)
	result, _ := q.RemoveFront()
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldRemoveStringThatWasAddedEarliest(t *testing.T) {
	want := "five"
	q := NewDoubleEndedQueue[string]()
	q.InsertFront("five")
	q.InsertFront("six")
	result, _ := q.RemoveFront()
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldRemoveStringThatWasAddedLatestWhenRemovingLast(t *testing.T) {
	q := NewDoubleEndedQueue[string]()
	q.InsertFront("five")
	q.InsertFront("six")
	result, _ := q.RemoveLast()
	want := "six"
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldRemoveStringThatWasAddedLatestToEndWhenRemovingFirst(t *testing.T) {
	q := NewDoubleEndedQueue[string]()
	q.InsertLast("six")
	q.InsertLast("five")
	result, _ := q.RemoveFirst()
	want := "six"
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}
