package main

import (
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2023, 8)

func main() {
	lines := utils.Lines(input)
	insts := lines[0]
	network := make(map[string][2]string)
	for _, line := range lines[2:] {
		network[line[0:3]] = [2]string{line[7:10], line[12:15]}
	}
	steps := 0
	for n := "AAA"; n != "ZZZ"; steps++ {
		n = network[n][strings.IndexByte("LR", insts[0])]
		insts = insts[1:] + insts[:1]
	}
	utils.Println(steps)

	var cycles []int
	for n := range network {
		if n[2] == 'A' {
			steps := 0
			insts := lines[0]
			for ; n[2] != 'Z'; steps++ {
				n = network[n][strings.IndexByte("LR", insts[0])]
				insts = insts[1:] + insts[:1]
			}
			cycles = append(cycles, steps)
		}
	}
	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}
	r := cycles[0]
	for _, c := range cycles[1:] {
		r = (r * c) / gcd(r, c)
	}
	utils.Println(r)
}
