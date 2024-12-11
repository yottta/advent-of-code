package main

import (
	"fmt"
	"strconv"
	"strings"

	aoc "github.com/yottta/advent-of-code/00_aoc"
	aoc2024 "github.com/yottta/advent-of-code/2024"
)

func main() {
	aoc2024.Run(part1, part2, false, "11")
}

func part1(content []string) {
	peb := parseContent(content[0])
	peb.print()
	fmt.Println()
	for i := 0; i < 25; i++ {
		peb.transform()
		//peb.print()
		//fmt.Println()
	}
	fmt.Println(peb.count())
}

func parseContent(content string) *pebble {
	var (
		first, last *pebble
	)
	for _, p := range strings.Split(content, " ") {
		val, err := strconv.Atoi(p)
		aoc.Must(err)
		newP := pebble{
			val:    val,
			strVal: p,
		}
		if first == nil {
			first = &newP
			last = first
			continue
		}
		last.next = &newP
		last = last.next
	}
	return first
}

func part2(content []string) {

}

type pebble struct {
	val    int
	strVal string
	next   *pebble
}

func (p *pebble) transform() {
	var (
		newNext *pebble
	)
	switch {
	case p.val == 0:
		p.val = 1
		p.strVal = strconv.Itoa(p.val)
	case len(p.strVal)%2 == 0:
		currVal, err := strconv.Atoi(p.strVal[:len(p.strVal)/2])
		aoc.Must(err)
		newVal, err := strconv.Atoi(p.strVal[len(p.strVal)/2:])
		aoc.Must(err)
		p.val = currVal
		p.strVal = strconv.Itoa(currVal)
		newNext = &pebble{
			val:    newVal,
			strVal: strconv.Itoa(newVal),
		}
	default:
		p.val = p.val * 2024
		p.strVal = strconv.Itoa(p.val)
	}
	if p.val == 0 {
		p.val = 1
	}
	if p.next != nil {
		p.next.transform()
	}
	if newNext != nil {
		newNext.next = p.next
		p.next = newNext
	}
}

func (p *pebble) print() {
	fmt.Print(p.strVal)
	if p.next == nil {
		return
	}
	fmt.Print(" ")
	p.next.print()
}

func (p *pebble) count() int {
	if p == nil {
		return 0
	}
	return 1 + p.next.count()
}
