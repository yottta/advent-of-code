package main

import (
	"bufio"
	"fmt"
	aoc "github.com/yottta/aoc2022/00_aoc"
	"os"
	"strconv"
	"strings"
)

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

func (s *stack) push(c string) {
	*s = append(*s, c)
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

func (s *stack) top() (string, bool) {
	if s.isEmpty() {
		return "", false
	}
	idx := len(*s) - 1
	return (*s)[idx], true
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)
	const (
		crateSpacing = 4
	)
	var (
		end bool
		//total  int
		stacks []stack
	)
	for {
		if end {
			break
		}
		r, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading string", err)
			end = true
		}
		r = strings.ReplaceAll(r, "\n", "")
		if len(r) == 0 {
			continue
		}
		if stacks == nil {
			stacks = make([]stack, (len(r)+1)/crateSpacing)
		}

		// stacks description
		if strings.Contains(r, "[") {
			for i := 0; i < len(stacks); i++ {
				endIdx := (i+1)*crateSpacing - 1
				if endIdx >= len(r) {
					endIdx = len(r) - 1
				}
				crateContent := r[i*crateSpacing : endIdx]
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
			continue
		}

		// move
		if strings.Contains(r, "move") {
			actions := strings.Split(r, " ")
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
	}
	var b strings.Builder
	for _, s := range stacks {
		top, ok := s.top()
		if ok {
			b.WriteString(top)
		}
	}
	fmt.Println(b.String())
}
