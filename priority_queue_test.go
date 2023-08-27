package godsa

import "testing"

func TestShouldBeSizeZeroForNewlyInstantiatedPriorityQueue(t *testing.T) {
	q := NewPriorityQueue[int]()
	want := 0
	result := q.Size()
	if want != result {
		t.Fatalf("Got %v. Want %v\n", result, want)
	}
}

func TestShouldBeSizeOneWhenItemAddedToNewlyInstantiatedPriorityQueue(t *testing.T) {
	q := NewPriorityQueue[int]()
	q.Add(5)
	want := 1
	result := q.Size()
	if want != result {
		t.Fatalf("Got %v. Want %v\n", result, want)
	}
}

func TestShouldBeSizeTwoWhenAddingTwoItems(t *testing.T) {
	q := NewPriorityQueue[int]()
	q.Add(5)
	q.Add(6)
	want := 2
	result := q.Size()
	if want != result {
		t.Fatalf("Got %v. Want %v\n", result, want)
	}
}

func TestShouldBeSizeTwoWhenAddingTwoIdenticalItems(t *testing.T) {
	q := NewPriorityQueue[int]()
	q.Add(5)
	q.Add(5)
	want := 2
	result := q.Size()
	if want != result {
		t.Fatalf("Got %v. Want %v\n", result, want)
	}
}

func TestShouldBeSizeZeroAfterAddingAndRemovingItem(t *testing.T) {
	q := NewPriorityQueue[int]()
	q.Add(5)
	_ = q.Remove(5)
	want := 0
	result := q.Size()
	if want != result {
		t.Fatalf("Got %v. Want %v\n", result, want)
	}
}

func TestShouldBeTrueWhenRemovingItemThatisInPriorityQueue(t *testing.T) {
	q := NewPriorityQueue[int]()
	q.Add(5)
	ok := q.Remove(5)
	want := true
	if ok != want {
		t.Fatalf("Got %v. Want %v.\n", ok, want)
	}
}

func TestShouldBeFalseWhenRemovingItemThatIsNotInPriorityQueue(t *testing.T) {
	q := NewPriorityQueue[int]()
	ok := q.Remove(5)
	want := false
	if ok != want {
		t.Fatalf("Got %v. Want %v.\n", ok, want)
	}
}

func TestShouldBeHighestPriorityItemWhenPollingPriorityQueue(t *testing.T) {
	q := NewPriorityQueue[int]()
	q.Add(5)
	q.Add(3)
	result, _ := q.Poll()
	want := 5
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldBeTrueWhenPollingNonEmptyPriorityQueue(t *testing.T) {
	q := NewPriorityQueue[int]()
	q.Add(5)
	_, ok := q.Poll()
	want := true
	if want != ok {
		t.Fatalf("Got %v. Want %v.\n", ok, want)
	}
}

func TestShouldBeFalseWhenPollingEmptyPriorityQueue(t *testing.T) {
	q := NewPriorityQueue[int]()
	_, ok := q.Poll()
	want := false
	if want != ok {
		t.Fatalf("Got %v. Want %v.\n", ok, want)
	}
}

func TestShouldBeHighestPriorityItemWhenPeekingPriorityQueue(t *testing.T) {
	q := NewPriorityQueue[int]()
	q.Add(5)
	q.Add(3)
	result, _ := q.Peek()
	want := 5
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldBeTrueWhenPeekingNonEmptyPriorityQueue(t *testing.T) {
	q := NewPriorityQueue[int]()
	q.Add(5)
	_, ok := q.Peek()
	want := true
	if want != ok {
		t.Fatalf("Got %v. Want %v.\n", ok, want)
	}
}

func TestShouldBeFalseWhenPeekingEmptyPriorityQueue(t *testing.T) {
	q := NewPriorityQueue[int]()
	_, ok := q.Peek()
	want := false
	if want != ok {
		t.Fatalf("Got %v. Want %v.\n", ok, want)
	}
}

func TestShouldBeSameSizeAfterPeekingNonEmptyPriorityQueue(t *testing.T) {
	q := NewPriorityQueue[int]()
	q.Add(5)
	_, _ = q.Peek()
	want := 1
	result := q.Size()
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldBeSizeZeroAfterPollingPriorityQueueOfSizeOne(t *testing.T) {
	q := NewPriorityQueue[int]()
	q.Add(5)
	_, _ = q.Poll()
	want := 0
	result := q.Size()
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}
