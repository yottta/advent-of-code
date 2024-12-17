package main

import (
	"strconv"

	aoc "github.com/yottta/advent-of-code/00_aoc"
)

func transform(val uint64, blinks int) uint64 {
	count := uint64(0) // the initial pebble
	if blinks == 0 {
		return max(1, count)
	}
	for {
		if blinks == 0 {
			return max(1, count)
		}
		if val == 0 {
			val = 1
		} else if strVal := strconv.FormatUint(val, 10); len(strVal)%2 == 0 {
			currVal, err := strconv.ParseUint(strVal[:len(strVal)/2], 10, 0)
			aoc.Must(err)
			newVal, err := strconv.ParseUint(strVal[len(strVal)/2:], 10, 0)
			aoc.Must(err)
			count += transform(currVal, blinks-1)
			count += transform(newVal, blinks-1)
			break
		} else {
			val = val * 2024
		}
		blinks--
	}
	return count
}

//
//func printPebs(pebs []*pebble) {
//	for _, p := range pebs {
//		var (
//			next = p
//		)
//		for next != nil {
//			fmt.Printf("%d ", next.val)
//			next = next.next
//		}
//	}
//
//	fmt.Println()
//}
//
//func count(first *pebble) int64 {
//	var (
//		p   = first
//		sum int64
//	)
//	for p != nil {
//		sum++
//		p = p.next
//	}
//	return sum
//}

func maxUint64(v1, v2 uint64) uint64 {
	if v1 > v2 {
		return v1
	}
	return v2
}
