package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	aoc "github.com/yottta/advent-of-code/00_aoc"
)

func main() {
	aoc.BasicRun(part1, part2)
}

type mappingData struct {
	name     string
	mappings []mapping
}

type mapping struct {
	destinationRange int
	sourceRange      int
	rangeLength      int
}

type plantingAlmanac struct {
	seeds    []int
	mappings []mappingData
}

func part1(content []string) {
	alamanc := loadAlamanc(content)
	minLocation := math.MaxInt64
	for _, seed := range alamanc.seeds {
		newLoc := findValueMapping(seed, alamanc.mappings)
		if newLoc < minLocation {
			minLocation = newLoc
		}
	}
	fmt.Println(minLocation)
}

func findValueMapping(value int, mappings []mappingData) int {
	if len(mappings) == 0 {
		return value
	}
	for _, m := range mappings[0].mappings {
		if value >= m.sourceRange && value < m.sourceRange+m.rangeLength {
			newValue := m.destinationRange + max(value, m.sourceRange) - min(value, m.sourceRange)
			return findValueMapping(newValue, mappings[1:])
		}
	}
	return findValueMapping(value, mappings[1:])
}

func loadAlamanc(content []string) plantingAlmanac {
	almanac := plantingAlmanac{
		seeds: []int{},
	}

	var currentMapData mappingData
	for _, line := range content {
		if strings.TrimSpace(line) == "" {
			continue
		}
		if strings.HasPrefix(line, "seeds:") {
			almanac.seeds = parseSeeds(line)
			continue
		}
		if strings.HasSuffix(line, "map:") {
			if len(currentMapData.mappings) > 0 {
				almanac.mappings = append(almanac.mappings, currentMapData)
			}
			currentMapData = mappingData{
				name: strings.Split(line, " ")[0],
			}
			continue
		}
		mappings := strings.Split(line, " ")
		if len(mappings) != 3 {
			continue
		}
		destinationRange, err := strconv.Atoi(strings.TrimSpace(mappings[0]))
		aoc.Must(err)
		sourceRange, err := strconv.Atoi(strings.TrimSpace(mappings[1]))
		aoc.Must(err)
		rangeLength, err := strconv.Atoi(strings.TrimSpace(mappings[2]))
		aoc.Must(err)
		currentMapData.mappings = append(currentMapData.mappings, mapping{
			destinationRange: destinationRange,
			sourceRange:      sourceRange,
			rangeLength:      rangeLength,
		})
	}
	almanac.mappings = append(almanac.mappings, currentMapData)
	return almanac
}

func parseSeeds(line string) []int {
	var out []int
	seeds := strings.Split(strings.TrimPrefix(line, "seeds: "), " ")
	for _, seed := range seeds {
		seedNo, err := strconv.Atoi(strings.TrimSpace(seed))
		aoc.Must(err)
		out = append(out, seedNo)
	}
	return out
}

// 229165891 too high
func part2(content []string) {
	alamanc := loadAlamanc(content)
	minLocation := math.MaxInt64
	for i := 0; i < len(alamanc.seeds); i += 2 {
		lo := alamanc.seeds[i]
		plusRange := alamanc.seeds[i+1]
		newLoc := findRangeMapping(lo, plusRange, alamanc.mappings)
		if newLoc < minLocation {
			minLocation = newLoc
		}
	}
	fmt.Println(minLocation)
}

func findRangeMapping(lo, plusRange int, mappings []mappingData) int {
	if len(mappings) == 0 {
		return lo
	}
	currentMapping := mappings[0]
	mappings = mappings[1:]
	var results []int
	for _, m := range currentMapping.mappings {
		inLoBound := lo
		inHiBound := lo + plusRange - 1
		sourceLoBound := m.sourceRange
		sourceHiBound := m.sourceRange + m.rangeLength - 1
		if inLoBound > sourceHiBound || sourceLoBound >= inHiBound {
			continue // no overlap
		}
		loDiff := 0
		tempLo := max(sourceLoBound, inLoBound)
		tempHi := min(sourceHiBound, inHiBound)
		if tempLo == inLoBound {
			loDiff = inLoBound - sourceLoBound
		}
		newDestinationLo := m.destinationRange + loDiff
		newDestinationRange := tempHi - tempLo + 1
		results = append(results, findRangeMapping(newDestinationLo, newDestinationRange, mappings))
	}
	if len(results) == 0 {
		results = append(results, findRangeMapping(lo, plusRange, mappings))
	}
	slices.Sort(results)
	return results[0]
}
