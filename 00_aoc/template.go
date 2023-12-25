package _0_aoc

import (
	"flag"
	"fmt"
	"time"
)

var verbose bool

func BasicRun(part1, part2 func([]string)) {
	var (
		dataFilePath string
		partToRun    string
	)
	flag.StringVar(&dataFilePath, "d", "./input.txt", "The path of the file containing the data for the current problem")
	flag.StringVar(&partToRun, "p", "-1", "The part of the problem to run, in case the problem has more than one parts")
	flag.BoolVar(&verbose, "v", false, "Add this for more information during running if available")
	flag.Parse()

	Verbose(verbose)

	content, err := ReadFile(dataFilePath)
	Must(err)

	switch partToRun {
	case "1":
		runPart1(part1, content)
	case "2":
		runPart2(part2, content)
	default:
		runPart1(part1, content)
		runPart2(part2, content)
	}
}

func runPart1(part1 func([]string), content []string) {
	startedAt := time.Now()
	fmt.Println("--> Part 1 started at", startedAt.Format(time.RFC3339))
	part1(content)
	fmt.Println("--> Part 1 finished at", time.Now().Format(time.RFC3339), "took", time.Since(startedAt))
}

func runPart2(part2 func([]string), content []string) {
	startedAt := time.Now()
	fmt.Println("--> Part 2 started at", startedAt.Format(time.RFC3339))
	part2(content)
	fmt.Println("--> Part 2 finished at", time.Now().Format(time.RFC3339), "took", time.Since(startedAt))
}
