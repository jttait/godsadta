// Package arrayminheap provides an implementation of minheap
package arrayminheap

import (
	"cmp"
)

type ArrayMinHeap[T cmp.Ordered] struct {
	array []T
}

// NewMinHeap instantiates a min heap and returns a pointer to it. The variadic argument can be used
// to insert items into the heap after it is created.
func NewArrayMinHeap[T cmp.Ordered](i ...T) *ArrayMinHeap[T] {
	m := ArrayMinHeap[T]{}
	for _, j := range i {
		m.Insert(j)
	}
	return &m
}

func findParentIndex(index int) int {
	return (index - 1) / 2
}

// Insert inserts the item into the min heap.
func (m *ArrayMinHeap[T]) Insert(i T) {
	m.array = append(m.array, i)
	index := len(m.array) - 1
	parentIndex := findParentIndex(index)
	for parentIndex >= 0 && m.array[index] < m.array[parentIndex] {
		swap(index, parentIndex, &m.array)
		index = parentIndex
		parentIndex = findParentIndex(index)
	}
}

// Extract returns the minimum item from the min heap and swaps the elements as necessary to
// restore the heap property in the data structure.
func (m *ArrayMinHeap[T]) Extract() T {
	result := m.array[0]
	m.array[0] = m.array[len(m.array)-1]
	m.array = m.array[:len(m.array)-1]

	index := 0
	childIndex1, childIndex2 := findChildIndices(index)
	for index < len(m.array) && ((childIndex1 < len(m.array) && m.array[index] > m.array[childIndex1]) || (childIndex2 < len(m.array) && m.array[index] > m.array[childIndex2])) {
		if m.array[childIndex1] < m.array[childIndex2] {
			swap(index, childIndex1, &m.array)
		} else {
			swap(index, childIndex2, &m.array)
		}
		index = childIndex1
		childIndex1, childIndex2 = findChildIndices(index)
	}
	return result
}

func swap[T cmp.Ordered](i int, j int, array *[]T) {
	temp := (*array)[i]
	(*array)[i] = (*array)[j]
	(*array)[j] = temp
}

func findChildIndices(index int) (int, int) {
	return (index * 2) + 1, (index * 2) + 2
}

func (m *ArrayMinHeap[T]) Size() int {
	return len(m.array)
}

func (m *ArrayMinHeap[T]) Peek() T {
	return m.array[0]
}
