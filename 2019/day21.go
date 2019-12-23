package main

import (
	"github.com/lukechampine/advent/2019/intcode"
	"github.com/lukechampine/advent/utils"
)

var input = utils.ReadInput("day21_input.txt")
var prog = utils.ExtractInts(input)

func main() {
	// part 1
	m := intcode.New(prog)
	for !m.AwaitingInput {
		m.Run()
	}
	// D && (!A || !C)
	script := `
NOT A J
NOT C T
OR T J
AND D J
WALK
`
	m.Run(intcode.ASCII(script[1:])...)
	for !m.Halted {
		out := m.Run()
		if out > 128 {
			utils.Println(out)
		}
	}

	// part 2
	m = intcode.New(prog)
	for !m.AwaitingInput {
		m.Run()
	}

	// D && (!A || (!B && H) || (!C && (E || H)))
	script = `
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
`
	m.Run(intcode.ASCII(script[1:])...)
	for !m.Halted {
		out := m.Run()
		if out > 128 {
			utils.Println(out)
		}
	}
}
