package binarytree

import "testing"

func TestShouldInstantiateRootWithNilChildNodes(t *testing.T) {
	b := NewBinaryTree[int]()
	if b.Root.Left != nil || b.Root.Right != nil {
		t.Fatalf("Root left and right children are not nil.")
	}
}
