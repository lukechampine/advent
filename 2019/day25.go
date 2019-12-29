package main

import (
	"bufio"
	"os"

	"github.com/lukechampine/advent/2019/intcode"
	"github.com/lukechampine/advent/utils"
)

var input = utils.ReadInput("day25_input.txt")
var prog = utils.ExtractInts(input)

func main() {
	// part 1
	stdin := bufio.NewReader(os.Stdin)
	m := intcode.New(prog)
	utils.Print(intcode.DecodeASCII(m.Run()))
	for !m.Halted {
		line, _, _ := stdin.ReadLine()
		input := string(line) + "\n"
		utils.Print(intcode.DecodeASCII(m.RunASCII(input)))
	}
}
