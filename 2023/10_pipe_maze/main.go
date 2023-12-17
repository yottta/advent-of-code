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

type tile struct {
	x, y            int
	value           rune
	strVal          string
	next            *tile
	previous        *tile
	explodedElement bool
	outerTile       bool
	counted         bool
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

func findPathToTileVal(startPoint *tile, m [][]*tile, target rune) int {
	firstTiles := validTilesToGo(m, startPoint, startPoint.previous)
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
			row = append(row, &tile{x: x, y: y, value: c, strVal: string(c)})
			if c == startChar {
				startTile = row[x]
			}
		}
		m = append(m, row)
	}
	return m, startTile
}

func validTilesToGo(tilesMap [][]*tile, point *tile, exclude ...*tile) []*tile {
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

// 564 too high
func part2(content []string) {
	tilesMap, startPoint := generateTiles(content, start)

	printMap(tilesMap)
	explodedMap, startPoint := explode(tilesMap)

	fmt.Println("Exploded")
	printExplodedMap(explodedMap, true)

	depth := findPathToTileVal(startPoint, explodedMap, start)
	fmt.Println("Path: ", depth/2)
	printPath(startPoint, start)

	elementsInPath := pathToMap(startPoint, start)

	floodOuterTiles(explodedMap, elementsInPath)

	fmt.Println("Flooded")
	innerElements := countInnerNonExploded(explodedMap, elementsInPath)
	printExplodedMap(explodedMap, false)

	fmt.Println("Part 2 ", innerElements)
}

func explode(m [][]*tile) ([][]*tile, *tile) {
	exploded := make([][]*tile, len(m)*2+1)
	var startPoint *tile
	var explodedY int
	exploded[explodedY] = explodedRow(len(m[0])*2+1, explodedY)
	explodedY++
	for y := 0; y < len(m); y++ {
		prevRow := exploded[explodedY-1]
		bottomRow := explodedRow(len(m[0])*2+1, explodedY+1)

		// explode the row elements
		currentRow := make([]*tile, len(m[0])*2+1)
		var explodedX int
		currentRow[explodedX] = &tile{x: explodedX, y: explodedY, value: '.', strVal: ".", explodedElement: true}
		explodedX++
		for x := 0; x < len(m[y]); x++ {
			xExplodedChar := '.'
			yExplodedChar := '.'
			// left changes
			if m[y][x].value == '7' || m[y][x].value == 'J' {
				currentRow[explodedX-1].value = '-'
				currentRow[explodedX-1].strVal = string('-')
			}
			if m[y][x].value == '-' {
				xExplodedChar = '-'
				currentRow[explodedX-1].value = '-'
				currentRow[explodedX-1].strVal = string('-')
			}
			// right changes
			if m[y][x].value == 'F' || m[y][x].value == 'L' {
				xExplodedChar = '-'
			}
			// vertical changes (up and down)
			if m[y][x].value == '|' {
				yExplodedChar = '|'
				prevRow[explodedX].value = yExplodedChar
				prevRow[explodedX].strVal = string(yExplodedChar)
				bottomRow[explodedX].value = yExplodedChar
				bottomRow[explodedX].strVal = string(yExplodedChar)
			}
			if m[y][x].value == 'J' || m[y][x].value == 'L' {
				yExplodedChar = '|'
				prevRow[explodedX].value = yExplodedChar
				prevRow[explodedX].strVal = string(yExplodedChar)
			}
			if m[y][x].value == 'F' || m[y][x].value == '7' {
				yExplodedChar = '|'
				bottomRow[explodedX].value = yExplodedChar
				bottomRow[explodedX].strVal = string(yExplodedChar)
			}
			// real tile
			currentRow[explodedX] = &tile{
				x:      explodedX,
				y:      explodedY,
				value:  m[y][x].value,
				strVal: m[y][x].strVal,
			}
			if m[y][x].value == start {
				startPoint = &tile{
					x:      explodedX,
					y:      explodedY,
					value:  m[y][x].value,
					strVal: m[y][x].strVal,
				}
			}
			explodedX++
			// exploded tile
			currentRow[explodedX] = &tile{x: explodedX, y: explodedY, value: xExplodedChar, strVal: string(xExplodedChar), explodedElement: true}

			explodedX++
		}
		exploded[explodedY] = currentRow
		explodedY++
		exploded[explodedY] = bottomRow
		explodedY++
	}
	return exploded, startPoint
}

func explodedRow(size int, y int) []*tile {
	var row []*tile
	for x := 0; x < size; x++ {
		row = append(row, &tile{
			x: x, y: y, value: '.', strVal: ".", explodedElement: true,
		})
	}
	return row
}

func floodOuterTiles(m [][]*tile, outOfBounds map[string]struct{}) {
	floodTile(m, 0, 0, outOfBounds)
}

func floodTile(m [][]*tile, x, y int, outOfBounds map[string]struct{}) {
	for _, dir := range allDirections {
		xDelta, yDelta := dir.Points()
		x := x + xDelta
		y := y + yDelta
		if x < 0 || y < 0 || x >= len(m[0]) || y >= len(m) {
			continue
		}
		if m[y][x].outerTile {
			continue
		}
		if _, ok := outOfBounds[fmt.Sprintf("%d,%d", x, y)]; ok {
			continue
		}
		m[y][x].outerTile = true
		floodTile(m, x, y, outOfBounds)
	}
}

func pathToMap(start *tile, target rune) map[string]struct{} {
	out := make(map[string]struct{})
	current := start
	out[fmt.Sprintf("%d,%d", current.x, current.y)] = struct{}{}
	for current.next != nil && current.next.value != target {
		current = current.next
		out[fmt.Sprintf("%d,%d", current.x, current.y)] = struct{}{}
	}
	return out
}

func countInnerNonExploded(m [][]*tile, pipeNotToCount map[string]struct{}) int {
	var count int
	for _, row := range m {
		for _, t := range row {
			if t.outerTile {
				continue
			}
			if t.explodedElement {
				continue
			}
			if _, ok := pipeNotToCount[fmt.Sprintf("%d,%d", t.x, t.y)]; ok {
				continue
			}
			t.counted = true
			count++
		}
	}
	return count
}

// general printing
func printMap(m [][]*tile) {
	for _, row := range m {
		var printed bool
		for _, t := range row {
			printed = true
			fmt.Printf("%c", t.value)
		}
		if printed {
			fmt.Println()
		}
	}
}
func printExplodedMap(m [][]*tile, includeExploded bool) {
	for _, row := range m {
		var printed bool
		for _, t := range row {
			if t.explodedElement && !includeExploded {
				//fmt.Printf("%c", 'x')
				continue
			}
			if t.outerTile {
				printed = true
				fmt.Printf("%c", 'O')
				continue
			}
			if t.counted {
				printed = true
				fmt.Printf("%c", 'X')
				continue
			}
			printed = true
			fmt.Printf("%c", t.value)
		}
		if printed {
			fmt.Println()
		}
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
	fmt.Printf("%c", current.next.value)
	fmt.Println()
}
