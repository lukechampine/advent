package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2021, 11)
var inputLines = utils.Lines(input)

func update(grid [][]byte) (nFlashed int) {
	g := utils.Grid{X: len(grid[0]) - 1, Y: len(grid) - 1}

	var toFlash []utils.Pos
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == '9' {
				toFlash = append(toFlash, utils.Pos{x, y})
			}
			grid[y][x]++
		}
	}

	for ; len(toFlash) > 0; nFlashed++ {
		p := toFlash[0]
		toFlash = toFlash[1:]
		adj := p.ValidNumpad(g)
		for _, a := range adj {
			if grid[a.Y][a.X] == '9' {
				toFlash = append(toFlash, a)
			}
			grid[a.Y][a.X]++
		}
	}

	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] > '9' {
				grid[y][x] = '0'
			}
		}
	}

	return nFlashed
}

func main() {
	var flashed int
	grid := utils.ToByteGrid(inputLines)
	var step int
	for step = 1; step <= 100; step++ {
		flashed += update(grid)
	}
	utils.Println(flashed)

	for ; ; step++ {
		if update(grid) == len(grid)*len(grid[0]) {
			utils.Println(step)
			return
		}
	}
}
