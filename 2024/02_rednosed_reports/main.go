package main

import (
	"fmt"
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
			aoc.Log(fmt.Sprintf("%d report is safe", i))
			safe++
			//} else {
			//	v.print()
		}
	}
	fmt.Println(safe)
}

// 569 too low
// 574 too low
func part2(content []string) {
	return
	reports := parseContent(content, 0)
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

func parseContent(content []string, allowedFailures int) []*report {
	reports := make([]*report, len(content))
	for ri, l := range content {
		var lvls []*level
		rl := strings.Split(l, " ")
		for i := 0; i < len(rl); i++ {
			val, err := strconv.Atoi(rl[i])
			aoc.Must(err)
			lvls = append(lvls, &level{val: val})
		}
		rep := &report{levels: lvls}
		rep.process(allowedFailures)
		reports[ri] = rep
	}
	return reports
}

type report struct {
	levels []*level
	isSafe bool
}

func (r *report) print() {
	var b strings.Builder
	for i := 0; i < len(r.levels); i++ {
		b.WriteString(strconv.Itoa(r.levels[i].val))
		b.WriteString(" ")
	}
	fmt.Println(b.String())
}

func (r *report) safe() bool {
	return r.isSafe
}

func (r *report) process(allowedFailures int) {
	//r.print()
	diffs := r.diffs()
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

	//fmt.Println(diffs)

}

type level struct {
	val int
}

func (r *report) diffs() []int {
	diffs := make([]int, len(r.levels)-1)
	for i := 0; i < len(r.levels)-1; i++ {
		diffs[i] = r.levels[i].val - r.levels[i+1].val
	}
	return diffs
}

//
//func (l *level) nextInRange(minVal, maxVal int, sign int) bool {
//	if l.next == nil {
//		return true
//	}
//
//	delta := l.val - l.next.val
//	if sign == 0 {
//		internalDelta := int(math.Abs(float64(delta)))
//		if internalDelta < minVal || internalDelta > maxVal {
//			return false
//		}
//	} else {
//		internalMin := sign * minVal
//		internalMax := sign * maxVal
//		if delta < min(internalMin, internalMax) || delta > max(internalMin, internalMax) {
//			return false
//		}
//	}
//	if sign == 0 {
//		if delta > 0 {
//			sign = 1
//		} else {
//			sign = -1
//		}
//	}
//	return l.next.nextInRange(minVal, maxVal, sign)
//}

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

func abs[T int | int32 | int64](v1 T) T {
	if v1 >= 0 {
		return v1
	}
	return v1 * -1
}
