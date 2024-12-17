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
	var sum int
	for i := 0; i < 6; i++ {
		pebs = transform(pebs)
	}

	for s, c := range pebs {
		fmt.Println(s, c)
		sum += c
	}

	fmt.Println(sum)
}

// 223767208675409 too low
func part2(content []string) {
	pebs := parseContent(content[0])
	var sum int
	for i := 0; i < 75; i++ {
		pebs = transform(pebs)
	}

	for s, c := range pebs {
		fmt.Println(s, c)
		sum += c
	}

	fmt.Println(sum)
}

func parseContent(content string) map[uint64]int {
	var (
		pebs = map[uint64]int{}
	)
	for _, p := range strings.Split(content, " ") {
		val, err := strconv.ParseUint(p, 10, 0)
		aoc.Must(err)
		exists := pebs[val]
		pebs[val] = exists + 1
	}
	return pebs
}

func transform(in map[uint64]int) map[uint64]int {
	temp := map[uint64]int{}
	for val, c := range in {
		if val == 0 {
			cc := temp[1]
			temp[1] = cc + c
			continue
		}
		if strVal := strconv.FormatUint(val, 10); len(strVal)%2 == 0 {
			currVal, err := strconv.ParseUint(strVal[:len(strVal)/2], 10, 0)
			aoc.Must(err)
			newVal, err := strconv.ParseUint(strVal[len(strVal)/2:], 10, 0)
			aoc.Must(err)
			currValC := temp[currVal]
			temp[currVal] = currValC + c
			newValC := temp[newVal]
			temp[newVal] = newValC + c
			continue
		}
		v := val * 2024
		ex := temp[v]
		temp[v] = c + ex
	}
	return temp
}
