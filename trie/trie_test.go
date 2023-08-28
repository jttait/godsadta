package trie

import "testing"

func TestShouldReturnTrueIfTrieContainsItem(t *testing.T) {
	trie := NewTrie[byte]()
	_ = trie.Insert([]byte{'c', 'a', 't'})
	result := trie.Contains([]byte{'c', 'a', 't'})
	want := true
	if result != want {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldReturnFalseIfTrieDoesNotContainItem(t *testing.T) {
	trie := NewTrie[byte]()
	result := trie.Contains([]byte{'c', 'a', 't'})
	want := false
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldRemoveExistingItemFromTrie(t *testing.T) {
	trie := NewTrie[byte]()
	_ = trie.Insert([]byte{'c', 'a', 't'})
	_ = trie.Remove([]byte{'c', 'a', 't'})
	result := trie.Contains([]byte{'c', 'a', 't'})
	want := false
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldReturnFalseWhenRemovingNonExistingItemFromTrie(t *testing.T) {
	trie := NewTrie[byte]()
	result := trie.Remove([]byte{'c', 'a', 't'})
	want := false
	if result != want {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldReturnTrueWhenRemovingExistingWordFromTrie(t *testing.T) {
	trie := NewTrie[byte]()
	_ = trie.Insert([]byte{'c', 'a', 't'})
	result := trie.Remove([]byte{'c', 'a', 't'})
	want := true
	if result != want {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldContainItemAfterAddingRemovingAddingToTrie(t *testing.T) {
	trie := NewTrie[byte]()
	_ = trie.Insert([]byte{'c', 'a', 't'})
	_ = trie.Remove([]byte{'c', 'a', 't'})
	_ = trie.Insert([]byte{'c', 'a', 't'})
	result := trie.Contains([]byte{'c', 'a', 't'})
	want := true
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldBeTrueWhenInsertingItemNotAlreadyExistingInTrie(t *testing.T) {
	trie := NewTrie[byte]()
	result := trie.Insert([]byte{'c', 'a', 't'})
	want := true
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldBeFalseWhenInsertingItemAlreadyExistingInTrie(t *testing.T) {
	trie := NewTrie[byte]()
	_ = trie.Insert([]byte{'c', 'a', 't'})
	result := trie.Insert([]byte{'c', 'a', 't'})
	want := false
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}
