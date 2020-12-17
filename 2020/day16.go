package main

import (
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2020, 16)

func check(i int, ranges [][]int) bool {
	for _, r := range ranges {
		if (r[0] <= i && i <= r[1]) || (r[2] <= i && i <= r[3]) {
			return true
		}
	}
	return false
}

func main() {
	groups := utils.Split(input, "\n\n")
	rules, nearby := utils.Lines(groups[0]), utils.Lines(groups[2])[1:]

	var ranges [][]int
	for _, l := range rules {
		ranges = append(ranges, utils.ExtractInts(utils.Replace(l, "-", " ")))
	}

	var rate int
	for _, l := range nearby {
		for _, i := range utils.ExtractInts(l) {
			if !check(i, ranges) {
				rate += i
			}
		}
	}
	utils.Println(rate)

	var valid [][]int
	for _, l := range nearby {
		ticket := utils.ExtractInts(l)
		if utils.All(len(ticket), func(i int) bool { return check(ticket[i], ranges) }) {
			valid = append(valid, ticket)
		}
	}

	possible := make([][]int, len(ranges))
	for ri := range ranges {
		for i := range valid[0] {
			if utils.All(len(valid), func(j int) bool { return check(valid[j][i], ranges[ri:ri+1]) }) {
				possible[ri] = append(possible[ri], i)
			}
		}
	}
	certain := make(map[int]int)
	for done := false; !done; {
		done = true
		for i, p := range possible {
			if len(p) == 1 {
				certain[i] = p[0]
				for j := range possible {
					utils.DeleteSliceValue(&possible[j], p[0])
				}
				done = false
				break
			}
		}
	}

	ours := utils.ExtractInts(groups[1])
	product := 1
	for ri, index := range certain {
		if strings.HasPrefix(rules[ri], "departure") {
			product *= ours[index]
		}
	}
	utils.Println(product)
}
