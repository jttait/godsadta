package stack

import "testing"

func TestShouldBeSizeZeroForNewlyInstantiatedStackLinkedList(t *testing.T) {
	stack := NewStackLinkedList[int]()
	result := stack.Size()
	assertEqual(result, 0, t)
}

func TestShouldBeSizeOneForItemPushedToNewlyInstantiatedStackLinkedList(t *testing.T) {
	stack := NewStackLinkedList[int]()
	stack.Push(5)
	result := stack.Size()
	assertEqual(result, 1, t)
}

func TestShouldPopLatestPushedIntegerFromStackLinkedList(t *testing.T) {
	stack := NewStackLinkedList[int]()
	stack.Push(5)
	result, ok := stack.Pop()
	assertTrue(ok, t)
	assertEqual(result, 5, t)
}

func TestShouldPopLatestPushedStringFromStackLinkedList(t *testing.T) {
	stack := NewStackLinkedList[string]()
	stack.Push("five")
	result, ok := stack.Pop()
	assertTrue(ok, t)
	assertEqual(result, "five", t)
}

func TestShouldReturnFalseWhenPoppingEmptyStackLinkedList(t *testing.T) {
	stack := NewStackLinkedList[int]()
	_, ok := stack.Pop()
	assertFalse(ok, t)
}

func TestShouldReturnLatestItemWhenPeekingStackLinkedList(t *testing.T) {
	stack := NewStackLinkedList[int]()
	stack.Push(5)
	result, _ := stack.Peek()
	assertEqual(result, 5, t)
}

func TestShouldReturnFalseWhenPeekingEmptyStackLinkedList(t *testing.T) {
	stack := NewStackLinkedList[int]()
	_, ok := stack.Peek()
	assertFalse(ok, t)
}

func assertTrue(i bool, t *testing.T) {
	if !i {
		t.Fatalf("Got false. Want true.")
	}
}

func assertFalse(i bool, t *testing.T) {
	if i {
		t.Fatalf("Got true. Want false.")
	}
}

func assertEqual[T comparable](got T, want T, t *testing.T) {
	if got != want {
		t.Fatalf("Got %v. Want %v.\n", got, want)
	}
}
