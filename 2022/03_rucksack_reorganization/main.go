package main

import (
	"fmt"

	aoc "github.com/yottta/advent-of-code/00_aoc"
)

func main() {
	aoc.BasicRun(part1, part2)
}

func part1(content []string) {
	var total int
	for _, line := range content {
		firstCompartment := make(map[uint8]struct{})
		fmt.Println(line)
		for i := 0; i < len(line)/2; i++ {
			item := line[i]
			firstCompartment[item] = struct{}{}
		}
		common := make(map[uint8]struct{})
		for i := len(line) / 2; i < len(line); i++ {
			_, ok := firstCompartment[line[i]]
			if ok {
				common[line[i]] = struct{}{}
			}
		}
		var rt int
		for item := range common {
			itemPrio := priority(item)
			fmt.Println("common ", string(item), itemPrio)
			rt += itemPrio
		}
		total += rt
		fmt.Println()
	}
	fmt.Println(total)
}

func part2(fileContent []string) {
	var (
		total       int
		commonGroup map[uint8]struct{}
	)
	for ln, line := range fileContent {
		content := make(map[uint8]struct{})
		for i := 0; i < len(line); i++ {
			content[line[i]] = struct{}{}
		}

		if ln%3 == 0 {
			total += groupItemsSum(commonGroup)
			commonGroup = content
			continue
		}
		// remove items from the common group that are not found in the current rucksack
		for c := range commonGroup {
			_, ok := content[c]
			if !ok {
				delete(commonGroup, c)
			}
		}
	}
	total += groupItemsSum(commonGroup)
	fmt.Println(total)
}

func groupItemsSum(commonGroup map[uint8]struct{}) int {
	var gc int
	for item := range commonGroup {
		itemPrio := priority(item)
		gc += itemPrio
	}
	return gc
}

func priority(r uint8) int {
	if r >= 97 {
		return int(r - 96) // lower case
	}
	return int(r - 38) // upper case
}
