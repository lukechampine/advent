package main

import (
	"sort"

	"github.com/lukechampine/advent/2019/intcode"
	"github.com/lukechampine/advent/utils"
)

var input = utils.ReadInput("day19_input.txt")
var prog = utils.ExtractInts(input)

func inBeam(x, y int) bool {
	return intcode.New(prog).Run(x, y) == 1
}

func fitsSquare(x, y int) bool {
	return inBeam(x, y) && inBeam(x-99, y+99)
}

func rightEdge(y int) (x int) {
	x = y // heuristic
	for !inBeam(x, y) || inBeam(x+1, y) {
		x++
	}
	return
}

func main() {
	// part 1
	utils.Println(utils.Count(50*50, func(i int) bool {
		return inBeam(i%50, i/50)
	}))

	// part 2
	y := sort.Search(10000, func(y int) bool {
		return fitsSquare(rightEdge(y), y)
	})
	x := rightEdge(y) - 99
	utils.Println(x*10000 + y)
	return
}
