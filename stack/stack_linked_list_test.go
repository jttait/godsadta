package stack

import "testing"

func TestShouldBeSizeZeroForNewlyInstantiatedStackLinkedList(t *testing.T) {
	want := 0
	stack := NewStackLinkedList[int]()
	result := stack.Size()
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldBeSizeOneForItemPushedToNewlyInstantiatedStackLinkedList(t *testing.T) {
	want := 1
	stack := NewStackLinkedList[int]()
	stack.Push(5)
	result := stack.Size()
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldPopLatestPushedIntegerFromStackLinkedList(t *testing.T) {
	want := 5
	stack := NewStackLinkedList[int]()
	stack.Push(5)
	result, ok := stack.Pop()
	if !ok {
		t.Fatalf("Was false")
	}
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldPopLatestPushedStringFromStackLinkedList(t *testing.T) {
	want := "five"
	stack := NewStackLinkedList[string]()
	stack.Push("five")
	result, ok := stack.Pop()
	if !ok {
		t.Fatalf("Was false")
	}
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldReturnFalseWhenPoppingEmptyStackLinkedList(t *testing.T) {
	want := false
	stack := NewStackLinkedList[int]()
	_, ok := stack.Pop()
	if ok {
		t.Fatalf("Want %v. Got %v\n", want, ok)
	}
}

func TestShouldReturnLatestItemWhenPeekingStackLinkedList(t *testing.T) {
	want := 5
	stack := NewStackLinkedList[int]()
	stack.Push(5)
	result, _ := stack.Peek()
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldReturnFalseWhenPeekingEmptyStackLinkedList(t *testing.T) {
	want := false
	stack := NewStackLinkedList[int]()
	_, ok := stack.Peek()
	if ok {
		t.Fatalf("Want %v. Got %v\n", want, ok)
	}
}
