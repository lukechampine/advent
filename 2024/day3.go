package main

import (
	"regexp"
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2024, 3)

func main() {
	var sum1, sum2 int
	enabled := true
	for _, match := range regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`).FindAllString(input, -1) {
		if strings.HasPrefix(match, "do") {
			enabled = match == "do()"
			continue
		}
		ints := utils.ExtractInts(match)
		sum1 += ints[0] * ints[1]
		sum2 += utils.BoolToInt(enabled) * ints[0] * ints[1]
	}
	utils.Println(sum1)
	utils.Println(sum2)
}
