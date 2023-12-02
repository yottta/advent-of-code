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

func part1(content []string) {
	c := initCpu(map[string]*register{
		"X": {
			value: 1,
		},
	})

	valuesPerCycle := map[int]int{
		20:  -1,
		60:  -1,
		100: -1,
		140: -1,
		180: -1,
		220: -1,
	}
	c.cycleListeners = append(c.cycleListeners, func(cycleNo int) {
		_, ok := valuesPerCycle[cycleNo]
		if !ok {
			return
		}
		valuesPerCycle[cycleNo] = c.registers["X"].value
	})

	c.processInput(content)

	var total int
	for cy, regVal := range valuesPerCycle {
		total += cy * regVal
	}
	fmt.Println(total)
}

func part2(content []string) {
	cr := newCrt(240)
	c := initCpu(map[string]*register{
		"X": {
			value: 1,
		},
	})
	c.cycleListeners = append(c.cycleListeners, cr.litBit)
	c.registerChangeListener = append(c.registerChangeListener, cr.repositionSprite)

	c.processInput(content)
}

func initCpu(regs map[string]*register) *cpu {
	return &cpu{
		cyclesDone: 1,
		registers:  regs,
	}
}

type cpu struct {
	cyclesDone             int
	registers              map[string]*register
	cycleListeners         []func(int)
	registerChangeListener []func(regVal int)
}

type register struct {
	value int
}

func (c *cpu) processInput(inputCmds []string) {
	for _, line := range inputCmds {
		commandAndArgs := strings.Split(line, " ")
		c.runCommand(commandAndArgs[0], commandAndArgs[1:]...)
	}
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
		c.notifyRegisterChange(c.registers[regName].value)
	}
}

func (c *cpu) notifyRegisterChange(newVal int) {
	for _, l := range c.registerChangeListener {
		l(newVal)
	}
}

func (c *cpu) notifyCycle(noOfCycles int) {
	for _, l := range c.cycleListeners {
		l(noOfCycles)
	}
}

func (c *cpu) increaseCycle() {
	c.cyclesDone++
	c.notifyCycle(c.cyclesDone)
}

type crt struct {
	bits       []string
	currentBit int
	currentRow int
	spritePos  int
}

func newCrt(size int) *crt {
	return &crt{
		bits:       make([]string, size),
		currentBit: 0,
	}
}

func (cr *crt) litBit(_ int) {
	upp := cr.spritePos - 1 + 2
	low := cr.spritePos - 1
	bit := cr.currentBit - cr.currentRow*40
	if inRange(bit, low, upp) {
		cr.bits[cr.currentBit] = "X"
	} else {
		cr.bits[cr.currentBit] = " "
	}

	if cr.currentBit%40 == 0 {
		fmt.Println()
	}
	fmt.Printf(cr.bits[cr.currentBit])

	cr.currentBit++
	cr.currentRow = cr.currentBit / 40
}

func (cr *crt) repositionSprite(start int) {
	cr.spritePos = start
}

func inRange(val, low, up int) bool {
	if val < low {
		return false
	}
	if val > up {
		return false
	}
	return true
}
