// Package doublylinkedlist provides the doubly-linked list data structure and associated methods
package linkedlist

// DoublyLinkedListNode is a data structure that contains a value and pointers to the previous and
// next nodes in the list.
type DoublyLinkedListNode[T any] struct {
	Val  T
	Prev *DoublyLinkedListNode[T]
	Next *DoublyLinkedListNode[T]
}

// NewDoublyLinkedListNode instantiates a doubly-linked list node and returns a pointer to it.
func NewDoublyLinkedListNode[T any](i T) *DoublyLinkedListNode[T] {
	d := DoublyLinkedListNode[T]{}
	d.Val = i
	return &d
}

// Double-ended queue enables inserting and polling from either end
type DoublyLinkedList[T any] struct {
	dummyHead *DoublyLinkedListNode[T]
	dummyTail *DoublyLinkedListNode[T]
}

// NewDoublyLinkedList instantiates a new doubly-linked list and returns a pointer to it.
func NewDoublyLinkedList[T any](items ...T) *DoublyLinkedList[T] {
	q := DoublyLinkedList[T]{}
	var zeroValue T
	q.dummyHead = NewDoublyLinkedListNode[T](zeroValue)
	q.dummyTail = NewDoublyLinkedListNode[T](zeroValue)
	q.dummyHead.Next = q.dummyTail
	q.dummyTail.Prev = q.dummyHead
	for _, i := range items {
		q.InsertLast(i)
	}
	return &q
}

// Size returns the number of items in the list.
func (q *DoublyLinkedList[T]) Size() int {
	result := 0
	current := q.dummyHead.Next
	for current != q.dummyTail {
		result += 1
		current = current.Next
	}
	return result
}

// InsertFront inserts a new item at the front of the list.
func (q *DoublyLinkedList[T]) InsertFront(i T) {
	insertedNode := NewDoublyLinkedListNode[T](i)
	currentHead := q.dummyHead.Next

	insertedNode.Prev = q.dummyHead
	insertedNode.Next = currentHead
	currentHead.Prev = insertedNode
	q.dummyHead.Next = insertedNode
}

// InsertLast inserts a new item at the end of the list.
func (q *DoublyLinkedList[T]) InsertLast(i T) {
	insertedNode := NewDoublyLinkedListNode[T](i)
	currentTail := q.dummyTail.Prev

	insertedNode.Next = q.dummyTail
	insertedNode.Prev = currentTail
	currentTail.Next = insertedNode
	q.dummyTail.Prev = insertedNode
}

func (q *DoublyLinkedList[T]) PeekFront() (T, bool) {
	if q.dummyHead.Next == q.dummyTail {
		var zeroValue T
		return zeroValue, false
	}
	return q.dummyHead.Next.Val, true
}

func (q *DoublyLinkedList[T]) PeekLast() (T, bool) {
	if q.dummyHead.Next == q.dummyTail {
		var zeroValue T
		return zeroValue, false
	}
	current := q.dummyHead
	for current.Next != q.dummyTail {
		current = current.Next
	}
	return current.Val, true
}

// RemoveFront removes and returns the item at the front of the doubly-linked list.
func (q *DoublyLinkedList[T]) RemoveFront() (T, bool) {
	if q.dummyHead.Next == q.dummyTail {
		var zeroValue T
		return zeroValue, false
	}
	result := q.dummyHead.Next.Val
	head := q.dummyHead.Next
	newHead := head.Next
	q.dummyHead.Next = newHead
	newHead.Prev = q.dummyHead
	return result, true
}

// RemoveLast removes and returns the item at the end of the doubly-linked list.
func (q *DoublyLinkedList[T]) RemoveLast() (T, bool) {
	if q.dummyHead.Next == q.dummyTail {
		var zeroValue T
		return zeroValue, false
	}
	result := q.dummyTail.Prev.Val
	tail := q.dummyTail.Prev
	newTail := tail.Prev
	q.dummyTail.Prev = newTail
	newTail.Next = q.dummyTail
	return result, true
}

func (l *DoublyLinkedList[T]) Insert(index int, item T) bool {
	return true
}

func (l *DoublyLinkedList[T]) Remove(index int) bool {
	return true
}

func (d *DoublyLinkedList[T]) Get(index int) (T, bool) {
	if d.dummyHead.Next == d.dummyTail {
		var zeroValue T
		return zeroValue, false
	}
	current := d.dummyHead.Next
	currentIndex := 0
	for currentIndex < index {
		currentIndex += 1
		current = current.Next
	}
	if current == d.dummyTail {
		var zeroValue T
		return zeroValue, false
	}
	return current.Val, true
}

func (d *DoublyLinkedList[T]) Map(f func(T) T) LinkedList[T] {
	result := NewDoublyLinkedList[T]()
	current := d.dummyHead.Next
	for current != nil {
		result.InsertLast(f(current.Val))
		current = current.Next
	}
	return result
}

func (d *DoublyLinkedList[T]) Filter(f func(T) bool) LinkedList[T] {
	result := NewDoublyLinkedList[T]()
	current := d.dummyHead.Next
	for current != nil {
		if f(current.Val) {
			result.InsertLast(current.Val)
		}
		current = current.Next
	}
	return result
}
