package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	aoc "github.com/yottta/advent-of-code/00_aoc"
	aoc2024 "github.com/yottta/advent-of-code/2024"
)

func main() {
	aoc2024.Run(part1, part2, false, "05")
}

func part1(content []string) {
	ordering, updates := readInputInOrderingRulesAndUpdates(content)
	var validUpdates [][]int
	for _, update := range updates {
		if validUpdate(update, ordering) {
			validUpdates = append(validUpdates, update)
		}
	}
	var sum int
	for _, update := range validUpdates {
		fmt.Println(update)
		sum += update[len(update)/2]
	}
	fmt.Println(sum)
}

func validUpdate(update []int, ordering map[int][]int) bool {
	var failed int
	for i := 0; i < len(update); i++ {
		iPage := update[i]
		afterPageX := ordering[iPage]
		for j := i + 1; j < len(update); j++ {
			jPage := update[j]
			if slices.Contains(afterPageX, jPage) {
				continue
			}
			afterPageY := ordering[jPage]
			if slices.Contains(afterPageY, iPage) {
				failed++
				continue
			}
		}
	}
	return failed == 0
}

func part2(content []string) {
}

func readInputInOrderingRulesAndUpdates(content []string) (map[int][]int, [][]int) {
	var (
		processUpdates bool
		ordering       = map[int][]int{}
		updates        [][]int
	)
	for _, s := range content {
		if strings.TrimSpace(s) == "" {
			processUpdates = true
			continue
		}
		if !processUpdates {
			parseOrder(s, ordering)
			continue
		}
		updates = append(updates, parseUpdate(s))
	}
	return ordering, updates
}

func parseOrder(s string, alreadyExisting map[int][]int) {
	parts := strings.Split(s, "|")
	before, err := strconv.Atoi(parts[0])
	aoc.Must(err)
	after, err := strconv.Atoi(parts[1])
	aoc.Must(err)
	afterPages, _ := alreadyExisting[before]
	afterPages = append(afterPages, after)
	alreadyExisting[before] = afterPages
}

func parseUpdate(s string) []int {
	parts := strings.Split(s, ",")
	out := make([]int, len(parts))
	for i, page := range parts {
		p, err := strconv.Atoi(page)
		aoc.Must(err)
		out[i] = p
	}
	return out
}
