package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2022, 12)

func main() {
	g := utils.ToByteGrid(utils.Lines(input))
	m := utils.Maze{
		Grid:   utils.Grid{X: len(g[0]) - 1, Y: len(g) - 1},
		IsWall: func(p utils.Pos) bool { return false },
		CanMove: func(p, d utils.Pos) bool {
			p, d = d, p // moving backwards

			pc, dc := g[p.Y][p.X], g[d.Y][d.X]
			if pc == 'S' {
				pc = 'a'
			}
			if dc == 'E' {
				dc = 'z'
			}
			return dc <= pc+1
		},
	}

	dists := m.DistancesFrom(utils.Locate(g, 'E'))
	best := dists[utils.Locate(g, 'S')]
	utils.Println(best)

	m.Grid.ForEach(func(p utils.Pos) {
		if g[p.Y][p.X] == 'a' {
			if d, ok := dists[p]; ok && d < best {
				best = d
			}
		}
	})
	utils.Println(best)
}
