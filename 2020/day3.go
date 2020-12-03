package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2020, 3)
var lines = utils.Lines(input)
var maxX = len(lines[0])

func countTrees(dx, dy int) (n int) {
	x, y := 0, 0
	for y < len(lines) {
		if lines[y][x] == '#' {
			n++
		}
		x = (x + dx) % maxX
		y += dy
	}
	return
}

func main() {
	utils.Println(countTrees(3, 1))
	utils.Println(countTrees(1, 1) *
		countTrees(3, 1) *
		countTrees(5, 1) *
		countTrees(7, 1) *
		countTrees(1, 2))
}
