package stack

import "testing"

func TestShouldBeSizeZeroForNewlyInstantiatedStack(t *testing.T) {
	want := 0
	stack := NewStack[int]()
	result := stack.Size()
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldBeSizeOneForItemPushedToNewlyInstantiatedStack(t *testing.T) {
	want := 1
	stack := NewStack[int]()
	stack.Push(5)
	result := stack.Size()
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldPopLatestPushedInteger(t *testing.T) {
	want := 5
	stack := NewStack[int]()
	stack.Push(5)
	result, ok := stack.Pop()
	if !ok {
		t.Fatalf("Was false")
	}
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldPopLatestPushedString(t *testing.T) {
	want := "five"
	stack := NewStack[string]()
	stack.Push("five")
	result, ok := stack.Pop()
	if !ok {
		t.Fatalf("Was false")
	}
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldReturnFalseWhenPoppingEmptyStack(t *testing.T) {
	want := false
	stack := NewStack[int]()
	_, ok := stack.Pop()
	if ok {
		t.Fatalf("Want %v. Got %v\n", want, ok)
	}
}

func TestShouldReturnLatestItemWhenPeeking(t *testing.T) {
	want := 5
	stack := NewStack[int]()
	stack.Push(5)
	result, _ := stack.Peek()
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldReturnFalseWhenPeekingEmptyStack(t *testing.T) {
	want := false
	stack := NewStack[int]()
	_, ok := stack.Peek()
	if ok {
		t.Fatalf("Want %v. Got %v\n", want, ok)
	}
}
