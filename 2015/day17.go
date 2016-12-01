package main

import (
	"github.com/lukechampine/advent/utils"
)

const input = `43
3
4
10
21
44
4
6
47
41
34
17
17
44
36
31
46
9
27
38`

func combinations(liters int, cs []int) [][]int {
	var combos [][]int
	for i, c := range cs {
		switch {
		case c > liters:
			// can't use this container
		case c == liters:
			combos = append(combos, []int{c})
		case c < liters:
			// we want combinations, not permutations, so don't reuse cs[:i]
			for _, combo := range combinations(liters-c, cs[i+1:]) {
				combos = append(combos, append(combo, c))
			}
		}
	}
	return combos
}

func main() {
	// part 1
	var containers []int
	for _, str := range utils.Lines(input) {
		containers = append(containers, utils.Atoi(str))
	}
	combos := combinations(150, containers)
	println(len(combos))

	// part 2
	minLen := utils.Minimum(len(combos), func(i int) int {
		return len(combos[i])
	})
	count := 0
	for _, c := range combos {
		if len(c) == minLen {
			count++
		}
	}
	println(count)
}
