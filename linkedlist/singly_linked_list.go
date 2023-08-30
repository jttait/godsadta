package linkedlist

// SinglyLinkedListNode is a data structure that contains a value and a pointer to the next node in
// the list.
type SinglyLinkedListNode[T any] struct {
	Val  T
	Next *SinglyLinkedListNode[T]
}

// NewSinglyLinkedListNode instantiates a singly-linked list node and returns a pointer to it.
func NewSinglyLinkedListNode[T any](i T) *SinglyLinkedListNode[T] {
	n := SinglyLinkedListNode[T]{}
	n.Val = i
	return &n
}

type SinglyLinkedList[T any] struct {
	dummyHead *SinglyLinkedListNode[T]
	size      int
}

// NewSinglyLinkedList instantiates a new singly-linked list and returns a pointer to it.
func NewSinglyLinkedList[T any]() *SinglyLinkedList[T] {
	q := SinglyLinkedList[T]{}
	var zeroValue T
	q.dummyHead = NewSinglyLinkedListNode[T](zeroValue)
	q.size = 0
	return &q
}

// Size returns the number of items in the list.
func (q *SinglyLinkedList[T]) Size() int {
	return q.size
}

// Insert inserts a new item at the front of the list.
func (q *SinglyLinkedList[T]) InsertFront(i T) {
	if q.dummyHead.Next == nil {
		q.dummyHead.Next = NewSinglyLinkedListNode[T](i)
	} else {
		temp := q.dummyHead.Next
		q.dummyHead.Next = NewSinglyLinkedListNode[T](i)
		q.dummyHead.Next.Next = temp
	}
	q.size += 1
}

func (q *SinglyLinkedList[T]) InsertLast(i T) {
	current := q.dummyHead
	for current.Next != nil {
		current = current.Next
	}
	current.Next = NewSinglyLinkedListNode[T](i)
	q.size += 1
}

// Remove removes and returns the item at the front of the singly-linked list.
func (q *SinglyLinkedList[T]) RemoveFront() (T, bool) {
	if q.size == 0 || q.dummyHead.Next == nil {
		var zeroValue T
		return zeroValue, false
	}
	result := q.dummyHead.Next.Val
	q.dummyHead.Next = q.dummyHead.Next.Next
	q.size -= 1
	return result, true
}

func (q *SinglyLinkedList[T]) RemoveLast() (T, bool) {
	current := q.dummyHead
	for current.Next.Next != nil {
		current = current.Next
	}
	result := current.Next.Val
	current.Next = nil
	q.size -= 1
	return result, true
}

func (q *SinglyLinkedList[T]) PeekFront() (T, bool) {
	if q.size == 0 {
		var zeroValue T
		return zeroValue, false
	}
	return q.dummyHead.Next.Val, true
}

func (q *SinglyLinkedList[T]) PeekLast() (T, bool) {
	if q.size == 0 {
		var zeroValue T
		return zeroValue, false
	}
	current := q.dummyHead
	for current.Next != nil {
		current = current.Next
	}
	return current.Val, true
}
