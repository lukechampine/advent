package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2024, 8)

func main() {
	grid := utils.ToByteGrid(utils.Lines(input))
	g := utils.Grid{X: len(grid[0]) - 1, Y: len(grid) - 1}
	an := make(map[utils.Pos]bool)
	for y := range grid {
		for x, c := range grid[y] {
			if grid[y][x] == '.' {
				continue
			}
			a := utils.Pos{X: x, Y: y}
			g.ForEach(func(p utils.Pos) {
				if p == a || grid[p.Y][p.X] != c {
					return
				}
				d := p.Rel(a)
				if anti := p.Add(d.X, d.Y); g.Valid(anti) {
					an[anti] = true
				}
			})
		}
	}
	utils.Println(len(an))

	for y := range grid {
		for x, c := range grid[y] {
			if grid[y][x] == '.' {
				continue
			}
			a := utils.Pos{X: x, Y: y}
			g.ForEach(func(p utils.Pos) {
				if p == a || grid[p.Y][p.X] != c {
					return
				}
				dx, dy := p.X-a.X, p.Y-a.Y
				for anti := a.Add(dx, dy); g.Valid(anti); anti = anti.Add(dx, dy) {
					an[anti] = true
				}
			})
		}
	}
	utils.Println(len(an))
}
