package binarysearchtree

import (
	"testing"

	"github.com/jttait/godsa/assert"
)

func TestShouldInstantiateBinarySearchTreeWithNilRootNode(t *testing.T) {
	b := NewBinarySearchTree[int]()
	if b.Root != nil {
		t.Fatalf("Root node is not nil.")
	}
}

func TestShouldInsertIntoEmptyBinarySearchTree(t *testing.T) {
	b := NewBinarySearchTree[int]()
	_ = b.Insert(5)
	result := b.Root.Val
	assert.AssertEqual(result, 5, t)
}

func TestShouldInsertIntoNonEmptyBinarySearchTree(t *testing.T) {
	b := NewBinarySearchTree[int]()
	_ = b.Insert(5)
	_ = b.Insert(10)
	result := b.Root.Right.Val
	assert.AssertEqual(result, 10, t)
}

func TestShouldTrueWhenInsertingItemThatIsNotAlreadyInBinarySearchTree(t *testing.T) {
	b := NewBinarySearchTree[int]()
	_ = b.Insert(5)
	ok := b.Insert(10)
	assert.AssertTrue(ok, t)
}

func TestShouldFalseWhenInsertingItemThatIsAlreadyInBinarySearchTree(t *testing.T) {
	b := NewBinarySearchTree[int]()
	_ = b.Insert(5)
	ok := b.Insert(5)
	assert.AssertFalse(ok, t)
}

func TestShouldRemoveItemFromBinarySearchTreeOfOneNode(t *testing.T) {
	b := NewBinarySearchTree[int]()
	_ = b.Insert(5)
	_ = b.Remove(5)
	if b.Root != nil {
		t.Fatalf("Root node should be nil.")
	}
}

func TestShouldRemoveItemFromBinarySearchTreeOfThreeNodes(t *testing.T) {
	b := NewBinarySearchTree[int]()
	_ = b.Insert(2)
	_ = b.Insert(1)
	_ = b.Insert(3)
	_ = b.Remove(3)
	result := b.Root.Right
	if result != nil {
		t.Fatalf("Want nil. Got %v.\n", result)
	}
}

func TestShouldBeTrueWhenInsertingIntoEmptyBinarySearchTree(t *testing.T) {
	b := NewBinarySearchTree[int]()
	ok := b.Insert(5)
	assert.AssertTrue(ok, t)
}

func TestShouldBeFalseWhenInsertingItemThatAlreadyExistsInBinarySearchTree(t *testing.T) {
	b := NewBinarySearchTree[int]()
	_ = b.Insert(5)
	ok := b.Insert(5)
	assert.AssertFalse(ok, t)
}

func TestShouldBeFalseForContainsOnEmptyBinarySearchTree(t *testing.T) {
	b := NewBinarySearchTree[int]()
	result := b.Contains(5)
	assert.AssertFalse(result, t)
}

func TestShouldBeTrueForContainsOnBinarySearchTreeThatContainsItem(t *testing.T) {
	b := NewBinarySearchTree[int]()
	_ = b.Insert(5)
	result := b.Contains(5)
	assert.AssertTrue(result, t)
}
