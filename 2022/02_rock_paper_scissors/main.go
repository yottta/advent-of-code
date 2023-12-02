package main

import (
	"fmt"
	"strings"

	aoc "github.com/yottta/advent-of-code/00_aoc"
)

func main() {
	aoc.BasicRun(part1, part2)
}

func part1(content []string) {
	var totalPoints int
	for _, line := range content {
		split := strings.Split(line, " ")
		enemyMove := parseMove(split[0])
		myMove := parseMove(split[1])
		totalPoints += myMove.fightPoints(enemyMove) + int(myMove)
	}

	fmt.Println(totalPoints)
}

func part2(content []string) {
	var totalPoints int
	for _, line := range content {
		split := strings.Split(line, " ")
		enemyMove := parseMove(split[0])
		myMove := moveToPerform(enemyMove, split[1])
		totalPoints += myMove.fightPoints(enemyMove) + int(myMove)
	}

	fmt.Println(totalPoints)
}

type Move int

const (
	Rock Move = iota + 1
	Paper
	Scissor
)

// key move beats the value move
var moveBeatMove = map[Move]Move{
	Rock:    Scissor,
	Paper:   Rock,
	Scissor: Paper,
}

func (m Move) fightPoints(mo Move) int {
	// draw
	if m == mo {
		return 3
	}

	winsAgainst, ok := moveBeatMove[m]
	if !ok {
		panic(fmt.Errorf("%d has no other move that it wins against. wrongly configured", m))
	}
	// win
	if winsAgainst == mo {
		return 6
	}
	// loss
	return 0
}

func parseMove(s string) Move {
	if s == "A" || s == "X" {
		return Rock
	}
	if s == "B" || s == "Y" {
		return Paper
	}
	if s == "C" || s == "Z" {
		return Scissor
	}
	panic(fmt.Errorf("invalid value %s", s))
}

func moveToPerform(enemy Move, indication string) Move {
	move, ok := moveBeatMove[enemy]
	if !ok {
		panic(fmt.Errorf("wrong input. no winning move for %d", enemy))
	}
	switch indication {
	case "X":
		return move
	case "Y":
		return enemy
	case "Z":
		for k, v := range moveBeatMove {
			if v == enemy {
				return k
			}
		}
		panic(fmt.Errorf("misconfiguration. there is no move that beats %d", enemy))
	default:
		panic(fmt.Errorf("wrong indication %s", indication))
	}
}
