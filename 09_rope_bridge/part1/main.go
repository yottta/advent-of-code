package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	aoc "github.com/yottta/aoc2022/00_aoc"
)

type point struct {
	x int
	y int
}

func main() {
	var test bool
	file, err := os.Open("../input.txt")
	//file, err := os.Open("./input_test.txt")
	//test = true

	aoc.Must(err)

	reader := bufio.NewReader(file)
	var (
		end          bool
		tailPosition = point{x: 0, y: 0}
		headPosition = point{x: 0, y: 0}
		points       = []point{tailPosition}
	)
	for {
		if end {
			break
		}
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading string", err)
			end = true
		}
		line = strings.ReplaceAll(line, "\n", "")
		if len(line) == 0 {
			continue
		}
		parts := strings.Split(line, " ")
		steps, err := strconv.Atoi(parts[1])
		aoc.Must(err)
		if test {
			fmt.Println(line)
		}
		for i := 0; i < steps; i++ {
			oldHeadPos := headPosition
			switch parts[0] {
			case "R":
				headPosition.x++
				if distance(headPosition, tailPosition) > 1 {
					tailPosition.x = oldHeadPos.x
					tailPosition.y = oldHeadPos.y
					points = append(points, tailPosition)
				}
			case "L":
				headPosition.x--
				if distance(headPosition, tailPosition) > 1 {
					tailPosition.y = oldHeadPos.y
					tailPosition.x = oldHeadPos.x
					points = append(points, tailPosition)
				}
			case "U":
				headPosition.y++
				if distance(headPosition, tailPosition) > 1 {
					tailPosition.y = oldHeadPos.y
					tailPosition.x = oldHeadPos.x
					points = append(points, tailPosition)
				}
			case "D":
				headPosition.y--
				if distance(headPosition, tailPosition) > 1 {
					tailPosition.y = oldHeadPos.y
					tailPosition.x = oldHeadPos.x
					points = append(points, tailPosition)
				}
			default:
				panic("invalid instruction")
			}
			if test {
				drawPoints(headPosition, tailPosition, 6)
			}
		}
	}
	drawPath(points)

	res := make(map[point]struct{})
	for _, p := range points {
		res[p] = struct{}{}
	}
	fmt.Println(len(res))
}

func drawPath(points []point) {
	var minX, minY int
	for _, p := range points {
		if p.x < minX {
			minX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
	}
	for i := 0; i < len(points); i++ {
		points[i].x += absInt(minX)
		points[i].y += absInt(minY)
	}
	var maxX, maxY int
	for _, p := range points {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}
	grid := make([][]string, maxY+1)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]string, maxX+1)
	}
	for i, p := range points {
		if i == 0 {
			grid[p.y][p.x] = "s"
			continue
		}
		if grid[p.y][p.x] == "s" {
			continue
		}
		grid[p.y][p.x] = "x"
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			v := grid[i][j]
			if v == "" {
				fmt.Printf("%2s", ".")
				continue
			}
			fmt.Printf("%2s", v)
		}
		fmt.Println()
	}
}

func drawPoints(p1, p2 point, gridSize int) {
	grid := make([][]string, gridSize)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]string, gridSize)
	}
	grid[p1.y][p1.x] = "H"
	grid[p2.y][p2.x] = "T"

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			v := grid[i][j]
			if v == "" {
				fmt.Printf("%2s", ".")
				continue
			}
			if j == 0 && i == 0 {
				fmt.Printf("%2s", "s")
				continue
			}
			fmt.Printf("%2s", v)
		}
		fmt.Println()
	}
	fmt.Println()
}

func distance(p1, p2 point) int {
	return maxInt(absInt(p1.x-p2.x), absInt(p1.y-p2.y))
}

func absInt(i int) int {
	if i >= 0 {
		return i
	}
	return i * -1
}

func maxInt(i1, i2 int) int {
	if i1 > i2 {
		return i1
	}
	return i2
}
