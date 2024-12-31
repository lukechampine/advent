package main

import (
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2024, 25)

func main() {
	groups := strings.Split(input, "\n\n")
	var locks [][5]int
	var keys [][5]int
	for _, group := range groups {
		lines := utils.Lines(group)
		cols := utils.Transpose(lines).([]string)
		var a [5]int
		for i := range a {
			a[i] = strings.Count(cols[i], "#")
		}
		if utils.CharCounts(lines[0])['#'] == 5 {
			locks = append(locks, a)
		} else {
			keys = append(keys, a)
		}
	}

	var n int
	for _, l := range locks {
		for _, k := range keys {
			n += utils.BoolToInt(utils.All(5, func(i int) bool {
				return l[i]+k[i] <= 7
			}))
		}
	}
	utils.Println(n)
}
