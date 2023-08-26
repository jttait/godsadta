package godsa

type Trie struct {
	root TrieNode
}

type TrieNode struct {
	children map[byte]*TrieNode
	terminal bool
}

func NewTrie() *Trie {
	trie := Trie{}
	trie.root.children = make(map[byte]*TrieNode)
	return &trie
}

func (trie *Trie) Insert(array []byte) bool {
	current := &(trie.root)
	for _, v := range array {
		val, ok := current.children[v]
		if ok {
			current = val
		} else {
			child := TrieNode{}
			child.children = make(map[byte]*TrieNode)
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

func (trie *Trie) Contains(array []byte) bool {
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

func (trie *Trie) Remove(array []byte) bool {
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
