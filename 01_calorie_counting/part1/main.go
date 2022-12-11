package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)
	var (
		curVal, curIdx, maxVal, maxIdx int
	)
	curIdx++
	for {
		r, err := reader.ReadString('\n')
		if err != nil {
			if curVal > maxVal {
				maxVal = curVal
				maxIdx = curIdx
			}
			fmt.Println("error reading string", err)
			break
		}
		r = strings.ReplaceAll(r, "\n", "")
		if len(strings.TrimSpace(r)) == 0 {
			if curVal > maxVal {
				maxVal = curVal
				maxIdx = curIdx
			}
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
	fmt.Println(maxIdx, maxVal)
}
