package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	aoc "github.com/yottta/aoc2022/00_aoc"
)

type treeScenery struct {
	top    int
	left   int
	bottom int
	right  int
}

func (s treeScenery) score() int {
	return s.top * s.left * s.bottom * s.right
}

type tree struct {
	height  int
	scenery treeScenery
	line    int
	col     int
}

func (t *tree) same(t2 *tree) bool {
	if t.col != t2.col {
		return false
	}
	if t.line != t2.line {
		return false
	}
	return t.height == t2.height
}

type forest struct {
	gridSize int
	trees    [][]*tree
	// from top to bottom to figure out if a new tree is visible from the top
	colHighest []*tree
	// from left to right to figure out if a new tree is visible from the left
	lineHighest []*tree
}

func (f *forest) addTree(h int, newRow bool) {
	e := &tree{
		height: h,
	}
	// add new row if asked for
	if newRow {
		f.trees = append(f.trees, []*tree{})
		f.lineHighest = append(f.lineHighest, e)
	}
	// generate the tree position on the grid
	col := f.trees[len(f.trees)-1]
	col = append(col, e)
	f.trees[len(f.trees)-1] = col

	e.line = len(f.trees) - 1
	e.col = len(col) - 1
	if len(f.colHighest) < len(col) {
		f.colHighest = append(f.colHighest, e)
	}
	f.updateLeftTopScenery(e)
	f.updateHighest(e)
}

func (f *forest) updateHighest(t *tree) {
	// update the highest looking from the top
	if colHigh := f.colHighest[t.col]; !t.same(colHigh) && t.height >= colHigh.height {
		f.colHighest[t.col] = t
	}
	// update the highest looking from the left
	if lineHigh := f.lineHighest[t.line]; !t.same(lineHigh) && t.height >= lineHigh.height {
		f.lineHighest[t.line] = t
	}
}

func (f *forest) updateLeftTopScenery(e *tree) {
	// update left scenery
	lineHighest := f.lineHighest[e.line]
	if e.height == lineHighest.height {
		e.scenery.left = e.col - lineHighest.col
	} else if e.height > lineHighest.height {
		e.scenery.left = e.col
	} else {
		for i := e.col; i > 0; i-- {
			e.scenery.left++
			if f.trees[e.line][i].height >= e.height {
				break
			}
		}
	}
	// update top
	colHighest := f.colHighest[e.col]
	if e.height == colHighest.height {
		e.scenery.top = e.line - colHighest.line
	} else if e.height > colHighest.height {
		e.scenery.top = e.line
	} else {
		for i := e.col; i > 0; i-- {
			e.scenery.top++
			if f.trees[e.line][i].height >= e.height {
				break
			}
		}
	}
}

func (f *forest) updateRightBottomScenery(in *tree) {
	for i := in.col + 1; i < f.gridSize; i++ {
		in.scenery.right++
		t := f.trees[in.line][i]
		if in.height <= t.height {
			break
		}
	}
	for i := in.line + 1; i < f.gridSize; i++ {
		in.scenery.bottom++
		t := f.trees[i][in.col]
		if in.height <= t.height {
			break
		}
	}
}

func main() {
	file, err := os.Open("../input.txt")
	aoc.Must(err)

	reader := bufio.NewReader(file)
	var (
		end bool
		f   forest
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
		f.gridSize = len(line)
		for i, c := range line {
			height, err := strconv.Atoi(string(c))
			aoc.Must(err)
			// add tree as it's height and create a new line every time a new line is read
			f.addTree(height, i == 0)
		}
	}
	var maxScenery int
	for i := 0; i < len(f.trees); i++ {
		for j := 0; j < len(f.trees); j++ {
			t := f.trees[i][j]
			f.updateRightBottomScenery(t)
			scenery := t.scenery.score()
			if scenery > maxScenery {
				maxScenery = scenery
			}
			fmt.Printf("%d(%00d) ", t.height, scenery)
		}
		fmt.Println()
	}
	fmt.Println(maxScenery)
}
