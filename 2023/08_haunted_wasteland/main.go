package main

import (
	"fmt"
	aoc "github.com/yottta/advent-of-code/00_aoc"
	"strings"
)

func main() {
	aoc.BasicRun(part1, part2)
}

type node struct {
	name  string
	right string
	left  string
}

func part1(content []string) {
	nodes, instructions := parseInput(content)
	nextNode := "AAA"
	steps := walkToEnd(nextNode, nodes, func(s string) bool {
		return s == "ZZZ"
	}, instructions)
	fmt.Println(steps)
}

func part2(content []string) {
	nodes, directions := parseInput(content)

	allNodes := allNodesEndingInA(nodes)
	steps := make([]int, len(allNodes))
	for i, n := range allNodes {
		stepsNeeded := walkToEnd(n, nodes, func(s string) bool {
			return strings.HasSuffix(s, "Z")
		}, directions)
		steps[i] = stepsNeeded
	}

	fmt.Println(lcm(steps))
}

func lcm(nums []int) int {
	lowestMultiple := 1
	for _, n := range nums {
		lowestMultiple = (lowestMultiple * n) / gcd(lowestMultiple, n)
	}
	return lowestMultiple
}

func gcd(x int, y int) int {
	for y > 0 {
		t := y
		y = x % y
		x = t
	}
	return x
}

func walkToEnd(nextNode string, nodes map[string]node, endingCheck func(string) bool, directions string) int {
	var nextDirection int
	var steps int
	for !endingCheck(nextNode) {
		dir := directions[nextDirection]
		if dir == 'L' {
			nextNode = nodes[nextNode].left
		} else if dir == 'R' {
			nextNode = nodes[nextNode].right
		}
		steps++
		nextDirection++
		if nextDirection >= len(directions) {
			nextDirection = 0
		}
	}
	return steps
}

func allNodesEndingInA(in map[string]node) []string {
	var result []string
	for root := range in {
		if strings.HasSuffix(root, "A") {
			result = append(result, root)
		}
	}
	return result
}

func parseInput(content []string) (map[string]node, string) {
	nodes := map[string]node{}
	instructions := content[0]
	var startFrom string
	for _, connections := range content[2:] {
		if strings.TrimSpace(connections) == "" {
			continue
		}
		parts := strings.Split(connections, " = ")
		source := strings.TrimSpace(parts[0])

		n, ok := nodes[source]
		if !ok {
			n = node{name: source}
		}
		cleanedUp := strings.TrimSpace(parts[1])
		cleanedUp = strings.ReplaceAll(cleanedUp, "(", "")
		cleanedUp = strings.ReplaceAll(cleanedUp, ")", "")
		directions := strings.Split(cleanedUp, ", ")
		n.left = directions[0]
		n.right = directions[1]
		nodes[source] = n
		if startFrom == "" {
			startFrom = source
		}
	}
	return nodes, instructions
}
