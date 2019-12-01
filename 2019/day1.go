package main

import (
	"github.com/lukechampine/advent/utils"
)

var input = utils.ReadInput("day1_input.txt")
var inputInts = utils.ExtractInts(input)

func calcFuel(m int) int {
	m = (m / 3) - 2
	if m <= 0 {
		return 0
	}
	return m + calcFuel(m)
}

func main() {
	// part 1
	var sum int
	for _, m := range inputInts {
		sum += (m / 3) - 2
	}
	utils.Println(sum)

	// part 2
	sum = 0
	for _, m := range inputInts {
		sum += calcFuel(m)
	}
	utils.Println(sum)
}
