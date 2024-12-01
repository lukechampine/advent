package main

import (
	"sort"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2024, 1)

func main() {
	cols := utils.Transpose(utils.Window(utils.ExtractInts(input), 2)).([][]int)
	sort.Ints(cols[0])
	sort.Ints(cols[1])
	utils.Println(utils.Sum(len(cols[0]), func(i int) int {
		return utils.Abs(cols[0][i] - cols[1][i])
	}))
	utils.Println(utils.Sum(len(cols[0]), func(i int) int {
		return cols[0][i] * utils.Count(len(cols[1]), func(j int) bool {
			return cols[1][j] == cols[0][i]
		})
	}))
}
