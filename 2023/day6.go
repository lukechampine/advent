package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2023, 5)

type seedRange struct {
	start, end int
}

func (s seedRange) overlaps(o seedRange) bool {
	return s.start <= o.end && o.start <= s.end
}

func (s seedRange) overlap(o seedRange) seedRange {
	return seedRange{utils.Max(s.start, o.start), utils.Min(s.end, o.end)}
}

type rangeMapping struct {
	src, dst seedRange
}

func (o rangeMapping) overlap(m []rangeMapping) []rangeMapping {
	var matches []rangeMapping
	for _, r := range m {
		if o.dst.overlaps(r.dst) {
			matches = append(matches, rangeMapping{o.src, r.src})
		}
	}
	return matches
}

func lowestValid(r rangeMapping, maps [][]rangeMapping, seeds []seedRange) {
	if len(maps) == 0 {
		for _, s := range seeds {
			if r.src.overlaps(s) {
				panic(r.src.overlap(s).start)
			}
		}
		return
	}
	for _, match := range r.overlap(maps[len(maps)-1]) {
		lowestValid(match, maps[:len(maps)-1], seeds)
	}
}

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
	utils.Println(utils.Minimum(len(seeds), func(i int) int {
		return resolve(seeds[i], maps)
	}))

	var min int = 1e9
	for i := 0; i < len(seeds); i += 2 {
		for seed := seeds[i]; seed < seeds[i]+seeds[i+1]; seed++ {
			min = utils.Min(min, resolve(seed, maps))
		}
	}
	utils.Println(min)
}
