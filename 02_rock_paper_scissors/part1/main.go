package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

func parseMove(s string) (Move, error) {
	if s == "A" || s == "X" {
		return Rock, nil
	}
	if s == "B" || s == "Y" {
		return Paper, nil
	}
	if s == "C" || s == "Z" {
		return Scissor, nil
	}
	return 0, fmt.Errorf("invalid value %s", s)
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)
	var totalPoints int
	for {
		r, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading string", err)
			break
		}
		r = strings.ReplaceAll(r, "\n", "")
		split := strings.Split(r, " ")
		if len(split) != 2 {
			fmt.Printf("invalid input %s\n", r)
			continue
		}
		enemyMove, err := parseMove(split[0])
		if err != nil {
			panic(err)
		}
		myMove, err := parseMove(split[1])
		if err != nil {
			panic(err)
		}
		totalPoints += myMove.fightPoints(enemyMove) + int(myMove)
	}
	fmt.Println(totalPoints)
}
