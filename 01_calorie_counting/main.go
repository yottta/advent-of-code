package main

import (
	"flag"
	"fmt"
	"sort"
	"strconv"
	"strings"

	aoc "github.com/yottta/aoc2022/00_aoc"
)

func main() {
	var (
		dataFilePath string
		partToRun    string
	)
	flag.StringVar(&dataFilePath, "d", "./input.txt", "The path of the file containing the data for the current problem")
	flag.StringVar(&partToRun, "p", "1", "The part of the problem to run, in case the problem has more than one parts")
	flag.Parse()

	content, err := aoc.ReadFile(dataFilePath)
	aoc.Must(err)

	switch partToRun {
	case "1":
		part1(content)
	case "2":
		part2(content)
	default:
		panic(fmt.Errorf("no part '%s' configured", partToRun))
	}
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
