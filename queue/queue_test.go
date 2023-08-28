package queue

import "testing"

func TestShouldBeSizeZeroForNewlyInstantiatedQueue(t *testing.T) {
	want := 0
	q := NewQueue[int]()
	result := q.Size()
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldBeSizeOneWhenItemAddedToNewlyInstantiatedQueue(t *testing.T) {
	want := 1
	q := NewQueue[int]()
	q.Add(5)
	result := q.Size()
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldRemoveItemFromQueueOfSizeOne(t *testing.T) {
	want := 5
	q := NewQueue[int]()
	q.Add(5)
	result, _ := q.Remove()
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldBeFalseWhenRemovingFromEmptyQueue(t *testing.T) {
	want := false
	q := NewQueue[int]()
	_, ok := q.Remove()
	if want != ok {
		t.Fatalf("Want %v. Got %v\n", want, ok)
	}
}

func TestShouldBeTrueWhenRemovingFromNonEmptyQueue(t *testing.T) {
	want := true
	q := NewQueue[int]()
	q.Add(5)
	_, ok := q.Remove()
	if want != ok {
		t.Fatalf("Want %v. Got %v\n", want, ok)
	}
}

func TestShouldRemoveIntegerThatWasAddedEarliest(t *testing.T) {
	want := 5
	q := NewQueue[int]()
	q.Add(5)
	q.Add(6)
	result, _ := q.Remove()
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldRemoveStringThatWasAddedEarliest(t *testing.T) {
	want := "five"
	q := NewQueue[string]()
	q.Add("five")
	q.Add("six")
	result, _ := q.Remove()
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}
