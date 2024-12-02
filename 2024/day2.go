package main

import (
	"fmt"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2024, 2)

func safe(ints []int) bool {
	w := utils.SlidingWindow(ints, 2).([][]int)
	return (utils.All(len(w), func(i int) bool {
		return w[i][0]-w[i][1] > 0
	}) || utils.All(len(w), func(i int) bool {
		return w[i][0]-w[i][1] < 0
	})) && utils.All(len(w), func(i int) bool {
		d := utils.Abs(w[i][0] - w[i][1])
		return 0 < d && d < 4
	})
}

func main() {
	lines := utils.Lines(input)
	utils.Println(utils.Count(len(lines), func(i int) bool {
		return safe(utils.ExtractInts(lines[i]))
	}))
	utils.Println(utils.Count(len(lines), func(i int) bool {
		ints := utils.ExtractInts(lines[i])
		s := safe(ints) ||
			utils.Any(len(ints), func(i int) bool {
				del := append([]int(nil), ints...)
				utils.DeleteSliceIndex(&del, i)
				return safe(del)
			})
		if s {
			fmt.Println(i)
		}
		return s
	}))
}
