package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	aoc "github.com/yottta/advent-of-code/00_aoc"
)

func main() {
	aoc.BasicRun(part1, part2)
}

// TODO use this instead
const (
	byteZero, byteNine = 48, 57
)

type point struct {
	x, y int
}

type piece struct {
	value int
	id    int
}

func part1(content []string) {
	seen := map[int]struct{}{}
	pieces := findValidPieces(content)
	var sum int
	for _, pc := range pieces {
		if _, ok := seen[pc.id]; ok {
			continue
		}
		seen[pc.id] = struct{}{}
		sum += pc.value
	}

	fmt.Println(sum)
}

func findValidPieces(content []string) map[point]piece {
	out := map[point]piece{}
	schematic := buildSchematic(content)
	currentPoint := point{}
	var id int
	for {
		number, start, stop := findNumber(schematic, currentPoint)
		if number < 0 {
			break
		}
		if isValidPiece(schematic, start, stop) {
			pc := piece{
				value: number,
				id:    id,
			}
			out[start] = pc
			out[stop] = pc
		}
		currentPoint = stop
		currentPoint.x++
		id++
	}
	return out
}

func findNumber(schematic [][]string, p point) (int, point, point) {
	var (
		number    int
		numBuf    bytes.Buffer
		start     = point{-1, -1}
		stop      = point{-1, -1}
		prevPoint = point{-1, -1}
	)
	for {
		if p.x >= len(schematic[p.y]) {
			p.x = 0
			p.y++
			if p.y >= len(schematic) {
				return -1, start, stop
			}
		}
		if _, err := strconv.Atoi(schematic[p.y][p.x]); err != nil { // TODO use bytes comparison here
			if start.x != -1 {
				stop = prevPoint
				break
			}
			p.x++
			continue
		}
		numBuf.WriteString(schematic[p.y][p.x])
		if start.x == -1 {
			start = p
		}
		prevPoint = p
		p.x++
	}
	number, err := strconv.Atoi(numBuf.String())
	aoc.Must(err)
	return number, start, stop
}

func isValidPiece(schematic [][]string, start, stop point) bool {
	// check top
	if start.y > 0 {
		for i := start.x - 1; i <= stop.x+1; i++ {
			if i < 0 || i >= len(schematic[start.y-1]) {
				continue
			}
			if _, err := strconv.Atoi(schematic[start.y-1][i]); err == nil {
				continue
			}
			if schematic[start.y-1][i] != "." {
				return true
			}
		}
	}
	// check bottom
	if start.y < len(schematic)-1 {
		for i := start.x - 1; i <= stop.x+1; i++ {
			if i < 0 || i >= len(schematic[start.y+1]) {
				continue
			}
			if _, err := strconv.Atoi(schematic[start.y+1][i]); err == nil {
				continue
			}
			if schematic[start.y+1][i] != "." {
				return true
			}
		}
	}
	// check right
	if stop.x < len(schematic[stop.y])-1 {
		// number cannot be due to previous parsing
		if schematic[stop.y][stop.x+1] != "." {
			return true
		}
	}
	// check left
	if start.x > 0 {
		// number cannot be due to previous parsing
		if schematic[start.y][start.x-1] != "." {
			return true
		}
	}
	return false
}
func buildSchematic(content []string) [][]string {
	var schematic [][]string
	for _, line := range content {
		schematic = append(schematic, strings.Split(line, ""))
	}
	return schematic
}

func part2(content []string) {
	pieces := findValidPieces(content)
	gears := findGears(content)
	fmt.Println(aoc.SumList(calculateGearRatio(pieces, gears)))
}

func findGears(content []string) []point {
	var gears []point
	for y, line := range content {
		for x, char := range line {
			if char == '*' {
				gears = append(gears, point{x, y})
			}
		}
	}
	return gears
}

func calculateGearRatio(pieces map[point]piece, gears []point) []int {
	var ratios []int
	var seenPieces = map[int]struct{}{}
	for _, gear := range gears {
		gearPieces := findAdjacentGearPieces(gear, pieces, seenPieces)
		if len(gearPieces) != 2 {
			continue
		}
		gearRation := 1
		for _, pc := range gearPieces {
			gearRation *= pc.value
		}
		ratios = append(ratios, gearRation)
	}
	return ratios
}

func findAdjacentGearPieces(gearPoint point, pieces map[point]piece, seenPieces map[int]struct{}) []piece {
	var out []piece
	// check top
	for i := gearPoint.x - 1; i <= gearPoint.x+1; i++ {
		if i < 0 {
			continue
		}
		if pc, ok := pieces[point{i, gearPoint.y - 1}]; ok {
			if _, ok := seenPieces[pc.id]; ok {
				continue
			}
			seenPieces[pc.id] = struct{}{}
			out = append(out, pc)
		}
	}
	// check bottom
	for i := gearPoint.x - 1; i <= gearPoint.x+1; i++ {
		if i < 0 {
			continue
		}
		if pc, ok := pieces[point{i, gearPoint.y + 1}]; ok {
			if _, ok := seenPieces[pc.id]; ok {
				continue
			}
			seenPieces[pc.id] = struct{}{}
			out = append(out, pc)
		}
	}
	// check right
	if pc, ok := pieces[point{gearPoint.x + 1, gearPoint.y}]; ok {
		if _, ok := seenPieces[pc.id]; !ok {
			seenPieces[pc.id] = struct{}{}
			out = append(out, pc)
		}

	}
	// check left
	if pc, ok := pieces[point{gearPoint.x - 1, gearPoint.y}]; ok {
		if _, ok := seenPieces[pc.id]; !ok {
			seenPieces[pc.id] = struct{}{}
			out = append(out, pc)
		}
	}
	return out
}
