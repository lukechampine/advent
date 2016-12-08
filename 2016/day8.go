package main

import (
	"strings"

	"github.com/lukechampine/advent/utils"
)

const input = `rect 1x1
rotate row y=0 by 2
rect 1x1
rotate row y=0 by 5
rect 1x1
rotate row y=0 by 3
rect 1x1
rotate row y=0 by 3
rect 2x1
rotate row y=0 by 5
rect 1x1
rotate row y=0 by 5
rect 4x1
rotate row y=0 by 2
rect 1x1
rotate row y=0 by 2
rect 1x1
rotate row y=0 by 5
rect 4x1
rotate row y=0 by 3
rect 2x1
rotate row y=0 by 5
rect 4x1
rotate row y=0 by 2
rect 1x2
rotate row y=1 by 6
rotate row y=0 by 2
rect 1x2
rotate column x=32 by 1
rotate column x=23 by 1
rotate column x=13 by 1
rotate row y=0 by 6
rotate column x=0 by 1
rect 5x1
rotate row y=0 by 2
rotate column x=30 by 1
rotate row y=1 by 20
rotate row y=0 by 18
rotate column x=13 by 1
rotate column x=10 by 1
rotate column x=7 by 1
rotate column x=2 by 1
rotate column x=0 by 1
rect 17x1
rotate column x=16 by 3
rotate row y=3 by 7
rotate row y=0 by 5
rotate column x=2 by 1
rotate column x=0 by 1
rect 4x1
rotate column x=28 by 1
rotate row y=1 by 24
rotate row y=0 by 21
rotate column x=19 by 1
rotate column x=17 by 1
rotate column x=16 by 1
rotate column x=14 by 1
rotate column x=12 by 2
rotate column x=11 by 1
rotate column x=9 by 1
rotate column x=8 by 1
rotate column x=7 by 1
rotate column x=6 by 1
rotate column x=4 by 1
rotate column x=2 by 1
rotate column x=0 by 1
rect 20x1
rotate column x=47 by 1
rotate column x=40 by 2
rotate column x=35 by 2
rotate column x=30 by 2
rotate column x=10 by 3
rotate column x=5 by 3
rotate row y=4 by 20
rotate row y=3 by 10
rotate row y=2 by 20
rotate row y=1 by 16
rotate row y=0 by 9
rotate column x=7 by 2
rotate column x=5 by 2
rotate column x=3 by 2
rotate column x=0 by 2
rect 9x2
rotate column x=22 by 2
rotate row y=3 by 40
rotate row y=1 by 20
rotate row y=0 by 20
rotate column x=18 by 1
rotate column x=17 by 2
rotate column x=16 by 1
rotate column x=15 by 2
rotate column x=13 by 1
rotate column x=12 by 1
rotate column x=11 by 1
rotate column x=10 by 1
rotate column x=8 by 3
rotate column x=7 by 1
rotate column x=6 by 1
rotate column x=5 by 1
rotate column x=3 by 1
rotate column x=2 by 1
rotate column x=1 by 1
rotate column x=0 by 1
rect 19x1
rotate column x=44 by 2
rotate column x=40 by 3
rotate column x=29 by 1
rotate column x=27 by 2
rotate column x=25 by 5
rotate column x=24 by 2
rotate column x=22 by 2
rotate column x=20 by 5
rotate column x=14 by 3
rotate column x=12 by 2
rotate column x=10 by 4
rotate column x=9 by 3
rotate column x=7 by 3
rotate column x=3 by 5
rotate column x=2 by 2
rotate row y=5 by 10
rotate row y=4 by 8
rotate row y=3 by 8
rotate row y=2 by 48
rotate row y=1 by 47
rotate row y=0 by 40
rotate column x=47 by 5
rotate column x=46 by 5
rotate column x=45 by 4
rotate column x=43 by 2
rotate column x=42 by 3
rotate column x=41 by 2
rotate column x=38 by 5
rotate column x=37 by 5
rotate column x=36 by 5
rotate column x=33 by 1
rotate column x=28 by 1
rotate column x=27 by 5
rotate column x=26 by 5
rotate column x=25 by 1
rotate column x=23 by 5
rotate column x=22 by 1
rotate column x=21 by 2
rotate column x=18 by 1
rotate column x=17 by 3
rotate column x=12 by 2
rotate column x=11 by 2
rotate column x=7 by 5
rotate column x=6 by 5
rotate column x=5 by 4
rotate column x=3 by 5
rotate column x=2 by 5
rotate column x=1 by 3
rotate column x=0 by 4`

type screen [6][50]bool

func (s *screen) rect(a, b int) {
	for i := 0; i < a; i++ {
		for j := 0; j < b; j++ {
			s[j][i] = true
		}
	}
}

func (s *screen) rotateRow(a, b int) {
	buf := make([]bool, len(s[a]))
	for i := range s[a] {
		buf[(i+b)%len(buf)] = s[a][i]
	}
	for i := range s[a] {
		s[a][i] = buf[i]
	}
}

func (s *screen) rotateCol(a, b int) {
	buf := make([]bool, len(s))
	for i := range s {
		buf[(i+b)%len(buf)] = s[i][a]
	}
	for i := range s {
		s[i][a] = buf[i]
	}
}

type cmd struct {
	fn   func(*screen, int, int)
	a, b int
}

func parse(str string) []cmd {
	var cmds []cmd
	for _, line := range utils.Lines(str) {
		fields := strings.Fields(line)
		if fields[0] == "rotate" {
			if fields[1] == "column" {
				a := utils.Atoi(strings.TrimPrefix(fields[2], "x="))
				b := utils.Atoi(fields[4])
				cmds = append(cmds, cmd{(*screen).rotateCol, a, b})
			} else {
				a := utils.Atoi(strings.TrimPrefix(fields[2], "y="))
				b := utils.Atoi(fields[4])
				cmds = append(cmds, cmd{(*screen).rotateRow, a, b})
			}
		} else if fields[0] == "rect" {
			var a, b int
			utils.Sscanf(fields[1], "%dx%d", &a, &b)
			cmds = append(cmds, cmd{(*screen).rect, a, b})
		}
	}
	return cmds
}

func main() {
	cmds := parse(input)
	s := new(screen)
	for _, cmd := range cmds {
		cmd.fn(s, cmd.a, cmd.b)
	}
	// part 2
	count := 0
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			if s[i][j] {
				print("#")
				count++
			} else {
				print(".")
			}
		}
		println()
	}
	// part 1
	println(count)
}
