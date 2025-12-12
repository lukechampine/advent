package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2025, 7)

func timelines(grid [][]byte, p utils.Pos, memo map[utils.Pos]int) int {
	for p.Y < len(grid)-1 && grid[p.Y+1][p.X] != '^' {
		p.Y++
	}
	if p.Y == len(grid)-1 {
		return 1
	} else if v, ok := memo[p]; ok {
		return v
	}
	l := timelines(grid, utils.Pos{X: p.X - 1, Y: p.Y}, memo)
	r := timelines(grid, utils.Pos{X: p.X + 1, Y: p.Y}, memo)
	memo[p] = l + r
	return l + r
}

func main() {
	grid := utils.ToByteGrid(utils.Lines(input))
	start := utils.Locate(grid, 'S')
	grid[start.Y][start.X] = '|'
	splits := 0
	for y := 1; y < len(grid)-1; y++ {
		for x, c := range grid[y] {
			if grid[y-1][x] == '|' {
				if c == '.' {
					grid[y][x] = '|'
				} else if c == '^' {
					grid[y][x-1] = '|'
					grid[y][x+1] = '|'
					splits++
				}
			}
		}
	}
	utils.Println(splits)

	grid = utils.ToByteGrid(utils.Lines(input))
	utils.Println(timelines(grid, start, make(map[utils.Pos]int)))
}
