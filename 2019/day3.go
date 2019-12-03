package main

import (
	"strings"

	"github.com/lukechampine/advent/utils"
)

var input = utils.ReadInput("day3_input.txt")
var ils = utils.Lines(input)

func main() {
	// part 1
	var dirs [2][]string
	dirs[0] = strings.Split(ils[0], ",")
	dirs[1] = strings.Split(ils[1], ",")

	var p1, p2 utils.Pos
	steps := 0
	seen := make(map[utils.Pos]int)
	for _, d := range dirs[0] {
		dir := utils.DirFromUDLR(d[0])
		n := utils.Atoi(d[1:])
		p1 = p1.Tread(dir, n, func(p utils.Pos) {
			steps++
			seen[p] = steps
		})
	}
	steps = 0
	intersections := make(map[utils.Pos]int)
	for _, d := range dirs[1] {
		dir := utils.DirFromUDLR(d[0])
		n := utils.Atoi(d[1:])
		p2 = p2.Tread(dir, n, func(p utils.Pos) {
			steps++
			if seen[p] != 0 {
				intersections[p] = steps + seen[p]
			}
		})
	}
	var bestDist int = 1e9
	for inter := range intersections {
		bestDist = utils.Min(bestDist, inter.Dist(utils.Pos{0, 0}))
	}
	utils.Println(bestDist)

	// part 2
	var bestSteps int = 1e9
	for _, steps := range intersections {
		bestSteps = utils.Min(bestSteps, steps)
	}
	utils.Println(bestSteps)
}
