package main

import (
	"bytes"
	"fmt"
	"strconv"

	aoc "github.com/yottta/advent-of-code/00_aoc"
)

func main() {
	aoc.BasicRun(part1, part2)
}

var shifts = []int{-1, 1}

const (
	byteZero, byteNine = '0', '9'
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

func findNumber(schematic [][]byte, p point) (int, point, point) {
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
		if schematic[p.y][p.x] < byteZero || schematic[p.y][p.x] > byteNine {
			if start.x != -1 {
				stop = prevPoint
				break
			}
			p.x++
			continue
		}
		numBuf.WriteByte(schematic[p.y][p.x])
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

func isValidPiece(schematic [][]byte, start, stop point) bool {
	// check top and bottom
	for _, shiftToRow := range shifts { // row shift
		for x := start.x - 1; x <= stop.x+1; x++ {
			if x < 0 || x >= len(schematic[0]) {
				continue
			}
			y := start.y + shiftToRow
			if y < 0 || y >= len(schematic) {
				continue
			}
			if schematic[y][x] >= byteZero && schematic[y][x] <= byteNine { // is number
				continue
			}
			if schematic[y][x] != '.' {
				return true
			}
		}
	}
	// check right
	if stop.x < len(schematic[stop.y])-1 {
		// number cannot be due to previous parsing
		if schematic[stop.y][stop.x+1] != '.' {
			return true
		}
	}
	// check left
	if start.x > 0 {
		// number cannot be due to previous parsing
		if schematic[start.y][start.x-1] != '.' {
			return true
		}
	}
	return false
}

func buildSchematic(content []string) [][]byte {
	var schematic [][]byte
	for _, line := range content {
		schematic = append(schematic, []byte(line))
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

	// check top and bottom
	for _, rowShift := range shifts { // row shift
		for x := gearPoint.x - 1; x <= gearPoint.x+1; x++ {
			if x < 0 {
				continue
			}
			y := gearPoint.y + rowShift
			if pc, ok := pieces[point{x, y}]; ok {
				if _, ok := seenPieces[pc.id]; ok {
					continue
				}
				seenPieces[pc.id] = struct{}{}
				out = append(out, pc)
			}
		}
	}
	// check left and right
	for _, columnShift := range shifts { // column shift
		x := gearPoint.x + columnShift
		y := gearPoint.y
		if pc, ok := pieces[point{x, y}]; ok {
			if _, ok := seenPieces[pc.id]; !ok {
				seenPieces[pc.id] = struct{}{}
				out = append(out, pc)
			}
		}
	}
	return out
}
