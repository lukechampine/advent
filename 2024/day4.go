package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2024, 4)

func countMatches(grid [][]byte, shape []utils.Pos, f func(string) bool) (n int) {
	for y := range grid {
		for x := range grid[y] {
			var b []byte
			for _, p := range shape {
				if (utils.Grid{X: len(grid[0]) - 1, Y: len(grid) - 1}).Valid(p.Add(x, y)) {
					b = append(b, grid[y+p.Y][x+p.X])
				}
			}
			n += utils.BoolToInt(len(b) == len(shape) && f(string(b)))
		}
	}
	return
}

func main() {
	grid := utils.ToByteGrid(utils.Lines(input))

	var found int
	for _, shape := range [][]utils.Pos{
		[]utils.Pos{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}, {X: 3, Y: 0}},
		[]utils.Pos{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: 2}, {X: 0, Y: 3}},
		[]utils.Pos{{X: 0, Y: 0}, {X: 1, Y: 1}, {X: 2, Y: 2}, {X: 3, Y: 3}},
		[]utils.Pos{{X: 0, Y: 0}, {X: 1, Y: -1}, {X: 2, Y: -2}, {X: 3, Y: -3}},
	} {
		found += countMatches(grid, shape, func(s string) bool {
			return s == "XMAS" || s == "SAMX"
		})
	}
	utils.Println(found)

	cross := []utils.Pos{{X: 0, Y: 0}, {X: 2, Y: 0}, {X: 1, Y: 1}, {X: 0, Y: 2}, {X: 2, Y: 2}}
	utils.Println(countMatches(grid, cross, func(s string) bool {
		return s == "MSAMS" || s == "SMASM" || s == "MMASS" || s == "SSAMM"
	}))
}
