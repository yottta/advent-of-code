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
		total, i    int
		end         bool
		commonGroup map[uint8]struct{}
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
		content := make(map[uint8]struct{})
		fmt.Println(r)
		for i := 0; i < len(r); i++ {
			content[r[i]] = struct{}{}
		}

		if i%3 == 0 {
			var gc int
			for item := range commonGroup {
				itemPrio := priority(item)
				gc += itemPrio
			}
			total += gc
			commonGroup = content
			i++
			continue
		}
		for c := range commonGroup {
			_, ok := content[c]
			if !ok {
				delete(commonGroup, c)
			}
		}
		i++
	}
	fmt.Println(total)
}
