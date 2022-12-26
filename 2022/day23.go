package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2022, 23)

func main() {
	grid := utils.Lines(input)
	g := utils.Grid{len(grid[0]) - 1, len(grid) - 1}

	elves := make(map[utils.Pos]struct{})
	g.ForEach(func(p utils.Pos) {
		if grid[p.Y][p.X] == '#' {
			elves[p] = struct{}{}
		}
	})

	noElves := func(ps []utils.Pos) bool {
		for _, p := range ps {
			if _, ok := elves[p]; ok {
				return false
			}
		}
		return true
	}

	for round := 1; ; round++ {
		moving := make(map[utils.Pos][]utils.Pos)
		for e := range elves {
			adj := e.Numpad()
			if noElves(adj) {
				continue
			}
			opts := [4][]utils.Pos{
				{adj[3], adj[5], adj[0]},
				{adj[4], adj[7], adj[2]},
				{adj[1], adj[0], adj[2]},
				{adj[6], adj[5], adj[7]},
			}
			for i := range opts {
				os := opts[(i+round-1)%len(opts)]
				d := os[0]
				if noElves(os) {
					moving[d] = append(moving[d], e)
					break
				}
			}
		}
		var anyMoved bool
		for d, es := range moving {
			if len(es) == 1 {
				delete(elves, es[0])
				elves[d] = struct{}{}
				anyMoved = true
			}
		}

		if round == 10 {
			var min, max utils.Pos
			min.X, min.Y = int(1e9), int(1e9)
			max.X, max.Y = int(-1e9), int(-1e9)
			for e := range elves {
				min.X = utils.Min(min.X, e.X)
				min.Y = utils.Min(min.Y, e.Y)
				max.X = utils.Max(max.X, e.X)
				max.Y = utils.Max(max.Y, e.Y)
			}
			utils.Println((max.X-min.X+1)*(max.Y-min.Y+1) - len(elves))
		}
		if !anyMoved {
			utils.Println(round)
			break
		}
	}
}
