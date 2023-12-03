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
}
