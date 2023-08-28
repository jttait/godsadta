// Package maxheap provides the max heap data structure and associated methods
package maxheap

import (
	"cmp"
)

type MaxHeap[T cmp.Ordered] struct {
	array []T
}

// NewMaxHeap instantiates a max heap and returns a pointer to it. The variadic argument can be used
// to insert items into the heap after it is created.
func NewMaxHeap[T cmp.Ordered](i ...T) *MaxHeap[T] {
	m := MaxHeap[T]{}
	for _, j := range i {
		m.Insert(j)
	}
	return &m
}

func findParentIndex(index int) int {
	if index%2 == 0 {
		return (index - 2) / 2
	}
	return (index - 1) / 2
}

// Insert inserts the item into the max heap.
func (m *MaxHeap[T]) Insert(i T) {
	m.array = append(m.array, i)
	index := len(m.array) - 1
	parentIndex := findParentIndex(index)
	for parentIndex >= 0 && m.array[index] > m.array[parentIndex] {
		temp := m.array[index]
		m.array[index] = m.array[parentIndex]
		m.array[parentIndex] = temp
		index = parentIndex
		parentIndex = findParentIndex(index)
	}
}

// Extract returns the maximum item from the max heap and swaps the elements as necessary to
// restore the heap property in the data structure.
func (m *MaxHeap[T]) Extract() T {
	result := m.array[0]
	m.array[0] = m.array[len(m.array)-1]
	m.array = m.array[:len(m.array)-1]

	index := 0
	childIndex1 := (index * 2) + 1
	childIndex2 := (index * 2) + 2
	for index < len(m.array) && ((childIndex1 < len(m.array) && m.array[index] < m.array[childIndex1]) || (childIndex2 < len(m.array) && m.array[index] < m.array[childIndex2])) {
		if m.array[childIndex1] > m.array[childIndex2] {
			temp := m.array[childIndex1]
			m.array[childIndex1] = m.array[index]
			m.array[index] = temp
			index = childIndex1
		} else {
			temp := m.array[childIndex2]
			m.array[childIndex2] = m.array[index]
			m.array[index] = temp
			index = childIndex1
		}
		childIndex1 = (index * 2) + 1
		childIndex2 = (index * 2) + 2
	}
	return result
}

func (m *MaxHeap[T]) Size() int {
	return len(m.array)
}

func (m *MaxHeap[T]) Peek() T {
	return m.array[0]
}