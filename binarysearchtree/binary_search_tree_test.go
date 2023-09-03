package binarysearchtree

import "testing"

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
	want := 5
	if result != want {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldInsertIntoNonEmptyBinarySearchTree(t *testing.T) {
	b := NewBinarySearchTree[int]()
	_ = b.Insert(5)
	_ = b.Insert(10)
	result := b.Root.Right.Val
	want := 10
	if result != want {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldTrueWhenInsertingItemThatIsNotAlreadyInBinarySearchTree(t *testing.T) {
	b := NewBinarySearchTree[int]()
	_ = b.Insert(5)
	ok := b.Insert(10)
	if !ok {
		t.Fatalf("Got false. Want true.\n")
	}
}

func TestShouldFalseWhenInsertingItemThatIsAlreadyInBinarySearchTree(t *testing.T) {
	b := NewBinarySearchTree[int]()
	_ = b.Insert(5)
	ok := b.Insert(5)
	if ok {
		t.Fatalf("Got true. Want false.\n")
	}
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
	if !ok {
		t.Fatalf("Want true. Got false.")
	}
}

func TestShouldBeFalseWhenInsertingItemThatAlreadyExistsInBinarySearchTree(t *testing.T) {
	b := NewBinarySearchTree[int]()
	_ = b.Insert(5)
	ok := b.Insert(5)
	if ok {
		t.Fatalf("Want false. Got true.")
	}
}

func TestShouldBeFalseForContainsOnEmptyBinarySearchTree(t *testing.T) {
	b := NewBinarySearchTree[int]()
	result := b.Contains(5)
	if result {
		t.Fatalf("Want false. Got true.")
	}
}

func TestShouldBeTrueForContainsOnBinarySearchTreeThatContainsItem(t *testing.T) {
	b := NewBinarySearchTree[int]()
	_ = b.Insert(5)
	result := b.Contains(5)
	if !result {
		t.Fatalf("Want true. Got false.")
	}
}
