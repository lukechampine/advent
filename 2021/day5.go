package main

import (
	"fmt"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2021, 5)
var inputLines = utils.Lines(input)

func main() {
	seen := make(map[utils.Pos]int)
	seenHV := make(map[utils.Pos]int)
	for _, line := range inputLines {
		var start, end utils.Pos
		utils.Sscanf(line, "%d,%d -> %d,%d", &start.X, &start.Y, &end.X, &end.Y)
		p := start
		for seen[p]++; p != end; seen[p]++ {
			p = p.StrideTowards(end)
		}
		if start.X == end.X || start.Y == end.Y {
			p := start
			for seenHV[p]++; p != end; seenHV[p]++ {
				p = p.StrideTowards(end)
			}
		}
	}
	// part 1
	var sumHV int
	for _, n := range seenHV {
		if n > 1 {
			sumHV++
		}
	}
	fmt.Println(sumHV)

	// part 2
	var sum int
	for _, n := range seen {
		if n > 1 {
			sum++
		}
	}
	fmt.Println(sum)
}
