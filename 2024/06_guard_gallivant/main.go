package main

import (
	"fmt"

	aoc2024 "github.com/yottta/advent-of-code/2024"
)

func main() {
	aoc2024.Run(part1, part2, false, "06")
}

func part1(content []string) {
	_map, sold := parseContent(content)
	res := walk(_map, sold)
	//printMap(_map, sold)
	fmt.Println(len(res))

}
func part2(content []string) {
}

func walk(_map [][]position, sold *soldier) map[position]struct{} {
	var (
		out = map[position]struct{}{}
	)

	for {
		nextSoldPos := sold.nextPosition()
		// got out of the map to the left or right
		if nextSoldPos.x < 0 || nextSoldPos.x >= len(_map) {
			break
		}
		// got out of the map to the top or to the bottom
		if nextSoldPos.y < 0 || nextSoldPos.y >= len(_map[0]) {
			break
		}
		// if it's an obstacle ahead, rotate and reiterate
		if _map[nextSoldPos.y][nextSoldPos.x].obstacle {
			sold.rotate()
			continue
		}
		sold.currentPosition = nextSoldPos
		out[nextSoldPos] = struct{}{} // register the place where the soldier was
	}
	return out
}

func parseContent(content []string) ([][]position, *soldier) {
	var (
		_map [][]position
		sold *soldier
	)

	for y := 0; y < len(content); y++ {
		row := make([]position, 0, len(content[y]))
		for x := 0; x < len(content[y]); x++ {
			val := content[y][x]
			pos := position{
				x:        x,
				y:        y,
				obstacle: val == '#',
			}
			row = append(row, pos)
			if val == '^' {
				sold = &soldier{
					currentPosition: pos,
					deltaX:          0,
					deltaY:          -1,
				}
			}
		}
		_map = append(_map, row)
	}
	return _map, sold
}

func printMap(_map [][]position, sold *soldier) {
	fmt.Println(sold)
	for x := 0; x < len(_map); x++ {
		fmt.Println()
		for y := 0; y < len(_map[x]); y++ {
			p := _map[x][y]
			val := "."
			if p.obstacle {
				val = "#"
			}
			if p == sold.currentPosition {
				val = "S"
			}
			fmt.Printf("%s", val)
		}
	}
	fmt.Println()
}

type position struct {
	x, y     int
	obstacle bool
	passes   int
}

type soldier struct {
	currentPosition position
	deltaX          int
	deltaY          int
}

func (s *soldier) nextPosition() position {
	return position{
		x: s.currentPosition.x + s.deltaX,
		y: s.currentPosition.y + s.deltaY,
	}
}

func (s *soldier) rotate() {
	if s.deltaX == 0 && s.deltaY == -1 {
		s.deltaX = 1
		s.deltaY = 0
		return
	}
	if s.deltaX == 1 && s.deltaY == 0 {
		s.deltaY = 1
		s.deltaX = 0
		return
	}
	if s.deltaX == 0 && s.deltaY == 1 {
		s.deltaX = -1
		s.deltaY = 0
		return
	}
	if s.deltaX == -1 && s.deltaY == 0 {
		s.deltaX = 0
		s.deltaY = -1
		return
	}
}
