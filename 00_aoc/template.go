package _0_aoc

import (
	"flag"
	"fmt"
)

func BasicRun(part1, part2 func([]string)) {
	var (
		dataFilePath string
		partToRun    string
	)
	flag.StringVar(&dataFilePath, "d", "./input.txt", "The path of the file containing the data for the current problem")
	flag.StringVar(&partToRun, "p", "1", "The part of the problem to run, in case the problem has more than one parts")
	flag.Parse()

	content, err := ReadFile(dataFilePath)
	Must(err)

	switch partToRun {
	case "1":
		part1(content)
	case "2":
		part2(content)
	default:
		panic(fmt.Errorf("no part '%s' configured", partToRun))
	}
}
