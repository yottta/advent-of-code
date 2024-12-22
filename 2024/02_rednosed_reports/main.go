package main

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	aoc "github.com/yottta/advent-of-code/00_aoc"
	aoc2024 "github.com/yottta/advent-of-code/2024"
)

const (
	minAllowedDistance = 1
	maxAllowedDistance = 3
)

func main() {
	aoc2024.Run(part1, part2, false, "02")
}

func part1(content []string) {
	reports := parseContent(content, 0)
	var safe int
	for i, v := range reports {
		if v.safe() {
			slog.Debug("%d report is safe", i)
			safe++
		}
	}
	fmt.Println(safe)
}

// 569 too low
// 574 too low
func part2(content []string) {
	reports := parseContent(content, 1)
	var safe int
	for i, v := range reports {
		if v.safe() {
			slog.Debug("%d report is safe", i)
			safe++
		}
	}
	fmt.Println(safe)
}

func parseContent(content []string, allowedFailures int) []*report {
	reports := make([]*report, len(content))
	for ri, l := range content {
		var lvls []int
		rl := strings.Split(l, " ")
		for i := 0; i < len(rl); i++ {
			val, err := strconv.Atoi(rl[i])
			aoc.Must(err)
			lvls = append(lvls, val)
		}
		rep := &report{levels: lvls}
		rep.process(allowedFailures)
		reports[ri] = rep
	}
	return reports
}

type report struct {
	levels []int
	isSafe bool
}

func (r *report) print() {
	var b strings.Builder
	for i := 0; i < len(r.levels); i++ {
		b.WriteString(strconv.Itoa(r.levels[i]))
		b.WriteString(" ")
	}
	fmt.Println(b.String())
}

func (r *report) safe() bool {
	return r.isSafe
}

func (r *report) process(allowedFailures int) {
	for i := -1; i < len(r.levels); i++ {
		diffs := r.diffs(i)
		var (
			negative, positive, inRange int
		)
		for _, diff := range diffs {
			if diff < 0 {
				negative++
			} else if diff > 0 {
				positive++
			}
			if absDiff := abs(diff); absDiff >= minAllowedDistance && absDiff <= maxAllowedDistance {
				inRange++
			}
		}
		r.isSafe = (negative == len(diffs) || positive == len(diffs)) && inRange == len(diffs)
		if r.isSafe {
			return
		}
		if allowedFailures == 0 {
			return
		}
	}
}

func (r *report) diffs(removeIdx int) []int {
	alteredLevels := make([]int, len(r.levels))
	copy(alteredLevels, r.levels)
	if removeIdx >= 0 {
		alteredLevels = append(alteredLevels[:removeIdx], alteredLevels[removeIdx+1:]...)
	}
	var diffs []int
	for i := 0; i < len(alteredLevels)-1; i++ {
		diffs = append(diffs, alteredLevels[i]-alteredLevels[i+1])
	}
	return diffs
}

func abs[T int | int32 | int64](v1 T) T {
	if v1 >= 0 {
		return v1
	}
	return v1 * -1
}
