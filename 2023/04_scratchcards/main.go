package main

import (
	"fmt"
	"strconv"
	"strings"

	aoc "github.com/yottta/advent-of-code/00_aoc"
)

func main() {
	aoc.BasicRun(part1, part2)
}

type card struct {
	id             int
	winningNumbers map[int]struct{}
	playerNumbers  map[int]struct{}
	wonPoints      int
	matchedNumbers int
}

func part1(content []string) {
	cards := buildCards(content)
	allPoints := 0
	for _, c := range cards {
		allPoints += c.wonPoints
	}
	fmt.Println(allPoints)

}

func buildCards(content []string) (cards []card) {
	for i, line := range content {
		parts := strings.Split(line, ":")
		numbers := strings.Split(parts[1], "|")
		winningNumbers := strings.Split(strings.TrimSpace(numbers[0]), " ")
		playerNumbers := strings.Split(strings.TrimSpace(numbers[1]), " ")
		c := card{
			id:             i + 1,
			winningNumbers: map[int]struct{}{},
			playerNumbers:  map[int]struct{}{},
		}
		for _, n := range winningNumbers {
			if n == "" {
				continue
			}
			winningNumber, err := strconv.Atoi(strings.TrimSpace(n))
			aoc.Must(err)
			c.winningNumbers[winningNumber] = struct{}{}
		}

		for _, n := range playerNumbers {
			if n == "" {
				continue
			}
			playerNumber, err := strconv.Atoi(strings.TrimSpace(n))
			aoc.Must(err)
			c.playerNumbers[playerNumber] = struct{}{}
			_, ok := c.winningNumbers[playerNumber]
			if ok {
				c.matchedNumbers++
				if c.wonPoints == 0 {
					c.wonPoints = 1
				} else {
					c.wonPoints *= 2
				}
			}
		}
		cards = append(cards, c)

	}
	return cards
}

func part2(content []string) {
	cards := buildCards(content)
	newWonCards := map[int]int{}
	var totalCards = 0
	for _, c := range cards {
		totalCards++
		countWonCards(c, newWonCards)
		copies, ok := newWonCards[c.id]
		if !ok {
			continue
		}
		totalCards += copies
		if c.matchedNumbers == 0 {
			continue
		}
		for i := 0; i < copies; i++ {
			countWonCards(c, newWonCards)
		}
	}
	fmt.Println(totalCards)
}

func countWonCards(c card, countIn map[int]int) {
	nextCardID := c.id + 1
	for i := 0; i < c.matchedNumbers; i++ {
		alreadyWon, _ := countIn[nextCardID+i]
		alreadyWon++
		countIn[nextCardID+i] = alreadyWon
	}
}
