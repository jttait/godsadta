package godsa

import (
	"testing"

	"github.com/jttait/godsa/assert"
)

func TestShouldBeSizeZeroForNewlyInstantiatedSet(t *testing.T) {
	s := NewMapSet[int]()
	result := s.Size()
	assert.AssertEqual(result, 0, t)
}

func TestShouldBeSizeOneAfterItemIsAddedToNewlyInstantiatedSet(t *testing.T) {
	s := NewMapSet[int]()
	_ = s.Insert(5)
	result := s.Size()
	assert.AssertEqual(result, 1, t)
}

func TestShouldBeSizeTwoAfterAddingTwoDifferentItems(t *testing.T) {
	s := NewMapSet[int]()
	_ = s.Insert(5)
	_ = s.Insert(6)
	result := s.Size()
	assert.AssertEqual(result, 2, t)
}

func TestShouldBeSizeOneAfterAddingTwoSameItems(t *testing.T) {
	s := NewMapSet[int]()
	_ = s.Insert(5)
	_ = s.Insert(5)
	result := s.Size()
	assert.AssertEqual(result, 1, t)
}

func TestShouldBeSizeZeroAfterAddingAndRemovingItemFromSet(t *testing.T) {
	s := NewMapSet[int]()
	_ = s.Insert(5)
	_ = s.Remove(5)
	result := s.Size()
	assert.AssertEqual(result, 0, t)
}

func TestShouldBeTrueIfAddingItemThatIsNotAlreadyInSet(t *testing.T) {
	s := NewMapSet[int]()
	ok := s.Insert(5)
	assert.AssertTrue(ok, t)
}

func TestShouldBeFalseIfAddingItemThatIsAlreadyInSet(t *testing.T) {
	s := NewMapSet[int](5)
	ok := s.Insert(5)
	assert.AssertFalse(ok, t)
}

func TestShouldBeFalseIfRemovingItemThatIsNotInSet(t *testing.T) {
	s := NewMapSet[int]()
	ok := s.Remove(5)
	assert.AssertFalse(ok, t)
}

func TestShouldBeTrueIfRemovingItemThatIsInSet(t *testing.T) {
	s := NewMapSet[int](5)
	ok := s.Remove(5)
	assert.AssertTrue(ok, t)
}

func TestShouldBeTrueIfSetContainsItem(t *testing.T) {
	s := NewMapSet[int](5)
	result := s.Contains(5)
	assert.AssertTrue(result, t)
}

func TestShouldBeFalseIfSetDoesNotContainItem(t *testing.T) {
	s := NewMapSet[int]()
	result := s.Contains(5)
	assert.AssertFalse(result, t)
}

func TestShouldBeEqualForTwoEmptySets(t *testing.T) {
	s := NewMapSet[int]()
	u := NewMapSet[int]()
	result := s.Equals(u)
	assert.AssertTrue(result, t)
}

func TestShouldBeEqualForTwoIdenticalSets(t *testing.T) {
	s := NewMapSet[int](5)
	u := NewMapSet[int](5)
	result := s.Equals(u)
	assert.AssertTrue(result, t)
}

func TestShouldBeNotEqualForTwoDifferentSets(t *testing.T) {
	s := NewMapSet[int](5)
	u := NewMapSet[int](6)
	result := s.Equals(u)
	assert.AssertFalse(result, t)
}

func TestShouldContainAllItemsFromBothSetsAfterUnion(t *testing.T) {
	s := NewMapSet[int](1, 2)
	u := NewMapSet[int](3, 4)
	result := s.Union(u)
	want := NewMapSet[int](1, 2, 3, 4)
	if !result.Equals(want) {
		t.Fatalf("Union does not contain all items.")
	}
}

func TestShouldContainOnlyItemsThatAreInBothSetsAfterIntersection(t *testing.T) {
	s := NewMapSet[int](1, 2, 3, 4, 5)
	u := NewMapSet[int](4, 5, 6, 7, 8)
	result := s.Intersection(u)
	want := NewMapSet[int](4, 5)
	if !result.Equals(want) {
		t.Fatalf("Intersection does not only contain items in both sets.")
	}
}

func TestShouldContainOnlyItemsThatAreInFirstSetButNotSecondAfterDifference(t *testing.T) {
	s := NewMapSet[int](1, 2, 3, 4, 5)
	u := NewMapSet[int](4, 5, 6, 7, 8)
	result := s.Difference(u)
	want := NewMapSet[int](1, 2, 3)
	if !result.Equals(want) {
		t.Fatalf("Difference does not only contain items in first set but not second.")
	}
}
