package main

import (
	"flag"
	"fmt"
	"path"
	"strconv"
	"strings"

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

	content, err := aoc.ReadFile(dataFilePath)
	aoc.Must(err)

	switch partToRun {
	case "1":
		part1(content)
	case "2":
		part2(content)
	default:
		panic(fmt.Errorf("no part '%s' configured", partToRun))
	}
}

func part1(content []string) {
	sess := buildSession(content)
	var total int
	for _, s := range sess.sizes {
		if s > 100_000 {
			continue
		}
		total += s
	}
	fmt.Println("total", total)
}

func part2(content []string) {
	sess := buildSession(content)

	const (
		diskSize   = 70000000
		neededSize = 30000000
	)

	var (
		used   = sess.sizes["/"]
		empty  = diskSize - used
		needed = neededSize - empty
	)

	var min int
	for _, size := range sess.sizes {
		if size >= needed {
			if min == 0 {
				min = size
				continue
			}
			if min > size {
				min = size
				continue
			}
		}
	}
	fmt.Println(min)
}

func buildSession(content []string) *session {
	s := &session{
		dirs:  make(map[string][]file),
		sizes: make(map[string]int),
	}
	for _, line := range content {
		if strings.HasPrefix(line, "$ ") {
			processCommand(line, s)
			continue
		}
		if strings.HasPrefix(line, "dir ") {
			processDir(line, s)
			continue
		}
		processFile(line, s)
	}
	return s
}

func processCommand(cmd string, s *session) {
	command := strings.TrimPrefix(cmd, "$ ")
	commandParts := strings.Split(command, " ")
	switch commandParts[0] {
	case "cd":
		aoc.Must(s.cd(commandParts[1]))
	}
}

func processDir(dirDesc string, s *session) {
	lsDirParts := strings.Split(dirDesc, " ")
	s.addFile(file{
		name: lsDirParts[1],
		dir:  true,
		size: 0,
	})
}

func processFile(fileDesc string, s *session) {
	fileParts := strings.Split(fileDesc, " ")
	size, err := strconv.Atoi(fileParts[0])
	aoc.Must(err)
	s.addFile(file{
		name: fileParts[1],
		dir:  false,
		size: size,
	})
}

type session struct {
	currDir *string

	dirs  map[string][]file
	sizes map[string]int
}

type file struct {
	name string
	dir  bool
	size int
}

func (s *session) cd(dirName string) error {
	switch dirName {
	case "/":
		_, ok := s.dirs[dirName]
		if !ok {
			s.dirs[dirName] = []file{}
		}
		s.currDir = &dirName
		return nil
	case "..":
		if s.currDir == nil {
			return fmt.Errorf("no parent dir because nothing set yet")
		}
		if s.currDir == nil || *s.currDir == "/" {
			return nil
		}

		base := path.Dir(*s.currDir)
		_, ok := s.dirs[base]
		if !ok {
			return fmt.Errorf("no dir found for %s", base)
		}
		s.currDir = &base
		return nil
	default:
		_, ok := s.dirs[dirName]
		if !ok {
			s.dirs[dirName] = []file{}
		}
		fullPath := path.Join(*s.currDir, dirName)
		s.currDir = &fullPath
		return nil
	}
}

func (s *session) addFile(f file) {
	if s.currDir == nil {
		aoc.Must(s.cd("/"))
	}
	files := s.dirs[*s.currDir]
	files = append(files, f)
	s.dirs[*s.currDir] = files
	cDir := *s.currDir
	if !f.dir {
		for {
			s.sizes[cDir] += f.size
			if cDir == "/" {
				break
			}
			cDir = path.Dir(cDir)
		}
	}
}
