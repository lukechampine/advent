package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2024, 12)

func main() {
	grid := utils.ToByteGrid(utils.Lines(input))
	g := utils.Grid{X: len(grid[0]) - 1, Y: len(grid) - 1}
	seen := make(map[utils.Pos]bool)
	var sum int
	g.ForEach(func(p utils.Pos) {
		if seen[p] {
			return
		}
		c := grid[p.Y][p.X]
		var area, perimeter int
		utils.Flood(p, func(p utils.Pos) (next []utils.Pos) {
			seen[p] = true
			for _, adj := range p.ValidMoves(g) {
				if grid[adj.Y][adj.X] == c {
					next = append(next, adj)
				}
			}
			area++
			perimeter += 4 - len(next)
			return
		})
		sum += area * perimeter
	})
	utils.Println(sum)

	seen = make(map[utils.Pos]bool)
	sum = 0
	g.ForEach(func(p utils.Pos) {
		if seen[p] {
			return
		}
		c := grid[p.Y][p.X]
		var plot []utils.Pos
		utils.Flood(p, func(p utils.Pos) (next []utils.Pos) {
			seen[p] = true
			plot = append(plot, p)
			for _, adj := range p.ValidMoves(g) {
				if grid[adj.Y][adj.X] == c {
					next = append(next, adj)
				}
			}
			return
		})

		min, max := utils.BoundingBox(plot)
		inPlot := func(x, y int) bool {
			return utils.Any(len(plot), func(i int) bool { return plot[i] == utils.Pos{X: x, Y: y} })
		}
		sides := 0
		for x := min.X; x <= max.X; x++ {
			// edges to the left
			for y := min.Y; y <= max.Y; y++ {
				if inPlot(x, y) && !inPlot(x-1, y) {
					sides++
					for inPlot(x, y) && !inPlot(x-1, y) {
						y++
					}
				}
			}
			// edges to the right
			for y := min.Y; y <= max.Y; y++ {
				if inPlot(x, y) && !inPlot(x+1, y) {
					sides++
					for inPlot(x, y) && !inPlot(x+1, y) {
						y++
					}
				}
			}
		}
		for y := min.Y; y <= max.Y; y++ {
			// edges above
			for x := min.X; x <= max.X; x++ {
				if inPlot(x, y) && !inPlot(x, y-1) {
					sides++
					for inPlot(x, y) && !inPlot(x, y-1) {
						x++
					}
				}
			}
			// edges below
			for x := min.X; x <= max.X; x++ {
				if inPlot(x, y) && !inPlot(x, y+1) {
					sides++
					for inPlot(x, y) && !inPlot(x, y+1) {
						x++
					}
				}
			}
		}

		sum += len(plot) * sides
	})
	utils.Println(sum)
}
