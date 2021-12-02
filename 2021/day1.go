package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2021, 1)
var inputInts = utils.ExtractInts(input)

func increasing(ints []int) int {
	w := utils.SlidingWindow(ints, 2).([][]int)
	return utils.Count(len(w), func(i int) bool { return w[i][0] < w[i][1] })
}

func main() {
	// part 1
	utils.Println(increasing(inputInts))

	// part 2
	w := utils.SlidingWindow(inputInts, 3).([][]int)
	summed := make([]int, len(w))
	for i := range summed {
		summed[i] = utils.IntSum(w[i])
	}
	utils.Println(increasing(summed))
}
