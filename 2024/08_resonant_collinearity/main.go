package main

import (
	"fmt"

	aoc "github.com/yottta/advent-of-code/00_aoc"
	aoc2024 "github.com/yottta/advent-of-code/2024"
)

func main() {
	aoc2024.Run(part1, part2, false, "08")
}

type antennasAntinode struct {
	a1, a2 aoc.Point
}

// 441 - too high
func part1(content []string) {
	return
	coords := antennasCoordinates(content)
	antinodes := calculateAntinodes(
		coords,
		aoc.Point{
			X: 0,
			Y: 0,
		},
		aoc.Point{
			X: len(content),
			Y: len(content[0]),
		},
		1,
	)
	//printMap(content, antinodes)
	//fmt.Println(antinodes)
	fmt.Println(len(antinodes))
}

// 1257 - too low
func part2(content []string) {
	coords := antennasCoordinates(content)
	antinodes := calculateAntinodes(
		coords,
		aoc.Point{
			X: 0,
			Y: 0,
		},
		aoc.Point{
			X: len(content),
			Y: len(content[0]),
		},
		1000, // Too high but better safe than sorry. It's breaking the loop nonetheless when it's getting out of the map
	)

	printMap(content, antinodes)
	fmt.Println(len(antinodes))
}

func antennasCoordinates(content []string) map[rune][]aoc.Point {
	out := map[rune][]aoc.Point{}
	for y := 0; y < len(content); y++ {
		for x, v := range content[y] {
			if v == '.' {
				continue
			}
			sameFreqAntennas := out[v]
			sameFreqAntennas = append(sameFreqAntennas, aoc.Point{
				X: x,
				Y: y,
			})
			out[v] = sameFreqAntennas
		}
	}
	return out
}

func calculateAntinodes(antennas map[rune][]aoc.Point, minP, maxP aoc.Point, maxAntinodes int) map[aoc.Point][]antennasAntinode {
	out := map[aoc.Point][]antennasAntinode{}
	for _, sameFreqAtennas := range antennas {
		for i, a1 := range sameFreqAtennas {
			for j := i + 1; j < len(sameFreqAtennas); j++ {
				a2 := sameFreqAtennas[j]
				for _, aa := range calculateAtennasAntinodes(a1, a2, minP, maxP, maxAntinodes) {
					alreadyAntinode := out[aa]
					alreadyAntinode = append(alreadyAntinode, antennasAntinode{
						a1: a1,
						a2: a2,
					})
					out[aa] = alreadyAntinode
				}
			}
		}
	}
	return out
}

func calculateAtennasAntinodes(a1, a2 aoc.Point, minP, maxP aoc.Point, maxAntinodes int) []aoc.Point {
	var res []aoc.Point

	dx := a2.X - a1.X
	dy := a2.Y - a1.Y

	anti1 := aoc.Point{
		X: a1.X - dx,
		Y: a1.Y - dy,
	}
	for i := 0; i < maxAntinodes; i++ {
		if !inBoundaries(minP, maxP, anti1) {
			break
		}
		// For the T example from the 2nd part, it doesn't work correctly. Still confused by the statements for this day
		if maxAntinodes > 10 { // HERE IS WHERE THE BUG IS
			res = append(res, a1)
		}
		res = append(res, anti1)
		anti1.X = anti1.X - dx
		anti1.Y = anti1.Y - dy
	}
	anti2 := aoc.Point{
		X: a2.X + dx,
		Y: a2.Y + dy,
	}
	for i := 0; i < maxAntinodes; i++ {
		if !inBoundaries(minP, maxP, anti2) {
			break
		}
		if maxAntinodes > 10 { // HERE IS WHERE THE BUG IS
			res = append(res, a2)
		}
		res = append(res, anti2)
		anti2.X = anti2.X + dx
		anti2.Y = anti2.Y + dy
	}
	return res
}

func printMapDebug(input []string, antinodes map[aoc.Point][]antennasAntinode) {
	fmt.Print("  ")
	for i := 0; i < len(input[0]); i++ {
		fmt.Printf("%3d", i)
	}
	fmt.Println()
	for y := 0; y < len(input); y++ {
		fmt.Printf("%2d", y)
		for x := 0; x < len(input[y]); x++ {
			if _, ok := antinodes[aoc.Point{
				X: x,
				Y: y,
			}]; ok {
				fmt.Printf("%3s", "#")
				continue
			}
			fmt.Printf("%3c", input[y][x])
		}
		fmt.Println()
	}
}

func printMap(input []string, antinodes map[aoc.Point][]antennasAntinode) {
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			if _, ok := antinodes[aoc.Point{
				X: x,
				Y: y,
			}]; ok {
				fmt.Printf("%s", "#")
				continue
			}
			fmt.Printf("%c", input[y][x])
		}
		fmt.Println()
	}
}

func inBoundaries(minP, maxP, p aoc.Point) bool {
	if p.X < minP.X {
		return false
	}
	if p.Y < minP.Y {
		return false
	}
	if p.X >= maxP.X {
		return false
	}
	if p.Y >= maxP.Y {
		return false
	}
	return true
}
