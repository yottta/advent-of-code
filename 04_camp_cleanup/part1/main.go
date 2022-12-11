package main

import (
	"bufio"
	"fmt"
	aoc "github.com/yottta/aoc2022/00_aoc"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	start int
	stop  int
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

func fullyOverlap(r1 Range, r2 Range) bool {
	if r1.start > r2.start {
		return r1.stop <= r2.stop
	}
	if r2.start > r1.start {
		return r2.stop <= r1.stop
	}
	return true
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)
	var (
		end   bool
		total int
	)
	for {
		if end {
			break
		}
		r, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading string", err)
			end = true
		}
		r = strings.ReplaceAll(r, "\n", "")
		if len(r) == 0 {
			break
		}
		split := strings.Split(r, ",")
		if len(split) != 2 {
			panic(fmt.Errorf("wrong input %s", r))
		}

		r1 := parseRange(split[0])
		r2 := parseRange(split[1])

		if fullyOverlap(r1, r2) {
			total++
		}
	}
	fmt.Println(total)
}
