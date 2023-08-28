package maxheap

import "testing"

func TestShouldInstantiateMaxHeap(t *testing.T) {
	m := NewMaxHeap[int]()
	if m == nil {
		t.Fatalf("Should not be nil.")
	}
}

func TestShouldInsertItemToMaxHeap(t *testing.T) {
	m := NewMaxHeap[int]()
	m.Insert(5)
}

func TestShouldExtractFromMaxHeapWithOneItem(t *testing.T) {
	m := NewMaxHeap[int]()
	m.Insert(5)
	result := m.Extract()
	want := 5
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldExtractFromMaxHeapWithTwoItems(t *testing.T) {
	m := NewMaxHeap[int](5, 15, 10)
	result := m.Extract()
	want := 15
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldExtractMultipleTimesFromMaxHeap(t *testing.T) {
	m := NewMaxHeap[int](78, 39, 24, 98, 72, 77, 71, 68, 53, 41, 60, 54, 92, 36, 75, 47, 33, 80, 9, 61)
	result := m.Extract()
	want := 98
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
	result = m.Extract()
	want = 92
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
	result = m.Extract()
	want = 80
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}
