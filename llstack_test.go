package godsa

import (
	"testing"

	"github.com/jttait/godsa/assert"
)

func TestShouldBeSizeZeroForNewlyInstantiatedLLStack(t *testing.T) {
	stack := NewLLStack[int]()
	result := stack.Size()
	assert.AssertEqual(result, 0, t)
}

func TestShouldBeSizeOneForItemPushedToNewlyInstantiatedLLStack(t *testing.T) {
	stack := NewLLStack[int]()
	stack.Push(5)
	result := stack.Size()
	assert.AssertEqual(result, 1, t)
}

func TestShouldPopLatestPushedIntegerFromLLStack(t *testing.T) {
	stack := NewLLStack[int]()
	stack.Push(5)
	result, ok := stack.Pop()
	assert.AssertTrue(ok, t)
	assert.AssertEqual(result, 5, t)
}

func TestShouldPopLatestPushedStringFromLLStack(t *testing.T) {
	stack := NewLLStack[string]()
	stack.Push("five")
	result, ok := stack.Pop()
	assert.AssertTrue(ok, t)
	assert.AssertEqual(result, "five", t)
}

func TestShouldReturnFalseWhenPoppingEmptyLLStack(t *testing.T) {
	stack := NewLLStack[int]()
	_, ok := stack.Pop()
	assert.AssertFalse(ok, t)
}

func TestShouldReturnLatestItemWhenPeekingLLStack(t *testing.T) {
	stack := NewLLStack[int]()
	stack.Push(5)
	result, _ := stack.Peek()
	assert.AssertEqual(result, 5, t)
}

func TestShouldReturnFalseWhenPeekingEmptyLLStack(t *testing.T) {
	stack := NewLLStack[int]()
	_, ok := stack.Peek()
	assert.AssertFalse(ok, t)
}
