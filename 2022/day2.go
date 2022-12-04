package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2022, 2)

var beats = map[byte]byte{
	'A': 'Z',
	'B': 'X',
	'C': 'Y',
	'X': 'C',
	'Y': 'A',
	'Z': 'B',
}

func score(a, b byte) int {
	var outcome int
	var choice int = int(b - 'X' + 1)
	switch {
	case beats[a] == b:
		outcome = 0
	default:
		outcome = 3
	case beats[b] == a:
		outcome = 6
	}
	return outcome + choice
}

func score2(a, b byte) int {
	var outcome int = 3 * int(b-'X')
	var choice byte
	switch b {
	case 'X': // lose
		choice = beats[a]
	case 'Y': // tie
		choice = 'X' + (a - 'A')
	case 'Z': // win
		for k, v := range beats {
			if v == a {
				choice = k
				break
			}
		}
	}
	return outcome + int(choice-'X'+1)
}

func main() {
	// part 1
	var sum int
	for _, line := range utils.Lines(input) {
		sum += score(line[0], line[2])
	}
	utils.Println(sum)

	// part 2
	sum = 0
	for _, line := range utils.Lines(input) {
		sum += score2(line[0], line[2])
	}
	utils.Println(sum)
}
