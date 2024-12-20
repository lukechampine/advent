package main

import (
	"sort"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2024, 18)

func path(bytes []utils.Pos) int {
	corrupt := make(map[utils.Pos]bool)
	for _, p := range bytes {
		corrupt[p] = true
	}
	g := utils.Grid{X: 70, Y: 70}
	queue := []utils.Pos{{X: 0, Y: 0}}
	seen := make(map[utils.Pos]int)
	best := 999999
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		if p == (utils.Pos{X: g.X, Y: g.Y}) {
			best = utils.Min(best, seen[p])
			continue
		}
		for _, m := range p.ValidMoves(g) {
			if s, ok := seen[m]; (!ok || s > seen[p]+1) && !corrupt[m] {
				queue = append(queue, m)
				seen[m] = seen[p] + 1
			}
		}
	}
	return best
}

func main() {
	var bytes []utils.Pos
	ints := utils.ExtractInts(input)
	for i := 0; i < len(ints); i += 2 {
		bytes = append(bytes, utils.Pos{X: ints[i], Y: ints[i+1]})
	}
	utils.Println(path(bytes[:1024]))

	n := sort.Search(len(bytes), func(n int) bool {
		return path(bytes[:n]) == 999999
	})
	utils.Print(bytes[n-1].X, ",", bytes[n-1].Y, "\n", n)
}
