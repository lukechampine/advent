package main

import (
	"github.com/lukechampine/advent/utils"
)

var input = utils.ReadInput("day5_input.txt")
var prog = utils.ExtractInts(input)

func runMachine(p []int, input int) (output int) {
	p = append([]int(nil), p...) // mutate a copy
	for i := 0; i < len(p); {
		op, flags := p[i]%100, utils.Digits(utils.Itoa(p[i]/100))
		getArg := func(n int) *int {
			imm := p[i+n]
			if len(flags) < n || flags[len(flags)-n] == 0 {
				return &p[imm]
			}
			return &imm
		}
		switch op {
		case 1:
			*getArg(3) = *getArg(1) + *getArg(2)
			i += 4
		case 2:
			*getArg(3) = *getArg(1) * *getArg(2)
			i += 4
		case 3:
			*getArg(1) = input
			i += 2
		case 4:
			return *getArg(1)
			i += 2
		case 5:
			if *getArg(1) != 0 {
				i = *getArg(2)
			} else {
				i += 3
			}
		case 6:
			if *getArg(1) == 0 {
				i = *getArg(2)
			} else {
				i += 3
			}
		case 7:
			*getArg(3) = utils.BoolToInt(*getArg(1) < *getArg(2))
			i += 4
		case 8:
			*getArg(3) = utils.BoolToInt(*getArg(1) == *getArg(2))
			i += 4
		case 99:
			return
		default:
			i++
		}
	}
	panic("unreachable")
}

func main() {
	// part 1
	utils.Println(runMachine(prog, 1))
	// part 2
	utils.Println(runMachine(prog, 5))

}
