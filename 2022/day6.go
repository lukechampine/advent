package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2022, 6)

func main() {
	for _, n := range []int{4, 14} {
		for i, group := range utils.SlidingWindow(input, n).([]string) {
			if len(utils.CharCounts(group)) == n {
				utils.Println(i + n)
				break
			}
		}
	}
}
