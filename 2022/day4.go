package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2022, 4)

func main() {
	var ranges [][4]int
	for _, line := range utils.Lines(input) {
		var r [4]int
		utils.Sscanf(line, "%d-%d,%d-%d", &r[0], &r[1], &r[2], &r[3])
		ranges = append(ranges, r)
	}

	utils.Println(utils.Count(len(ranges), func(i int) bool {
		r := ranges[i]
		return (r[0] <= r[2] && r[1] >= r[3]) || (r[2] <= r[0] && r[3] >= r[1])
	}))

	utils.Println(utils.Count(len(ranges), func(i int) bool {
		r := ranges[i]
		return (r[2] <= r[0] && r[0] <= r[3]) ||
			(r[2] <= r[1] && r[1] <= r[3]) ||
			(r[0] <= r[2] && r[2] <= r[1]) ||
			(r[0] <= r[3] && r[3] <= r[1])
	}))
}
