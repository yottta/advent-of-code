package trie

type trieNode struct {
	children map[rune]*trieNode
	value    rune
	isWord   bool
}

type Trie interface {
	Insert(s string)
	Find(s string) (bool, bool)
}

type trie struct {
	root *trieNode
}

func NewTrie() Trie {
	return &trie{root: &trieNode{children: make(map[rune]*trieNode)}}
}

func (t *trie) Insert(s string) {
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
func (t *trie) Find(s string) (bool, bool) {
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
