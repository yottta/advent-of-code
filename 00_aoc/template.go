package _0_aoc

import (
	"flag"
	"fmt"
	"time"
)

var verbose bool

// Deprecated - use BasiccRunV2 instead
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
		run(part1, "1", content)
	case "2":
		run(part2, "2", content)
	default:
		run(part1, "1", content)
		run(part2, "2", content)
	}
}

func BasicRunV2(part1, part2 func([]string), testData bool, problemIdx, year string) {
	var (
		dataFilePath string
		partToRun    string
	)
	if len(problemIdx) == 1 {
		problemIdx = "0" + problemIdx
	}
	inputFileName := "input"
	if testData {
		inputFileName = "test_" + inputFileName
	}
	defaultDataFile := fmt.Sprintf("../../../advent-of-code-data/%s/%s/%s", year, problemIdx, inputFileName)
	flag.StringVar(&dataFilePath, "d", defaultDataFile, "The path of the file containing the data for the current problem")
	flag.StringVar(&partToRun, "p", "-1", "The part of the problem to run, in case the problem has more than one parts")
	flag.BoolVar(&verbose, "v", false, "Add this for more information during running if available")
	flag.Parse()

	Verbose(verbose)

	content, err := ReadFile(dataFilePath)
	Must(err)

	switch partToRun {
	case "1":
		run(part1, "1", content)
	case "2":
		run(part2, "2", content)
	default:
		run(part1, "1", content)
		run(part2, "2", content)
	}
}

type RunCfg struct {
	RunName  string
	FileName string
	F        func([]string)
}

func RunV3(year, dataDir string, cfg ...RunCfg) {
	var (
		dataFilePathDir string
		dataFileName    string
	)
	if len(dataDir) == 1 {
		dataDir = "0" + dataDir
	}
	flag.StringVar(&dataFilePathDir, "d", fmt.Sprintf("../../../advent-of-code-data/%s/%s", year, dataDir), "The path of the file containing the data for the current problem")
	flag.StringVar(&dataFileName, "n", "", "The name of the input data file. This will be combined with the 'd' argument which will generate the full path of the input data file. If not given, the name will be used from the received RunCfg")
	flag.BoolVar(&verbose, "v", false, "Add this for more information during running if available")
	flag.Parse()

	Verbose(verbose)

	var (
		content []string
		err     error
	)
	if len(dataFileName) > 0 {
		content, err = ReadFile(fmt.Sprintf("%s/%s", dataFilePathDir, dataFileName))
		Must(err)
	}

	for _, r := range cfg {
		runContent := content
		if len(runContent) == 0 {
			runContent, err = ReadFile(fmt.Sprintf("%s/%s", dataFilePathDir, r.FileName))
			Must(err)
		}
		run(r.F, r.RunName, runContent)
	}
}

func run(f func([]string), testName string, content []string) {
	startedAt := time.Now()
	fmt.Printf("--> Part '%s' started at %s\n", testName, startedAt.Format(time.RFC3339))
	f(content)
	fmt.Printf("--> Part '%s' finished at %s and took %s\n", testName, time.Now().Format(time.RFC3339), time.Since(startedAt))
}
