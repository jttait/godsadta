package maxheap

import "cmp"

type MaxHeapArray[T cmp.Ordered] struct {
	array []T
}

// NewMaxHeap instantiates a max heap and returns a pointer to it. The variadic argument can be used
// to insert items into the heap after it is created.
func NewMaxHeapArray[T cmp.Ordered](i ...T) *MaxHeapArray[T] {
	m := MaxHeapArray[T]{}
	for _, j := range i {
		m.Insert(j)
	}
	return &m
}

func findParentIndex(index int) int {
	return (index - 1) / 2
}

// Insert inserts the item into the max heap.
func (m *MaxHeapArray[T]) Insert(i T) {
	m.array = append(m.array, i)
	index := len(m.array) - 1
	parentIndex := findParentIndex(index)
	for parentIndex >= 0 && m.array[index] > m.array[parentIndex] {
		swap(index, parentIndex, &m.array)
		index = parentIndex
		parentIndex = findParentIndex(index)
	}
}

// Extract returns the maximum item from the max heap and swaps the elements as necessary to
// restore the heap property in the data structure.
func (m *MaxHeapArray[T]) Extract() (T, bool) {
	if m.Size() == 0 {
		var zeroValue T
		return zeroValue, false
	}
	result := m.array[0]
	m.array[0] = m.array[len(m.array)-1]
	m.array = m.array[:len(m.array)-1]

	index := 0
	childIndex1, childIndex2 := findChildIndices(index)
	for index < len(m.array) && ((childIndex1 < len(m.array) && m.array[index] < m.array[childIndex1]) || (childIndex2 < len(m.array) && m.array[index] < m.array[childIndex2])) {
		if m.array[childIndex1] > m.array[childIndex2] {
			swap(index, childIndex1, &m.array)
		} else {
			swap(index, childIndex2, &m.array)
		}
		index = childIndex1
		childIndex1, childIndex2 = findChildIndices(index)
	}
	return result, true
}

func swap[T cmp.Ordered](i int, j int, array *[]T) {
	temp := (*array)[i]
	(*array)[i] = (*array)[j]
	(*array)[j] = temp
}

func findChildIndices(index int) (int, int) {
	return (index * 2) + 1, (index * 2) + 2
}

func (m *MaxHeapArray[T]) Size() int {
	return len(m.array)
}

func (m *MaxHeapArray[T]) Peek() (T, bool) {
	if m.Size() == 0 {
		var zeroValue T
		return zeroValue, false
	}
	return m.array[0], true
}
