package main

import (
	"fmt"

	aoc2024 "github.com/yottta/advent-of-code/2024"
)

func main() {
	aoc2024.Run(part1, part2, false, "06")
}

func part1(content []string) {
	//return
	_map, sold := parseContent(content)
	res, _ := walk(_map, sold)
	//printMap(_map, sold)
	fmt.Println(len(res))

}

// 1656 - too high
func part2(content []string) {
	_map, sold := parseContent(content)
	soldBak := *sold
	res, _ := walk(_map, sold)
	sold = &soldBak
	var validObstacles int
	for pos := range res {
		testSold := soldBak
		if pos.x == soldBak.currentPosition.x && pos.y == soldBak.currentPosition.y {
			continue
		}
		_map[pos.y][pos.x].obstacle = true
		if _, ok := walk(_map, &testSold); !ok {
			validObstacles++
		}
		_map[pos.y][pos.x].obstacle = false

	}
	fmt.Println(validObstacles)
}

func walk(_map [][]position, sold *soldier) (map[position]struct{}, bool) {
	var (
		out    = map[position]struct{}{}
		passed = map[string]struct{}{}
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
		passedKey := fmt.Sprintf("%d#%d#%d#%d", sold.currentPosition.x, sold.currentPosition.y, sold.deltaX, sold.deltaY)
		if _, ok := passed[passedKey]; ok {
			return nil, false
		}
		// if it's an obstacle ahead, rotate and reiterate
		if _map[nextSoldPos.y][nextSoldPos.x].obstacle {
			sold.rotate()
			continue
		}
		sold.currentPosition = nextSoldPos
		passed[passedKey] = struct{}{} // register the position and the direction the soldier is passing the slot
		out[nextSoldPos] = struct{}{}  // register the place where the soldier was
	}
	return out, true
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

func printPath(_map [][]position, steps map[position]struct{}, sold soldier) {
	for x := 0; x < len(_map); x++ {
		fmt.Println()
		for y := 0; y < len(_map[x]); y++ {
			p := _map[x][y]
			val := "."
			if p.obstacle {
				val = "#"
			}
			if _, ok := steps[p]; ok {
				val = "x"
			}
			if p == sold.currentPosition {
				val = "^"
			}
			fmt.Printf("%s", val)
		}
	}
	fmt.Println()
}

type position struct {
	x, y     int
	obstacle bool
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
