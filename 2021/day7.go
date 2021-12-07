package main

import (
	"sort"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2021, 7)
var inputInts = utils.ExtractInts(input)

func main() {
	sort.Ints(inputInts)
	max := inputInts[len(inputInts)-1]

	utils.Println(utils.Minimum(max, func(i int) int {
		var sum int
		for _, v := range inputInts {
			sum += utils.Abs(v - i)
		}
		return sum
	}))

	utils.Println(utils.Minimum(max, func(i int) int {
		var sum int
		for _, v := range inputInts {
			delta := utils.Abs(v - i)
			sum += (delta + delta*delta) / 2
		}
		return sum
	}))
}
