package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2025, 4)

func main() {
	grid := utils.ToByteGrid(utils.Lines(input))
	g := utils.Grid{X: len(grid[0]) - 1, Y: len(grid) - 1}
	accessible := func(p utils.Pos) bool {
		paper := 0
		for _, n := range p.ValidNumpad(g) {
			if grid[n.Y][n.X] == '@' {
				paper++
			}
		}
		return grid[p.Y][p.X] == '@' && paper < 4
	}
	var part1 int
	g.ForEach(func(p utils.Pos) {
		part1 += utils.BoolToInt(accessible(p))
	})
	utils.Println(part1)

	var part2 int
	for {
		prev := part2
		g.ForEach(func(p utils.Pos) {
			if accessible(p) {
				part2++
				grid[p.Y][p.X] = 'x'
			}
		})
		if part2 == prev {
			break
		}
	}
	utils.Println(part2)
}
