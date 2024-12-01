package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2023, 13)

func diff(a, b string) int {
	var d int
	for i := range a {
		if a[i] != b[i] {
			d++
		}
	}
	return d
}

func horizontal(grid []string, smudges int) (int, bool) {
	isMirror := func(i int, smudges int) bool {
		if i == len(grid)-1 {
			return false
		}
		above := i
		below := i + 1
		for above >= 0 && below < len(grid) {
			if grid[above] != grid[below] {
				if smudges > 0 && diff(grid[above], grid[below]) == 1 {
					smudges--
				} else {
					return false
				}
			}
			above--
			below++
		}
		return smudges == 0
	}
	for i := range grid {
		if isMirror(i, smudges) {
			return i + 1, true
		}
	}
	return 0, false
}

func transpose(grid []string) []string {
	t := make([]string, len(grid[0]))
	for i := range t {
		for _, row := range grid {
			t[i] += string(row[i])
		}
	}
	return t
}

func vertical(grid []string, smudges int) (int, bool) {
	return horizontal(transpose(grid), smudges)
}

func main() {
	if false {
		input = `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`
	}

	for smudges := 0; smudges < 2; smudges++ {
		var sum int
		for _, group := range utils.Split(input, "\n\n") {
			grid := utils.Lines(group)
			if h, ok := horizontal(grid, smudges); ok {
				sum += 100 * h
			} else if v, ok := vertical(grid, smudges); ok {
				sum += v
			} else {
				panic("unreachable")
			}
		}
		utils.Println(sum)
	}
}
