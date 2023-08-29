// Package doublylinkedlist provides the doubly-linked list data structure and associated methods
package singlylinkedlist

// Double-ended queue enables inserting and polling from either end
type SinglyLinkedList[T any] struct {
	dummyHead *SinglyLinkedListNode[T]
	size      int
}

// NewDoublyLinkedList instantiates a new doubly-linked list and returns a pointer to it.
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
func (q *SinglyLinkedList[T]) Insert(i T) {
	if q.dummyHead.Next == nil {
		q.dummyHead.Next = NewSinglyLinkedListNode[T](i)
	} else {
		temp := q.dummyHead.Next
		q.dummyHead.Next = NewSinglyLinkedListNode[T](i)
		q.dummyHead.Next.Next = temp
	}
	q.size += 1
}

// Remove removes and returns the item at the front of the singly-linked list.
func (q *SinglyLinkedList[T]) Remove() (T, bool) {
	if q.size == 0 || q.dummyHead.Next == nil {
		var zeroValue T
		return zeroValue, false
	}
	result := q.dummyHead.Next.Val
	q.dummyHead.Next = q.dummyHead.Next.Next
	q.size -= 1
	return result, true
}

func (q *SinglyLinkedList[T]) Peek() (T, bool) {
	if q.size == 0 {
		var zeroValue T
		return zeroValue, false
	}
	return q.dummyHead.Next.Val, true
}
