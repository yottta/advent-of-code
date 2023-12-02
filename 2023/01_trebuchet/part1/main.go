package main

import (
	"fmt"
	"strconv"
	"unicode"

	aoc "github.com/yottta/advent-of-code/00_aoc"
)

func main() {
	content, err := aoc.ReadFile("./input.txt")
	aoc.Must(err)

	var sum int
	for _, line := range content {
		sum += findCalibrationValue(line)
	}
	fmt.Println(sum)
}

func findCalibrationValue(in string) int {
	var firstDigit, lastDigit string
	for i, j := 0, len(in)-1; i < len(in) && j >= 0; i, j = i+1, j-1 {
		if firstDigit == "" && unicode.IsDigit(rune(in[i])) {
			firstDigit = string(in[i])
		}
		if lastDigit == "" && unicode.IsDigit(rune(in[j])) {
			lastDigit = string(in[j])
		}
		if firstDigit != "" && lastDigit != "" {
			break
		}
	}

	out, err := strconv.Atoi(firstDigit + lastDigit)
	aoc.Must(err)
	return out
}
