package main

import (
	"fmt"

	smLib "github.com/yottta/advent-of-code/00_aoc/state_machine"
	aoc2024 "github.com/yottta/advent-of-code/2024"
)

func main() {
	aoc2024.Run(part1, part2, false, "04")
}

func part1(content []string) {
	sm := smLib.New(
		smLib.RuneMatcher('X', smLib.RuneMatcher('M', smLib.RuneMatcher('A', smLib.StopOnRune('S')))),
		false,
	)
	m := toMatrix(content)

	ingest := func(r rune) {
		res := sm.Consume(r)
		switch res {
		case smLib.StateMachineCmdReset:
			sm.Consume(r)
		default:
			break
		}
	}

	// left to right
	for y := 0; y < len(m); y++ {
		sm.Consume('#') // force reset
		for x := 0; x < len(m[y]); x++ {
			r := m[y][x]
			ingest(r)
		}
	}

	// right to left
	for y := 0; y < len(m); y++ {
		sm.Consume('#') // force reset
		for x := len(m[y]) - 1; x >= 0; x-- {
			r := m[y][x]
			ingest(r)
		}
	}

	// top to bottom
	for x := 0; x < len(m[0]); x++ {
		sm.Consume('#') // force reset
		for y := 0; y < len(m); y++ {
			r := m[y][x]
			ingest(r)
		}
	}

	// bottom to top
	for x := 0; x < len(m[0]); x++ {
		sm.Consume('#') // force reset
		for y := len(m) - 1; y >= 0; y-- {
			r := m[y][x]
			ingest(r)
		}
	}

	////////////// All diagonals from top left to bottom right (and backwards)
	// top left - bottom right - bottom left corner (contains main diagonal)
	for y2 := 0; y2 < len(m)-3; y2++ {
		sm.Consume('#') // force reset
		for y, x := y2, 0; y < len(m) && x < len(m[y2]); y, x = y+1, x+1 {
			r := m[y][x]
			ingest(r)
		}
	}

	// top left - bottom right - top right corner (without main diagonal)
	for x2 := 1; x2 < len(m[0])-3; x2++ {
		sm.Consume('#') // force reset
		for y, x := 0, x2; y < len(m) && x < len(m[x2]); y, x = y+1, x+1 {
			r := m[y][x]
			ingest(r)
		}
	}

	// bottom right - top left - top right corner (contains main diagonal as well) (backwards of top left - bottom right - top right corner)
	for y2 := len(m) - 1; y2 >= 3; y2-- {
		sm.Consume('#') // force reset
		for y, x := y2, len(m[y2])-1; y >= 0 && x >= 0; y, x = y-1, x-1 {
			r := m[y][x]
			ingest(r)
		}
	}

	// bottom right - top left - bottom left corner (no main diagonal) (backwards of top left - bottom right - bottom left corner)
	for x2 := len(m[0]) - 2; x2 >= 3; x2-- {
		sm.Consume('#') // force reset
		for y, x := len(m)-1, x2; y >= 0 && x >= 0; y, x = y-1, x-1 {
			r := m[y][x]
			ingest(r)
		}
	}

	// top right - bottom left - bottom right corner (contains main diagonal)
	for y2 := 0; y2 < len(m)-3; y2++ {
		sm.Consume('#') // force reset
		for y, x := y2, len(m[y2])-1; y < len(m) && x >= 0; y, x = y+1, x-1 {
			r := m[y][x]
			ingest(r)
		}
	}

	// top right - bottom left - top left corner (without main diagonal)
	for x2 := len(m[0]) - 2; x2 >= 3; x2-- {
		sm.Consume('#') // force reset
		for y, x := 0, x2; y < len(m) && x >= 0; y, x = y+1, x-1 {
			r := m[y][x]
			ingest(r)
		}
	}

	//bottom left - top right - top left corner (contains main diagonal)
	for y2 := len(m) - 1; y2 >= 3; y2-- {
		sm.Consume('#') // force reset
		for y, x := y2, 0; y >= 0 && x < len(m[y2]); y, x = y-1, x+1 {
			r := m[y][x]
			ingest(r)
		}
	}

	// bottom left - top right - bottom right corner (no main diagonal)
	for x2 := 1; x2 < len(m[0])-3; x2++ {
		sm.Consume('#') // force reset
		for y, x := len(m)-1, x2; y >= 0 && x < len(m[0]); y, x = y-1, x+1 {
			r := m[y][x]
			ingest(r)
		}
	}
	fmt.Println(len(sm.Results()))
}

func part2(content []string) {
}

func toMatrix(content []string) [][]rune {
	matrix := make([][]rune, len(content))
	for y, line := range content {
		matrixCol := make([]rune, len(line))
		for x, c := range line {
			matrixCol[x] = c
		}
		matrix[y] = matrixCol
	}
	return matrix
}
