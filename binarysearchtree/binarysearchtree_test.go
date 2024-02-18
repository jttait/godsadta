package binarysearchtree

import (
	"testing"

	"github.com/jttait/godsadta/assert"
)

func getBinarySearchTreeImplementations() []BinarySearchTree[int] {
	return []BinarySearchTree[int]{
		NewUnbalancedBinarySearchTree[int](),
	}
}

func TestShouldInstantiateBinarySearchTreeWithNilRootNode(t *testing.T) {
	for _, b := range getBinarySearchTreeImplementations() {
		if b == nil {
			t.Fatalf("Binary search tree should not be nil.")
		}
	}
}

func TestShouldInsertIntoEmptyBinarySearchTree(t *testing.T) {
	for _, b := range getBinarySearchTreeImplementations() {
		_ = b.Insert(5)
		assert.AssertTrue(b.Contains(5), t)
	}
}

func TestShouldInsertIntoNonEmptyBinarySearchTree(t *testing.T) {
	for _, b := range getBinarySearchTreeImplementations() {
		_ = b.Insert(5)
		_ = b.Insert(10)
		assert.AssertTrue(b.Contains(10), t)
	}
}

func TestShouldTrueWhenInsertingItemThatIsNotAlreadyInBinarySearchTree(t *testing.T) {
	for _, b := range getBinarySearchTreeImplementations() {
		_ = b.Insert(5)
		ok := b.Insert(10)
		assert.AssertTrue(ok, t)
	}
}

func TestShouldFalseWhenInsertingItemThatIsAlreadyInBinarySearchTree(t *testing.T) {
	for _, b := range getBinarySearchTreeImplementations() {
		_ = b.Insert(5)
		ok := b.Insert(5)
		assert.AssertFalse(ok, t)
	}
}

func TestShouldRemoveItemFromBinarySearchTreeOfOneNode(t *testing.T) {
	for _, b := range getBinarySearchTreeImplementations() {
		_ = b.Insert(5)
		_ = b.Remove(5)
		assert.AssertFalse(b.Contains(5), t)
	}
}

func TestShouldRemoveItemFromBinarySearchTreeOfThreeNodes(t *testing.T) {
	for _, b := range getBinarySearchTreeImplementations() {
		_ = b.Insert(2)
		_ = b.Insert(1)
		_ = b.Insert(3)
		_ = b.Remove(3)
		assert.AssertFalse(b.Contains(3), t)
	}
}

func TestShouldBeTrueWhenInsertingIntoEmptyBinarySearchTree(t *testing.T) {
	for _, b := range getBinarySearchTreeImplementations() {
		ok := b.Insert(5)
		assert.AssertTrue(ok, t)
	}
}

func TestShouldBeFalseWhenInsertingItemThatAlreadyExistsInBinarySearchTree(t *testing.T) {
	for _, b := range getBinarySearchTreeImplementations() {
		_ = b.Insert(5)
		ok := b.Insert(5)
		assert.AssertFalse(ok, t)
	}
}

func TestShouldBeFalseForContainsOnEmptyBinarySearchTree(t *testing.T) {
	for _, b := range getBinarySearchTreeImplementations() {
		assert.AssertFalse(b.Contains(5), t)
	}
}

func TestShouldBeTrueForContainsOnBinarySearchTreeThatContainsItem(t *testing.T) {
	for _, b := range getBinarySearchTreeImplementations() {
		_ = b.Insert(5)
		assert.AssertTrue(b.Contains(5), t)
	}
}
