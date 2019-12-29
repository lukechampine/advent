package main

import (
	"github.com/lukechampine/advent/2019/intcode"
	"github.com/lukechampine/advent/utils"
)

var input = utils.ReadInput("day9_input.txt")
var prog = utils.ExtractInts(input)

func main() {
	// part 1
	utils.Println(intcode.New(prog).Run(1)[0])
	// part 2
	utils.Println(intcode.New(prog).Run(2)[0])
}
