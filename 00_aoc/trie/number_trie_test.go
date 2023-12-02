package trie_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	aoc "github.com/yottta/advent-of-code/00_aoc"
	trie2 "github.com/yottta/advent-of-code/00_aoc/trie"
)

func TestTrie2023_01_p2(t *testing.T) {
	var stringToDigit = map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	trie := trie2.NewTrie()
	for k, v := range stringToDigit {
		trie.Insert(k)
		trie.Insert(aoc.StringReverse(k))
		trie.Insert(v)
	}
	for k, v := range stringToDigit {
		t.Run(fmt.Sprintf("%s->%d", k, v), func(t *testing.T) {
			for i := len(k); i > 0; i-- {
				input := k[:i]
				t.Run(input, func(t *testing.T) {
					partial, word := trie.Find(input)
					assert.True(t, partial)
					t.Logf("main word: %s; current test: %s; isPartial: %v; isWord: %v; expectedWord: %v", k, input, partial, word, len(input) == len(k))
					assert.Equal(t, len(input) == len(k), word)
				})
			}

			reverse := aoc.StringReverse(k)
			for i := len(reverse); i > 0; i-- {
				input := reverse[:i]
				t.Run(input, func(t *testing.T) {
					partial, word := trie.Find(input)
					assert.True(t, partial)
					t.Logf("main word: %s; current test: %s; isPartial: %v; isWord: %v; expectedWord: %v", reverse, input, partial, word, len(input) == len(k))
					assert.Equal(t, len(input) == len(k), word)
				})
			}

			t.Run(v, func(t *testing.T) {
				partial, word := trie.Find(v)
				assert.True(t, partial)
				assert.True(t, word)
			})
		})
	}
}
