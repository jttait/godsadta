package set

import (
	"testing"
)

func TestShouldBeSizeZeroForNewlyInstantiatedSet(t *testing.T) {
	s := NewSet[int]()
	result := s.Size()
	want := 0
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldBeSizeOneAfterItemIsAddedToNewlyInstantiatedSet(t *testing.T) {
	s := NewSet[int]()
	_ = s.Add(5)
	result := s.Size()
	want := 1
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldBeSizeTwoAfterAddingTwoDifferentItems(t *testing.T) {
	s := NewSet[int]()
	_ = s.Add(5)
	_ = s.Add(6)
	result := s.Size()
	want := 2
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldBeSizeOneAfterAddingTwoSameItems(t *testing.T) {
	s := NewSet[int]()
	_ = s.Add(5)
	_ = s.Add(5)
	result := s.Size()
	want := 1
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldBeSizeZeroAfterAddingAndRemovingItemFromSet(t *testing.T) {
	s := NewSet[int]()
	_ = s.Add(5)
	_ = s.Remove(5)
	result := s.Size()
	want := 0
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldBeTrueIfAddingItemThatIsNotAlreadyInSet(t *testing.T) {
	s := NewSet[int]()
	ok := s.Add(5)
	want := true
	if want != ok {
		t.Fatalf("Want %v. Got %v\n", want, ok)
	}
}

func TestShouldBeFalseIfAddingItemThatIsAlreadyInSet(t *testing.T) {
	s := NewSet[int](5)
	ok := s.Add(5)
	want := false
	if want != ok {
		t.Fatalf("Want %v. Got %v\n", want, ok)
	}
}

func TestShouldBeFalseIfRemovingItemThatIsNotInSet(t *testing.T) {
	s := NewSet[int]()
	ok := s.Remove(5)
	want := false
	if want != ok {
		t.Fatalf("Want %v. Got %v\n", want, ok)
	}
}

func TestShouldBeTrueIfRemovingItemThatIsInSet(t *testing.T) {
	s := NewSet[int](5)
	ok := s.Remove(5)
	want := true
	if want != ok {
		t.Fatalf("Want %v. Got %v\n", want, ok)
	}
}

func TestShouldBeTrueIfSetContainsItem(t *testing.T) {
	s := NewSet[int](5)
	result := s.Contains(5)
	want := true
	if want != result {
		t.Fatalf("Want %v. Got %v.\n", result, want)
	}
}

func TestShouldBeFalseIfSetDoesNotContainItem(t *testing.T) {
	s := NewSet[int]()
	result := s.Contains(5)
	want := false
	if want != result {
		t.Fatalf("Want %v. Got %v.\n", result, want)
	}
}

func TestShouldBeEqualForTwoEmptySets(t *testing.T) {
	s := NewSet[int]()
	u := NewSet[int]()
	result := s.Equals(u)
	want := true
	if want != result {
		t.Fatalf("Want %v. Got %v.\n", result, want)
	}
}

func TestShouldBeEqualForTwoIdenticalSets(t *testing.T) {
	s := NewSet[int](5)
	u := NewSet[int](5)
	result := s.Equals(u)
	want := true
	if want != result {
		t.Fatalf("Want %v. Got %v.\n", result, want)
	}
}

func TestShouldBeNotEqualForTwoDifferentSets(t *testing.T) {
	s := NewSet[int](5)
	u := NewSet[int](6)
	result := s.Equals(u)
	want := false
	if want != result {
		t.Fatalf("Want %v. Got %v.\n", result, want)
	}
}

func TestShouldContainAllItemsFromBothSetsAfterUnion(t *testing.T) {
	s := NewSet[int](1, 2)
	u := NewSet[int](3, 4)
	result := s.Union(u)
	want := NewSet[int](1, 2, 3, 4)
	if !result.Equals(want) {
		t.Fatalf("Union does not contain all items.")
	}
}

func TestShouldContainOnlyItemsThatAreInBothSetsAfterIntersection(t *testing.T) {
	s := NewSet[int](1, 2, 3, 4, 5)
	u := NewSet[int](4, 5, 6, 7, 8)
	result := s.Intersection(u)
	want := NewSet[int](4, 5)
	if !result.Equals(want) {
		t.Fatalf("Intersection does not only contain items in both sets.")
	}
}

func TestShouldContainOnlyItemsThatAreInFirstSetButNotSecondAfterDifference(t *testing.T) {
	s := NewSet[int](1, 2, 3, 4, 5)
	u := NewSet[int](4, 5, 6, 7, 8)
	result := s.Difference(u)
	want := NewSet[int](1, 2, 3)
	if !result.Equals(want) {
		t.Fatalf("Difference does not only contain items in first set but not second.")
	}

}
