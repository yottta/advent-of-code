package main

import (
	"strconv"

	aoc "github.com/yottta/advent-of-code/00_aoc"
)

type cacheEntry struct {
	val          uint64
	additionlVal *uint64
}

type cache struct {
	s map[uint64]cacheEntry
}

var c = &cache{map[uint64]cacheEntry{}}

func transform(val uint64, blinks int) uint64 {
	count := uint64(0) // the initial pebble
	if blinks == 0 {
		return 1
	}
	for {
		if blinks == 0 {
			return max(1, count)
		}
		if val == 0 {
			val = 1
			blinks--
			continue
		}
		e, ok := fromCache(val)
		if ok {
			if e.additionlVal == nil {
				val = e.val
			} else {
				count += transform(e.val, blinks-1)
				count += transform(*e.additionlVal, blinks-1)
				break
			}
		} else {
			if strVal := strconv.FormatUint(val, 10); len(strVal)%2 == 0 {
				currVal, err := strconv.ParseUint(strVal[:len(strVal)/2], 10, 0)
				aoc.Must(err)
				newVal, err := strconv.ParseUint(strVal[len(strVal)/2:], 10, 0)
				aoc.Must(err)
				count += transform(currVal, blinks-1)
				count += transform(newVal, blinks-1)
				toCache(val, cacheEntry{
					val:          currVal,
					additionlVal: &newVal,
				})
				break
			} else {
				newVal := val * 2024
				toCache(val, cacheEntry{
					val: newVal,
				})
				val = newVal
			}
		}
		blinks--
	}
	return count
}

func fromCache(val uint64) (cacheEntry, bool) {
	v, ok := c.s[val]
	return v, ok
}

func toCache(val uint64, entry cacheEntry) {
	c.s[val] = entry
}
