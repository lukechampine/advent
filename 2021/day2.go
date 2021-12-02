package main

import (
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2021, 2)
var inputInts = utils.ExtractInts(input)

func main() {
	// part 1
	var pos utils.Pos
	for _, line := range utils.Lines(input) {
		fs := strings.Fields(line)
		n := utils.Atoi(fs[1])
		switch fs[0] {
		case "forward":
			pos.X += n
		case "up":
			pos.Y -= n
		case "down":
			pos.Y += n
		}
	}
	utils.Println(pos.X * pos.Y)

	// part 2
	pos = utils.Pos{0, 0}
	var aim int
	for _, line := range utils.Lines(input) {
		fs := strings.Fields(line)
		n := utils.Atoi(fs[1])
		switch fs[0] {
		case "forward":
			pos.X += n
			pos.Y += aim * n
		case "up":
			aim -= n
		case "down":
			aim += n
		}
	}
	utils.Println(pos.X * pos.Y)
}
