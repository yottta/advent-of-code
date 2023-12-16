package main

import (
	"fmt"
	aoc "github.com/yottta/advent-of-code/00_aoc"
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

const start = 'S'

// - is a horizontal pipe connecting east and west.
// L is a 90-degree bend connecting north and east.
// J is a 90-degree bend connecting north and west.
// 7 is a 90-degree bend connecting south and west.
// F is a 90-degree bend connecting south and east.
// . is ground; there is no pipe in this tile.
// S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has
var possiblePipeMovement = map[rune]map[direction][]rune{
	'|': {
		up:   {'|', 'F', 'S', '7'},
		down: {'|', 'L', 'J', 'S'},
	},
	'-': {
		left:  {'-', 'F', 'S', 'L'},
		right: {'-', 'J', 'S', '7'},
	},
	'L': {
		up:    {'|', 'F', 'S', '7'},
		right: {'-', 'J', 'S', '7'},
	},
	'J': {
		up:   {'|', 'F', 'S', '7'},
		left: {'-', 'F', 'S', 'L'},
	},
	'7': {
		down: {'|', 'L', 'J', 'S'},
		left: {'-', 'F', 'S', 'L'},
	},
	'F': {
		down:  {'|', 'L', 'J', 'S'},
		right: {'-', 'J', 'S', '7'},
	},
	'S': {
		up:    {'|', 'F', '7'},
		down:  {'|', 'L', 'J'},
		left:  {'-', 'F', 'L'},
		right: {'-', 'J', '7'},
	},
}

func main() {
	aoc.BasicRun(part1, part2)
}

func part1(content []string) {
	tilesMap, startPoint := generateTiles(content, start)
	depth := findPathToTileVal(startPoint, tilesMap, start)
	fmt.Println(depth / 2)
	printPath(startPoint, start)
}

func part2(content []string) {
	tilesMap, startPoint := generateTiles(content, start)
	printMap(tilesMap)
	depth := findPathToTileVal(startPoint, tilesMap, start)
	printPath(startPoint, start)
	fmt.Println()
	fmt.Println(depth / 2)
	emptyMap := make([][]*tile, len(tilesMap))
	for i := range emptyMap {
		emptyMap[i] = make([]*tile, len(tilesMap[0]))
		for j := range emptyMap[i] {
			emptyMap[i][j] = &tile{j, i, '.', ".", nil, nil}
		}
	}
	emptyMap[startPoint.y][startPoint.x] = startPoint
	next := startPoint.next
	for {
		if next.value == start {
			break
		}
		emptyMap[next.y][next.x] = next
		next = next.next
	}
	printMap(emptyMap)

}

type tile struct {
	x, y     int
	value    rune
	strVal   string
	next     *tile
	previous *tile
}

func findPathToTileVal(startPoint *tile, m [][]*tile, target rune) int {
	firstTiles := findPossibleTiles(m, startPoint, startPoint.previous)
	var depth int
	for _, t := range firstTiles {
		startPoint.next = t
		t.previous = startPoint
		if t.value == target {
			return 1
		}
		res := findPathToTileVal(t, m, target)
		if res == -1 {
			startPoint.next = nil
			t.previous = nil
			continue
		}
		depth = res + 1
		break
	}

	return depth
}

func generateTiles(content []string, startChar rune) ([][]*tile, *tile) {
	var m [][]*tile
	var startTile *tile
	for y, line := range content {
		var row []*tile
		for x, c := range line {
			row = append(row, &tile{x, y, c, string(c), nil, nil})
			if c == startChar {
				startTile = row[x]
			}
		}
		m = append(m, row)
	}
	return m, startTile
}

func findPossibleTiles(tilesMap [][]*tile, point *tile, exclude ...*tile) []*tile {
	var possibleTiles []*tile
	for _, dir := range allDirections {
		xDelta, yDelta := dir.Points()
		x := point.x + xDelta
		y := point.y + yDelta
		if x < 0 || y < 0 || x >= len(tilesMap[0]) || y >= len(tilesMap) {
			continue
		}
		if tileExcluded(exclude, tilesMap[y][x]) {
			continue
		}
		if !isTileValidFromTile(dir, point, tilesMap[y][x]) {
			continue
		}

		possibleTiles = append(possibleTiles, tilesMap[y][x])
	}
	return possibleTiles
}

func tileExcluded(tiles []*tile, exclude *tile) bool {
	for _, t := range tiles {
		if t == nil {
			continue
		}
		if t.x == exclude.x && t.y == exclude.y {
			return true
		}
	}
	return false
}

func isTileValidFromTile(dir direction, from, to *tile) bool {
	possibleDirections, ok := possiblePipeMovement[from.value]
	if !ok {
		return false
	}
	possibleValues, ok := possibleDirections[dir]
	if !ok {
		return false
	}
	for _, v := range possibleValues {
		if v == to.value {
			return true
		}
	}

	return false
}

func printMap(m [][]*tile) {
	for _, row := range m {
		for _, t := range row {
			fmt.Printf("%c", t.value)
		}
		fmt.Println()
	}
}

func printPath(start *tile, target rune) {
	fmt.Printf("%c -> ", start.value)
	if start.next == nil {
		fmt.Println()
		fmt.Println("END")
		return
	}

	current := start
	for current.next != nil && current.next.value != target {
		current = current.next
		fmt.Printf("%c -> ", current.value)
	}
}
