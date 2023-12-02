package main

import (
	"bytes"
	"fmt"
	"strconv"

	aoc "github.com/yottta/advent-of-code/00_aoc"
)

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

// 54283 too low
// 54343 too low
// 54353 too low
func main() {
	forwardTrie, reverseTrie := createTrie()
	// debug
	// good example why the "backtrack" is needed
	//fmt.Println(findCalibrationValue(forwardTrie, reverseTrie, "6jxgvknvbbbmqbjkkbnineninenj"))        // expected: 69; resulted:66

	content, err := aoc.ReadFile("input.txt")
	aoc.Must(err)

	var sum int
	for i, line := range content {
		value := findCalibrationValue(forwardTrie, reverseTrie, line)
		sum += value
		fmt.Println(i, line, value)
	}
	fmt.Println(sum)
}

func findCalibrationValue(forwardTrie, reverseTrie *trie, in string) int {
	var firstDigit, firstPartial, lastPartial, lastDigit string

	for i, j := 0, len(in)-1; i < len(in) && j >= 0; i, j = i+1, j-1 {
		if firstDigit == "" {
			firstPartial += string(in[i])
			partial, word := forwardTrie.find(firstPartial)
			if !partial {
				tmp := firstPartial[1:]
				// backtrack
				if ok, okW := forwardTrie.find(tmp); ok && len(tmp) > 0 {
					firstPartial = tmp
					if okW {
						firstDigit = firstPartial
					}
				} else {
					firstPartial = string(in[i])
					if _, ok := forwardTrie.find(firstPartial); ok {
						firstDigit = firstPartial
					}
				}
			} else if word {
				firstDigit = firstPartial
			}
		}

		if lastDigit == "" {
			lastPartial += string(in[j])
			partial, word := reverseTrie.find(lastPartial)
			if !partial {
				tmp := lastPartial[1:]
				// backtrack
				if ok, okW := reverseTrie.find(tmp); ok && len(tmp) > 0 {
					lastPartial = tmp
					if okW {
						lastDigit = lastPartial
					}
				} else {
					lastPartial = string(in[j])
					if _, ok := reverseTrie.find(lastPartial); ok {
						lastDigit = lastPartial
					}
				}
			} else if word {
				lastDigit = lastPartial
			}
		}
		if firstDigit != "" && lastDigit != "" {
			break
		}
	}

	if len(firstDigit) > 1 {
		firstDigit = stringToDigit[firstDigit]
	}
	if len(lastDigit) > 1 {
		lastDigit = stringToDigit[stringReverse(lastDigit)]
	}
	out, err := strconv.Atoi(firstDigit + lastDigit)
	aoc.Must(err)
	return out
}

func createTrie() (*trie, *trie) {
	t := newTrie()
	reverseT := newTrie()
	for k, v := range stringToDigit {
		t.insert(k)
		reverseT.insert(stringReverse(k))
		t.insert(v)
		reverseT.insert(v)
	}
	return t, reverseT
}

func stringReverse(s string) string {
	var b bytes.Buffer
	for i := len(s) - 1; i >= 0; i-- {
		b.WriteByte(s[i])
	}
	return b.String()
}
