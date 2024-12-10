package main

import (
	"fmt"
	"maps"
	"strconv"

	aoc "github.com/yottta/advent-of-code/00_aoc"
	aoc2024 "github.com/yottta/advent-of-code/2024"
)

func main() {
	aoc2024.Run(part1, part2, false, "10")
}

func part1(content []string) {
	m, trailheads := parseMap(content)

	var sum int
	for _, trailhead := range trailheads {
		ends := trailhead.walk(m)
		sum += len(ends)
	}
	fmt.Println(sum)
}

func part2(content []string) {
}

func parseMap(content []string) (out [][]step, trailheads []step) {
	for y := 0; y < len(content); y++ {
		var row []step
		for x := 0; x < len(content[y]); x++ {
			v, err := strconv.Atoi(string(content[y][x]))
			aoc.Must(err)
			elems := step{
				p: aoc.Point{
					X: x,
					Y: y,
				},
				val: v,
			}
			row = append(row, elems)
			if elems.val == 0 {
				trailheads = append(trailheads, elems)
			}

		}
		out = append(out, row)
	}
	return out, trailheads
}

type step struct {
	p   aoc.Point
	val int
}

func (s *step) walk(m [][]step) map[step]struct{} {
	visitedPeaks := map[step]struct{}{}
	nextVal := s.val + 1
	for _, point := range possibleMovements(s.p, len(m[0]), len(m)) {
		nextStep := m[point.Y][point.X]
		if nextStep.val != nextVal {
			continue
		}
		if nextStep.val == 9 {
			visitedPeaks[nextStep] = struct{}{}
			continue
		}
		r := nextStep.walk(m)
		maps.Copy(visitedPeaks, r)
	}
	return visitedPeaks
}

func possibleMovements(p aoc.Point, maxX, maxY int) []aoc.Point {
	var nextSteps []aoc.Point
	for _, d := range allDeltas {
		if nextP := d.applyToPoint(p); validPoint(nextP, maxX, maxY) {
			nextSteps = append(nextSteps, nextP)
		}
	}
	return nextSteps
}

func validPoint(p aoc.Point, maxX, maxY int) bool {
	if p.X < 0 || p.X >= maxX {
		return false
	}
	if p.Y < 0 || p.Y >= maxY {
		return false
	}
	return true
}
