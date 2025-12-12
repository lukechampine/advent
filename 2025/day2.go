package main

import (
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2025, 2)

func isRepeated(s string) bool {
	for i := 1; i <= len(s)/2; i++ {
		if len(s)%i == 0 && s == strings.Repeat(s[:i], len(s)/i) {
			return true
		}
	}
	return false
}

func main() {
	var part1, part2 int
	for _, r := range utils.Split(input, ",") {
		var start, end int
		utils.Sscanf(r, "%d-%d", &start, &end)
		for i := start; i <= end; i++ {
			s := utils.Itoa(i)
			if s[:len(s)/2] == s[len(s)/2:] {
				part1 += i
			}
			if isRepeated(s) {
				part2 += i
			}
		}
	}
	utils.Println(part1)
	utils.Println(part2)
}
