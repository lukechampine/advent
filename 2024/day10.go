package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2024, 10)

func main() {
	grid := utils.ToByteGrid(utils.Lines(input))
	g := utils.Grid{X: len(grid[0]) - 1, Y: len(grid) - 1}
	sum := 0
	g.ForEach(func(p utils.Pos) {
		if grid[p.Y][p.X] != '0' {
			return
		}
		queue := []utils.Pos{p}
		seen := make(map[utils.Pos]bool)
		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]
			if seen[p] {
				continue
			}
			seen[p] = true
			if grid[p.Y][p.X] == '9' {
				sum++
				continue
			}
			for _, m := range p.ValidMoves(g) {
				if grid[m.Y][m.X] == grid[p.Y][p.X]+1 {
					queue = append(queue, m)
				}
			}
		}
	})
	utils.Println(sum)

	sum = 0
	g.ForEach(func(p utils.Pos) {
		if grid[p.Y][p.X] != '0' {
			return
		}
		queue := []utils.Pos{p}
		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]
			if grid[p.Y][p.X] == '9' {
				sum++
				continue
			}
			for _, m := range p.ValidMoves(g) {
				if grid[m.Y][m.X] == grid[p.Y][p.X]+1 {
					queue = append(queue, m)
				}
			}
		}
	})
	utils.Println(sum)
}
