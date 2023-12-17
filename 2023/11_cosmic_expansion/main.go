package main

import (
	"fmt"
	aoc "github.com/yottta/advent-of-code/00_aoc"
	"github.com/yottta/advent-of-code/00_aoc/queue"
)

// x, y
type direction int

const (
	up direction = iota
	right
	down
	left
)

var (
	allDirections = []direction{up, right, down, left}
)

// Points returns the x and y values of the direction
func (d direction) Points() (int, int) {
	switch d {
	case up:
		return 0, -1
	case right:
		return 1, 0
	case down:
		return 0, 1
	case left:
		return -1, 0
	}
	return 0, 0
}

type table struct {
	elements     [][]element
	emptyRows    []int
	emptyColumns []int
	noOfGalaxies int
}

type element struct {
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
	printImage(inputTable)
	expandedTable := expand(inputTable, 10)
	fmt.Println("After expansion: ")
	printImage(expandedTable)

	// find min distance between all galaxies
	var sumOfDistances int
	for i := 0; i < expandedTable.noOfGalaxies; i++ {
		for j := i + 1; j < expandedTable.noOfGalaxies; j++ {
			distance := minDistance(expandedTable, i+1, j+1)
			//fmt.Printf("%d -> %d: %d\n", i+1, j+1, distance)
			sumOfDistances += distance
		}
	}
	fmt.Println("Part1: ", sumOfDistances)
}

func expand(in table, factor int) table {
	out := table{
		elements:     in.elements,
		noOfGalaxies: in.noOfGalaxies,
	}
	for i, targetColumn := range in.emptyColumns {
		replacement := out.elements
		var columExpansion []element
		for f := 0; f < factor; f++ {
			columExpansion = append(columExpansion, element{
				strVal: ".",
				idx:    -1,
			})
		}
		for rIdx := 0; rIdx < len(replacement); rIdx++ {
			res := make([]element, len(replacement[rIdx])+factor)
			copy(res, replacement[rIdx][:targetColumn+(i*factor)])
			copy(res[targetColumn+(i*factor):], columExpansion)
			copy(res[targetColumn+(i*factor)+factor:], replacement[rIdx][targetColumn+(i*factor):])
			replacement[rIdx] = res
		}
		out.elements = replacement
	}
	for i, targetRow := range in.emptyRows {
		rowsExpansion := make([][]element, factor)
		er := emptyRow(len(out.elements[0]))
		for e := range rowsExpansion {
			rowsExpansion[e] = er
		}
		res := make([][]element, len(out.elements)+factor)
		copy(res, out.elements[:targetRow+(i*factor)])
		copy(res[targetRow+(i*factor):], rowsExpansion)
		copy(res[targetRow+(i*factor)+factor:], out.elements[targetRow+(i*factor):])
		out.elements = res
	}
	return out
}

func emptyRow(size int) []element {
	var row []element
	for i := 0; i < size; i++ {
		row = append(row, element{
			strVal: ".",
			idx:    -1,
		})
	}
	return row
}

type target struct {
	e             element
	otherElements []target
	x, y          int
	dist          int
}

func minDistance(t table, from, to int) int {
	q := queue.New[target]()
	visited := make([][]bool, len(t.elements))
	for i := range visited {
		visited[i] = make([]bool, len(t.elements[0]))
	}
	// find source
	var source target
sourceFind:
	for y, line := range t.elements {
		for x, el := range line {
			if el.idx == from {
				source = target{
					x:    x,
					y:    y,
					dist: 0,
				}
				break sourceFind
			}
		}
	}
	q.Push(source)
	visited[source.y][source.x] = true
	for q.Length() > 0 {
		src, ok := q.Pop()
		if !ok {
			break
		}
		if t.elements[src.y][src.x].idx == to {
			return src.dist
		}
		for _, dir := range allDirections {
			dx, dy := dir.Points()
			deltaX := src.x + dx
			deltaY := src.y + dy
			// check if target is reachable
			if deltaY < 0 || deltaX < 0 || deltaY >= len(t.elements) || deltaX >= len(t.elements[0]) {
				continue
			}
			if visited[deltaY][deltaX] {
				continue
			}
			// check if target is reachable
			q.Push(target{
				x:    deltaX,
				y:    deltaY,
				dist: src.dist + 1,
			})
			visited[deltaY][deltaX] = true
		}
	}
	return -1
}

func part2(content []string) {

}

// utils
func readInput(content []string) table {
	out := table{}
	var lastGalaxy int
	for i, line := range content {
		var rowContainsGalaxies bool
		var currentLine []element
		for _, char := range line {
			el := element{
				strVal: string(char),
				idx:    -1,
			}
			if el.strVal == "#" {
				lastGalaxy++
				el.isGalaxy = true
				el.idx = lastGalaxy
				rowContainsGalaxies = true
			}
			currentLine = append(currentLine, el)
		}
		if !rowContainsGalaxies {
			out.emptyRows = append(out.emptyRows, i)
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
	out.noOfGalaxies = lastGalaxy
	return out
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
