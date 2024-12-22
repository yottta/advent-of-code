package main

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	aoc "github.com/yottta/advent-of-code/00_aoc"
)

var verbose bool

func main() {
	aoc.BasicRun(part1, part2)
}

func part1(content []string) {
	var (
		knotsPositions = buildKnots(2)
		points         []point
	)
	points = append(points, knotsPositions[len(knotsPositions)-1])
	for _, line := range content {
		parts := strings.Split(line, " ")
		steps, err := strconv.Atoi(parts[1])
		aoc.Must(err)
		slog.Debug(line)
		for i := 0; i < steps; i++ {
			oldHeadPos := knotsPositions[0]
			switch parts[0] {
			case "R":
				knotsPositions[0].x++
				if distance(knotsPositions[0], knotsPositions[1]) > 1 {
					knotsPositions[1].x = oldHeadPos.x
					knotsPositions[1].y = oldHeadPos.y
					points = append(points, knotsPositions[1])
				}
			case "L":
				knotsPositions[0].x--
				if distance(knotsPositions[0], knotsPositions[1]) > 1 {
					knotsPositions[1].y = oldHeadPos.y
					knotsPositions[1].x = oldHeadPos.x
					points = append(points, knotsPositions[1])
				}
			case "U":
				knotsPositions[0].y++
				if distance(knotsPositions[0], knotsPositions[1]) > 1 {
					knotsPositions[1].y = oldHeadPos.y
					knotsPositions[1].x = oldHeadPos.x
					points = append(points, knotsPositions[1])
				}
			case "D":
				knotsPositions[0].y--
				if distance(knotsPositions[0], knotsPositions[1]) > 1 {
					knotsPositions[1].y = oldHeadPos.y
					knotsPositions[1].x = oldHeadPos.x
					points = append(points, knotsPositions[1])
				}
			default:
				panic("invalid instruction")
			}
			drawPoints(knotsPositions...)
		}
	}
	drawPath(points)

	res := make(map[point]struct{})
	for _, p := range points {
		res[p] = struct{}{}
	}
	fmt.Println(len(res))
}

func part2(content []string) {
	var (
		knotsPositions = buildKnots(10)
		// traveled by the last knot
		points []point
	)

	points = append(points, knotsPositions[len(knotsPositions)-1])
	for _, line := range content {
		parts := strings.Split(line, " ")
		steps, err := strconv.Atoi(parts[1])
		aoc.Must(err)

		slog.Debug(line)
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

		slog.Debug(line)
		drawPoints(knotsPositions...)
		slog.Debug("")
	}

	drawPath(points)

	res := make(map[point]struct{})
	for _, p := range points {
		res[p] = struct{}{}
	}
	fmt.Println(len(res))
}

func buildKnots(size int) []point {
	res := make([]point, size)
	res[0] = point{
		name: "H",
	}
	for i := 1; i < size; i++ {
		res[i] = point{
			name: strconv.Itoa(i),
		}
	}
	return res
}

type point struct {
	name string
	x    int
	y    int
}

func drawPath(points []point) {
	if !verbose {
		return
	}
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
	if !verbose {
		return
	}

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
