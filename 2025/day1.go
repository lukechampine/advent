package main

import (
	"fmt"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2025, 1)

func main() {
	ints := utils.ExtractInts(utils.Replace(input, "R", "+", "L", "-"))
	var part1 int
	pos := 50
	for _, n := range ints {
		pos = ((pos+n)%100 + 100) % 100
		if pos == 0 {
			part1++
		}
	}
	utils.Println(part1)

	var part2 int
	pos = 50
	for _, n := range ints[:100] {
		for i := 0; i < utils.Abs(n); i++ {
			pos = ((pos+utils.Signum(n))%100 + 100) % 100
			if pos == 0 {
				part2++
			}
		}
		fmt.Println("pos =", pos, "n =", n, "part2 =", part2)
	}
	utils.Println(part2)
}
