package main

import (
	"fmt"

	_0_aoc "github.com/yottta/advent-of-code/00_aoc"
	"github.com/yottta/advent-of-code/00_aoc/queue"
	aoc2024 "github.com/yottta/advent-of-code/2024"
)

func main() {
	aoc2024.Run(part1, part2, false, "12")
}

func part1(content []string) {
	m := parseContent(content)
	//printMap(m)
	toBeProcessed := queue.New[*plot]()
	toBeProcessed.Push(m[0][0])
	for {
		next, ok := toBeProcessed.Pop()
		if !ok {
			break
		}
		walk(next, m, toBeProcessed)
	}
	noRegions := markRegions(m)
	fmt.Printf("found %d regions\n", noRegions)
	surfaces := gatherSurfaces(m)
	var sum int
	for _, s := range surfaces {
		price := s.area * s.perim
		//fmt.Printf("%s (area: %d; perim: %d) price is %d\n", s.plant, s.area, s.perim, price)
		sum += price
	}
	fmt.Println(sum)
}

func part2(content []string) {
}

type regionSurface struct {
	plant string // NOTE: just for debugging
	perim int
	area  int
}

func walk(p *plot, m [][]*plot, q queue.Queue[*plot]) {
	if p.processed {
		return
	}
	// check perimeter
	// NOTE: area is calculated separately
	if p.coord.X == 0 || p.coord.X == len(m[0])-1 {
		p.perimeter++
	}
	if p.coord.Y == 0 || p.coord.Y == len(m)-1 {
		p.perimeter++
	}
	p.processed = true
	adjacentPlots := p.allAdjacentPlots(m)
	for _, adj := range adjacentPlots {
		if adj.plant == p.plant {
			p.nextInRegion = append(p.nextInRegion, adj)
		}
		if adj.processed {
			continue
		}
		q.Push(adj)
		if p.plant != adj.plant {
			p.perimeter++
			adj.perimeter++
			continue
		}
	}
}

func markRegions(m [][]*plot) int {
	lastRegionId := 1
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			regionStartPlot := m[y][x]
			if regionStartPlot.regionId != 0 {
				continue
			}
			regionStartPlot.regionId = lastRegionId
			lastRegionId++
			nextInRegion := regionStartPlot.nextInRegion
			for len(nextInRegion) > 0 {
				var collectNextInRegion []*plot
				for _, n := range nextInRegion {
					if n.regionId != 0 {
						continue
					}
					n.regionId = regionStartPlot.regionId
					collectNextInRegion = append(collectNextInRegion, n.nextInRegion...)
				}
				nextInRegion = collectNextInRegion
			}
		}
	}
	return lastRegionId - 1
}

func gatherSurfaces(m [][]*plot) map[int]regionSurface {
	res := map[int]regionSurface{}
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			p := m[y][x]
			s := res[p.regionId]
			s.area++
			s.perim += p.perimeter
			s.plant = p.plant
			res[p.regionId] = s
		}
	}
	return res
}

func parseContent(content []string) [][]*plot {
	out := make([][]*plot, len(content))

	for y := 0; y < len(content); y++ {
		row := make([]*plot, len(content[y]))
		for x := 0; x < len(content[y]); x++ {
			row[x] = &plot{
				coord: _0_aoc.Point{
					X: x,
					Y: y,
				},
				plant: string(content[y][x]),
			}
		}
		out[y] = row
	}
	return out
}

func printMap(m [][]*plot) {
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			fmt.Print(m[y][x].plant)
		}
		fmt.Println()
	}
}
