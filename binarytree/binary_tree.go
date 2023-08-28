// Package binary tree provides the binary tree data structure and associated methods
package binarytree

// BinaryTree contains a pointer to the root node.
type BinaryTree[T any] struct {
	Root *BinaryTreeNode[T]
}

// BinaryTreeNode contains a value, a pointer to the left child node, and a pointer to the right
// child node.
type BinaryTreeNode[T any] struct {
	Val   T
	Left  *BinaryTreeNode[T]
	Right *BinaryTreeNode[T]
}

// NewBinaryTree instantiates a binary tree and returns a pointer to it.
func NewBinaryTree[T any]() *BinaryTree[T] {
	b := BinaryTree[T]{}
	b.Root = &BinaryTreeNode[T]{}
	return &b
}
