package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	aoc "github.com/yottta/advent-of-code/00_aoc"
	aoc2024 "github.com/yottta/advent-of-code/2024"
)

func main() {
	aoc2024.Run(part1, part2, false, "01")
}

func part1(content []string) {
	left, right := parseContent(content)
	slices.Sort(left)
	slices.Sort(right)

	var sum int64
	for i := 0; i < len(left); i++ {
		sum += int64(math.Abs(float64(left[i] - right[i])))
	}
	fmt.Println(sum)
}

func part2(content []string) {
	left, right := parseContent(content)
	f := map[int]int{}
	for _, x := range right {
		c, _ := f[x]
		c++
		f[x] = c
	}
	var sum int64
	for _, v := range left {
		c, _ := f[v] 	
		sum += int64(v * c)
	}
	fmt.Println(sum)
}

func parseContent(content []string) ([]int, []int) {
	left, right := make([]int, len(content)), make([]int, len(content))
	for i, s := range content {
		sides := strings.Split(s, "  ")
		leftNo, err := strconv.Atoi(strings.TrimSpace(sides[0]))
		aoc.Must(err)
		rightNo, err := strconv.Atoi(strings.TrimSpace(sides[1]))
		aoc.Must(err)
		left[i] = leftNo
		right[i] = rightNo
	}

	return left, right
}
