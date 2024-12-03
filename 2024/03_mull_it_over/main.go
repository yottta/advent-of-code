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
	c := stateMachine{
		startMatcher: runeMatcher('m', runeMatcher('u', runeMatcher('l', runeMatcher('(', numberAndComma(numberAndEnding()))))),
	}
	for _, line := range content {
		for _, s := range line {
			c.consume(s)
		}
	}

	var sum int
	for _, s := range c.res {
		sum += executeMul(s)
	}
	fmt.Println(sum)
}

func part2(content []string) {
	c := stateMachine{
		startMatcher: or(
			runeMatcher('m', runeMatcher('u', runeMatcher('l', runeMatcher('(', numberAndComma(numberAndEnding()))))),
			runeMatcher('d', runeMatcher('o', or(
				runeMatcher('n', runeMatcher('\'', runeMatcher('t', runeMatcher('(', lockOnRune(')'))))),
				runeMatcher('(', unlockOnRune(')')))),
			),
		),
	}
	for _, line := range content {
		for _, s := range line {
			c.consume(s)
		}
	}

	var sum int
	for _, s := range c.res {
		sum += executeMul(s)
	}
	fmt.Println(sum)
}

var r = strings.NewReplacer("mul(", "", ")", "")

func executeMul(mul string) int {
	res := r.Replace(mul)
	parts := strings.Split(res, ",")
	x, err := strconv.Atoi(parts[0])
	aoc.Must(err)
	y, err := strconv.Atoi(parts[1])
	aoc.Must(err)
	return x * y
}

type matcherF func(r rune) (stateMachineCmd, matcherF)

func or(matchers ...matcherF) matcherF {
	return func(r rune) (stateMachineCmd, matcherF) {
		for _, m := range matchers {
			if res, next := m(r); res == 1 {
				return res, next
			}
		}
		return stateMachineCmdReset, nil
	}
}

func runeMatcher(exp rune, next matcherF) matcherF {
	return func(r rune) (stateMachineCmd, matcherF) {
		if exp == r {
			return stateMachineCmdStoreToBuffer, next
		}
		return 0, nil
	}
}

func numberAndComma(next matcherF) matcherF {
	return func(r rune) (stateMachineCmd, matcherF) {
		if unicode.IsDigit(r) {
			return stateMachineCmdStoreToBuffer, nil
		}
		if r == ',' {
			return stateMachineCmdStoreToBuffer, next
		}
		return stateMachineCmdReset, nil
	}
}

func numberAndEnding() matcherF {
	return func(r rune) (stateMachineCmd, matcherF) {
		if unicode.IsDigit(r) {
			return stateMachineCmdStoreToBuffer, nil
		}
		if r == ')' {
			return stateMachineCmdWriteAndReset, nil
		}
		return stateMachineCmdReset, nil
	}
}

func lockOnRune(exp rune) matcherF {
	return func(r rune) (stateMachineCmd, matcherF) {
		if exp == r {
			return stateMachineCmdLock, nil
		}
		return stateMachineCmdReset, nil
	}
}

func unlockOnRune(exp rune) matcherF {
	return func(r rune) (stateMachineCmd, matcherF) {
		if exp == r {
			return stateMachineCmdUnlock, nil
		}
		return stateMachineCmdReset, nil
	}
}

type stateMachineCmd int

const (
	stateMachineCmdReset = iota
	stateMachineCmdStoreToBuffer
	stateMachineCmdWriteAndReset
	stateMachineCmdLock
	stateMachineCmdUnlock
)

type stateMachine struct {
	startMatcher matcherF
	nextMatcher  matcherF

	buf bytes.Buffer
	res []string

	locked bool
}

// 0 stopped; 1 running; 2 finished;
func (sm *stateMachine) consume(r rune) {
	if sm.nextMatcher == nil {
		sm.nextMatcher = sm.startMatcher
	}
	res, next := sm.nextMatcher(r)
	switch res {
	case stateMachineCmdReset:
		sm.buf.Reset()
		sm.nextMatcher = sm.startMatcher
	case stateMachineCmdStoreToBuffer:
		sm.buf.WriteRune(r)
		if next != nil {
			sm.nextMatcher = next
		}
	case stateMachineCmdWriteAndReset:
		sm.buf.WriteRune(r)
		if sm.locked {
			sm.nextMatcher = sm.startMatcher
			sm.buf.Reset()
			break
		}
		sm.res = append(sm.res, sm.buf.String())
		sm.buf.Reset()
		sm.nextMatcher = sm.startMatcher
	case stateMachineCmdLock:
		sm.buf.Reset()
		sm.nextMatcher = sm.startMatcher
		sm.locked = true
	case stateMachineCmdUnlock:
		sm.buf.Reset()
		sm.nextMatcher = sm.startMatcher
		sm.locked = false
	}
}
