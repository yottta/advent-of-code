package main

import _0_aoc "github.com/yottta/advent-of-code/00_aoc"

type plot struct {
	regionId        int
	coord           _0_aoc.Point
	plant           string
	area, perimeter int
	processed       bool

	nextInRegion []*plot
}

func (p *plot) allAdjacentPlots(m [][]*plot) []*plot {
	var res []*plot

	for _, point := range p.coord.AllAdjacent() {
		if !point.InRange(0, 0, len(m[0]), len(m)) {
			continue
		}
		res = append(res, m[point.Y][point.X])
	}
	return res
}
