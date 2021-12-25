package main

import (
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2021, 24)
var prog = utils.Lines(input)

func verify(model string) bool {
	vars := make(map[string]int)
	for _, line := range prog {
		parts := strings.Fields(line)
		inst := parts[0]
		a := parts[1]
		var b int
		if len(parts) > 2 {
			if strings.Contains("wxyz", parts[2]) {
				b = vars[parts[2]]
			} else {
				b = utils.Atoi(parts[2])
			}
		}
		switch inst {
		case "inp":
			vars[a] = utils.Atoi(model[:1])
			model = model[1:]
		case "add":
			vars[a] += b
		case "mul":
			vars[a] *= b
		case "div":
			vars[a] /= b
		case "mod":
			vars[a] %= b
		case "eql":
			vars[a] = utils.BoolToInt(vars[a] == b)
		}
	}
	return vars["z"] == 0
}

func main() {
	// By decompiling the program, the following series of deductions can be made:
	//
	// input[3]  = input[2] + 5
	// input[5]  = input[4] - 3
	// input[7]  = input[6] + 7
	// input[10] = input[9] - 1
	// input[11] = input[8] + 3
	// input[12] = input[1] + 6
	// input[13] = input[0]
	//
	// Bearing in mind the further constraint that each digit must be in [1,9], we can then freely
	// choose values for inputs 0, 1, 2, 4, 6, 8, and 9 that maximize (or minimize) the total.

	part1 := []byte("934X9X2X69XXXX")
	part1[3] = part1[2] + 5
	part1[5] = part1[4] - 3
	part1[7] = part1[6] + 7
	part1[10] = part1[9] - 1
	part1[11] = part1[8] + 3
	part1[12] = part1[1] + 6
	part1[13] = part1[0]
	if !verify(string(part1)) {
		panic("not a valid input")
	}
	utils.Println(string(part1))

	part2 := []byte("111X4X1X12XXXX")
	part2[3] = part2[2] + 5
	part2[5] = part2[4] - 3
	part2[7] = part2[6] + 7
	part2[10] = part2[9] - 1
	part2[11] = part2[8] + 3
	part2[12] = part2[1] + 6
	part2[13] = part2[0]
	if !verify(string(part2)) {
		panic("not a valid input")
	}
	utils.Println(string(part2))
}
