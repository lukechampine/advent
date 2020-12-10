package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2020, 9)
var inputInts = utils.ExtractInts(input)

func isSum(x int, prev []int) bool {
	for i := range prev {
		for j := i + 1; j < len(prev); j++ {
			if prev[i]+prev[j] == x {
				return true
			}
		}
	}
	return false
}

func main() {
	// part 1
	var n int
	for i, x := range inputInts[25:] {
		if !isSum(x, inputInts[:25+i]) {
			n = x
			break
		}
	}
	utils.Println(n)

	// part 2
outer:
	for i := range inputInts {
		for j := i + 2; j < len(inputInts); j++ {
			section := inputInts[i:j]
			if s := utils.IntSum(section); s == n {
				max := utils.Maximum(len(section), func(k int) int { return section[k] })
				min := utils.Minimum(len(section), func(k int) int { return section[k] })
				utils.Println(min + max)
				return
			} else if s > n {
				continue outer
			}
		}
	}
}
