package main

import (
	"fmt"
	aoc "github.com/yottta/advent-of-code/00_aoc"
	"strings"
)

func main() {
	aoc.BasicRun(part1, part2)
}

func part1(content []string) {
	in := "????.#...#..."

	out := strings.Trim(in, ".")
	for strings.Contains(out, "..") {
		out = strings.ReplaceAll(out, "..", ".")
	}

	fmt.Println(out)
	parts := strings.Split(out, ".")
	fmt.Println(len(parts) == 3)
}

func part2(content []string) {

}
