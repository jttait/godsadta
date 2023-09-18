package multiset

import (
	"testing"

	"github.com/jttait/godsa/assert"
)

func TestShouldAddSameItemsToMapMultiset(t *testing.T) {
	m := NewMapMultiset[int]()
	m.Add(5)
	m.Add(5)
	m.Add(5)
	result := m.Size()
	assert.AssertEqual(result, 3, t)
}

func TestShouldAddDifferentItemsToMapMultiset(t *testing.T) {
	m := NewMapMultiset[int]()
	m.Add(5)
	m.Add(6)
	m.Add(7)
	result := m.Size()
	assert.AssertEqual(result, 3, t)
}

func TestShouldRemoveSameItemsToMapMultiset(t *testing.T) {
	m := NewMapMultiset[int]()
	m.Add(5)
	m.Add(5)
	m.Add(5)
	m.Remove(5)
	m.Remove(5)
	result := m.Size()
	assert.AssertEqual(result, 1, t)
}

func TestShouldRemoveDifferentItemsToMapMultiset(t *testing.T) {
	m := NewMapMultiset[int]()
	m.Add(5)
	m.Add(6)
	m.Add(7)
	m.Remove(6)
	m.Remove(7)
	result := m.Size()
	assert.AssertEqual(result, 1, t)
}

func TestShouldCountSameItemsToMapMultiset(t *testing.T) {
	m := NewMapMultiset[int]()
	m.Add(5)
	m.Add(5)
	m.Add(5)
	result := m.Count(5)
	assert.AssertEqual(result, 3, t)
}

func TestShouldCountDifferentItemsToMapMultiset(t *testing.T) {
	m := NewMapMultiset[int]()
	m.Add(5)
	m.Add(6)
	m.Add(7)
	result := m.Count(6)
	assert.AssertEqual(result, 1, t)
}

func TestShouldCountZeroForItemThatIsNotInMapMultiset(t *testing.T) {
	m := NewMapMultiset[int]()
	result := m.Count(6)
	assert.AssertEqual(result, 0, t)
}
