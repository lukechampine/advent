package main

import (
	"sort"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2021, 9)
var inputLines = utils.Lines(input)

func main() {
	g := make([][]int, len(inputLines))
	for i, line := range inputLines {
		g[i] = utils.Digits(line)
	}
	grid := utils.Grid{len(g[0]) - 1, len(g) - 1}
	var sum int
	grid.ForEach(func(p utils.Pos) {
		c := g[p.Y][p.X]
		adj := p.ValidMoves(grid)
		if utils.All(len(adj), func(i int) bool { return c < g[adj[i].Y][adj[i].X] }) {
			sum += c + 1
		}
	})
	utils.Println(sum)

	seen := make(map[utils.Pos]bool)
	var flood func(p utils.Pos) []utils.Pos
	flood = func(p utils.Pos) []utils.Pos {
		if seen[p] || g[p.Y][p.X] == 9 {
			return nil
		}
		seen[p] = true
		basin := []utils.Pos{p}
		for _, m := range p.ValidMoves(grid) {
			basin = append(basin, flood(m)...)
		}
		return basin
	}
	var basins [][]utils.Pos
	for y, row := range g {
		for x, c := range row {
			p := utils.Pos{x, y}
			if seen[p] || c == 9 {
				continue
			}
			basins = append(basins, flood(p))
		}
	}
	sort.Slice(basins, func(i, j int) bool { return len(basins[i]) > len(basins[j]) })
	utils.Println(utils.Product(3, func(i int) int { return len(basins[i]) }))
}
