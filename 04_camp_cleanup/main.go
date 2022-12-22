package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	aoc "github.com/yottta/aoc2022/00_aoc"
)

func main() {
	var (
		dataFilePath string
		partToRun    string
	)
	flag.StringVar(&dataFilePath, "d", "./input.txt", "The path of the file containing the data for the current problem")
	flag.StringVar(&partToRun, "p", "1", "The part of the problem to run, in case the problem has more than one parts")
	flag.Parse()

	content, err := aoc.ReadFile(dataFilePath)
	aoc.Must(err)
	
	switch partToRun {
	case "1":
		part1(content)
	case "2":
		part2(content)
	default:
		panic(fmt.Errorf("no part '%s' configured", partToRun))
	}
}

func part1(content []string) {
	var total int
	for _, line := range content {
		split := strings.Split(line, ",")
		r1 := parseRange(split[0])
		r2 := parseRange(split[1])

		if r1.fullyOverlaps(r2) {
			fmt.Println("fully overlapping", r1, r2)
			total++
		}
	}
	fmt.Println(total)
}

func part2(content []string) {
	var total int
	for _, line := range content {
		split := strings.Split(line, ",")
		r1 := parseRange(split[0])
		r2 := parseRange(split[1])

		if r1.partiallyOverlaps(r2) {
			fmt.Println("partially overlapping", r1, r2)
			total++
		}
	}
	fmt.Println(total)
}

type Range struct {
	start int
	stop  int
}

func (r Range) fullyOverlaps(r2 Range) bool {
	if r.start > r2.start {
		return r.stop <= r2.stop
	}
	if r2.start > r.start {
		return r2.stop <= r.stop
	}
	return true
}

func (r Range) partiallyOverlaps(r2 Range) bool {
	if r.start > r2.start {
		return r2.stop >= r.start
	}
	if r.start < r2.start {
		return r.stop >= r2.start
	}
	return true
}

func parseRange(in string) Range {
	split := strings.Split(in, "-")
	if len(split) != 2 {
		panic(fmt.Errorf("wrong input: %s", in))
	}
	start, err := strconv.Atoi(split[0])
	aoc.Must(err)
	end, err := strconv.Atoi(split[1])
	aoc.Must(err)
	return Range{
		start: start,
		stop:  end,
	}
}
