package trie

import (
	"testing"

	"github.com/jttait/godsadta/assert"
)

func TestShouldReturnTrueIfTrieContainsItem(t *testing.T) {
	trie := NewTrie[byte]()
	_ = trie.Insert([]byte{'c', 'a', 't'})
	result := trie.Contains([]byte{'c', 'a', 't'})
	assert.AssertTrue(result, t)
}

func TestShouldReturnFalseIfTrieDoesNotContainItem(t *testing.T) {
	trie := NewTrie[byte]()
	result := trie.Contains([]byte{'c', 'a', 't'})
	assert.AssertFalse(result, t)
}

func TestShouldRemoveExistingItemFromTrie(t *testing.T) {
	trie := NewTrie[byte]()
	_ = trie.Insert([]byte{'c', 'a', 't'})
	_ = trie.Remove([]byte{'c', 'a', 't'})
	result := trie.Contains([]byte{'c', 'a', 't'})
	assert.AssertFalse(result, t)
}

func TestShouldReturnFalseWhenRemovingNonExistingItemFromTrie(t *testing.T) {
	trie := NewTrie[byte]()
	result := trie.Remove([]byte{'c', 'a', 't'})
	assert.AssertFalse(result, t)
}

func TestShouldReturnTrueWhenRemovingExistingWordFromTrie(t *testing.T) {
	trie := NewTrie[byte]()
	_ = trie.Insert([]byte{'c', 'a', 't'})
	result := trie.Remove([]byte{'c', 'a', 't'})
	assert.AssertTrue(result, t)
}

func TestShouldContainItemAfterAddingRemovingAddingToTrie(t *testing.T) {
	trie := NewTrie[byte]()
	_ = trie.Insert([]byte{'c', 'a', 't'})
	_ = trie.Remove([]byte{'c', 'a', 't'})
	_ = trie.Insert([]byte{'c', 'a', 't'})
	result := trie.Contains([]byte{'c', 'a', 't'})
	assert.AssertTrue(result, t)
}

func TestShouldBeTrueWhenInsertingItemNotAlreadyExistingInTrie(t *testing.T) {
	trie := NewTrie[byte]()
	result := trie.Insert([]byte{'c', 'a', 't'})
	assert.AssertTrue(result, t)
}

func TestShouldBeFalseWhenInsertingItemAlreadyExistingInTrie(t *testing.T) {
	trie := NewTrie[byte]()
	_ = trie.Insert([]byte{'c', 'a', 't'})
	result := trie.Insert([]byte{'c', 'a', 't'})
	assert.AssertFalse(result, t)
}
