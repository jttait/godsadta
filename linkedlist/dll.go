package linkedlist

import (
	"reflect"
)

// DoublyLinkedListNode is a data structure that contains a value and pointers to the previous and
// next nodes in the list.
type DLLNode[T any] struct {
	Val  T
	Prev *DLLNode[T]
	Next *DLLNode[T]
}

// NewDoublyLinkedListNode instantiates a doubly-linked list node and returns a pointer to it.
func NewDLLNode[T any](i T) *DLLNode[T] {
	d := DLLNode[T]{}
	d.Val = i
	return &d
}

// Double-ended queue enables inserting and polling from either end
type DLL[T any] struct {
	dummyHead *DLLNode[T]
	dummyTail *DLLNode[T]
}

// NewDoublyLinkedList instantiates a new doubly-linked list and returns a pointer to it.
func NewDLL[T any](items ...T) *DLL[T] {
	q := DLL[T]{}
	var zeroValue T
	q.dummyHead = NewDLLNode[T](zeroValue)
	q.dummyTail = NewDLLNode[T](zeroValue)
	q.dummyHead.Next = q.dummyTail
	q.dummyTail.Prev = q.dummyHead
	for _, i := range items {
		q.InsertLast(i)
	}
	return &q
}

// Size returns the number of items in the list.
func (q *DLL[T]) Size() int {
	result := 0
	current := q.dummyHead.Next
	for current != q.dummyTail {
		result += 1
		current = current.Next
	}
	return result
}

// InsertFront inserts a new item at the front of the list.
func (q *DLL[T]) InsertFront(i T) {
	insertedNode := NewDLLNode[T](i)
	currentHead := q.dummyHead.Next

	insertedNode.Prev = q.dummyHead
	insertedNode.Next = currentHead
	currentHead.Prev = insertedNode
	q.dummyHead.Next = insertedNode
}

func (q *DLL[T]) InsertNodeFront(insertedNode *DLLNode[T]) {
	currentHead := q.dummyHead.Next
	insertedNode.Prev = q.dummyHead
	insertedNode.Next = currentHead
	currentHead.Prev = insertedNode
	q.dummyHead.Next = insertedNode
}

// InsertLast inserts a new item at the end of the list.
func (q *DLL[T]) InsertLast(i T) {
	insertedNode := NewDLLNode[T](i)
	currentTail := q.dummyTail.Prev

	insertedNode.Next = q.dummyTail
	insertedNode.Prev = currentTail
	currentTail.Next = insertedNode
	q.dummyTail.Prev = insertedNode
}

// Insert inserts a new item at the given index of the list. It returns a Boolean if the index is
// true if the index was within the bounds of the list and it was possible to insert the item.
func (l *DLL[T]) Insert(index int, item T) bool {
	current := l.dummyHead.Next
	currentIndex := 0
	for current != l.dummyTail {
		if currentIndex == index {
			newNode := NewDLLNode[T](current.Val)
			newNode.Next = current.Next
			current.Next = newNode
			current.Val = item
			return true
		}
		index += 1
		current = current.Next
	}
	return false
}

func (q *DLL[T]) PeekFront() (T, bool) {
	if q.dummyHead.Next == q.dummyTail {
		var zeroValue T
		return zeroValue, false
	}
	return q.dummyHead.Next.Val, true
}

func (q *DLL[T]) PeekLast() (T, bool) {
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
func (q *DLL[T]) RemoveFront() (T, bool) {
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
func (l *DLL[T]) RemoveLast() (T, bool) {
	if l.dummyHead.Next == l.dummyTail {
		var zeroValue T
		return zeroValue, false
	}
	result := l.dummyTail.Prev.Val
	tail := l.dummyTail.Prev
	newTail := tail.Prev
	l.dummyTail.Prev = newTail
	newTail.Next = l.dummyTail
	return result, true
}

func (l *DLL[T]) Remove(index int) bool {
	current := l.dummyHead.Next
	currentIndex := 0
	for current != l.dummyTail {
		if currentIndex == index {
			current.Prev.Next = current.Next
			return true
		}
		currentIndex += 1
		current = current.Next
	}
	return false
}

func (l *DLL[T]) RemoveNode(ptr *DLLNode[T]) {
	ptr.Prev.Next = ptr.Next
}

func (d *DLL[T]) Get(index int) (T, bool) {
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

func (d *DLL[T]) Equal(f LinkedList[T]) bool {
	e := f.(*DLL[T])
	dCurrent := d.dummyHead.Next
	eCurrent := e.dummyHead.Next
	for dCurrent != d.dummyTail && eCurrent != e.dummyTail {
		if !reflect.DeepEqual(dCurrent.Val, eCurrent.Val) {
			return false
		}
		dCurrent = dCurrent.Next
		eCurrent = eCurrent.Next
	}
	return reflect.DeepEqual(dCurrent.Val, eCurrent.Val)
}

func (d *DLL[T]) Map(f func(T) T) LinkedList[T] {
	result := NewDLL[T]()
	current := d.dummyHead.Next
	for current != nil {
		result.InsertLast(f(current.Val))
		current = current.Next
	}
	return result
}

func (d *DLL[T]) Filter(f func(T) bool) LinkedList[T] {
	result := NewDLL[T]()
	current := d.dummyHead.Next
	for current != nil {
		if f(current.Val) {
			result.InsertLast(current.Val)
		}
		current = current.Next
	}
	return result
}
