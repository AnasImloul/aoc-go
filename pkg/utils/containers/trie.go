package containers

// Trie is a trie (prefix tree) for string operations.
type Trie struct {
	children map[rune]*Trie
	isEnd    bool
}

// NewTrie creates a new trie.
func NewTrie() *Trie {
	return &Trie{
		children: make(map[rune]*Trie),
		isEnd:    false,
	}
}

// Insert inserts a word into the trie.
func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		if node.children[ch] == nil {
			node.children[ch] = NewTrie()
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

// Search returns true if the word is in the trie.
func (t *Trie) Search(word string) bool {
	node := t
	for _, ch := range word {
		if node.children[ch] == nil {
			return false
		}
		node = node.children[ch]
	}
	return node.isEnd
}

// StartsWith returns true if there is any word in the trie that starts with the given prefix.
func (t *Trie) StartsWith(prefix string) bool {
	node := t
	for _, ch := range prefix {
		if node.children[ch] == nil {
			return false
		}
		node = node.children[ch]
	}
	return true
}

// Delete removes a word from the trie.
func (t *Trie) Delete(word string) bool {
	deleted, _ := t.deleteHelper(word, 0)
	return deleted
}

func (t *Trie) deleteHelper(word string, index int) (bool, bool) {
	if index == len(word) {
		if !t.isEnd {
			return false, false
		}
		t.isEnd = false
		return true, len(t.children) == 0
	}
	ch := rune(word[index])
	child, exists := t.children[ch]
	if !exists {
		return false, false
	}
	deleted, shouldDelete := child.deleteHelper(word, index+1)
	if shouldDelete {
		delete(t.children, ch)
	}
	return deleted, len(t.children) == 0 && !t.isEnd
}
