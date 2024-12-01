package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2023, 5)

func resolve(n int, maps [][][2][2]int) int {
	for _, m := range maps {
		for _, r := range m {
			if r[1][0] <= n && n < r[1][1] {
				n = n - r[1][0] + r[0][0]
				break
			}
		}
	}
	return n
}

func main() {
	groups := utils.Split(input, "\n\n")
	seeds := utils.ExtractInts(groups[0])
	maps := make([][][2][2]int, len(groups)-1)
	for i, g := range groups[1:] {
		for _, line := range utils.Lines(g)[1:] {
			ints := utils.ExtractInts(line)
			maps[i] = append(maps[i], [2][2]int{
				[2]int{ints[0], ints[0] + ints[2]},
				[2]int{ints[1], ints[1] + ints[2]},
			})
		}
	}
	var min int = 1e9
	for _, seed := range seeds {
		min = utils.Min(min, resolve(seed, maps))
	}
	utils.Println(min)

	min = 1e9
	for i := 0; i < len(seeds); i += 2 {
		for seed := seeds[i]; seed < seeds[i]+seeds[i+1]; seed++ {
			min = utils.Min(min, resolve(seed, maps))
		}
	}
	utils.Println(min)
}
