package priorityqueue

import "testing"

func TestShouldBeSizeZeroForNewlyInstantiatedPriorityQueue(t *testing.T) {
	q := NewPriorityQueue[int]()
	result := q.Size()
	assertEqual(result, 0, t)
}

func TestShouldBeSizeOneWhenItemAddedToNewlyInstantiatedPriorityQueue(t *testing.T) {
	q := NewPriorityQueue[int]()
	q.Add(5)
	result := q.Size()
	assertEqual(result, 1, t)
}

func TestShouldBeSizeTwoWhenAddingTwoItems(t *testing.T) {
	q := NewPriorityQueue[int]()
	q.Add(5)
	q.Add(6)
	result := q.Size()
	assertEqual(result, 2, t)
}

func TestShouldBeSizeTwoWhenAddingTwoIdenticalItems(t *testing.T) {
	q := NewPriorityQueue[int]()
	q.Add(5)
	q.Add(5)
	result := q.Size()
	assertEqual(result, 2, t)
}

func TestShouldBeHighestPriorityItemWhenPollingPriorityQueue(t *testing.T) {
	q := NewPriorityQueue[int]()
	q.Add(5)
	q.Add(3)
	result, _ := q.Poll()
	assertEqual(result, 5, t)
}

func TestShouldBeTrueWhenPollingNonEmptyPriorityQueue(t *testing.T) {
	q := NewPriorityQueue[int]()
	q.Add(5)
	_, ok := q.Poll()
	assertTrue(ok, t)
}

func TestShouldBeFalseWhenPollingEmptyPriorityQueue(t *testing.T) {
	q := NewPriorityQueue[int]()
	_, ok := q.Poll()
	assertFalse(ok, t)
}

func TestShouldBeHighestPriorityItemWhenPeekingPriorityQueue(t *testing.T) {
	q := NewPriorityQueue[int]()
	q.Add(5)
	q.Add(3)
	result, _ := q.Peek()
	assertEqual(result, 5, t)
}

func TestShouldBeTrueWhenPeekingNonEmptyPriorityQueue(t *testing.T) {
	q := NewPriorityQueue[int]()
	q.Add(5)
	_, ok := q.Peek()
	assertTrue(ok, t)
}

func TestShouldBeFalseWhenPeekingEmptyPriorityQueue(t *testing.T) {
	q := NewPriorityQueue[int]()
	_, ok := q.Peek()
	assertFalse(ok, t)
}

func TestShouldBeSameSizeAfterPeekingNonEmptyPriorityQueue(t *testing.T) {
	q := NewPriorityQueue[int]()
	q.Add(5)
	_, _ = q.Peek()
	result := q.Size()
	assertEqual(result, 1, t)
}

func TestShouldBeSizeZeroAfterPollingPriorityQueueOfSizeOne(t *testing.T) {
	q := NewPriorityQueue[int]()
	q.Add(5)
	_, _ = q.Poll()
	result := q.Size()
	assertEqual(result, 0, t)
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
