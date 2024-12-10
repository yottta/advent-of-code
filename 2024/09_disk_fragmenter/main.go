package main

import (
	"fmt"
	"strconv"

	aoc "github.com/yottta/advent-of-code/00_aoc"
	aoc2024 "github.com/yottta/advent-of-code/2024"
)

type block struct {
	fileId int // -1 is free space
	size   int
}

func main() {
	aoc2024.Run(part1, part2, false, "09")
}

func part1(content []string) {
	m := parseMemory(content[0])
	compressBlocks(m)
	fmt.Println(checksum(m))
}

func part2(content []string) {
	m := parseMemory(content[0])
	compressFiles(m)
	fmt.Println(checksum(m))
}

func parseMemory(content string) []block {
	var out []block
	nextId := 0
	for i, r := range content {
		id := -1
		if i%2 == 0 {
			id = nextId
		}
		count, err := strconv.Atoi(string(r))
		aoc.Must(err)
		for j := 0; j < count; j++ {
			out = append(out, block{fileId: id, size: count})
		}

		if id != -1 {
			nextId++
		}
	}
	return out
}

func compressBlocks(in []block) {
	var (
		nextToBeMovedMemoryBlockIdx = len(in) - 1
		nextToBeMovedFreeBlockIdx   = 0
	)
	for {
		for i := nextToBeMovedFreeBlockIdx; i < len(in); i++ {
			nextToBeMovedFreeBlockIdx = i
			if in[nextToBeMovedFreeBlockIdx].fileId == -1 {
				break
			}
		}
		for i := nextToBeMovedMemoryBlockIdx; i >= 0; i-- {
			nextToBeMovedMemoryBlockIdx = i
			if in[nextToBeMovedMemoryBlockIdx].fileId != -1 {
				break
			}
		}
		if nextToBeMovedMemoryBlockIdx < nextToBeMovedFreeBlockIdx {
			break
		}
		memBl := in[nextToBeMovedMemoryBlockIdx]
		in[nextToBeMovedMemoryBlockIdx] = in[nextToBeMovedFreeBlockIdx]
		in[nextToBeMovedFreeBlockIdx] = memBl
	}
}

func compressFiles(in []block) {
	i := len(in) - 1
	for {
		if i <= 0 {
			break
		}
		if in[i].fileId == -1 {
			i--
			continue
		}
		fSize := in[i].size
		idx := nextStartIdxOfFreeMemoryOfSize(in, fSize, i)
		if idx == -1 {
			i -= fSize
			continue
		}
		if idx+fSize >= len(in) {
			i -= fSize
			continue
		}
		for j := 0; j < fSize; j++ {
			mem := in[i-j]
			in[i-j] = in[idx+j]
			in[idx+j] = mem
		}
		i -= fSize
	}
}

func nextStartIdxOfFreeMemoryOfSize(mem []block, size int, notAfter int) int {
	foundStartIdx := -1
	freeSize := 0
	for i := 0; i < notAfter; i++ {
		bl := mem[i]
		if bl.fileId != -1 {
			foundStartIdx = -1
			freeSize = 0
			continue
		}
		if foundStartIdx == -1 {
			foundStartIdx = i
		}
		freeSize++
		if size == freeSize {
			return foundStartIdx
		}
	}
	return foundStartIdx
}

func checksum(in []block) int {
	var out int
	for i, b := range in {
		if b.fileId == -1 {
			continue
		}
		out += b.fileId * i
	}
	return out
}

func printMemory(in []block) {
	for i := 0; i < len(in); i++ {
		dat := in[i]
		if dat.fileId == -1 {
			fmt.Print(".")
			continue
		}
		fmt.Print(in[i].fileId)
	}
	fmt.Println()
}
