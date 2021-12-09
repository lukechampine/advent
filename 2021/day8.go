package main

import (
	"math/bits"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2021, 8)
var inputLines = utils.Lines(input)

func solve(seq string, output string) int {
	toMask := func(s string) uint8 {
		var mask uint8
		for _, c := range s {
			mask |= 1 << (c - 'a')
		}
		return mask
	}

	masks := make(map[int]uint8)
	for _, w := range utils.Split(seq, " ") {
		switch len(w) {
		case 2, 4:
			masks[len(w)] = toMask(w)
		}
	}
	var b []byte
	for _, w := range utils.Split(output, " ") {
		m := toMask(w)
		b = append(b, map[[3]int]byte{
			{6, 3, 2}: '0',
			{2, 2, 2}: '1',
			{5, 2, 1}: '2',
			{5, 3, 2}: '3',
			{4, 4, 2}: '4',
			{5, 3, 1}: '5',
			{6, 3, 1}: '6',
			{3, 2, 2}: '7',
			{7, 4, 2}: '8',
			{6, 4, 2}: '9',
		}[[3]int{bits.OnesCount8(m), bits.OnesCount8(m & masks[4]), bits.OnesCount8(m & masks[2])}])
	}
	return utils.Atoi(string(b))
}

func main() {
	var sum int
	for _, line := range inputLines {
		seq := utils.Split(line, " | ")[1]
		for _, w := range utils.Split(seq, " ") {
			switch len(w) {
			case 2, 4, 3, 7:
				sum++
			}
		}
	}
	utils.Println(sum)

	sum = 0
	for _, line := range inputLines {
		parts := utils.Split(line, " | ")
		sum += solve(parts[0], parts[1])
	}
	utils.Println(sum)
}
