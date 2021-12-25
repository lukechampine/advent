package main

import (
	"fmt"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2021, 25)

func doStep(grid [][]byte) [][]byte {
	next := utils.CloneGrid(grid)
	for y, row := range grid {
		for x, c := range row {
			if c == '>' {
				if grid[y][(x+1)%len(row)] == '.' {
					next[y][x] = '.'
					next[y][(x+1)%len(row)] = '>'
				}
			}
		}
	}
	next2 := utils.CloneGrid(next)
	for y, row := range next {
		for x, c := range row {
			if c == 'v' {
				if next[(y+1)%len(next)][x] == '.' {
					next2[y][x] = '.'
					next2[(y+1)%len(next2)][x] = 'v'
				}
			}
		}
	}
	return next2
}

func main() {
	grid := utils.ToByteGrid(utils.Lines(input))
	for step := 1; ; step++ {
		next := doStep(grid)
		if fmt.Sprint(next) == fmt.Sprint(grid) {
			utils.Println(step)
			return
		}
		grid = next
	}
}
