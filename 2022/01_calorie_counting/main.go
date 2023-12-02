package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	aoc "github.com/yottta/advent-of-code/00_aoc"
)

func main() {
	aoc.BasicRun(part1, part2)
}

func part1(content []string) {
	var (
		curr        elf
		maxCalories elf
	)

	for _, line := range content {
		if len(strings.TrimSpace(line)) == 0 {
			if curr.calories > maxCalories.calories {
				maxCalories = curr
			}
			curr = elf{idx: curr.idx + 1}
			continue
		}
		atoi, err := strconv.Atoi(line)
		aoc.Must(err)
		curr.calories += atoi
	}
	fmt.Println(maxCalories)
}

func part2(content []string) {
	var (
		elves []elf
		curr  elf
	)
	for _, line := range content {
		if len(strings.TrimSpace(line)) == 0 {
			elves = append(elves, curr)
			curr = elf{idx: curr.idx + 1}
			continue
		}
		atoi, err := strconv.Atoi(line)
		aoc.Must(err)
		curr.calories += atoi
	}
	sort.Slice(elves, func(i, j int) bool {
		return elves[i].calories < elves[j].calories
	})
	var total int
	for i := len(elves) - 3; i < len(elves); i++ {
		total += elves[i].calories
	}
	fmt.Println(total)
}

type elf struct {
	idx      int
	calories int
}
