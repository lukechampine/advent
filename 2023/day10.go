package main

import (
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2023, 10)

func main() {
	// surround grid with a 1-tile border
	lines := utils.Lines(input)
	border := strings.Repeat(".", len(lines[0])+2)
	for i := range lines {
		lines[i] = "." + lines[i] + "."
	}
	lines = append([]string{border}, append(lines, border)...)
	grid := utils.ToByteGrid(lines)
	g := utils.Grid{X: len(grid[0]) - 1, Y: len(grid) - 1}

	connects := func(a, b utils.Pos) bool {
		up, down, left, right := "S|LJ", "S|7F", "S-J7", "S-LF"
		goes := func(s string, p utils.Pos) bool { return strings.IndexByte(s, grid[p.Y][p.X]) != -1 }
		return map[utils.Dir]bool{
			utils.Up:    goes(up, a) && goes(down, b),
			utils.Down:  goes(down, a) && goes(up, b),
			utils.Left:  goes(left, a) && goes(right, b),
			utils.Right: goes(right, a) && goes(left, b),
		}[a.DirTo(b)]
	}
	walk := func(first utils.Pos, fn func(prev, cur utils.Pos)) {
		start := utils.Locate(grid, 'S')
		prev, cur := start, first
		for cur != start {
			fn(prev, cur)
			var next utils.Pos
			for _, next = range cur.ValidMoves(g) {
				if connects(cur, next) && next != prev {
					break
				}
			}
			prev, cur = cur, next
		}
		fn(prev, cur)
	}

	// part 1
	dists := make(map[utils.Pos]int)
	start := utils.Locate(grid, 'S')
	dists[start] = 0
	for _, first := range start.ValidMoves(g) {
		if !connects(start, first) {
			continue
		}
		walk(first, func(prev, cur utils.Pos) {
			dist := dists[prev] + 1
			if d, ok := dists[cur]; !ok || d > dist {
				dists[cur] = dist
			}
		})
	}
	var maxDist int
	for _, d := range dists {
		maxDist = utils.Max(maxDist, d)
	}
	utils.Println(maxDist)

	// part 2
	// follow loop, flood-filling both sides
	var flood func(utils.Pos, byte)
	flood = func(p utils.Pos, c byte) {
		if _, ok := dists[p]; !ok && g.Valid(p) && grid[p.Y][p.X] != c {
			grid[p.Y][p.X] = c
			for _, m := range p.ValidMoves(g) {
				flood(m, c)
			}
		}
	}
	for _, first := range start.ValidMoves(g) {
		if !connects(start, first) {
			continue
		}
		walk(first, func(prev, cur utils.Pos) {
			dir := prev.DirTo(cur)
			flood(cur.MoveArray(dir.TurnLeft(), 1), '*')
			flood(prev.MoveArray(dir.TurnLeft(), 1), '*')
			flood(cur.MoveArray(dir.TurnRight(), 1), '_')
			flood(prev.MoveArray(dir.TurnRight(), 1), '_')
		})
	}
	// use border to determine which side is inside
	if grid[0][0] == '*' {
		utils.Println(utils.CountGrid(grid, '_'))
	} else {
		utils.Println(utils.CountGrid(grid, '*'))
	}
}
