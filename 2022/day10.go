package main

import (
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2022, 10)

func main() {
	x := 1
	pc := 0
	sum := 0
	grid := utils.NewLightBoard(40, 6)
	inc := func() {
		if i, j := pc%40, pc/40; utils.Abs(i-x) <= 1 {
			grid.Set(i, j, true)
		}
		if pc++; (pc-20)%40 == 0 {
			sum += pc * x
		}
	}
	for _, line := range utils.Lines(input) {
		fields := strings.Fields(line)
		switch fields[0] {
		case "noop":
			inc()
		case "addx":
			inc()
			inc()
			x += utils.Atoi(fields[1])
		}
	}
	utils.Println(sum, x)
	grid.Print()
}
