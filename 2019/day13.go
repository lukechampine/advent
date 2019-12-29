package main

import (
	"github.com/lukechampine/advent/2019/intcode"
	"github.com/lukechampine/advent/utils"
)

var input = utils.ReadInput("day13_input.txt")
var prog = utils.ExtractInts(input)

func parseTiles(out []int) map[utils.Pos]int {
	tiles := make(map[utils.Pos]int)
	for len(out) > 0 {
		x, y, id := out[0], out[1], out[2]
		out = out[3:]
		tiles[utils.Pos{X: x, Y: y}] = id
	}
	return tiles
}

func main() {
	// part 1
	tiles := parseTiles(intcode.New(prog).Run(0))
	var n int
	for _, id := range tiles {
		if id == 2 {
			n++
		}
	}
	utils.Println(n)

	// part 2
	var tilt int
	var ballX int
	var paddleX int
	prog[0] = 2
	m := intcode.New(prog)
	for !m.Halted {
		tiles = parseTiles(m.Run(tilt))
		for p, id := range tiles {
			switch id {
			case 3:
				paddleX = p.X
			case 4:
				ballX = p.X
			}
		}
		switch {
		case ballX < paddleX:
			tilt = -1
		case ballX == paddleX:
			tilt = 0
		case ballX > paddleX:
			tilt = 1
		}
	}
	utils.Println(tiles[utils.Pos{X: -1, Y: 0}])
}
