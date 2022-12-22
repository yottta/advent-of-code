package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	aoc "github.com/yottta/aoc2022/00_aoc"
)

func main() {
	var (
		dataFilePath string
		partToRun    string
	)
	flag.StringVar(&dataFilePath, "d", "./input.txt", "The path of the file containing the data for the current problem")
	flag.StringVar(&partToRun, "p", "1", "The part of the problem to run, in case the problem has more than one parts")
	flag.Parse()

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
	initialState, movements := splitContent(content)
	stacks := buildInitialState(initialState)
	for _, line := range movements {
		if !strings.Contains(line, "move") {
			continue
		}
		actions := strings.Split(line, " ")
		noOfCrates, err := strconv.Atoi(actions[1])
		aoc.Must(err)
		sourceStack, err := strconv.Atoi(actions[3])
		aoc.Must(err)
		targetStack, err := strconv.Atoi(actions[5])
		aoc.Must(err)

		for i := 0; i < noOfCrates; i++ {
			obj, ok := stacks[sourceStack-1].pop()
			if ok {
				stacks[targetStack-1].push(obj)
			}
		}
	}
	printStacksTop(stacks)
}

func part2(content []string) {
	initialState, movements := splitContent(content)
	stacks := buildInitialState(initialState)

	for _, line := range movements {
		if !strings.Contains(line, "move") {
			continue
		}
		actions := strings.Split(line, " ")
		noOfCrates, err := strconv.Atoi(actions[1])
		aoc.Must(err)
		sourceStack, err := strconv.Atoi(actions[3])
		aoc.Must(err)
		targetStack, err := strconv.Atoi(actions[5])
		aoc.Must(err)

		content, ok := stacks[sourceStack-1].popRange(noOfCrates)
		if !ok {
			continue
		}
		stacks[targetStack-1].push(content...)
	}
	printStacksTop(stacks)
}

// splits file content in two parts: first is the initial state description and the second is the list of movements
func splitContent(content []string) ([]string, []string) {
	var initialStateEnd int
	for i, line := range content {
		if !strings.Contains(line, "[") {
			initialStateEnd = i
			break
		}
	}
	return content[:initialStateEnd], content[initialStateEnd+1:]
}

type stack []string

func (s *stack) insert(c string) {
	newStack := make([]string, len(*s)+1)
	newStack[0] = c
	copy(newStack[1:], *s)
	*s = newStack
}

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) push(c ...string) {
	*s = append(*s, c...)
}

func (s *stack) pop() (string, bool) {
	if s.isEmpty() {
		return "", false
	}
	idx := len(*s) - 1
	res := (*s)[idx]
	*s = (*s)[:idx]
	return res, true
}

func (s *stack) popRange(size int) ([]string, bool) {
	if s.isEmpty() {
		return nil, false
	}
	if len(*s) < size {
		return nil, false
	}
	start := len(*s) - size
	end := len(*s)
	res := (*s)[start:end]
	*s = (*s)[:start]
	return res, true
}

func (s *stack) top() (string, bool) {
	if s.isEmpty() {
		return "", false
	}
	idx := len(*s) - 1
	return (*s)[idx], true
}

const crateSpacing = 4

func buildInitialState(content []string) []stack {
	var stacks []stack
	for _, line := range content {
		if stacks == nil {
			stacks = make([]stack, (len(line)+1)/crateSpacing)
		}
		for i := 0; i < len(stacks); i++ {
			endIdx := (i+1)*crateSpacing - 1
			if endIdx >= len(line) {
				endIdx = len(line) - 1
			}
			crateContent := line[i*crateSpacing : endIdx]
			if !strings.Contains(crateContent, "[") {
				continue
			}
			crate := strings.TrimSpace(
				strings.ReplaceAll(
					strings.ReplaceAll(
						crateContent,
						"]",
						"",
					),
					"[",
					"",
				),
			)
			if len(crate) > 0 {
				stacks[i].insert(
					crate,
				)
			}
		}
	}
	return stacks
}

func printStacksTop(stacks []stack) {
	var b strings.Builder
	for _, s := range stacks {
		top, ok := s.top()
		if ok {
			b.WriteString(top)
		}
	}
	fmt.Println(b.String())
}
