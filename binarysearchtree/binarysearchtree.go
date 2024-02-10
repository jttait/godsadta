// Package binarysearchtree provides interface and implementations for binary search tree data
// structure
package binarysearchtree

import "cmp"

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

// BinarySearchTree contains a pointer to the root node.
type BinarySearchTree[T cmp.Ordered] struct {
	Root *BinaryTreeNode[T]
}

// NewBinarySearchTree instantiates a binary search tree and returns a pointer to it.
func NewBinarySearchTree[T cmp.Ordered](values ...T) *BinarySearchTree[T] {
	b := BinarySearchTree[T]{}
	for _, v := range values {
		_ = b.Insert(v)
	}
	return &b
}

// Insert inserts the given item into the binary search tree. It returns a Boolean that is true if
// the item was not already in the binary search tree.
func (b *BinarySearchTree[T]) Insert(i T) bool {
	var parent *BinaryTreeNode[T]
	parent = nil
	current := b.Root
	for current != nil {
		parent = current
		if i == current.Val {
			return false
		} else if i < current.Val {
			current = current.Left
		} else {
			current = current.Right
		}
	}
	if parent == nil {
		b.Root = NewBinaryTreeNode(i)
	} else if i < parent.Val {
		parent.Left = NewBinaryTreeNode(i)
	} else {
		parent.Right = NewBinaryTreeNode(i)
	}
	return true
}

// Remove removes the given item from the binary search tree. It returns a Boolean that is true if
// the item was in the binary search tree.
func (b *BinarySearchTree[T]) Remove(i T) bool {
	var parent *BinaryTreeNode[T]
	parent = nil
	current := b.Root
	for current != nil {
		if i == current.Val {
			if parent == nil {
				b.Root = nil
			} else if current.Left == nil && current.Right == nil {
				if i < parent.Val {
					parent.Left = nil
				} else {
					parent.Right = nil
				}
			} else if (current.Left == nil && current.Right != nil) || (current.Left != nil && current.Right == nil) {
				if current.Left != nil {
					current.Val = current.Left.Val
					current.Left = nil
				} else {
					current.Val = current.Right.Val
					current.Right = nil
				}
			} else {
				if i < parent.Val {
					parent.Left = current
				} else {
					parent.Right = current
				}
			}
			return true
		} else if i < current.Val {
			parent = current
			current = current.Left
		} else {
			parent = current
			current = current.Right
		}
	}
	return false
}

// Contains returns true if the given item is in the binary search tree.
func (b *BinarySearchTree[T]) Contains(i T) bool {
	current := b.Root
	for current != nil {
		if i == current.Val {
			return true
		} else if i < current.Val {
			current = current.Left
		} else {
			current = current.Right
		}
	}
	return false
}
