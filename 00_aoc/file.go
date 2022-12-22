package _0_aoc

import (
	"os"
	"strings"
)

func ReadFile(p string) ([]string, error) {
	c, err := os.ReadFile(p)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.TrimSpace(string(c)), "\n")
	return lines, nil
}

func ReadFileBytes(p string) ([]byte, error) {
	return os.ReadFile(p)
}
