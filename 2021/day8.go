package main

import (
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

	// determine easy digits
	var digits [10]uint8
	words := utils.Split(seq, " ")
	for _, w := range words {
		switch len(w) {
		case 2:
			digits[1] = toMask(w)
		case 3:
			digits[7] = toMask(w)
		case 4:
			digits[4] = toMask(w)
		case 7:
			digits[8] = toMask(w)
		}
	}

	// determine remaining digits
	covers := func(a, b uint8) bool { return a&b == a }
	bd := digits[4] &^ digits[1]
	for _, w := range words {
		switch len(w) {
		case 5: // 2, 3, 5
			if m := toMask(w); covers(digits[1], m) {
				digits[3] = m
			} else if covers(bd, m) {
				digits[5] = m
			} else {
				digits[2] = m
			}
		case 6: // 0, 6, 9
			if m := toMask(w); !covers(digits[1], m) {
				digits[6] = m
			} else if !covers(bd, m) {
				digits[0] = m
			} else {
				digits[9] = m
			}
		}
	}

	// translate output
	words = utils.Split(output, " ")
	b := make([]byte, len(words))
	for i, w := range words {
		m := toMask(w)
		for n, d := range digits {
			if d == m {
				b[i] = '0' + byte(n)
				break
			}
		}
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
