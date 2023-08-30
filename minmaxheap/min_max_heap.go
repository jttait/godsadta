package minmaxheap

import "cmp"

type MinMaxHeap[T cmp.Ordered] struct {
	array []T
}

func NewMinMaxHeap[T cmp.Ordered]() *MinMaxHeap[T] {
	m := MinMaxHeap[T]{}
	return &m
}

func swap[T cmp.Ordered](i int, j int, array *[]T) {
	temp := (*array)[i]
	(*array)[i] = (*array)[j]
	(*array)[j] = temp
}

func findChildIndices(index int) (int, int) {
	return (index * 2) + 1, (index * 2) + 2
}

func findLeftChildIndex(index int) int {
	return (index * 2) + 1
}

func findRightChildIndex(index int) int {
	return (index * 2) + 2
}

func findParentIndex(index int) int {
	return (index - 1) / 2
}

func findGrandparentIndex(index int) int {
	return findParentIndex(findParentIndex(index))
}

func hasChildren[T cmp.Ordered](index int, array []T) bool {
	return findLeftChildIndex(index) < len(array)
}

func findIndexOfSmallestChildOrGrandchild[T cmp.Ordered](index int, array []T) (int, bool) {
	isGrandchild := false
	leftChildIndex := findLeftChildIndex(index)
	rightChildIndex := findRightChildIndex(index)
	leftLeftChildIndex := findLeftChildIndex(leftChildIndex)
	leftRightChildIndex := findRightChildIndex(leftChildIndex)
	rightLeftChildIndex := findLeftChildIndex(rightChildIndex)
	rightRightChildIndex := findRightChildIndex(rightChildIndex)
	smallestValue := array[leftChildIndex]
	smallestIndex := leftChildIndex
	if rightChildIndex < len(array) && array[rightChildIndex] < smallestValue {
		smallestIndex = rightChildIndex
		smallestValue = array[smallestIndex]
	}
	if leftLeftChildIndex < len(array) && array[leftLeftChildIndex] < smallestValue {
		smallestIndex = leftLeftChildIndex
		smallestValue = array[smallestIndex]
		isGrandchild = true
	}
	if leftRightChildIndex < len(array) && array[leftRightChildIndex] < smallestValue {
		smallestIndex = leftRightChildIndex
		smallestValue = array[smallestIndex]
		isGrandchild = true
	}
	if rightLeftChildIndex < len(array) && array[rightLeftChildIndex] < smallestValue {
		smallestIndex = rightLeftChildIndex
		smallestValue = array[smallestIndex]
		isGrandchild = true
	}
	if rightRightChildIndex < len(array) && array[rightRightChildIndex] < smallestValue {
		smallestIndex = rightRightChildIndex
		smallestValue = array[smallestIndex]
		isGrandchild = true
	}
	return smallestIndex, isGrandchild
}

func pushDown[T cmp.Ordered](index int, array []T) {
	if isEvenLevel(index) {
		pushDownMin(index)
	} else {
		pushDownMax(index)
	}
}

func pushDownMin(index int, array []T) {
	for hasChildren(index) {
		m, isGrandchild := findIndexOfSmallestChildOrGrandchild(index)
		if isGrandchild {
			if array[m] < array[index] {
				swap(m, index, array)
				if array[m] > array[findParentIndex(m)] {
					swap(m, findParentIndex(m), array)
				}
				pushDown(m)
			}
		} else {
			swap(m, index, array)
		}
	}
}
