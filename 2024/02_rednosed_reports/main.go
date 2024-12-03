package main

import (
	"fmt"
	"math"
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
	reports := parseContent(content)
	var safe int
	for i, v := range reports {
		if v.safe() {
			aoc.Log(fmt.Sprintf("%d report is safe", i))
			safe++
		}
	}
	fmt.Println(safe)
}

// 569 too low
// 574 too low
func part2(content []string) {
	reports := parseContent(content)
	var safe int
	for i, v := range reports {
		if v.safe() {
			aoc.Log(fmt.Sprintf("%d report is safe", i))
			safe++
			//} else {
			//	v.print()
		}
	}
	fmt.Println(safe)
}

func parseContent(content []string) []report {
	reports := make([]report, len(content))
	for ri, l := range content {
		var (
			prev *level
		)

		rl := strings.Split(l, " ")
		for i := len(rl) - 1; i >= 0; i-- {
			val, err := strconv.Atoi(rl[i])
			aoc.Must(err)
			lvl := &level{
				val:  val,
				next: prev,
			}
			prev = lvl
		}
		reports[ri] = report{first: prev}
	}
	return reports
}

type report struct {
	first *level
}

func (r report) print() {
	var b strings.Builder
	l := r.first
	for l != nil {
		b.WriteString(strconv.Itoa(l.val))
		b.WriteString(" ")
		l = l.next
	}
	fmt.Println(b.String())
}

func (r report) safe() bool {
	return r.first.nextInRange(minAllowedDistance, maxAllowedDistance, 0)
}

type level struct {
	val  int
	next *level
}

func (l *level) nextInRange(minVal, maxVal int, sign int) bool {
	if l.next == nil {
		return true
	}

	delta := l.val - l.next.val
	if sign == 0 {
		internalDelta := int(math.Abs(float64(delta)))
		if internalDelta < minVal || internalDelta > maxVal {
			return false
		}
	} else {
		internalMin := sign * minVal
		internalMax := sign * maxVal
		if delta < min(internalMin, internalMax) || delta > max(internalMin, internalMax) {
			return false
		}
	}
	if sign == 0 {
		if delta > 0 {
			sign = 1
		} else {
			sign = -1
		}
	}
	return l.next.nextInRange(minVal, maxVal, sign)
}

func min[T int | int32 | int64](v1, v2 T) T {
	if v1 < v2 {
		return v1
	}
	return v2
}

func max[T int | int32 | int64](v1, v2 T) T {
	if v1 > v2 {
		return v1
	}
	return v2
}
