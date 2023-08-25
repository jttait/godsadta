package godsa

import "testing"

func TestShouldBeSizeZeroForNewlyInstantiatedStack(t *testing.T) {
	want := 0
	stack := NewStack()
	result := stack.Size()
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldBeSizeOneForItemPushedToNewlyInstantiatedStack(t *testing.T) {
	want := 1
	stack := NewStack()
	stack.Push(5)
	result := stack.Size()
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldPopLatestPushedItem(t *testing.T) {
	want := 5
	stack := NewStack()
	stack.Push(5)
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
	stack := NewStack()
	_, ok := stack.Pop()
	if ok {
		t.Fatalf("Want %v. Got %v\n", want, ok)
	}
}

func TestShouldReturnLatestItemWhenPeeking(t *testing.T) {
	want := 5
	stack := NewStack()
	stack.Push(5)
	result, _ := stack.Peek()
	if want != result {
		t.Fatalf("Want %v. Got %v\n", want, result)
	}
}

func TestShouldReturnFalseWhenPeekingEmptyStack(t *testing.T) {
	want := false
	stack := NewStack()
	_, ok := stack.Peek()
	if ok {
		t.Fatalf("Want %v. Got %v\n", want, ok)
	}
}
