package main

import (
	"sort"
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2021, 10)
var inputLines = utils.Lines(input)

func parse(line string) (illegal byte, missing string) {
	closer := []byte{'(': ')', '[': ']', '{': '}', '<': '>'}
	var expect string
	for ; len(line) > 0; line = line[1:] {
		switch c := line[0]; c {
		case '(', '[', '{', '<':
			expect = string(closer[c]) + expect
		case ')', ']', '}', '>':
			if !strings.HasPrefix(expect, string(c)) {
				return c, expect
			}
			expect = expect[1:]
		}
	}
	return 0, expect
}

func main() {
	var sum int
	var scores []int
	for _, line := range inputLines {
		illegal, missing := parse(line)
		if illegal != 0 {
			sum += []int{')': 3, ']': 57, '}': 1197, '>': 25137}[illegal]
		} else {
			var score int
			for _, c := range missing {
				score *= 5
				score += strings.IndexRune(" )]}>", c)
			}
			scores = append(scores, score)
		}
	}
	utils.Println(sum)
	sort.Ints(scores)
	utils.Println(scores[len(scores)/2])
}
