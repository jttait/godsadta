package minheap

import "testing"

func TestShouldInstantiateMinHeapArray(t *testing.T) {
	m := NewMinHeapArray[int]()
	if m == nil {
		t.Fatalf("Should not be nil.")
	}
}

func TestShouldInsertItemToMinHeapArray(t *testing.T) {
	m := NewMinHeapArray[int]()
	m.Insert(5)
}

func TestShouldExtractFromMinHeapArrayWithOneItem(t *testing.T) {
	m := NewMinHeapArray[int](5)
	result := m.Extract()
	want := 5
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldExtractFromMinHeapArrayWithTwoItems(t *testing.T) {
	m := NewMinHeapArray[int](5, 15, 10)
	result := m.Extract()
	want := 5
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldExtractMultipleTimesFromMinHeapArray(t *testing.T) {
	m := NewMinHeapArray[int](78, 39, 24, 98, 72, 77, 71, 68, 53, 41, 60, 54, 92, 36, 75, 47, 33, 80, 9, 61)
	result := m.Extract()
	want := 9
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
	result = m.Extract()
	want = 24
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
	result = m.Extract()
	want = 33
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldPeekTopItemFromMinHeapArray(t *testing.T) {
	m := NewMinHeapArray[int](5, 15, 10)
	result := m.Peek()
	want := 5
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldBeSameSizeAfterPeekingMinHeapArray(t *testing.T) {
	m := NewMinHeapArray[int](5, 15, 10)
	_ = m.Peek()
	result := m.Size()
	want := 3
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldBeSizeZeroForNewlyInstantiatedMinHeapArray(t *testing.T) {
	m := NewMinHeapArray[int]()
	result := m.Size()
	want := 0
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldBeSizeTwoForMinHeapArrayWithTwoItems(t *testing.T) {
	m := NewMinHeapArray[int](1, 2)
	result := m.Size()
	want := 2
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}
