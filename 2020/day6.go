package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2020, 6)

func main() {
	var part1, part2 int
	for _, g := range utils.Split(input, "\n\n") {
		people := utils.Lines(g)
		c := make(map[rune]int)
		for _, p := range people {
			for k := range utils.CharCounts(p) {
				c[k]++
			}
		}
		for _, v := range c {
			part1++
			part2 += utils.BoolToInt(v == len(people))
		}
	}
	utils.Println(part1)
	utils.Println(part2)
}
