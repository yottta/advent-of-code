package main

import (
	"fmt"
	aoc "github.com/yottta/aoc2022/00_aoc"
	"io"
	"os"
)

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

func main() {
	file, err := os.Open("./input.txt")
	aoc.Must(err)

	content, err := io.ReadAll(file)
	aoc.Must(err)

	queue := newSizedQueue[byte](14)

	for i, r := range content {
		queue.push(r)
		full := queue.full()
		dup := queue.containsDuplicate()
		char := string(r)
		fmt.Println(char)
		if full && !dup {
			fmt.Println(i + 1)
			return
		}

	}
	fmt.Println(-1)
}
