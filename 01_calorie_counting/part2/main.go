package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type elf struct {
	idx      int
	calories int
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)
	var (
		elves          []elf
		curVal, curIdx int
	)
	curIdx++
	for {
		r, err := reader.ReadString('\n')
		if err != nil {
			elves = append(elves, elf{
				idx:      curIdx,
				calories: curVal,
			})
			fmt.Println("error reading string", err)
			break
		}
		r = strings.ReplaceAll(r, "\n", "")
		if len(strings.TrimSpace(r)) == 0 {
			elves = append(elves, elf{
				idx:      curIdx,
				calories: curVal,
			})
			curVal = 0
			curIdx++
			continue
		}
		atoi, err := strconv.Atoi(r)
		if err != nil {
			fmt.Printf("failed to convert %s to int: %s\n", r, err)
			panic(err)
		}
		curVal += atoi
	}
	sort.Slice(elves, func(i, j int) bool {
		return elves[i].calories < elves[j].calories
	})
	var total int
	for i := len(elves) - 3; i < len(elves); i++ {
		total += elves[i].calories
	}
	fmt.Println(total)

}
