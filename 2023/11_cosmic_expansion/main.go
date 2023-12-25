package main

import (
	"fmt"
	aoc "github.com/yottta/advent-of-code/00_aoc"
)

type table struct {
	elements     [][]element
	emptyRows    []int
	emptyColumns []int
	galaxies     []element
}

type element struct {
	x, y     int
	strVal   string
	idx      int
	visited  bool
	isGalaxy bool
}

func main() {
	aoc.BasicRun(part1, part2)
}

func part1(content []string) {
	inputTable := readInput(content)

	var sumOfDistances int
	for i := 0; i < len(inputTable.galaxies); i++ {
		for j := i + 1; j < len(inputTable.galaxies); j++ {
			dist := distanceBetweenTwoGalaxies(inputTable, inputTable.galaxies[i], inputTable.galaxies[j], 1)
			sumOfDistances += dist
		}
	}
	fmt.Println(sumOfDistances)
}

func part2(content []string) {
	inputTable := readInput(content)

	var sumOfDistances int
	for i := 0; i < len(inputTable.galaxies); i++ {
		for j := i + 1; j < len(inputTable.galaxies); j++ {
			dist := distanceBetweenTwoGalaxies(inputTable, inputTable.galaxies[i], inputTable.galaxies[j], 999_999)
			sumOfDistances += dist
		}
	}
	fmt.Println(sumOfDistances)
}

func readInput(content []string) table {
	out := table{}
	var galaxyCount int
	for y, line := range content {
		var rowContainsGalaxies bool
		var currentLine []element
		for x, char := range line {
			el := element{
				x:      x,
				y:      y,
				strVal: string(char),
				idx:    -1,
			}
			if el.strVal == "#" {
				galaxyCount++
				el.isGalaxy = true
				el.idx = galaxyCount
				rowContainsGalaxies = true
				out.galaxies = append(out.galaxies, el)
			}
			currentLine = append(currentLine, el)
		}
		if !rowContainsGalaxies {
			out.emptyRows = append(out.emptyRows, y)
		}
		out.elements = append(out.elements, currentLine)
	}
	for x := 0; x < len(out.elements[0]); x++ {
		var columnContainsGalaxies bool
		for y := 0; y < len(out.elements); y++ {
			if out.elements[y][x].isGalaxy {
				columnContainsGalaxies = true
				break
			}
		}
		if !columnContainsGalaxies {
			out.emptyColumns = append(out.emptyColumns, x)
		}
	}
	return out
}

func distanceBetweenTwoGalaxies(t table, from, to element, factor int) int {
	var noOfEmptyCols, numOfEmptyRows int
	fromX := min(from.x, to.x)
	fromY := min(from.y, to.y)
	toX := max(from.x, to.x)
	toY := max(from.y, to.y)
	for _, col := range t.emptyColumns {
		if col > fromX && col < toX {
			noOfEmptyCols++
		}
	}
	for _, row := range t.emptyRows {
		if row > fromY && row < toY {
			numOfEmptyRows++
		}
	}
	return distanceBetweenTwoPoints(fromX, fromY, toX+noOfEmptyCols*factor, toY+numOfEmptyRows*factor)
}

func distanceBetweenTwoPoints(x1, y1, x2, y2 int) int {
	return abs(x2-x1) + abs(y2-y1)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func printImage(image table) {
	for _, line := range image.elements {
		for _, el := range line {
			if el.isGalaxy {
				fmt.Print(el.idx)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
