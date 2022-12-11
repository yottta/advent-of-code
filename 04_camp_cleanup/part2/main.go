package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	start int
	stop  int
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func parseRange(in string) Range {
	split := strings.Split(in, "-")
	if len(split) != 2 {
		panic(fmt.Errorf("wrong input: %s", in))
	}
	start, err := strconv.Atoi(split[0])
	must(err)
	end, err := strconv.Atoi(split[1])
	must(err)
	return Range{
		start: start,
		stop:  end,
	}
}

func partiallyOverlap(r1 Range, r2 Range) bool {
	if r1.start > r2.start {
		return r2.stop >= r1.start
	}
	if r1.start < r2.start {
		return r1.stop >= r2.start
	}
	return true
}

// 603
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

		if partiallyOverlap(r1, r2) {
			fmt.Println(r1, r2)
			total++
		}
	}
	fmt.Println(total)
}
