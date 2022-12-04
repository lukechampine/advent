package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2022, 3)

func common(ss ...string) rune {
	counts := make(map[rune]int)
	for _, s := range ss {
		for c := range utils.CharCounts(s) {
			counts[c]++
		}
	}
	for c, n := range counts {
		if n == len(ss) {
			return c
		}
	}
	panic("unreachable")
}

func priority(b rune) int {
	if b >= 'a' {
		return int(b-'a') + 1
	}
	return int(b-'A') + 27
}

func main() {
	var sum int
	for _, line := range utils.Lines(input) {
		sum += priority(common(line[:len(line)/2], line[len(line)/2:]))
	}
	utils.Println(sum)

	sum = 0
	for _, group := range utils.Window(utils.Lines(input), 3).([][]string) {
		sum += priority(common(group...))
	}
	utils.Println(sum)
}
