// Package trie provides an interface and implementations for the trie data structure
package trie

// Trie is a data structure for locating keys within a set but the key is broken into individual
// parts (e.g. a string broken into characters) and the key is access by traversing the trie
// depth-first.
type Trie[T comparable] struct {
	root TrieNode[T]
}

// TrieNode is a node within the trie and contains a map of child nodes and a Boolean to say
// whether a key terminates at this node.
type TrieNode[T comparable] struct {
	children map[T]*TrieNode[T]
	terminal bool
}

// NewTrie instantiates a trie and returns a pointer to it.
func NewTrie[T comparable]() *Trie[T] {
	trie := Trie[T]{}
	trie.root.children = make(map[T]*TrieNode[T])
	return &trie
}

// Insert inserts a new key into the trie. It returns a Boolean that is true if the key was not
// already present in the trie.
func (trie *Trie[T]) Insert(array []T) bool {
	current := &(trie.root)
	for _, v := range array {
		val, ok := current.children[v]
		if ok {
			current = val
		} else {
			child := TrieNode[T]{}
			child.children = make(map[T]*TrieNode[T])
			current.children[v] = &child
			current = &child
		}
	}
	if current.terminal == true {
		return false
	}
	current.terminal = true
	return true
}

// Contains returns a Boolean that is true if the given key exists in the trie.
func (trie *Trie[T]) Contains(array []T) bool {
	current := &(trie.root)
	for _, v := range array {
		_, ok := current.children[v]
		if !ok {
			return false
		}
		current = current.children[v]
	}
	return current.terminal
}

// Remove will remove the given key from the trie.
func (trie *Trie[T]) Remove(array []T) bool {
	current := &(trie.root)
	for _, v := range array {
		_, ok := current.children[v]
		if !ok {
			return false
		}
		current = current.children[v]
	}
	current.terminal = false
	return true
}
