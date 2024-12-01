package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2023, 11)

func main() {
	if false {
		input = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....
`
	}
	grid := utils.ToByteGrid(utils.Lines(input))
	//g := utils.Grid{X: len(grid[0]) - 1, Y: len(grid) - 1}

	expand := make(map[utils.Pos]bool)
	for y := range grid {
		expand[utils.Pos{X: 0, Y: y}] = utils.All(len(grid[y]), func(x int) bool { return grid[y][x] == '.' })
	}
	for x := range grid[0] {
		expand[utils.Pos{X: x, Y: 0}] = utils.All(len(grid), func(y int) bool { return grid[y][x] == '.' })
	}

	var galaxies []utils.Pos
	for y := range grid {
		for x, c := range grid[y] {
			if c == '#' {
				galaxies = append(galaxies, utils.Pos{X: x, Y: y})
			}
		}
	}

	shortestPath := func(g1, g2 utils.Pos, age int) (n int) {
		if g1.X > g2.X {
			g1, g2 = g2, g1
		}
		for x := g1.X; x < g2.X; x++ {
			if expand[utils.Pos{X: x, Y: 0}] {
				n += age
			} else {
				n++
			}
		}
		if g1.Y > g2.Y {
			g1, g2 = g2, g1
		}
		for y := g1.Y; y < g2.Y; y++ {
			if expand[utils.Pos{X: 0, Y: y}] {
				n += age
			} else {
				n++
			}
		}
		return
	}

	var sum int
	for i, g1 := range galaxies {
		for _, g2 := range galaxies[i+1:] {
			sum += shortestPath(g1, g2, 2)
		}
	}
	utils.Println(sum)

	sum = 0
	for i, g1 := range galaxies {
		for _, g2 := range galaxies[i+1:] {
			sum += shortestPath(g1, g2, 100000)
		}
	}
	utils.Println(sum)
}
