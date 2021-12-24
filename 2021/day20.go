package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2021, 20)

func adjacent(p utils.Pos) []utils.Pos {
	return []utils.Pos{
		{p.X + 1, p.Y + 1},
		{p.X + 0, p.Y + 1},
		{p.X - 1, p.Y + 1},
		{p.X + 1, p.Y + 0},
		{p.X + 0, p.Y + 0},
		{p.X - 1, p.Y + 0},
		{p.X + 1, p.Y - 1},
		{p.X + 0, p.Y - 1},
		{p.X - 1, p.Y - 1},
	}
}

func enhance(input [][]byte, alg string, inf byte) [][]byte {
	g := utils.Grid{len(input[0]) - 1, len(input) - 1}
	enhanced := make([][]byte, len(input)+4)
	for i := range enhanced {
		enhanced[i] = make([]byte, len(input[0])+4)
	}
	for y := range enhanced {
		for x := range enhanced[y] {
			var n int
			for i, a := range adjacent(utils.Pos{x - 2, y - 2}) {
				isSet := g.Valid(a) && input[a.Y][a.X] == '#' ||
					!g.Valid(a) && inf == '#'
				n |= utils.BoolToInt(isSet) << i
			}
			enhanced[y][x] = alg[n]
		}
	}
	return enhanced
}

func main() {
	parts := utils.Split(input, "\n\n")
	lines := utils.Lines(parts[1])
	grid := utils.ToByteGrid(lines)
	grid = enhance(grid, parts[0], '.')
	grid = enhance(grid, parts[0], '#')
	utils.Println(utils.CountGrid(grid, '#'))
	for i := 0; i < 24; i++ {
		grid = enhance(grid, parts[0], '.')
		grid = enhance(grid, parts[0], '#')
	}
	utils.Println(utils.CountGrid(grid, '#'))
}
