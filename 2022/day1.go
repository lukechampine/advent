package main

import (
	"sort"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2022, 1)

func main() {
	groups := utils.Split(input, "\n\n")
	sums := make([]int, len(groups))
	for i, g := range groups {
		sums[i] = utils.IntSum(utils.ExtractInts(g))
	}
	sort.Slice(sums, func(i, j int) bool {
		return sums[i] > sums[j]
	})
	utils.Println(sums[0])
	utils.Println(utils.IntSum(sums[:3]))
}
