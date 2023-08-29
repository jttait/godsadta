// Package doublylinkedlist provides the doubly-linked list data structure and associated methods
package doublylinkedlist

// Double-ended queue enables inserting and polling from either end
type DoublyLinkedList[T any] struct {
	dummyHead *DoublyLinkedListNode[T]
	dummyTail *DoublyLinkedListNode[T]
	size      int
}

// NewDoublyLinkedList instantiates a new doubly-linked list and returns a pointer to it.
func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	q := DoublyLinkedList[T]{}
	var zeroValue T
	q.dummyHead = NewDoublyLinkedListNode[T](zeroValue)
	q.dummyTail = NewDoublyLinkedListNode[T](zeroValue)
	q.dummyHead.Next = q.dummyTail
	q.dummyTail.Prev = q.dummyHead
	q.size = 0
	return &q
}

// Size returns the number of items in the list.
func (q *DoublyLinkedList[T]) Size() int {
	return q.size
}

// InsertFront inserts a new item at the front of the list.
func (q *DoublyLinkedList[T]) InsertFront(i T) {
	insertedNode := NewDoublyLinkedListNode[T](i)
	currentHead := q.dummyHead.Next

	insertedNode.Prev = q.dummyHead
	insertedNode.Next = currentHead
	currentHead.Prev = insertedNode
	q.dummyHead.Next = insertedNode

	q.size += 1
}

// InsertLast inserts a new item at the end of the list.
func (q *DoublyLinkedList[T]) InsertLast(i T) {
	insertedNode := NewDoublyLinkedListNode[T](i)
	currentTail := q.dummyTail.Prev

	insertedNode.Next = q.dummyTail
	insertedNode.Prev = currentTail
	currentTail.Next = insertedNode
	q.dummyTail.Prev = insertedNode

	q.size += 1
}

// RemoveFront removes and returns the item at the front of the doubly-linked list.
func (q *DoublyLinkedList[T]) RemoveFront() (T, bool) {
	if q.size == 0 {
		var zeroValue T
		return zeroValue, false
	}
	result := q.dummyHead.Next.Val
	head := q.dummyHead.Next
	newHead := head.Next
	q.dummyHead.Next = newHead
	newHead.Prev = q.dummyHead
	q.size -= 1
	return result, true
}

// RemoveLast removes and returns the item at the end of the doubly-linked list.
func (q *DoublyLinkedList[T]) RemoveLast() (T, bool) {
	if q.size == 0 {
		var zeroValue T
		return zeroValue, false
	}
	result := q.dummyTail.Prev.Val
	tail := q.dummyTail.Prev
	newTail := tail.Prev
	q.dummyTail.Prev = newTail
	newTail.Next = q.dummyTail
	q.size -= 1
	return result, true
}
