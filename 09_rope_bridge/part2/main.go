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
	name string
	x    int
	y    int
}

func main() {
	var test bool
	file, err := os.Open("./input.txt")
	test = false

	aoc.Must(err)

	reader := bufio.NewReader(file)
	var (
		end            bool
		knotsPositions = make([]point, 10)
		// tracking the 9th knot
		points []point
	)
	for i := 0; i < 10; i++ {
		if i == 0 {
			knotsPositions[i] = point{
				name: "H",
			}
			continue
		}
		knotsPositions[i] = point{
			name: strconv.Itoa(i),
		}
	}
	points = append(points, knotsPositions[len(knotsPositions)-1])
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
			var deltaX, deltaY int
			switch parts[0] {
			case "R":
				deltaX = 1
			case "L":
				deltaX = -1
			case "U":
				deltaY = 1
			case "D":
				deltaY = -1
			default:
				panic("invalid instruction")
			}

			knotsPositions[0].x += deltaX
			knotsPositions[0].y += deltaY
			for kp := 1; kp < len(knotsPositions); kp++ {
				diffX := knotsPositions[kp-1].x - knotsPositions[kp].x
				diffY := knotsPositions[kp-1].y - knotsPositions[kp].y
				var modified bool
				if absInt(diffX) > 1 && absInt(diffY) > 1 {
					modified = true
					knotsPositions[kp].x += diffX / 2
					knotsPositions[kp].y += diffY / 2
				} else if absInt(diffY) > 1 {
					modified = true
					knotsPositions[kp].x += diffX
					knotsPositions[kp].y += diffY / 2
				} else if absInt(diffX) > 1 {
					modified = true
					knotsPositions[kp].x += diffX / 2
					knotsPositions[kp].y += diffY
				}
				if modified && kp == 9 {
					points = append(points, knotsPositions[kp])
				}
			}
		}

		if test {
			fmt.Println(line)
			drawPoints(knotsPositions...)
			fmt.Println()
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

// just for local testing with a limited set of data points
func drawPoints(pts ...point) {
	xSize := 26
	ySize := 21
	startingPoint := point{"s", 11, 5}
	points := make([]point, len(pts))
	for i := 0; i < len(pts); i++ {
		points[i] = point{
			name: pts[i].name,
			x:    pts[i].x,
			y:    pts[i].y,
		}
		points[i].x += startingPoint.x
		points[i].y += startingPoint.y
	}
	grid := make([][]string, ySize)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]string, xSize)
	}
	grid[startingPoint.y][startingPoint.x] = "s"
	for i := len(points) - 1; i >= 0; i-- {
		grid[points[i].y][points[i].x] = points[i].name
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
	fmt.Println()
}

func absInt(i int) int {
	if i >= 0 {
		return i
	}
	return i * -1
}
