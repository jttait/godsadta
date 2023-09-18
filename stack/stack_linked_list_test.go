package stack

import (
	"testing"

	"github.com/jttait/godsa/assert"
)

func TestShouldBeSizeZeroForNewlyInstantiatedStackLinkedList(t *testing.T) {
	stack := NewStackLinkedList[int]()
	result := stack.Size()
	assert.AssertEqual(result, 0, t)
}

func TestShouldBeSizeOneForItemPushedToNewlyInstantiatedStackLinkedList(t *testing.T) {
	stack := NewStackLinkedList[int]()
	stack.Push(5)
	result := stack.Size()
	assert.AssertEqual(result, 1, t)
}

func TestShouldPopLatestPushedIntegerFromStackLinkedList(t *testing.T) {
	stack := NewStackLinkedList[int]()
	stack.Push(5)
	result, ok := stack.Pop()
	assert.AssertTrue(ok, t)
	assert.AssertEqual(result, 5, t)
}

func TestShouldPopLatestPushedStringFromStackLinkedList(t *testing.T) {
	stack := NewStackLinkedList[string]()
	stack.Push("five")
	result, ok := stack.Pop()
	assert.AssertTrue(ok, t)
	assert.AssertEqual(result, "five", t)
}

func TestShouldReturnFalseWhenPoppingEmptyStackLinkedList(t *testing.T) {
	stack := NewStackLinkedList[int]()
	_, ok := stack.Pop()
	assert.AssertFalse(ok, t)
}

func TestShouldReturnLatestItemWhenPeekingStackLinkedList(t *testing.T) {
	stack := NewStackLinkedList[int]()
	stack.Push(5)
	result, _ := stack.Peek()
	assert.AssertEqual(result, 5, t)
}

func TestShouldReturnFalseWhenPeekingEmptyStackLinkedList(t *testing.T) {
	stack := NewStackLinkedList[int]()
	_, ok := stack.Peek()
	assert.AssertFalse(ok, t)
}
