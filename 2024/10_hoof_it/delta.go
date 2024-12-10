package main

import aoc "github.com/yottta/advent-of-code/00_aoc"

type delta int

func (d delta) applyToPoint(p aoc.Point) aoc.Point {
	switch d {
	case deltaUp:
		return aoc.Point{
			X: p.X,
			Y: p.Y - 1,
		}
	case deltaRight:
		return aoc.Point{
			X: p.X + 1,
			Y: p.Y,
		}
	case deltaDown:
		return aoc.Point{
			X: p.X,
			Y: p.Y + 1,
		}
	case deltaLeft:
		return aoc.Point{
			X: p.X - 1,
			Y: p.Y,
		}
	}
	return p
}

const (
	deltaUp = iota + 1
	deltaRight
	deltaDown
	deltaLeft
)

var (
	allDeltas = []delta{
		deltaUp,
		deltaRight,
		deltaDown,
		deltaLeft,
	}
)
