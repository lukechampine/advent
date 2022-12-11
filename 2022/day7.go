package main

import (
	"path"
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2022, 7)

func main() {
	dirsizes := make(map[string]int)
	lines := utils.Lines(input)

	var dir string
	for i := 0; i < len(lines); {
		fields := strings.Fields(lines[i])
		switch fields[1] {
		case "cd":
			dir = path.Clean(path.Join(dir, fields[2]))
			i++
		case "ls":
			for i++; i < len(lines) && lines[i][0] != '$'; i++ {
				fields := strings.Fields(lines[i])
				if fields[0] != "dir" {
					for d := dir; ; d = path.Dir(d) {
						dirsizes[d] += utils.Atoi(fields[0])
						if d == "/" {
							break
						}
					}
				}
			}
		}
	}

	var sum int
	for _, sz := range dirsizes {
		if sz <= 100000 {
			sum += sz
		}
	}
	utils.Println(sum)

	toFree := dirsizes["/"] - 40000000
	best := dirsizes["/"]
	for _, sz := range dirsizes {
		if sz >= toFree && sz < best {
			best = sz
		}
	}
	utils.Println(best)
}
