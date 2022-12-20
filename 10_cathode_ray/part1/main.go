package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	aoc "github.com/yottta/aoc2022/00_aoc"
)

func main() {
	file, err := os.Open("../input.txt")
	//file, err := os.Open("../input_test.txt")
	aoc.Must(err)
	reader := bufio.NewReader(file)

	c := initCpu(map[string]*register{
		"X": {
			value: 1,
		},
	}, 20, 60, 100, 140, 180, 220)
	var (
		end bool
		idx int
	)
	for {
		idx++
		if end {
			break
		}
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading string", err)
			end = true
		}
		line = strings.ReplaceAll(line, "\n", "")
		if len(line) == 0 {
			continue
		}
		commandAndArgs := strings.Split(line, " ")
		c.runCommand(commandAndArgs[0], commandAndArgs[1:]...)
	}
	var total int
	for cy, s := range c.cycleSnapshot {
		total += cy * s["X"].value
	}
	fmt.Println(total)
}

func initCpu(regs map[string]*register, snapshotOnCycle ...int) *cpu {
	marks := map[int]map[string]register{}
	for _, i := range snapshotOnCycle {
		marks[i] = map[string]register{}
	}
	return &cpu{
		cyclesDone:    1,
		registers:     regs,
		cycleSnapshot: marks,
	}
}

type cpu struct {
	cyclesDone    int
	cycleSnapshot map[int]map[string]register
	registers     map[string]*register
}

type register struct {
	value int
}

func (c *cpu) runCommand(cmd string, args ...string) {
	if cmd == "noop" {
		c.increaseCycle()
		return
	}
	if strings.HasPrefix(cmd, "add") {
		c.increaseCycle()
		s, err := strconv.Atoi(args[0])
		aoc.Must(err)
		regName := strings.ToUpper(strings.TrimPrefix(cmd, "add"))
		c.registers[regName].value += s
		c.increaseCycle()
	}
}

func (c *cpu) takeSnapshotIfNeeded() {
	_, ok := c.cycleSnapshot[c.cyclesDone]
	if ok {
		snapshot := make(map[string]register, len(c.registers))
		for rn, r := range c.registers {
			snapshot[rn] = register{
				value: r.value,
			}
		}
		c.cycleSnapshot[c.cyclesDone] = snapshot
	}
}
func (c *cpu) increaseCycle() {
	c.cyclesDone++
	c.takeSnapshotIfNeeded()
}
