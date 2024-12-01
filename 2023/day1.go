package main

import (
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2023, 1)

func calibration(s string) int {
	var first, last int
	for i := range s {
		if '0' <= s[i] && s[i] <= '9' {
			if first == 0 {
				first = int(s[i] - '0')
			}
			last = int(s[i] - '0')
		}
	}
	return 10*first + last
}

func calibration2(s string) int {
	var first, last int
	for i := 0; i < len(s); i++ {
		if '0' <= s[i] && s[i] <= '9' {
			if first == 0 {
				first = int(s[i] - '0')
			}
			last = int(s[i] - '0')
			continue
		}
		for n, p := range []string{
			"one", "two", "three",
			"four", "five", "six",
			"seven", "eight", "nine",
		} {
			if strings.HasPrefix(s[i:], p) {
				if first == 0 {
					first = n + 1
				}
				last = n + 1
				break
			}
		}
	}
	return 10*first + last
}

func main() {
	lines := utils.Lines(input)
	utils.Println(utils.Sum(len(lines), func(i int) int {
		return calibration(lines[i])
	}))
	utils.Println(utils.Sum(len(lines), func(i int) int {
		return calibration2(lines[i])
	}))
}
