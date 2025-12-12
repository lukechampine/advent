package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2025, 3)

func best(prefix, rest string, n int) int {
	if n == 0 {
		return 0
	}
	i := utils.MaximumIndex(len(rest)-n+1, func(i int) int {
		return int(rest[i])
	})
	return best(prefix+string(rest[i]), rest[i+1:], n-1)
}

func main() {
	var part1, part2 int
	for _, line := range utils.Lines(input) {
		part1 += best("", line, 2)
		part2 += best("", line, 12)
	}
	utils.Println(part1)
	utils.Println(part2)
}
