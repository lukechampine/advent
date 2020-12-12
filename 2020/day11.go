package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2020, 11)

func simulate(grid [][]byte, seats func(utils.Pos, [][]byte) []utils.Pos, thresh int) int {
	anyChanged := true
	for anyChanged {
		anyChanged = false
		grid = utils.GameOfLife(grid, func(c byte, p utils.Pos, _ []utils.Pos) byte {
			switch c {
			case 'L':
				shouldFill := true
				for _, q := range seats(p, grid) {
					shouldFill = shouldFill && grid[q.Y][q.X] != '#'
				}
				if shouldFill {
					c = '#'
					anyChanged = true
				}
			case '#':
				n := 0
				for _, q := range seats(p, grid) {
					n += utils.BoolToInt(grid[q.Y][q.X] == '#')
				}
				if n >= thresh {
					c = 'L'
					anyChanged = true
				}
			}
			return c
		})
	}
	return utils.CountGrid(grid, '#')
}

func main() {
	grid := utils.ToByteGrid(utils.Lines(input))
	g := utils.Grid{len(grid[0]) - 1, len(grid) - 1}
	seats := func(p utils.Pos, grid [][]byte) []utils.Pos {
		return p.ValidNumpad(g)
	}
	utils.Println(simulate(grid, seats, 4))

	seats = func(p utils.Pos, grid [][]byte) []utils.Pos {
		moves := p.ValidNumpad(g)
		for i, m := range moves {
			r := m.Rel(p)
			for grid[m.Y][m.X] == '.' && g.Valid(m.Add(r.X, r.Y)) {
				m = m.Add(r.X, r.Y)
			}
			moves[i] = m
		}
		return moves
	}
	utils.Println(simulate(grid, seats, 5))
}
