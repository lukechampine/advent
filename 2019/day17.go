package main

import (
	"github.com/lukechampine/advent/2019/intcode"
	"github.com/lukechampine/advent/utils"
)

var input = utils.ReadInput("day17_input.txt")
var prog = utils.ExtractInts(input)

func main() {
	// part 1
	m := intcode.New(prog)
	grid := utils.Lines(intcode.DecodeASCII(m.Run()))
	g := utils.Grid{X: len(grid[0]) - 1, Y: len(grid) - 1}

	var sum int
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] != '#' {
				continue
			}
			moves := (utils.Pos{X: x, Y: y}).Moves()
			intersection := utils.All(len(moves), func(i int) bool {
				return g.Valid(moves[i]) && grid[moves[i].Y][moves[i].X] == '#'
			})
			if intersection {
				sum += x * y
			}
		}
	}
	utils.Println(sum)

	// part 2
	prog[0] = 2
	m = intcode.New(prog)
	out := m.RunASCII(`
A,B,A,C,B,C,A,B,A,C
R,6,L,10,R,8,R,8
R,12,L,8,L,10
R,12,L,10,R,6,L,10
n
`)
	utils.Println(out[len(out)-1])
}
