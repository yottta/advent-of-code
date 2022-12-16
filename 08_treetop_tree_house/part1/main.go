package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	aoc "github.com/yottta/aoc2022/00_aoc"
)

type treeVisibility struct {
	top    bool
	left   bool
	right  bool
	bottom bool
}

func (t treeVisibility) visible() bool {
	return t.top || t.left || t.right || t.bottom
}

type tree struct {
	height     int
	visibility treeVisibility
	line       int
	col        int
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
	f.updateTreeVisibility(e)
	f.updateNeighborsVisibility(e)
}

func (f *forest) updateTreeVisibility(t *tree) {
	if t.line == 0 {
		t.visibility.top = true
	}
	if t.col == 0 {
		t.visibility.left = true
	}
	// new trees are always visible from the bottom and right side
	t.visibility.bottom = true
	t.visibility.right = true
	// if it's inside the grid (not on an edge) figure out if it's visible from the left or top
	if inOpenRange(t.col, 0, f.gridSize-1) && inOpenRange(t.line, 0, f.gridSize-1) {
		leftTree := f.lineHighest[t.line]
		t.visibility.left = leftTree.visibility.left && t.height > leftTree.height
		topTree := f.colHighest[t.col]
		t.visibility.top = topTree.visibility.top && t.height > topTree.height
	}
	// update the highest looking from the top
	if colHigh := f.colHighest[t.col]; !t.same(colHigh) && t.height > colHigh.height {
		f.colHighest[t.col] = t
	}
	// update the highest looking from the left
	if lineHigh := f.lineHighest[t.line]; !t.same(lineHigh) && t.height > lineHigh.height {
		f.lineHighest[t.line] = t
	}
}

func (f *forest) updateNeighborsVisibility(newT *tree) {
	if newT.col == 0 || newT.line == 0 {
		return // do not update visibility of the left top edges
	}
	// update trees at the top only if the new tree is between (included) the 2nd line
	// and if the column is not the last
	if newT.line-1 > 0 && newT.col < f.gridSize-1 {
		for i := newT.line - 1; i > 0; i-- {
			t := f.trees[i][newT.col]
			t.visibility.bottom = t.visibility.bottom && t.height > newT.height
		}
	}
	// update trees to the left only if the new tree is between (included) the 2nd column
	// and if the line is not the last
	if newT.col-1 > 0 && newT.line < f.gridSize-1 {
		for i := newT.col - 1; i > 0; i-- {
			t := f.trees[newT.line][i]
			t.visibility.right = t.visibility.right && t.height > newT.height
		}
	}
}

func main() {
	file, err := os.Open("./input.txt")
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
	var totalVisible int
	for i := 0; i < len(f.trees); i++ {
		for j := 0; j < len(f.trees); j++ {
			if f.trees[i][j].visibility.visible() {
				totalVisible++
			}
			fmt.Printf("%d(%s) ", f.trees[i][j].height, string(strconv.FormatBool(f.trees[i][j].visibility.visible())[0]))
		}
		fmt.Println()
	}
	fmt.Println(totalVisible)
}

func inOpenRange(val, lower, upper int) bool {
	return val < upper && val > lower
}
