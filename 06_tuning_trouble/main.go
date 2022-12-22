package main

import (
	"flag"
	"fmt"

	aoc "github.com/yottta/aoc2022/00_aoc"
)

func main() {
	var (
		dataFilePath string
		partToRun    string
	)
	flag.StringVar(&dataFilePath, "d", "./input.txt", "The path of the file containing the data for the current problem")
	flag.StringVar(&partToRun, "p", "1", "The part of the problem to run, in case the problem has more than one parts")
	flag.Parse()

	content, err := aoc.ReadFileBytes(dataFilePath)
	aoc.Must(err)

	switch partToRun {
	case "1":
		solve(content, 4)
	case "2":
		solve(content, 14)
	default:
		panic(fmt.Errorf("no part '%s' configured", partToRun))
	}
}

func solve(content []byte, queueSize int) {
	queue := newSizedQueue[byte](queueSize)

	for i, r := range content {
		queue.push(r)
		if queue.full() && !queue.containsDuplicate() {
			fmt.Println(i + 1)
			return
		}

	}
	fmt.Println(-1)
}

type sqConstraint interface {
	rune | byte
}

type sizedQueue[T sqConstraint] struct {
	maxSize int
	content []T
}

func newSizedQueue[T sqConstraint](size int) *sizedQueue[T] {
	d := make([]T, 0, size)
	return &sizedQueue[T]{
		maxSize: size,
		content: d,
	}
}

func (sz *sizedQueue[T]) push(entry T) {
	if len(sz.content) == sz.maxSize {
		sz.content = sz.content[1:len(sz.content)]
	}
	sz.content = append(sz.content, entry)
}

func (sz *sizedQueue[T]) containsDuplicate() bool {
	found := map[T]struct{}{}
	for _, e := range sz.content {
		_, ok := found[e]
		if ok {
			return true
		}
		found[e] = struct{}{}
	}
	return false
}

func (sz *sizedQueue[T]) full() bool {
	return len(sz.content) == sz.maxSize
}
