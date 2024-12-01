package main

import (
	"fmt"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2023, 14)

func tiltNorth(grid [][]byte) {
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == 'O' {
				y2 := y
				for y2-1 >= 0 && grid[y2-1][x] == '.' {
					y2--
				}
				grid[y][x] = '.'
				grid[y2][x] = 'O'
			}
		}
	}
}

func tiltWest(grid [][]byte) {
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == 'O' {
				x2 := x
				for x2-1 >= 0 && grid[y][x2-1] == '.' {
					x2--
				}
				grid[y][x] = '.'
				grid[y][x2] = 'O'
			}
		}
	}
}

func tiltSouth(grid [][]byte) {
	for y := len(grid) - 1; y >= 0; y-- {
		for x := range grid[y] {
			if grid[y][x] == 'O' {
				y2 := y
				for y2+1 < len(grid) && grid[y2+1][x] == '.' {
					y2++
				}
				grid[y][x] = '.'
				grid[y2][x] = 'O'
			}
		}
	}
}

func tiltEast(grid [][]byte) {
	for y := range grid {
		for x := len(grid[y]) - 1; x >= 0; x-- {
			if grid[y][x] == 'O' {
				x2 := x
				for x2+1 < len(grid[y]) && grid[y][x2+1] == '.' {
					x2++
				}
				grid[y][x] = '.'
				grid[y][x2] = 'O'
			}
		}
	}
}

func load(grid [][]byte) (l int) {
	for y := range grid {
		for _, c := range grid[y] {
			if c == 'O' {
				l += len(grid) - y
			}
		}
	}
	return
}

func main() {
	grid := utils.ToByteGrid(utils.Lines(input))
	tiltNorth(grid)
	utils.Println(load(grid))

	grid = utils.ToByteGrid(utils.Lines(input))
	cycle := func() {
		tiltNorth(grid)
		tiltWest(grid)
		tiltSouth(grid)
		tiltEast(grid)
	}
	seen := make(map[string]int)
	for i := 0; ; i++ {
		s := fmt.Sprint(grid)
		if j, ok := seen[s]; ok {
			for rem := (1000000000 - i) % (i - j); rem > 0; rem-- {
				cycle()
			}
			break
		}
		seen[s] = i
		cycle()
	}
	utils.Println(load(grid))
}
