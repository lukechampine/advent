package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2020, 5)
var inputLines = utils.Lines(input)

func convertToBits(s string, set byte) (n int) {
	for i := range s {
		n |= (utils.BoolToInt(s[i] == set) << (len(s) - i - 1))
	}
	return
}

func main() {
	grid := make([][]bool, 128)
	for i := range grid {
		grid[i] = make([]bool, 8)
	}
	utils.Println(utils.Maximum(len(inputLines), func(i int) int {
		row := convertToBits(inputLines[i][:7], 'B')
		col := convertToBits(inputLines[i][7:], 'R')
		grid[row][col] = true
		return row*8 + col
	}))

	lastEmptyRow := 0
	for row := range grid {
		for col := range grid[row] {
			if !grid[row][col] {
				if row-lastEmptyRow > 1 {
					utils.Println(row*8 + col)
					return
				}
				lastEmptyRow = row
			}
		}
	}
}
