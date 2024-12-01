package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2023, 2)

func maxCubes(game string) (r, g, b int) {
	for _, set := range utils.Split(utils.Split(game, ": ")[1], "; ") {
		colors := make(map[string]int)
		for _, cubes := range utils.Split(set, ", ") {
			colors[utils.Split(cubes, " ")[1]] += utils.ExtractInts(cubes)[0]
		}
		r = utils.Max(r, colors["red"])
		g = utils.Max(g, colors["green"])
		b = utils.Max(b, colors["blue"])
	}
	return
}

func main() {
	lines := utils.Lines(input)
	var part1, part2 int
	for _, line := range lines {
		maxR, maxG, maxB := maxCubes(line)
		if maxR <= 12 && maxG <= 13 && maxB <= 14 {
			part1 += utils.ExtractInts(line)[0]
		}
		part2 += maxR * maxG * maxB
	}
	utils.Println(part1)
	utils.Println(part2)
}
