package main

import (
	"github.com/lukechampine/advent/utils"
)

const input = `####....#...######.###.#...##....#.###.#.###.......###.##..##........##..#.#.#..##.##...####.#..##.#`

const ruleStrings = `..#.. => .
#.#.# => #
#.### => #
.##.. => .
#.#.. => #
.#.#. => #
.###. => #
.#### => #
##... => #
#.##. => #
#..## => #
....# => .
###.# => .
##### => #
..... => .
..#.# => .
.#... => #
##.#. => .
.#.## => #
..##. => .
#...# => .
##.## => #
...#. => .
#..#. => .
..### => .
.##.# => .
#.... => .
.#..# => #
####. => .
...## => #
##..# => .
###.. => .`

func calcScore(pots []byte) int {
	var total int
	for i := range pots {
		if pots[i] == '#' {
			total += i - len(pots)/2
		}
	}
	return total
}

func main() {
	rules := make(map[string]byte)
	for _, line := range utils.Lines(ruleStrings) {
		rules[line[:5]] = line[len(line)-1]
	}

	var gen20score int
	pots := make([]byte, 5000)
	newpots := make([]byte, len(pots))
	for i := range pots {
		pots[i] = '.'
	}
	for i := range input {
		pots[i+len(pots)/2] = input[i]
	}
	// growth factor converges after ~100 generations
	const convergeGens = 500
	var last, diff int
	for gen := 0; gen < convergeGens; gen++ {
		for i := 2; i < len(pots)-2; i++ {
			newpots[i] = rules[string(pots[i-2:i+3])]
		}
		copy(pots, newpots)

		score := calcScore(pots)
		if gen == 19 {
			gen20score = score
		}
		last, diff = score, score-last
	}

	// part 1
	utils.Println(gen20score)

	// part 2
	utils.Println(last + diff*(50e9-convergeGens))
}
