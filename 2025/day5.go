package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2025, 5)

func main() {
	groups := utils.Split(input, "\n\n")
	var ranges [][2]int
	for _, line := range utils.Lines(groups[0]) {
		var lo, hi int
		utils.Sscanf(line, "%d-%d", &lo, &hi)
		ranges = append(ranges, [2]int{lo, hi})
	}
	ingredients := utils.ExtractInts(groups[1])
	utils.Println(utils.Count(len(ingredients), func(i int) bool {
		x := ingredients[i]
		for _, r := range ranges {
			if x >= r[0] && x <= r[1] {
				return true
			}
		}
		return false
	}))

	for {
		rem := ranges[:1]
		changed := false
	outer:
		for _, r := range ranges[1:] {
			for i, q := range rem {
				if r[0] <= q[1]+1 && r[1] >= q[0]-1 {
					rem[i][0] = utils.Min(rem[i][0], r[0])
					rem[i][1] = utils.Max(rem[i][1], r[1])
					changed = true
					continue outer
				}
			}
			rem = append(rem, r)
		}
		ranges = rem
		if !changed {
			break
		}
	}
	var total int
	for _, r := range ranges {
		total += r[1] - r[0] + 1
	}
	utils.Println(total)
}
