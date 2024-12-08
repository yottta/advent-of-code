package main

import (
	"fmt"
	"strconv"
	"strings"

	aoc "github.com/yottta/advent-of-code/00_aoc"
	aoc2024 "github.com/yottta/advent-of-code/2024"
)

func main() {
	aoc2024.Run(part1, part2, false, "07")
}

func part1(content []string) {
	results := parseData(content)
	var res int

	for _, r := range results {
		if r.valid() {
			res += r.val
		}
	}
	fmt.Println(res)
}
func part2(content []string) {

}

func parseData(content []string) []*result {
	var res []*result
	for _, line := range content {
		parts := strings.Split(line, ":")
		resVal, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		aoc.Must(err)
		r := &result{
			val: resVal,
		}
		var prevVal *value
		for _, val := range strings.Split(parts[1], " ") {
			val = strings.TrimSpace(val)
			if val == "" {
				continue
			}
			vVal, err := strconv.Atoi(val)
			aoc.Must(err)
			v := &value{
				val: vVal,
			}
			if prevVal != nil {
				prevVal.next = v
				prevVal = v
				continue
			}
			prevVal = v
			r.firstValue = prevVal
		}
		res = append(res, r)
	}
	return res
}

type operator int

const (
	opAdd = iota + 1
	opMultiply
)

var (
	allOps = []operator{opAdd, opMultiply}
)

type value struct {
	next *value
	val  int
}

func (v *value) calc(prevVal int, sum int) bool {
	if v.next == nil {
		for _, op := range allOps {
			val := applyOp(prevVal, v.val, op)
			if val == sum {
				return true
			}
		}
		return false
	}
	for _, op := range allOps {
		val := applyOp(prevVal, v.val, op)
		if v.next.calc(val, sum) {
			return true
		}
	}
	return false
}

type result struct {
	val        int
	firstValue *value
}

func (r *result) valid() bool {
	if r.firstValue == nil {
		return r.val == 0
	}
	return r.firstValue.calc(0, r.val)
}

func applyOp(v1, v2 int, op operator) int {
	switch op {
	case opAdd:
		return v1 + v2
	case opMultiply:
		if v1 == 0 {
			v1 = 1
		}
		return v1 * v2
	}
	panic(fmt.Errorf("invalid operator %v", op))
}
