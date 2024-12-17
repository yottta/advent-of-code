package main

import (
	"fmt"
	"strconv"
	"strings"

	aoc "github.com/yottta/advent-of-code/00_aoc"
	aoc2024 "github.com/yottta/advent-of-code/2024"
)

func main() {
	aoc2024.Run(part1, part2, false, "11")
}

func part1(content []string) {
	pebs := parseContent(content[0])
	var sum uint64
	for _, peb := range pebs {
		sum += transform(peb, 25)
	}

	fmt.Println(sum)
}

func part2(content []string) {
	return
	pebs := parseContent(content[0])
	var sum uint64
	for _, peb := range pebs {
		sum += transform(peb, 75)
	}

	fmt.Println(sum)
}

func parseContent(content string) []uint64 {
	var (
		pebs []uint64
	)
	for _, p := range strings.Split(content, " ") {
		val, err := strconv.ParseUint(p, 10, 0)
		aoc.Must(err)
		pebs = append(pebs, val)
	}
	return pebs
}
