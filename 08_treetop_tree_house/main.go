package main

import (
	"flag"
	"fmt"
	"strconv"
	
	aoc "github.com/yottta/aoc2022/00_aoc"
)

var verbose bool

func main() {
	var (
		dataFilePath string
		partToRun    string
	)
	flag.StringVar(&dataFilePath, "d", "./input.txt", "The path of the file containing the data for the current problem")
	flag.StringVar(&partToRun, "p", "1", "The part of the problem to run, in case the problem has more than one parts")
	flag.BoolVar(&verbose, "v", false, "Add this for more information during running if available")
	flag.Parse()

	aoc.Verbose(verbose)

	content, err := aoc.ReadFile(dataFilePath)
	aoc.Must(err)

	switch partToRun {
	case "1":
		part1(content)
	case "2":
		part2(content)
	default:
		panic(fmt.Errorf("no part '%s' configured", partToRun))
	}
}

func part1(content []string) {
	var f forest
	for _, line := range content {
		f.gridSize = len(line)
		for i, c := range line {
			height, err := strconv.Atoi(string(c))
			aoc.Must(err)
			// add tree as its height and create a new line every time a new line is read
			f.addTreeForVisibility(height, i == 0)
		}
	}

	var totalVisible int
	for i := 0; i < len(f.trees); i++ {
		for j := 0; j < len(f.trees); j++ {
			if f.trees[i][j].visibility.visible() {
				totalVisible++
			}
			aoc.Logf("%d(%s) ", f.trees[i][j].height, string(strconv.FormatBool(f.trees[i][j].visibility.visible())[0]))
		}
		aoc.Log("")
	}
	fmt.Println(totalVisible)
}

func part2(content []string) {
	var f forest
	for _, line := range content {
		f.gridSize = len(line)
		for i, c := range line {
			height, err := strconv.Atoi(string(c))
			aoc.Must(err)
			// add tree as its height and create a new line every time a new line is read
			f.addTreeForScenery(height, i == 0)
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
			aoc.Logf("%d(%00d) ", t.height, scenery)
		}
		aoc.Log("")
	}
	fmt.Println(maxScenery)
}

type forest struct {
	gridSize int
	trees    [][]*tree
	// from top to bottom to figure out if a new tree is visible from the top
	colHighest []*tree
	// from left to right to figure out if a new tree is visible from the left
	lineHighest []*tree
}

func (f *forest) addTreeForVisibility(h int, newRow bool) {
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

// for part2
func (f *forest) addTreeForScenery(h int, newRow bool) {
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

// for part2
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

// for part2
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

// for part2
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

type treeVisibility struct {
	top    bool
	left   bool
	right  bool
	bottom bool
}

func (t treeVisibility) visible() bool {
	return t.top || t.left || t.right || t.bottom
}

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
	height     int
	visibility treeVisibility
	scenery    treeScenery
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

func inOpenRange(val, lower, upper int) bool {
	return val < upper && val > lower
}
