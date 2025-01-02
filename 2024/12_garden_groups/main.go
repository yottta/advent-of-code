package main

import (
	"fmt"
	"log/slog"
	"math"

	_0_aoc "github.com/yottta/advent-of-code/00_aoc"
	"github.com/yottta/advent-of-code/00_aoc/queue"
	aoc2024 "github.com/yottta/advent-of-code/2024"
)

func main() {
	aoc2024.Run(part1, part2, false, "12")
}

func part1(content []string) {
	return
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
	noRegions := extractRegions(m)
	slog.Debug("found %d regions", noRegions)
	surfaces := gatherSurfaces(m)
	var sum int
	for _, s := range surfaces {
		price := s.area * s.perim
		slog.Debug("%s (area: %d; perim: %d) price is %d\n", s.plant, s.area, s.perim, price)
		sum += price
	}
	fmt.Println(sum)
}

func part2(content []string) {
	m := parseContent(content)
	toBeProcessed := queue.New[*plot]()
	toBeProcessed.Push(m[0][0])
	for {
		next, ok := toBeProcessed.Pop()
		if !ok {
			break
		}
		walk(next, m, toBeProcessed)
	}
	printMap(m)
	markRegions(m)
	regionMaps := extractRegions(m)
	regionsSurfaces := map[int]regionSurface{}
	for regId, regMap := range regionMaps {
		//fmt.Println(regId)
		perim := calculateRegionPerim(regMap)
		area := calculateRegionArea(regMap)
		regionsSurfaces[regId] = regionSurface{
			perim: perim,
			area:  area,
		}
	}

	var sum int
	for regId, s := range regionsSurfaces {
		price := s.area * s.perim
		fmt.Println("regionData:", "regionId", regId, "area", s.area, "perim", s.perim, "price", price)
		sum += price
	}
	fmt.Println(sum)
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
			fmt.Println("letter", regionStartPlot.plant, "regionId", lastRegionId)
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

type regionLimits struct {
	minX, minY, maxX, maxY int
}

func extractRegions(m [][]*plot) map[int][][]*plot {
	limits := map[int]regionLimits{}
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			p := m[y][x]
			regionLim, ok := limits[p.regionId]
			if !ok {
				regionLim = regionLimits{
					minX: math.MaxInt,
					minY: math.MaxInt,
					maxX: math.MinInt,
					maxY: math.MinInt,
				}
			}
			if y < regionLim.minY {
				regionLim.minY = y
			}
			if y > regionLim.maxY {
				regionLim.maxY = y
			}
			if x < regionLim.minX {
				regionLim.minX = x
			}
			if x > regionLim.maxX {
				regionLim.maxX = x
			}
			limits[p.regionId] = regionLim
		}
	}
	regionMaps := map[int][][]*plot{}
	for regionId, lims := range limits {
		mapXSize := lims.maxX - lims.minX + 1
		mapYSize := lims.maxY - lims.minY + 1
		regionMap := make([][]*plot, mapYSize+2)
		for i := 0; i < len(regionMap); i++ {
			regionMap[i] = make([]*plot, mapXSize+2)
		}
		// NOTE: put the region in its own man
		regionMapY := 0
		for y := lims.minY; y <= lims.maxY; y++ {
			regionMapY++
			regionMapX := 0
			for x := lims.minX; x <= lims.maxX; x++ {
				regionMapX++
				currP := m[y][x]
				if currP.regionId != regionId {
					continue
				}
				regionMap[regionMapY][regionMapX] = &plot{
					regionId: currP.regionId,
					coord: _0_aoc.Point{
						X: regionMapX,
						Y: regionMapY,
					},
					plant:     currP.plant,
					processed: false,
				}
			}
		}
		if regionId == 2 {
			regionMaps[regionId] = regionMap
			break
		}
	}
	return regionMaps
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
			p := m[y][x]
			if p == nil {
				fmt.Print(".")
				continue
			}
			fmt.Print(m[y][x].plant)
		}
		fmt.Println()
	}
}
