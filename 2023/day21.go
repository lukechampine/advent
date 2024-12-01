package main

import (
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2023, 21)

func dfs(m utils.Maze, distances map[utils.Pos]int, dist, maxDist int, cur utils.Pos) {
	if dist > maxDist {
		return
	} else if d, ok := distances[cur]; ok && d <= dist {
		return
	}
	distances[cur] = dist
	for _, p := range m.ValidMoves(cur) {
		dfs(m, distances, dist+1, maxDist, p)
	}
}

func count(grid [][]byte, steps int) (plots int) {
	m := utils.MakeSimpleMaze(grid, '#')
	distances := make(map[utils.Pos]int)
	dfs(m, distances, 0, steps, utils.Locate(grid, 'S'))
	for _, d := range distances {
		plots += utils.BoolToInt(steps%2 == d%2)
	}
	return plots
}

func main() {
	utils.Println(count(utils.ToByteGrid(utils.Lines(input)), 64))

	// tile input into 5x5
	lines := utils.Lines(strings.Repeat(strings.Replace(input, "S", ".", -1), 5))
	for i := range lines {
		lines[i] = strings.Repeat(lines[i], 5)
	}
	grid := utils.ToByteGrid(lines)
	grid[len(grid)/2][len(grid)/2] = 'S'

	y0 := count(grid, 65)
	y1 := count(grid, 65+131)
	y2 := count(grid, 65+2*131)

	d0 := y1 - y0
	dd0 := (y2 - y1) - (y1 - y0)

	reps := (26501365 - 65) / 131

	a := d0
	b := y0
	for i := 1; i < reps; i++ {
		a += dd0
		b += a
	}
	utils.Println(b)
}
