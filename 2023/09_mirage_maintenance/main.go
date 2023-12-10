package main

import (
	"fmt"
	aoc "github.com/yottta/advent-of-code/00_aoc"
	"strconv"
	"strings"
)

type entryData struct {
	initialData   []int
	generatedRows [][]int
}

func main() {
	aoc.BasicRun(part1, part2)
}

func part1(content []string) {
	historyEntries := parseInput(content)
	var res int
	for _, entry := range historyEntries {
		genRow := extrapolations(entry)
		entry.generatedRows = genRow
		prevRowLastElement := generateNextHistoryEntry(entry)
		res += prevRowLastElement
	}
	fmt.Println("Part1: ", res)
}

func generateNextHistoryEntry(h entryData) int {
	allData := [][]int{h.initialData}
	allData = append(allData, h.generatedRows...)
	var prevRowLastElement int
	for i := len(allData) - 1; i >= 0; i-- {
		currentRowVal := allData[i][len(allData[i])-1]
		prevRowLastElement += currentRowVal
	}
	return prevRowLastElement
}

func generatePreviousHistoryEntry(h entryData) int {
	allData := [][]int{h.initialData}
	allData = append(allData, h.generatedRows...)
	var prevRowFirstElement int
	for i := len(allData) - 1; i >= 0; i-- {
		currentRowVal := allData[i][0]
		prevRowFirstElement = currentRowVal - prevRowFirstElement
	}
	return prevRowFirstElement
}

func extrapolations(h entryData) [][]int {
	data := h.initialData
	var out [][]int
	for !allZeroed(data) {
		newData := make([]int, len(data)-1)
		for i := 0; i < len(data)-1; i++ {
			newData[i] = data[i+1] - data[i]
		}
		data = newData
		out = append(out, data)
	}
	return out
}

func allZeroed(in []int) bool {
	for _, val := range in {
		if val != 0 {
			return false
		}
	}
	return true
}

func part2(content []string) {
	historyEntries := parseInput(content)
	var res int
	for _, entry := range historyEntries {
		genRow := extrapolations(entry)
		entry.generatedRows = genRow
		prevRowFirstElement := generatePreviousHistoryEntry(entry)
		res += prevRowFirstElement
	}
	fmt.Println("Part2: ", res)

}

func parseInput(content []string) []entryData {
	var data []entryData
	for _, line := range content {
		elements := strings.Split(line, " ")
		initialData := make([]int, len(elements))
		for i, element := range elements {
			val, err := strconv.Atoi(element)
			aoc.Must(err)
			initialData[i] = val
		}
		data = append(data, entryData{initialData: initialData})
	}
	return data
}
