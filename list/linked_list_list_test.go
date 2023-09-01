package list

import "testing"

func TestShouldInstantiateLinkedListListOfSizeZero(t *testing.T) {
	l := NewLinkedListList[int]()
	result := l.Size()
	want := 0
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldAppendMultipleItemsToLinkedListList(t *testing.T) {
	l := NewLinkedListList[int]()
	l.Append(5)
	l.Append(6)
	l.Append(7)
	result, _ := l.Get(0)
	want := 5
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
	result, _ = l.Get(1)
	want = 6
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
	result, _ = l.Get(2)
	want = 7
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldIncreaseSizeWhenAppendingToLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	l.Append(5)
	result := l.Size()
	want := 1
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
	l.Append(6)
	result = l.Size()
	want = 2
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
	l.Append(7)
	result = l.Size()
	want = 3
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldPrependMultipleItemsToLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	l.Prepend(5)
	l.Prepend(6)
	l.Prepend(7)
	result, _ := l.Get(0)
	want := 7
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
	result, _ = l.Get(1)
	want = 6
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
	result, _ = l.Get(2)
	want = 5
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldIncreaseSizeWhenPrependingToLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	l.Prepend(5)
	result := l.Size()
	want := 1
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
	l.Prepend(6)
	result = l.Size()
	want = 2
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
	l.Prepend(7)
	result = l.Size()
	want = 3
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldBeFalseIfGetOnEmptyLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	_, ok := l.Get(0)
	if ok {
		t.Fatalf("Got true. Want false.")
	}
}

func TestShouldBeTrueIfGetIndexIsWithinLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	l.Append(5)
	_, ok := l.Get(0)
	if !ok {
		t.Fatalf("Got false. Want true.")
	}
}

func TestShouldBeFalseIfGetIndexIsOutsideLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	l.Append(5)
	_, ok := l.Get(1)
	if ok {
		t.Fatalf("Got true. Want false.")
	}
}

func TestShouldRemoveItemAtIndexOfLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	_ = l.Remove(1)
	result, _ := l.Get(1)
	want := 3
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
	_ = l.Remove(1)
	_, ok := l.Get(1)
	if ok {
		t.Fatalf("Got true. Want false.")
	}
}

func TestShouldBeFalseWhenRemovingFromEmptyLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	ok := l.Remove(0)
	if ok {
		t.Fatalf("Got true. Want false.")
	}
}

func TestShouldBeFalseWhenRemovingFromIndexOutsideLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	l.Append(5)
	ok := l.Remove(1)
	if ok {
		t.Fatalf("Got true. Want false.")
	}
}

func TestShouldBeTrueWhenRemovingFromIndexInsideLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	l.Append(5)
	ok := l.Remove(0)
	if !ok {
		t.Fatalf("Got false. Want true.")
	}
}

func TestShouldInsertItemAtIndexOfLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	l.Append(1)
	_ = l.Insert(0, 2)
	result, _ := l.Get(0)
	want := 2
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldShiftTheItemCurrentlyAtTheIndexToRightWithLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	l.Append(1)
	_ = l.Insert(0, 2)
	result, _ := l.Get(1)
	want := 1
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldBeFalseWhenInsertingIntoEmptyLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	ok := l.Insert(0, 1)
	if ok {
		t.Fatalf("Got true. Want false.")
	}
}

func TestShouldBeTrueWhenInsertingIntoLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	l.Append(1)
	ok := l.Insert(0, 2)
	if !ok {
		t.Fatalf("Got false. Want true.")
	}
}

func TestShouldBeFalseWhenInsertingIntoIndexOutsideLinkedListList(t *testing.T) {
	l := NewArrayList[int]()
	l.Append(1)
	ok := l.Insert(1, 2)
	if ok {
		t.Fatalf("Got true. Want false.")
	}
}
