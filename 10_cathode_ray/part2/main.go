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
	aoc.Must(err)
	reader := bufio.NewReader(file)

	cr := newCrt(240)
	c := initCpu(map[string]*register{
		"X": {
			value: 1,
		},
	})
	c.cycleListeners = append(c.cycleListeners, cr.litBit)
	c.registerChangeListener = append(c.registerChangeListener, cr.repositionSprite)
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
	fmt.Println(c.cyclesDone)
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
	cycleListeners         []func()
	registerChangeListener []func(regVal int)
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
		c.notifyRegisterChange(c.registers[regName].value)
	}
}

func (c *cpu) notifyRegisterChange(newVal int) {
	for _, l := range c.registerChangeListener {
		l(newVal)
	}
}

func (c *cpu) notifyCycle() {
	for _, l := range c.cycleListeners {
		l()
	}
}

func (c *cpu) increaseCycle() {
	c.cyclesDone++
	c.notifyCycle()
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

func (cr *crt) litBit() {
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
