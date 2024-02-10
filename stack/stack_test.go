package stack

import (
	"testing"

	"github.com/jttait/godsa/assert"
)

func getStackImplementationsInt() []Stack[int] {
	return []Stack[int]{
		NewLLStack[int](),
	}
}

func getStackImplementationsString() []Stack[string] {
	return []Stack[string]{
		NewLLStack[string](),
	}
}

func TestShouldBeSizeZeroForNewlyInstantiatedStack(t *testing.T) {
	for _, stack := range getStackImplementationsInt() {
		result := stack.Size()
		assert.AssertEqual(result, 0, t)
	}
}

func TestShouldBeSizeOneForItemPushedToNewlyInstantiatedStack(t *testing.T) {
	for _, stack := range getStackImplementationsInt() {
		stack.Push(5)
		result := stack.Size()
		assert.AssertEqual(result, 1, t)
	}
}

func TestShouldPopLatestPushedIntegerFromStack(t *testing.T) {
	for _, stack := range getStackImplementationsInt() {
		stack.Push(5)
		result, ok := stack.Pop()
		assert.AssertTrue(ok, t)
		assert.AssertEqual(result, 5, t)
	}
}

func TestShouldPopLatestPushedStringFromStack(t *testing.T) {
	for _, stack := range getStackImplementationsString() {
		stack.Push("five")
		result, ok := stack.Pop()
		assert.AssertTrue(ok, t)
		assert.AssertEqual(result, "five", t)
	}
}

func TestShouldReturnFalseWhenPoppingEmptyStack(t *testing.T) {
	for _, stack := range getStackImplementationsInt() {
		_, ok := stack.Pop()
		assert.AssertFalse(ok, t)
	}
}

func TestShouldReturnLatestItemWhenPeekingStack(t *testing.T) {
	for _, stack := range getStackImplementationsInt() {
		stack.Push(5)
		result, _ := stack.Peek()
		assert.AssertEqual(result, 5, t)
	}
}

func TestShouldReturnFalseWhenPeekingEmptyStack(t *testing.T) {
	for _, stack := range getStackImplementationsInt() {
		_, ok := stack.Peek()
		assert.AssertFalse(ok, t)
	}
}
