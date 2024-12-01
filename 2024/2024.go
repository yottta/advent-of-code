package _2024

import aoc "github.com/yottta/advent-of-code/00_aoc"

const Year = "2024"

func Run(part1, part2 func([]string), testData bool, problemIdx string) {
	aoc.RunV3(
		Year,
		problemIdx,
		aoc.RunCfg{
			RunName:  "1",
			FileName: "input",
			F:        part1,
		},
		aoc.RunCfg{
			RunName:  "2",
			FileName: "input",
			F:        part2,
		},
	)
}
