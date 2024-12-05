package main

import (
	"fmt"
	"sort"
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2024, 5)

func main() {
	rules, updates, _ := strings.Cut(input, "\n\n")
	var part1, part2 int
	for _, line := range utils.Lines(updates) {
		u := utils.ExtractInts(line)
		ruleSort := func(i, j int) bool {
			return strings.Contains(rules, fmt.Sprintf("%d|%d", u[i], u[j]))
		}
		if sort.SliceIsSorted(u, ruleSort) {
			part1 += u[len(u)/2]
		} else {
			sort.Slice(u, ruleSort)
			part2 += u[len(u)/2]
		}
	}
	utils.Println(part1)
	utils.Println(part2)
}
