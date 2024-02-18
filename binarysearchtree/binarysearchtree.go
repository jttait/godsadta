// Package binarysearchtree provides interface and implementations for binary search tree data
// structure
package binarysearchtree

import "cmp"

type BinarySearchTree[T cmp.Ordered] interface {
	Insert(T) bool
	Remove(T) bool
	Contains(T) bool
}
