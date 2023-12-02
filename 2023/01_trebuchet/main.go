package main

import (
	"fmt"
	"strconv"
	"unicode"

	aoc "github.com/yottta/advent-of-code/00_aoc"
	"github.com/yottta/advent-of-code/00_aoc/trie"
)

func main() {
	aoc.BasicRun(part1, part2)
}

func part1(content []string) {
	var sum int
	for _, line := range content {
		var firstDigit, lastDigit string
		for i, j := 0, len(line)-1; i < len(line) && j >= 0; i, j = i+1, j-1 {
			if firstDigit == "" && unicode.IsDigit(rune(line[i])) {
				firstDigit = string(line[i])
			}
			if lastDigit == "" && unicode.IsDigit(rune(line[j])) {
				lastDigit = string(line[j])
			}
			if firstDigit != "" && lastDigit != "" {
				break
			}
		}

		out, err := strconv.Atoi(firstDigit + lastDigit)
		aoc.Must(err)
		sum += out
	}
	fmt.Println(sum)
}

func part2(content []string) {
	forwardTrie, reverseTrie := createTrie()
	// debug
	// good example why the "backtrack" is needed
	//fmt.Println(findCalibrationValue(forwardTrie, reverseTrie, "6jxgvknvbbbmqbjkkbnineninenj"))        // expected: 69; resulted:66

	var sum int
	for i, line := range content {
		value := findCalibrationValue(forwardTrie, reverseTrie, line)
		sum += value
		fmt.Println(i, line, value)
	}
	fmt.Println(sum)
}

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
func findCalibrationValue(forwardTrie, reverseTrie trie.Trie, in string) int {
	var firstDigit, firstPartial, lastPartial, lastDigit string

	for i, j := 0, len(in)-1; i < len(in) && j >= 0; i, j = i+1, j-1 {
		if firstDigit == "" {
			firstPartial += string(in[i])
			partial, word := forwardTrie.Find(firstPartial)
			if !partial {
				tmp := firstPartial[1:]
				// backtrack
				if ok, okW := forwardTrie.Find(tmp); ok && len(tmp) > 0 {
					firstPartial = tmp
					if okW {
						firstDigit = firstPartial
					}
				} else {
					firstPartial = string(in[i])
					if _, ok := forwardTrie.Find(firstPartial); ok {
						firstDigit = firstPartial
					}
				}
			} else if word {
				firstDigit = firstPartial
			}
		}

		if lastDigit == "" {
			lastPartial += string(in[j])
			partial, word := reverseTrie.Find(lastPartial)
			if !partial {
				tmp := lastPartial[1:]
				// backtrack
				if ok, okW := reverseTrie.Find(tmp); ok && len(tmp) > 0 {
					lastPartial = tmp
					if okW {
						lastDigit = lastPartial
					}
				} else {
					lastPartial = string(in[j])
					if _, ok := reverseTrie.Find(lastPartial); ok {
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
		lastDigit = stringToDigit[aoc.StringReverse(lastDigit)]
	}
	out, err := strconv.Atoi(firstDigit + lastDigit)
	aoc.Must(err)
	return out
}

func createTrie() (trie.Trie, trie.Trie) {
	t := trie.NewTrie()
	reverseT := trie.NewTrie()
	for k, v := range stringToDigit {
		t.Insert(k)
		reverseT.Insert(aoc.StringReverse(k))
		t.Insert(v)
		reverseT.Insert(v)
	}
	return t, reverseT
}
