package set

import (
	"testing"

	"github.com/jttait/godsa/assert"
)

func getSetImplementations(values1 []int, values2 []int) [][]Set[int] {
	return [][]Set[int]{
		[]Set[int]{NewMapSet[int](values1...), NewMapSet[int](values2...)},
	}
}

func TestShouldBeSizeZeroForNewlyInstantiatedSet(t *testing.T) {
	for _, s := range getSetImplementations([]int{}, []int{}) {
		result := s[0].Size()
		assert.AssertEqual(result, 0, t)
	}
}

func TestShouldBeSizeOneAfterItemIsAddedToNewlyInstantiatedSet(t *testing.T) {
	for _, s := range getSetImplementations([]int{}, []int{}) {
		_ = s[0].Insert(5)
		result := s[0].Size()
		assert.AssertEqual(result, 1, t)
	}
}

func TestShouldBeSizeTwoAfterAddingTwoDifferentItems(t *testing.T) {
	for _, s := range getSetImplementations([]int{}, []int{}) {
		_ = s[0].Insert(5)
		_ = s[0].Insert(6)
		result := s[0].Size()
		assert.AssertEqual(result, 2, t)
	}
}

func TestShouldBeSizeOneAfterAddingTwoSameItems(t *testing.T) {
	for _, s := range getSetImplementations([]int{}, []int{}) {
		_ = s[0].Insert(5)
		_ = s[0].Insert(5)
		result := s[0].Size()
		assert.AssertEqual(result, 1, t)
	}
}

func TestShouldBeSizeZeroAfterAddingAndRemovingItemFromSet(t *testing.T) {
	for _, s := range getSetImplementations([]int{}, []int{}) {
		_ = s[0].Insert(5)
		_ = s[0].Remove(5)
		result := s[0].Size()
		assert.AssertEqual(result, 0, t)
	}
}

func TestShouldBeTrueIfAddingItemThatIsNotAlreadyInSet(t *testing.T) {
	for _, s := range getSetImplementations([]int{}, []int{}) {
		ok := s[0].Insert(5)
		assert.AssertTrue(ok, t)
	}
}

func TestShouldBeFalseIfAddingItemThatIsAlreadyInSet(t *testing.T) {
	for _, s := range getSetImplementations([]int{5}, []int{}) {
		ok := s[0].Insert(5)
		assert.AssertFalse(ok, t)
	}
}

func TestShouldBeFalseIfRemovingItemThatIsNotInSet(t *testing.T) {
	for _, s := range getSetImplementations([]int{}, []int{}) {
		ok := s[0].Remove(5)
		assert.AssertFalse(ok, t)
	}
}

func TestShouldBeTrueIfRemovingItemThatIsInSet(t *testing.T) {
	for _, s := range getSetImplementations([]int{5}, []int{}) {
		ok := s[0].Remove(5)
		assert.AssertTrue(ok, t)
	}
}

func TestShouldBeTrueIfSetContainsItem(t *testing.T) {
	for _, s := range getSetImplementations([]int{5}, []int{}) {
		result := s[0].Contains(5)
		assert.AssertTrue(result, t)
	}
}

func TestShouldBeFalseIfSetDoesNotContainItem(t *testing.T) {
	for _, s := range getSetImplementations([]int{}, []int{}) {
		result := s[0].Contains(5)
		assert.AssertFalse(result, t)
	}
}

func TestShouldBeEqualForTwoEmptySets(t *testing.T) {
	for _, s := range getSetImplementations([]int{}, []int{}) {
		result := s[0].Equals(s[1])
		assert.AssertTrue(result, t)
	}
}

func TestShouldBeEqualForTwoIdenticalSets(t *testing.T) {
	for _, s := range getSetImplementations([]int{5}, []int{5}) {
		result := s[0].Equals(s[1])
		assert.AssertTrue(result, t)
	}
}

func TestShouldBeNotEqualForTwoDifferentSets(t *testing.T) {
	for _, s := range getSetImplementations([]int{5}, []int{6}) {
		result := s[0].Equals(s[1])
		assert.AssertFalse(result, t)
	}
}

func TestShouldContainAllItemsFromBothSetsAfterUnion(t *testing.T) {
	for _, s := range getSetImplementations([]int{1, 2}, []int{3, 4}) {
		result := s[0].Union(s[1])
		want := NewMapSet[int](1, 2, 3, 4)
		if !result.Equals(want) {
			t.Fatalf("Union does not contain all items.")
		}
	}
}

func TestShouldContainOnlyItemsThatAreInBothSetsAfterIntersection(t *testing.T) {
	for _, s := range getSetImplementations([]int{1, 2, 3, 4, 5}, []int{4, 5, 6, 7, 8}) {
		result := s[0].Intersection(s[1])
		want := NewMapSet[int](4, 5)
		if !result.Equals(want) {
			t.Fatalf("Intersection does not only contain items in both sets.")
		}
	}
}

func TestShouldContainOnlyItemsThatAreInFirstSetButNotSecondAfterDifference(t *testing.T) {
	for _, s := range getSetImplementations([]int{1, 2, 3, 4, 5}, []int{4, 5, 6, 7, 8}) {
		result := s[0].Difference(s[1])
		want := NewMapSet[int](1, 2, 3)
		if !result.Equals(want) {
			t.Fatalf("Difference does not only contain items in first set but not second.")
		}
	}
}
