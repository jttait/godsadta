package godsa

import (
	"testing"

	"github.com/jttait/godsa/assert"
	"github.com/jttait/godsa/mapmultiset"
)

func getMultisetImplementations() []Multiset[int] {
	return []Multiset[int]{
		mapmultiset.NewMapMultiset[int](),
	}
}

func TestShouldAddSameItemsToMultiset(t *testing.T) {
	for _, m := range getMultisetImplementations() {
		m.Add(5)
		m.Add(5)
		m.Add(5)
		result := m.Size()
		assert.AssertEqual(result, 3, t)
	}
}

func TestShouldAddDifferentItemsToMultiset(t *testing.T) {
	for _, m := range getMultisetImplementations() {
		m.Add(5)
		m.Add(6)
		m.Add(7)
		result := m.Size()
		assert.AssertEqual(result, 3, t)
	}
}

func TestShouldRemoveSameItemsToMultiset(t *testing.T) {
	for _, m := range getMultisetImplementations() {
		m.Add(5)
		m.Add(5)
		m.Add(5)
		m.Remove(5)
		m.Remove(5)
		result := m.Size()
		assert.AssertEqual(result, 1, t)
	}
}

func TestShouldRemoveDifferentItemsToMultiset(t *testing.T) {
	for _, m := range getMultisetImplementations() {
		m.Add(5)
		m.Add(6)
		m.Add(7)
		m.Remove(6)
		m.Remove(7)
		result := m.Size()
		assert.AssertEqual(result, 1, t)
	}
}

func TestShouldCountSameItemsToMultiset(t *testing.T) {
	for _, m := range getMultisetImplementations() {
		m.Add(5)
		m.Add(5)
		m.Add(5)
		result := m.Count(5)
		assert.AssertEqual(result, 3, t)
	}
}

func TestShouldCountDifferentItemsToMultiset(t *testing.T) {
	for _, m := range getMultisetImplementations() {
		m.Add(5)
		m.Add(6)
		m.Add(7)
		result := m.Count(6)
		assert.AssertEqual(result, 1, t)
	}
}

func TestShouldCountZeroForItemThatIsNotInMultiset(t *testing.T) {
	for _, m := range getMultisetImplementations() {
		result := m.Count(6)
		assert.AssertEqual(result, 0, t)
	}
}
