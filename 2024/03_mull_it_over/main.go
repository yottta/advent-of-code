package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	aoc "github.com/yottta/advent-of-code/00_aoc"
	smpkg "github.com/yottta/advent-of-code/00_aoc/state_machine"
	aoc2024 "github.com/yottta/advent-of-code/2024"
)

func main() {
	aoc2024.Run(part1, part2, false, "03")
}

func part1(content []string) {
	rules := smpkg.RuneMatcher('m', smpkg.RuneMatcher('u', smpkg.RuneMatcher('l', smpkg.RuneMatcher('(', numberAndComma(numberAndEnding())))))
	sm := smpkg.New(rules, false)
	for _, line := range content {
		for _, s := range line {
			sm.Consume(s)
		}
	}

	var sum int
	for _, res := range sm.Results() {
		sum += executeMul(res)
	}
	fmt.Println(sum)
}

func part2(content []string) {
	rules := smpkg.Or(
		smpkg.RuneMatcher('m', smpkg.RuneMatcher('u', smpkg.RuneMatcher('l', smpkg.RuneMatcher('(', numberAndComma(numberAndEnding()))))),
		smpkg.RuneMatcher('d', smpkg.RuneMatcher('o', smpkg.Or(
			smpkg.RuneMatcher('n', smpkg.RuneMatcher('\'', smpkg.RuneMatcher('t', smpkg.RuneMatcher('(', smpkg.LockOnRune(')'))))),
			smpkg.RuneMatcher('(', smpkg.UnlockOnRune(')')))),
		),
	)
	sm := smpkg.New(rules, false)
	for _, line := range content {
		for _, s := range line {
			sm.Consume(s)
		}
	}

	var sum int
	for _, s := range sm.Results() {
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

func numberAndComma(next smpkg.MatcherFunc) smpkg.MatcherFunc {
	return func(r rune) (smpkg.StateMachineCmd, smpkg.MatcherFunc) {
		if unicode.IsDigit(r) {
			return smpkg.StateMachineCmdStoreToBuffer, nil
		}
		if r == ',' {
			return smpkg.StateMachineCmdStoreToBuffer, next
		}
		return smpkg.StateMachineCmdReset, nil
	}
}

func numberAndEnding() smpkg.MatcherFunc {
	return func(r rune) (smpkg.StateMachineCmd, smpkg.MatcherFunc) {
		if unicode.IsDigit(r) {
			return smpkg.StateMachineCmdStoreToBuffer, nil
		}
		if r == ')' {
			return smpkg.StateMachineCmdWriteAndReset, nil
		}
		return smpkg.StateMachineCmdReset, nil
	}
}
