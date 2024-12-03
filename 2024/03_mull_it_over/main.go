package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	aoc "github.com/yottta/advent-of-code/00_aoc"
	aoc2024 "github.com/yottta/advent-of-code/2024"
)

func main() {
	aoc2024.Run(part1, part2, false, "03")
}

func part1(content []string) {
	var all []string
	for _, s := range content {
		all = append(all, extractor(s)...)
	}
	r := strings.NewReplacer("mul(", "", ")", "")
	var sum int
	for _, s := range all {
		res := r.Replace(s)
		parts := strings.Split(res, ",")
		x, err := strconv.Atoi(parts[0])
		aoc.Must(err)
		y, err := strconv.Atoi(parts[1])
		aoc.Must(err)
		sum += x * y
	}
	fmt.Println(sum)

}

func part2(content []string) {
}

func extractor(line string) []string {
	c := stateMachine{
		startMatcher: runeMatcher('m', runeMatcher('u', runeMatcher('l', runeMatcher('(', numberAndComma(numberAndEnding()))))),
	}
	for _, s := range line {
		c.consume(s)
	}
	return c.res
}

type matcherF func(r rune) (int, matcherF)

func runeMatcher(exp rune, next matcherF) matcherF {
	return func(r rune) (int, matcherF) {
		if exp == r {
			return 1, next
		}
		return 0, nil
	}
}

func numberAndComma(next matcherF) matcherF {
	return func(r rune) (int, matcherF) {
		if unicode.IsDigit(r) {
			return 1, nil
		}
		if r == ',' {
			return 1, next
		}
		return 0, nil
	}
}
func numberAndEnding() matcherF {
	return func(r rune) (int, matcherF) {
		if unicode.IsDigit(r) {
			return 1, nil
		}
		if r == ')' {
			return 2, nil
		}
		return 0, nil
	}
}

type stateMachine struct {
	startMatcher matcherF
	nextMatcher  matcherF

	buf bytes.Buffer
	res []string
}

// 0 stopped; 1 running; 2 finished
func (sm *stateMachine) consume(r rune) {
	if sm.nextMatcher == nil {
		sm.nextMatcher = sm.startMatcher
	}
	res, next := sm.nextMatcher(r)
	switch res {
	case 0:
		sm.buf.Reset()
		sm.nextMatcher = sm.startMatcher
	case 1:
		sm.buf.WriteRune(r)
		if next != nil {
			sm.nextMatcher = next
		}
	case 2:
		sm.buf.WriteRune(r)
		sm.res = append(sm.res, sm.buf.String())
		sm.buf.Reset()
		sm.nextMatcher = sm.startMatcher
	}
}
