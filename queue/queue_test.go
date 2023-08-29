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
	q.Insert(5)
	result := q.Size()
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldRemoveItemFromQueueOfSizeOne(t *testing.T) {
	want := 5
	q := NewQueue[int]()
	q.Insert(5)
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
	q.Insert(5)
	_, ok := q.Remove()
	if want != ok {
		t.Fatalf("Want %v. Got %v\n", want, ok)
	}
}

func TestShouldRemoveIntegerThatWasAddedEarliest(t *testing.T) {
	want := 5
	q := NewQueue[int]()
	q.Insert(5)
	q.Insert(6)
	result, _ := q.Remove()
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldRemoveStringThatWasAddedEarliest(t *testing.T) {
	want := "five"
	q := NewQueue[string]()
	q.Insert("five")
	q.Insert("six")
	result, _ := q.Remove()
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldRemoveMultipleStringsFromQueue(t *testing.T) {
	q := NewQueue[string]()
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
