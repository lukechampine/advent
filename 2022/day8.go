package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2022, 8)

var moves = []func(utils.Pos) utils.Pos{
	func(p utils.Pos) utils.Pos { return p.Add(0, 1) },
	func(p utils.Pos) utils.Pos { return p.Add(1, 0) },
	func(p utils.Pos) utils.Pos { return p.Add(0, -1) },
	func(p utils.Pos) utils.Pos { return p.Add(-1, 0) },
}

func isVisible(grid []string, x, y int) bool {
	g := utils.Grid{X: len(grid[0]) - 1, Y: len(grid) - 1}
	for _, move := range moves {
		p := move(utils.Pos{x, y})
		for g.Valid(p) && grid[p.Y][p.X] < grid[y][x] {
			p = move(p)
		}
		if !g.Valid(p) {
			return true
		}
	}
	return false
}

func score(grid []string, x, y int) int {
	g := utils.Grid{X: len(grid[0]) - 1, Y: len(grid) - 1}
	score := 1
	for _, move := range moves {
		p := utils.Pos{x, y}
		n := 0
		for p = move(p); g.Valid(p); p = move(p) {
			n++
			if grid[p.Y][p.X] >= grid[y][x] {
				break
			}
		}
		score *= n
	}
	return score
}

func main() {
	g := utils.Lines(input)
	n := 0
	for y := range g {
		for x := range g[y] {
			if isVisible(g, x, y) {
				n++
			}
		}
	}
	utils.Println(n)

	best := 0
	for y := range g {
		for x := range g[y] {
			if s := score(g, x, y); s > best {
				best = s
			}
		}
	}
	utils.Println(best)
}
