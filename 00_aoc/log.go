package _0_aoc

import "fmt"

var l logger

func init() {
	l = logger{}
}

type logger struct {
	verbose bool
}

func (l logger) infof(template string, args ...any) {
	if !l.verbose {
		return
	}
	fmt.Printf(template, args...)
}

func (l logger) info(msg any) {
	if !l.verbose {
		return
	}
	fmt.Println(msg)
}

func Verbose(v bool) {
	l.verbose = v
}

func Logf(template string, args ...any) {
	l.infof(template, args...)
}

func Log(o any) {
	l.info(o)
}
