package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2024, 20)

func main() {
	grid := utils.ToByteGrid(utils.Lines(input))
	m := utils.MakeSimpleMaze(grid, '#')
	dists := m.DistancesFrom(utils.Locate(grid, 'E'))

	n := 0
	for p := range dists {
		for m := range dists {
			d := p.Dist(m)
			n += utils.BoolToInt(d <= 2 && dists[p]-(dists[m]+d) >= 100)
		}
	}
	utils.Println(n)

	n = 0
	for p := range dists {
		for m := range dists {
			d := p.Dist(m)
			n += utils.BoolToInt(d <= 20 && dists[p]-(dists[m]+d) >= 100)
		}
	}
	utils.Println(n)
}
