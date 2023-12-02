package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrieWords(t *testing.T) {
	trie := newTrie()
	for k, v := range stringToDigit {
		trie.insert(k)
		trie.insert(stringReverse(k))
		trie.insert(v)
	}
	for k, v := range stringToDigit {
		t.Run(fmt.Sprintf("%s->%d", k, v), func(t *testing.T) {
			for i := len(k); i > 0; i-- {
				input := k[:i]
				t.Run(input, func(t *testing.T) {
					partial, word := trie.find(input)
					assert.True(t, partial)
					t.Logf("main word: %s; current test: %s; isPartial: %v; isWord: %v; expectedWord: %v", k, input, partial, word, len(input) == len(k))
					assert.Equal(t, len(input) == len(k), word)
				})
			}

			reverse := stringReverse(k)
			for i := len(reverse); i > 0; i-- {
				input := reverse[:i]
				t.Run(input, func(t *testing.T) {
					partial, word := trie.find(input)
					assert.True(t, partial)
					t.Logf("main word: %s; current test: %s; isPartial: %v; isWord: %v; expectedWord: %v", reverse, input, partial, word, len(input) == len(k))
					assert.Equal(t, len(input) == len(k), word)
				})
			}

			t.Run(v, func(t *testing.T) {
				partial, word := trie.find(v)
				assert.True(t, partial)
				assert.True(t, word)
			})
		})
	}
}
