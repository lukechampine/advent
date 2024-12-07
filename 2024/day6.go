package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2024, 6)

func walk(grid [][]byte, fn func(utils.Agent) bool) bool {
	g := utils.Grid{X: len(grid[0]) - 1, Y: len(grid) - 1}
	start := utils.Locate(grid, '^')
	a := utils.NewAgent(start.X, start.Y, utils.Up)
	for {
		a.MoveForwardArray(1)
		if !g.Valid(a.Pos) {
			return false
		} else if grid[a.Pos.Y][a.Pos.X] == '#' {
			a.MoveForwardArray(-1)
			a.TurnRight()
		}
		if stop := fn(a); stop {
			return true
		}
	}
}

func main() {
	grid := utils.ToByteGrid(utils.Lines(input))

	seen := make(map[utils.Pos]bool)
	walk(grid, func(a utils.Agent) bool {
		seen[a.Pos] = true
		return false
	})
	utils.Println(len(seen))

	causesCycle := func(x, y int) bool {
		if grid[y][x] != '.' {
			return false
		}
		grid[y][x] = '#'
		defer func() { grid[y][x] = '.' }()

		seen := make(map[utils.Pos]utils.Dir)
		return walk(grid, func(a utils.Agent) bool {
			a.TurnLeft()
			a.MoveForwardArray(1)
			if grid[a.Pos.Y][a.Pos.X] == '#' {
				if d, ok := seen[a.Pos]; ok && d == a.Dir {
					return true
				}
				seen[a.Pos] = a.Dir
			}
			return false
		})
	}
	n := 0
	for p := range seen {
		if causesCycle(p.X, p.Y) {
			n++
		}
	}
	utils.Println(n)
}
