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

func TestShouldBeSizeZeroAfterAddingAndRemovingItemFromSet(t *testing.T) {
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

func TestShouldBeTrueIfSetContainsItem(t *testing.T) {
	s := NewSet()
	_ = s.Add(5)
	result := s.Contains(5)
	want := true
	if want != result {
		t.Fatalf("Want %v. Got %v.\n", result, want)
	}
}

func TestShouldBeFalseIfSetDoesNotContainItem(t *testing.T) {
	s := NewSet()
	result := s.Contains(5)
	want := false
	if want != result {
		t.Fatalf("Want %v. Got %v.\n", result, want)
	}
}

func TestShouldBeEqualForTwoEmptySets(t *testing.T) {
	s := NewSet()
	u := NewSet()
	result := s.Equals(u)
	want := true
	if want != result {
		t.Fatalf("Want %v. Got %v.\n", result, want)
	}
}

func TestShouldBeEqualForTwoIdenticalSets(t *testing.T) {
	s := NewSet()
	u := NewSet()
	_ = s.Add(5)
	_ = u.Add(5)
	result := s.Equals(u)
	want := true
	if want != result {
		t.Fatalf("Want %v. Got %v.\n", result, want)
	}
}

func TestShouldBeNotEqualForTwoDifferentSets(t *testing.T) {
	s := NewSet()
	u := NewSet()
	_ = s.Add(5)
	_ = u.Add(6)
	result := s.Equals(u)
	want := false
	if want != result {
		t.Fatalf("Want %v. Got %v.\n", result, want)
	}
}
