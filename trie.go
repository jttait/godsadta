package godsa

type Trie[T comparable] struct {
	root TrieNode[T]
}

type TrieNode[T comparable] struct {
	children map[T]*TrieNode[T]
	terminal bool
}

func NewTrie[T comparable]() *Trie[T] {
	trie := Trie[T]{}
	trie.root.children = make(map[T]*TrieNode[T])
	return &trie
}

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
