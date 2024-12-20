package main

import (
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2024, 19)

func possible(design string, patterns []string) bool {
	return utils.Any(len(patterns), func(i int) bool {
		rem := strings.TrimPrefix(design, patterns[i])
		return len(rem) < len(design) && (len(rem) == 0 || possible(rem, patterns))
	})
}

var memo = make(map[string]int)

func possibilities(design string, patterns []string) int {
	if _, ok := memo[design]; !ok {
		memo[design] = utils.Sum(len(patterns), func(i int) int {
			if p := patterns[i]; !strings.HasPrefix(design, p) {
				return 0
			} else if len(design) == len(p) {
				return 1
			} else {
				return possibilities(design[len(p):], patterns)
			}
		})
	}
	return memo[design]
}

func main() {
	groups := utils.Split(input, "\n\n")
	patterns := utils.Split(groups[0], ", ")
	designs := utils.Lines(groups[1])
	utils.Println(utils.Count(len(designs), func(i int) bool {
		return possible(designs[i], patterns)
	}))

	utils.Println(utils.Sum(len(designs), func(i int) int {
		return possibilities(designs[i], patterns)
	}))
}
