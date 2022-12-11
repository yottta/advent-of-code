package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func priority(r uint8) int {
	if r >= 97 {
		return int(r - 96) // lower case
	}
	return int(r - 38) // upper case
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)
	var (
		total int
		end   bool
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
			break
		}
		firstCompartment := make(map[uint8]struct{})
		fmt.Println(r)
		for i := 0; i < len(r)/2; i++ {
			item := r[i]
			firstCompartment[item] = struct{}{}
		}
		common := make(map[uint8]struct{})
		for i := len(r) / 2; i < len(r); i++ {
			_, ok := firstCompartment[r[i]]
			if ok {
				common[r[i]] = struct{}{}
			}
		}
		var rt int
		for item := range common {
			itemPrio := priority(item)
			fmt.Println("common ", string(item), itemPrio)
			rt += itemPrio
		}
		total += rt
		fmt.Println()
	}
	fmt.Println(total)
}
