package main

import (
	"strconv"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2021, 3)
var inputLines = utils.Lines(input)

func oxy(lines []string, i int) byte {
	var n0, n1 int
	for _, line := range lines {
		n0 += utils.BoolToInt(line[i] == '0')
		n1 += utils.BoolToInt(line[i] == '1')
	}
	if n1 >= n0 {
		return '1'
	}
	return '0'
}

func co2(lines []string, i int) byte {
	return oxy(lines, i) ^ 1
}

func filter(lines []string, i int, c byte) []string {
	var rem []string
	for _, line := range lines {
		if line[i] == c {
			rem = append(rem, line)
		}
	}
	return rem
}

func fromBinary(s string) int {
	i, _ := strconv.ParseInt(s, 2, 64)
	return int(i)
}

func main() {
	// part 1
	var gamma, epsilon string
	for i := 0; i < len(inputLines[0]); i++ {
		m := oxy(inputLines, i)
		gamma += string(m)
		epsilon += string(m ^ 1)
	}
	utils.Println(fromBinary(gamma) * fromBinary(epsilon))

	// part 2
	rem := inputLines
	for i := 0; len(rem) > 1; i++ {
		rem = filter(rem, i, oxy(rem, i))
	}
	oxyrating := fromBinary(rem[0])

	rem = inputLines
	for i := 0; len(rem) > 1; i++ {
		rem = filter(rem, i, co2(rem, i))
	}
	co2rating := fromBinary(rem[0])
	utils.Println(oxyrating * co2rating)
}
