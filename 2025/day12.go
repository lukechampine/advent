package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2025, 12)

func main() {
	groups := utils.Split(input, "\n\n")
	shapes := make([][][]byte, len(groups)-1)
	for i := range shapes {
		shapes[i] = utils.ToByteGrid(utils.Lines(groups[i])[1:])
	}

	part1 := 0
	for _, line := range utils.Lines(groups[len(groups)-1]) {
		ints := utils.ExtractInts(line)
		var total int
		for i, c := range ints[2:] {
			total += c * utils.CountGrid(shapes[i], '#')
		}
		part1 += utils.BoolToInt(total < ints[0]*ints[1])
	}
	utils.Println(part1)
}
