// Package binarysearchtree provides interface and implementations for binary search tree data
// structure
package binarysearchtree

import "cmp"

type BinarySearchTree[T cmp.Ordered] interface {
	Insert(T) bool
	Remove(T) bool
	Contains(T) bool
}

// BinaryTreeNode contains a value, a pointer to the left child node, and a pointer to the right
// child node.
type BinaryTreeNode[T cmp.Ordered] struct {
	Val   T
	Left  *BinaryTreeNode[T]
	Right *BinaryTreeNode[T]
}

func NewBinaryTreeNode[T cmp.Ordered](i T) *BinaryTreeNode[T] {
	b := BinaryTreeNode[T]{}
	b.Val = i
	return &b
}

