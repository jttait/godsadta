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
	q := createDoubleEndedQueueOfIntegers()
	want := true
	_, ok := q.RemoveFront()
	if want != ok {
		t.Fatalf("Want %v. Got %v\n", want, ok)
	}
}

func TestShouldRemoveFrontFromDoubleEndedQueueOfIntergers(t *testing.T) {
	q := createDoubleEndedQueueOfIntegers()
	want := 1
	result, _ := q.RemoveFront()
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldRemoveLastFromDoubleEndedQueueOfStrings(t *testing.T) {
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

func TestShouldRemoveFrontFromDoubleEndedQueueOfStrings(t *testing.T) {
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

func createDoubleEndedQueueOfIntegers() *DoubleEndedQueue[int] {
	q := NewDoubleEndedQueue[int]()
	q.InsertFront(5)
	q.InsertFront(4)
	q.InsertFront(3)
	q.InsertFront(2)
	q.InsertFront(1)
	return q
}

func createDoubleEndedQueueOfStrings() *DoubleEndedQueue[string] {
	q := NewDoubleEndedQueue[string]()
	q.InsertFront("five")
	q.InsertFront("four")
	q.InsertFront("three")
	q.InsertFront("two")
	q.InsertFront("one")
	return q
}
