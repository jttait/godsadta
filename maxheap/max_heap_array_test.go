package maxheap

import "testing"

func TestShouldInstantiateMaxHeap(t *testing.T) {
	m := NewMaxHeapArray[int]()
	if m == nil {
		t.Fatalf("Should not be nil.")
	}
}

func TestShouldInsertItemToMaxHeap(t *testing.T) {
	m := NewMaxHeapArray[int]()
	m.Insert(5)
}

func TestShouldExtractFromMaxHeapWithOneItem(t *testing.T) {
	m := NewMaxHeapArray[int]()
	m.Insert(5)
	result, _ := m.Extract()
	want := 5
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldExtractFromMaxHeapWithTwoItems(t *testing.T) {
	m := NewMaxHeapArray[int](5, 15, 10)
	result, _ := m.Extract()
	want := 15
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldExtractMultipleTimesFromMaxHeap(t *testing.T) {
	m := NewMaxHeapArray[int](78, 39, 24, 98, 72, 77, 71, 68, 53, 41, 60, 54, 92, 36, 75, 47, 33, 80, 9, 61)
	result, _ := m.Extract()
	want := 98
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
	result, _ = m.Extract()
	want = 92
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
	result, _ = m.Extract()
	want = 80
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldPeekTopItemFromMaxHeap(t *testing.T) {
	m := NewMaxHeapArray[int](5, 15, 10)
	result, _ := m.Peek()
	want := 15
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldBeSameSizeAfterPeekingMaxHeap(t *testing.T) {
	m := NewMaxHeapArray[int](5, 15, 10)
	_, _ = m.Peek()
	result := m.Size()
	want := 3
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldBeSizeZeroForNewlyInstantiatedMaxHeap(t *testing.T) {
	m := NewMaxHeapArray[int]()
	result := m.Size()
	want := 0
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldBeSizeTwoForMaxHeapWithTwoItems(t *testing.T) {
	m := NewMaxHeapArray[int](1, 2)
	result := m.Size()
	want := 2
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}
