package main

import (
	"github.com/lukechampine/advent/utils"
)

var input = utils.ReadInput("day2_input.txt")
var prog = utils.ExtractInts(input)

func runMachine(p []int, noun, verb int) int {
	// mutate a copy
	p = append([]int(nil), p...)
	p[1] = noun
	p[2] = verb
	for i := 0; i < len(p); i += 4 {
		x, y, z := p[i+1], p[i+2], p[i+3]
		switch p[i] {
		case 1:
			p[z] = p[x] + p[y]
		case 2:
			p[z] = p[x] * p[y]
		case 99:
			return p[0]
		}
	}
	panic("unreachable")
}

func main() {
	// part 1
	output := runMachine(prog, 12, 2)
	utils.Println(output)

	// part 2
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			output := runMachine(prog, noun, verb)
			if output == 19690720 {
				utils.Println(100*noun + verb)
			}
		}
	}
}
