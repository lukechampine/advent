package main

import (
	"github.com/lukechampine/advent/2019/intcode"
	"github.com/lukechampine/advent/utils"
)

var input = utils.ReadInput("day21_input.txt")
var prog = utils.ExtractInts(input)

func main() {
	// part 1
	// D && (!A || !C)
	out := intcode.New(prog).RunASCII(`
NOT A J
NOT C T
OR T J
AND D J
WALK
`)
	utils.Println(out[len(out)-1])

	// part 2
	// D && (!A || (!B && H) || (!C && (E || H)))
	out = intcode.New(prog).RunASCII(`
NOT C T
OR E J
OR H J
AND T J
NOT B T
AND H T
OR T J
NOT A T
OR T J
AND D J
RUN
`)
	utils.Println(out[len(out)-1])
}
