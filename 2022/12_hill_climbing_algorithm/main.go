package main

import (
	"fmt"
	"strconv"

	aoc "github.com/yottta/advent-of-code/00_aoc"
)

var verbose bool

func main() {
	aoc.BasicRun(part1, part2)
}

func part1(content []string) {
	m := parseMountainHeights(content)
	printMountain(m)
	lastStep := m.navigate(false)
	if lastStep.stepNo < 0 {
		panic("no solution")
	}
	printSolution(m, &lastStep)
	fmt.Println(lastStep.stepNo)
}

func part2(content []string) {
	m := parseMountainHeights(content)
	printMountain(m)
	lastStep := m.navigate(true)
	if lastStep.stepNo < 0 {
		panic("no solution")
	}
	printSolution(m, &lastStep)
	fmt.Println(lastStep.stepNo)
}

func parseMountainHeights(content []string) mountain {
	var heights [][]*step
	var end *step
	var start *step
	for lineNo, line := range content {
		row := make([]*step, len(line))
		for colNo, h := range line {
			s := &step{
				point: point{
					x: colNo,
					y: lineNo,
				},
				height: h,
				value:  h,
			}
			if h == 'S' {
				s.height = 'a'
				start = s
			}
			if h == 'E' {
				s.height = 'z'
				end = s
			}
			row[colNo] = s
		}
		heights = append(heights, row)
	}
	return mountain{
		steps: heights,
		start: start,
		end:   end,
	}
}

type point struct {
	x int
	y int
}

type step struct {
	point
	height int32
	value  int32
}

type mountain struct {
	steps [][]*step
	start *step
	end   *step
}

func (m *mountain) navigate(backwards bool) navigateStep {
	var (
		queue []navigateStep
		seen  = map[point]struct{}{}
		diff  = [4][2]int{
			{-1, 0},
			{1, 0},
			{0, -1},
			{0, 1},
		}
	)

	if backwards {
		queue = append(queue, navigateStep{
			step:   m.end,
			stepNo: 0,
		})
	} else {
		queue = append(queue, navigateStep{
			step:   m.start,
			stepNo: 0,
		})
	}

	for len(queue) > 0 {
		currentStep := queue[0]
		queue = queue[1:]
		if backwards && currentStep.step.height == 'a' {
			return currentStep
		} else if !backwards && currentStep.step == m.end {
			return currentStep
		}
		if _, ok := seen[currentStep.point]; ok {
			continue
		}
		seen[currentStep.point] = struct{}{}
		for _, d := range diff {
			nextX := currentStep.x + d[0]
			nextY := currentStep.y + d[1]
			if nextX >= 0 && nextY >= 0 && nextX < len(m.steps[0]) && nextY < len(m.steps) {
				nextStep := m.steps[nextY][nextX]
				if _, ok := seen[nextStep.point]; ok {
					continue
				}
				if backwards && nextStep.height-currentStep.height >= -1 {
					queue = append(queue, navigateStep{
						step:     nextStep,
						stepNo:   currentStep.stepNo + 1,
						previous: &currentStep,
					})
				} else if !backwards && nextStep.height-currentStep.height <= 1 {
					queue = append(queue, navigateStep{
						step:     nextStep,
						stepNo:   currentStep.stepNo + 1,
						previous: &currentStep,
					})
				}
			}
		}
	}

	return navigateStep{stepNo: -1}
}

func printMountain(m mountain) {
	if !verbose {
		return
	}
	for _, c := range m.steps {
		for _, r := range c {
			fmt.Printf("%s", string(r.value))
		}
		fmt.Println()
	}
}

func printSolution(m mountain, end *navigateStep) {
	if !verbose {
		return
	}

	res := make([][]string, len(m.steps))
	for x, row := range m.steps {
		dat := make([]string, len(row))
		for y := range row {
			dat[y] = "."
		}
		res[x] = dat
	}

	var idx int

	for {
		res[end.y][end.x] = strconv.Itoa(idx)
		if end.previous == nil {
			break
		}
		end = end.previous
		idx++
	}
	for _, c := range res {
		for _, r := range c {
			fmt.Printf("%5s", r)
		}
		fmt.Println()
	}
}

type navigateStep struct {
	*step
	stepNo   int
	previous *navigateStep
}
