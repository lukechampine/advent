package main

import (
	"bufio"
	"os"

	"github.com/lukechampine/advent/2019/intcode"
	"github.com/lukechampine/advent/utils"
)

var input = utils.ReadInput("day25_input.txt")
var prog = utils.ExtractInts(input)

const test = ""

func main() {
	// part 1
	stdin := bufio.NewReader(os.Stdin)
	m := intcode.New(prog)
	for !m.Halted {
		for !m.AwaitingInput && !m.Halted {
			out := m.Run()
			if out > 0 && out < 128 {
				utils.Print(string(out))
			}
		}
		line, _, _ := stdin.ReadLine()
		out := m.Run(intcode.ASCII(string(line) + "\n")...)
		if out > 0 && out < 128 {
			utils.Print(string(out))
		}
	}

	utils.Println()

	// part 2

	utils.Println()
}
