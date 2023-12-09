package main

import (
	"fmt"
	"strconv"
	"strings"

	aoc "github.com/yottta/advent-of-code/00_aoc"
)

func main() {
	aoc.BasicRun(part1, part2)
}

func part1(content []string) {
	races := parseRaces(content)
	res := 1
	for _, r := range races {
		res *= noOfRecordBreakingCombinations(r)
	}
	fmt.Println(res)
}

func noOfRecordBreakingCombinations(r race) int {
	var res int
	for i := 1; i < r.time; i++ {
		remainingTime := r.time - i
		finalDistance := i * remainingTime
		if finalDistance > r.distance {
			res++
			continue
		}
	}
	return res
}

func part2(content []string) {
	parasedRaces := parseRaces(content)
	var bTime, bDist strings.Builder
	for _, r := range parasedRaces {
		bTime.WriteString(fmt.Sprintf("%d", r.time))
		bDist.WriteString(fmt.Sprintf("%d", r.distance))
	}
	time, err := strconv.Atoi(bTime.String())
	aoc.Must(err)
	dist, err := strconv.Atoi(bDist.String())
	aoc.Must(err)
	// obviously, not the most performant one. I should instead check from both ends until I find the first record breaking combination
	// and then just subtract the number of combinations that are not record breaking
	fmt.Println(noOfRecordBreakingCombinations(race{time, dist}))
}

type race struct {
	time     int
	distance int
}

func parseRaces(content []string) []race {
	times := strings.Split(strings.TrimSpace(strings.TrimPrefix(content[0], "Time: ")), " ")
	var result []race
	for _, timeStr := range times {
		if strings.TrimSpace(timeStr) == "" {
			continue
		}
		time, err := strconv.Atoi(strings.TrimSpace(timeStr))
		aoc.Must(err)
		result = append(result, race{time: time})
	}
	distances := strings.Split(strings.TrimSpace(strings.TrimPrefix(content[1], "Distance: ")), " ")
	var idx int
	for _, distanceStr := range distances {
		if strings.TrimSpace(distanceStr) == "" {
			continue
		}
		distance, err := strconv.Atoi(strings.TrimSpace(distanceStr))
		aoc.Must(err)
		result[idx].distance = distance
		idx++
	}
	return result
}
