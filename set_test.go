package godsa

import "testing"

func TestShouldBeSizeZeroForNewlyInstantiatedSet(t *testing.T) {
	s := NewSet()
	result := s.Size()
	want := 0
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldBeSizeOneAfterItemIsAddedToNewlyInstantiatedSet(t *testing.T) {
	s := NewSet()
	_ = s.Add(5)
	result := s.Size()
	want := 1
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldBeSizeTwoAfterAddingTwoDifferentItems(t *testing.T) {
	s := NewSet()
	_ = s.Add(5)
	_ = s.Add(6)
	result := s.Size()
	want := 2
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldBeSizeOneAfterAddingTwoSameItems(t *testing.T) {
	s := NewSet()
	_ = s.Add(5)
	_ = s.Add(5)
	result := s.Size()
	want := 1
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldBeSizeZeroAfterAddingAndRemovingItem(t *testing.T) {
	s := NewSet()
	_ = s.Add(5)
	_ = s.Remove(5)
	result := s.Size()
	want := 0
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldBeTrueIfAddingItemThatIsNotAlreadyInSet(t *testing.T) {
	s := NewSet()
	ok := s.Add(5)
	want := true
	if want != ok {
		t.Fatalf("Want %v. Got %v\n", want, ok)
	}
}

func TestShouldBeFalseIfAddingItemThatIsAlreadyInSet(t *testing.T) {
	s := NewSet()
	_ = s.Add(5)
	ok := s.Add(5)
	want := false
	if want != ok {
		t.Fatalf("Want %v. Got %v\n", want, ok)
	}
}

func TestShouldBeFalseIfRemovingItemThatIsNotInSet(t *testing.T) {
	s := NewSet()
	ok := s.Remove(5)
	want := false
	if want != ok {
		t.Fatalf("Want %v. Got %v\n", want, ok)
	}
}

func TestShouldBeTrueIfRemovingItemThatIsInSet(t *testing.T) {
	s := NewSet()
	_ = s.Add(5)
	ok := s.Remove(5)
	want := true
	if want != ok {
		t.Fatalf("Want %v. Got %v\n", want, ok)
	}
}
