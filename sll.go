package godsa

// SinglyLinkedListNode is a data structure that contains a value and a pointer to the next node in
// the list.
type SLLNode[T any] struct {
	Val  T
	Next *SLLNode[T]
}

// NewSinglyLinkedListNode instantiates a singly-linked list node and returns a pointer to it.
func NewSLLNode[T any](i T) *SLLNode[T] {
	n := SLLNode[T]{}
	n.Val = i
	return &n
}

type SLL[T any] struct {
	dummyHead *SLLNode[T]
}

// NewSinglyLinkedList instantiates a new singly-linked list and returns a pointer to it.
func NewSLL[T any](items ...T) *SLL[T] {
	l := SLL[T]{}
	var zeroValue T
	l.dummyHead = NewSLLNode[T](zeroValue)
	for _, i := range items {
		l.InsertLast(i)
	}
	return &l
}

// Size returns the number of items in the list.
func (l *SLL[T]) Size() int {
	if l.dummyHead.Next == nil {
		return 0
	}
	result := 0
	current := l.dummyHead
	for current.Next != nil {
		current = current.Next
		result += 1
	}
	return result
}

// InsertFront inserts a new item at the front of the list.
func (l *SLL[T]) InsertFront(i T) {
	if l.dummyHead.Next == nil {
		l.dummyHead.Next = NewSLLNode[T](i)
	} else {
		temp := l.dummyHead.Next
		l.dummyHead.Next = NewSLLNode[T](i)
		l.dummyHead.Next.Next = temp
	}
}

// InsertLast inserts a new item at the end of the list.
func (l *SLL[T]) InsertLast(i T) {
	current := l.dummyHead
	for current.Next != nil {
		current = current.Next
	}
	current.Next = NewSLLNode[T](i)
}

// Insert inserts a new item at the given index of the list. It returns a Boolean if the index is
// true if the index was within the bounds of the list and it was possible to insert the item.
func (l *SLL[T]) Insert(index int, item T) bool {
	current := l.dummyHead.Next
	currentIndex := 0
	for current != nil {
		if currentIndex == index {
			newNode := NewSLLNode[T](current.Val)
			newNode.Next = current.Next
			current.Val = item
			current.Next = newNode
			return true
		}
		index += 1
		current = current.Next
	}
	return false
}

func (l *SLL[T]) Remove(index int) bool {
	previous := l.dummyHead
	current := l.dummyHead.Next
	currentIndex := 0
	for current != nil {
		if currentIndex == index {
			previous.Next = current.Next
		}
		index += 1
		current = current.Next
		previous = previous.Next
		return true
	}
	return false
}

// RemoveFront removes and returns the item at the front of the singly-linked list.
func (l *SLL[T]) RemoveFront() (T, bool) {
	if l.dummyHead.Next == nil {
		var zeroValue T
		return zeroValue, false
	}
	result := l.dummyHead.Next.Val
	l.dummyHead.Next = l.dummyHead.Next.Next
	return result, true
}

// RemoveLast removes and returns the item at the end of the list.
func (l *SLL[T]) RemoveLast() (T, bool) {
	if l.dummyHead.Next == nil {
		var zeroValue T
		return zeroValue, false
	}
	current := l.dummyHead
	for current.Next.Next != nil {
		current = current.Next
	}
	result := current.Next.Val
	current.Next = nil
	return result, true
}

func (l *SLL[T]) Equal(m LinkedList[T]) bool {
	for l != nil && m != nil {
		if l != m {
			return false
		}
	}
	return l == m
}

// PeekFront returns the item at the front of the list.
func (l *SLL[T]) PeekFront() (T, bool) {
	if l.dummyHead.Next == nil {
		var zeroValue T
		return zeroValue, false
	}
	return l.dummyHead.Next.Val, true
}

// PeekLast returns the item at the end of the list.
func (l *SLL[T]) PeekLast() (T, bool) {
	if l.dummyHead.Next == nil {
		var zeroValue T
		return zeroValue, false
	}
	current := l.dummyHead
	for current.Next != nil {
		current = current.Next
	}
	return current.Val, true
}

// Get returns the item at a given index of the list.
func (l *SLL[T]) Get(index int) (T, bool) {
	if l.dummyHead.Next == nil {
		var zeroValue T
		return zeroValue, false
	}
	current := l.dummyHead.Next
	currentIndex := 0
	for currentIndex < index {
		currentIndex += 1
		current = current.Next
	}
	if current == nil {
		var zeroValue T
		return zeroValue, false
	}
	return current.Val, true
}

// Map applies a given function to each item in the list.
func (l *SLL[T]) Map(f func(T) T) LinkedList[T] {
	result := NewSLL[T]()
	current := l.dummyHead.Next
	for current != nil {
		result.InsertLast(f(current.Val))
		current = current.Next
	}
	return result
}

// Filter applies a given predicate function to each item in the list and returns a linked list of
// all items for which the predicate is true.
func (l *SLL[T]) Filter(f func(T) bool) LinkedList[T] {
	result := NewSLL[T]()
	current := l.dummyHead.Next
	for current != nil {
		if f(current.Val) {
			result.InsertLast(current.Val)
		}
		current = current.Next
	}
	return result
}
