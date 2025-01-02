package main

import (
	"fmt"
	"slices"

	_0_aoc "github.com/yottta/advent-of-code/00_aoc"
)

type perimWalker struct {
	prevStep _0_aoc.Point
	// visited is for the visited positions. The value is actually a list of deltas to know in which direction a position was visited
	visited map[_0_aoc.Point][]_0_aoc.Point
	// seen is for the seen plots (not walker positions)
	seen            map[_0_aoc.Point]struct{}
	currentPosition _0_aoc.Point
	deltaX, deltaY  int
	turns           int
	touchedRegion   bool
	touchedPos      _0_aoc.Point
}

type walkerOpt func(walker *perimWalker)

func withInitialPos(point _0_aoc.Point) walkerOpt {
	return func(walker *perimWalker) {
		walker.currentPosition = point
	}
}

func newWalker(opts ...walkerOpt) *perimWalker {
	p := &perimWalker{
		visited: map[_0_aoc.Point][]_0_aoc.Point{},
		seen:    map[_0_aoc.Point]struct{}{},
		currentPosition: _0_aoc.Point{
			X: 0,
			Y: 0,
		},
		deltaX:        1,
		turns:         0,
		touchedRegion: false,
	}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

func (pw *perimWalker) turn(m [][]*plot) {
	if pw.touchedRegion { // count turns only when it already is around the region
		pw.turns++
	}
	// always try to turn right, if the one at the right was ever seen, turn left
	turn := 1 // 1 - right, 2 - left
	rightPlot, rightCoord := pw.plotAtTheRight(m)
	if rightPlot != nil {
		_, seen := pw.seen[rightCoord]
		if seen {
			turn = 2
		}
	} else {
		_, visited := pw.visited[rightCoord]
		if pw.prevStep == rightCoord && visited {
			turn = 2 // don't return
		}
	}
	switch turn {
	case 1: // turn right
		if pw.deltaX == 0 && pw.deltaY == -1 {
			pw.deltaX = 1
			pw.deltaY = 0
			return
		}
		if pw.deltaX == 1 && pw.deltaY == 0 {
			pw.deltaY = 1
			pw.deltaX = 0
			return
		}
		if pw.deltaX == 0 && pw.deltaY == 1 {
			pw.deltaX = -1
			pw.deltaY = 0
			return
		}
		if pw.deltaX == -1 && pw.deltaY == 0 {
			pw.deltaX = 0
			pw.deltaY = -1
			return
		}
	case 2: // turn left
		if pw.deltaX == 0 && pw.deltaY == -1 {
			pw.deltaX = -1
			pw.deltaY = 0
			return
		}
		if pw.deltaX == 1 && pw.deltaY == 0 {
			pw.deltaY = -1
			pw.deltaX = 0
			return
		}
		if pw.deltaX == 0 && pw.deltaY == 1 {
			pw.deltaX = 1
			pw.deltaY = 0
			return
		}
		if pw.deltaX == -1 && pw.deltaY == 0 {
			pw.deltaX = 0
			pw.deltaY = 1
			return
		}
	}
}

func (pw *perimWalker) plotAtTheRight(m [][]*plot) (*plot, _0_aoc.Point) {
	if pw.deltaX == 0 && pw.deltaY == -1 { // going up
		return m[pw.currentPosition.Y][pw.currentPosition.X+1], _0_aoc.Point{X: pw.currentPosition.X + 1, Y: pw.currentPosition.Y}
	}
	if pw.deltaX == 1 && pw.deltaY == 0 { // going right
		return m[pw.currentPosition.Y+1][pw.currentPosition.X], _0_aoc.Point{X: pw.currentPosition.X, Y: pw.currentPosition.Y + 1}
	}
	if pw.deltaX == 0 && pw.deltaY == 1 { // going down
		return m[pw.currentPosition.Y][pw.currentPosition.X-1], _0_aoc.Point{X: pw.currentPosition.X - 1, Y: pw.currentPosition.Y}
	}
	if pw.deltaX == -1 && pw.deltaY == 0 { // going left
		return m[pw.currentPosition.Y-1][pw.currentPosition.X], _0_aoc.Point{X: pw.currentPosition.X, Y: pw.currentPosition.Y - 1}
	}
	return nil, _0_aoc.Point{}
}

func (pw *perimWalker) plotAtTheLeft(m [][]*plot) *plot {
	if pw.deltaX == 0 && pw.deltaY == -1 { // going up
		return m[pw.currentPosition.Y][pw.currentPosition.X-1]
	}
	if pw.deltaX == 1 && pw.deltaY == 0 { // going right
		return m[pw.currentPosition.Y-1][pw.currentPosition.X]
	}
	if pw.deltaX == 0 && pw.deltaY == 1 { // going down
		return m[pw.currentPosition.Y][pw.currentPosition.X+1]
	}
	if pw.deltaX == -1 && pw.deltaY == 0 { // going left
		return m[pw.currentPosition.Y+1][pw.currentPosition.X]
	}
	return nil
}

func (pw *perimWalker) plotInFront(m [][]*plot) *plot {
	newX := pw.currentPosition.X + pw.deltaX
	newY := pw.currentPosition.Y + pw.deltaY
	if newY < 0 || newY >= len(m) {
		return nil
	}
	if newX < 0 || newX >= len(m[0]) {
		return nil
	}
	return m[newY][newX]
}

func (pw *perimWalker) moveForward() {
	pw.prevStep = pw.currentPosition // needed for not turning again into the same point but also being able to end the walking
	pw.addVisited()
	pw.currentPosition.X = pw.currentPosition.X + pw.deltaX
	pw.currentPosition.Y = pw.currentPosition.Y + pw.deltaY
}

func (pw *perimWalker) addVisited() {
	dirs := pw.visited[pw.currentPosition]
	direction := _0_aoc.Point{X: pw.deltaX, Y: pw.deltaY}
	if slices.Contains(dirs, direction) {
		return
	}
	dirs = append(dirs, direction)
	pw.visited[pw.currentPosition] = dirs
}

// return true when it's getting back to an already visited spot
func (pw *perimWalker) doStep(m [][]*plot) bool {
	frontPlot := pw.plotInFront(m)
	rightPlot, rightPlotCoord := pw.plotAtTheRight(m)
	if !pw.touchedRegion {
		pw.touchedRegion = rightPlot != nil || frontPlot != nil // because the maps are having the buffer around the regions, only at the right can find a plot
		if pw.touchedRegion {
			pw.touchedPos = pw.currentPosition
		}
		if frontPlot == nil {
			pw.moveForward()
			if pw.currentPosition.X >= len(m[0]) { // NOTE: this is for the inner walks when the spawn is right next to a right edge
				return true
			}
		}
		return false
	}

	if rightPlot != nil {
		pw.seen[rightPlotCoord] = struct{}{}
		if frontPlot != nil { // if front is busy, turn
			pw.addVisited()
			pw.turn(m)
			pw.seen[frontPlot.coord] = struct{}{}
		} else { // if nothing in front
			pw.moveForward()
		}
	} else {
		_, visited := pw.visited[rightPlotCoord]
		if rightPlotCoord == pw.prevStep && visited {
			pw.moveForward()
		} else {
			pw.addVisited()
			pw.turn(m)
		}
	}

	dirs := pw.visited[pw.currentPosition]
	samePosAsTouch := pw.currentPosition == pw.touchedPos && pw.deltaY == 0 && pw.deltaX == 1 // same position and same direction
	return slices.Contains(dirs, _0_aoc.Point{
		X: pw.deltaX,
		Y: pw.deltaY,
	}) && samePosAsTouch
}

func calculateRegionPerim(m [][]*plot) int {
	//printMap(m)
	outerWalker := newWalker()
	perim := runWalker(outerWalker, m)
	visitedInnerPos := map[_0_aoc.Point]struct{}{}
	for visitedP := range outerWalker.visited {
		visitedInnerPos[visitedP] = struct{}{}
	}
	for y := 1; y < len(m)-1; y++ { // 1 and -1 because I want to ignore the buffer area around
		for x := 1; x < len(m[y])-1; x++ { // 1 and -1 because I want to ignore the buffer area around
			p := m[y][x]
			if p != nil {
				continue
			}
			innerP := _0_aoc.Point{X: x, Y: y}
			if _, visited := visitedInnerPos[innerP]; visited {
				continue
			}
			visitedInnerPos[innerP] = struct{}{}
			innerW := newWalker(withInitialPos(innerP))
			runWalker(innerW, m)
			if mapContainsAnyFromMap(innerW.visited, outerWalker.visited) {
				continue
			}
			for visitedP := range innerW.visited {
				visitedInnerPos[visitedP] = struct{}{}
			}
			perim += innerW.turns
		}
	}
	return perim
}

func mapContainsAnyFromMap(m1, m2 map[_0_aoc.Point][]_0_aoc.Point) bool {
	for k := range m1 {
		if _, found := m2[k]; found {
			return true
		}
	}
	return false
}

func runWalker(w *perimWalker, m [][]*plot) int {
	for {
		// do the outer part
		if w.doStep(m) {
			break
		}
		printMapWithWalker(m, w)
		fmt.Println()
	}
	return w.turns
}

func calculateRegionArea(m [][]*plot) int {
	var total int
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			p := m[y][x]
			if p != nil {
				total++
			}
		}
	}
	return total
}

func printMapWithWalker(m [][]*plot, w *perimWalker) {
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			p := m[y][x]
			if w.currentPosition.X == x && w.currentPosition.Y == y {
				fmt.Print("#")
				continue
			}
			if p == nil {
				fmt.Print(".")
				continue
			}
			fmt.Print(m[y][x].plant)
		}
		fmt.Println()
	}
}
