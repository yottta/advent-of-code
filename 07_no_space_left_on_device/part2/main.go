package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	
	aoc "github.com/yottta/aoc2022/00_aoc"
)

const (
	diskSize   = 70000000
	neededSize = 30000000
)

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

func main() {
	f, err := os.Open("./input.txt")
	aoc.Must(err)

	reader := bufio.NewReader(f)
	s := session{
		dirs:  make(map[string][]file),
		sizes: make(map[string]int),
	}
	var (
		end bool
	)
	for {
		if end {
			break
		}
		r, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading string", err)
			end = true
		}
		r = strings.ReplaceAll(r, "\n", "")
		if len(r) == 0 {
			continue
		}

		if strings.HasPrefix(r, "$ ") {
			command := strings.TrimPrefix(r, "$ ")
			commandParts := strings.Split(command, " ")
			if commandParts[0] == "cd" {
				aoc.Must(s.cd(commandParts[1]))
				continue
			}
			if commandParts[0] == "ls" {
				continue
			}
		}
		if strings.HasPrefix(r, "dir ") {
			lsDirParts := strings.Split(r, " ")
			s.addFile(file{
				name: lsDirParts[1],
				dir:  true,
				size: 0,
			})
			continue
		}
		fileParts := strings.Split(r, " ")
		size, err := strconv.Atoi(fileParts[0])
		aoc.Must(err)
		s.addFile(file{
			name: fileParts[1],
			dir:  false,
			size: size,
		})
	}

	var (
		used   = s.sizes["/"]
		empty  = diskSize - used
		needed = neededSize - empty
	)

	var min int
	for _, size := range s.sizes {
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
