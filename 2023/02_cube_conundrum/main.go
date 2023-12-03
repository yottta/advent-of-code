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

var maxCubes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

// 346 too low
func part1(content []string) {
	var possibleIds []int
	for i, s := range content {
		possible := true
		index := strings.Index(s, ":")
		sets := strings.Split(s[index+1:], ";")
		for _, set := range sets {
			if !possible {
				break
			}
			setCubes := strings.Split(set, ",")
			for _, cube := range setCubes {
				parts := strings.Split(strings.TrimSpace(cube), " ")
				noOfCubes, err := strconv.Atoi(strings.TrimSpace(parts[0]))
				aoc.Must(err)
				color := strings.TrimSpace(parts[1])
				maxCubesForColor, ok := maxCubes[color]
				if !ok {
					panic(fmt.Errorf("color %s not found", color))
				}
				if noOfCubes > maxCubesForColor {
					possible = false
					break
				}
			}
		}
		if possible {
			possibleIds = append(possibleIds, i+1)
		}
	}
	fmt.Println(aoc.SumList(possibleIds))
}

func part2(content []string) {
	var gamesPower []int
	for _, s := range content {
		var minimumNeededCubes = map[string]int{}
		index := strings.Index(s, ":")
		sets := strings.Split(s[index+1:], ";")
		for _, set := range sets {
			setCubes := strings.Split(set, ",")
			for _, cube := range setCubes {
				parts := strings.Split(strings.TrimSpace(cube), " ")
				noOfCubes, err := strconv.Atoi(strings.TrimSpace(parts[0]))
				aoc.Must(err)
				color := strings.TrimSpace(parts[1])
				current, _ := minimumNeededCubes[color]
				if current < noOfCubes {
					minimumNeededCubes[color] = noOfCubes
				}
			}
		}
		gamesPower = append(gamesPower, multiplyMapValues(minimumNeededCubes))
	}
	fmt.Println(aoc.SumList(gamesPower))
}

func multiplyMapValues(m map[string]int) int {
	total := 1
	for _, v := range m {
		total *= v
	}
	return total
}
