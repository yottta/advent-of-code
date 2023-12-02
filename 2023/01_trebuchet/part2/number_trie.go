package main

type trieNode struct {
	children map[rune]*trieNode
	value    rune
	isWord   bool
}

type trie struct {
	root *trieNode
}

func newTrie() *trie {
	return &trie{root: &trieNode{children: make(map[rune]*trieNode)}}
}

func (t *trie) insert(s string) {
	node := t.root
	for _, r := range s {
		next, ok := node.children[r]
		if !ok {
			next = &trieNode{children: make(map[rune]*trieNode), value: r}
			node.children[r] = next
		}
		node = next
	}
	node.isWord = true
}

// first bool = partial match
// second bool = full match aka is a word
func (t *trie) find(s string) (bool, bool) {
	node := t.root
	for _, r := range s {
		next, ok := node.children[r]
		if !ok {
			return false, false
		}
		node = next
	}
	return true, node.isWord
}
