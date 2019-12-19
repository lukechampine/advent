package main

import (
	"github.com/lukechampine/advent/utils"
)

var input = utils.ReadInput("day18_input.txt")
var grid = utils.ToByteGrid(utils.Lines(input))

type keyring struct {
	s string
}

func (kr keyring) add(k byte) keyring {
	if !utils.ContainsByte(kr.s, k) {
		kr.s = utils.SortString(kr.s + string(k))
	}
	return kr
}

func (kr keyring) opens(d byte) bool {
	return utils.ContainsByte(kr.s, d+('a'-'A'))
}

func isKey(c byte) bool  { return 'a' <= c && c <= 'z' }
func isDoor(c byte) bool { return 'A' <= c && c <= 'Z' }

type State struct {
	p    utils.Pos
	keys keyring
	dist int
}

func shortestRoute(m utils.Maze, grid [][]byte, start State, fullRing keyring) int {
	seen := make(map[State]struct{}, 1<<23)
	queue := []State{start}
	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]

		// have we been here before, with the same set of keys?
		if _, ok := seen[State{q.p, q.keys, 0}]; ok {
			continue
		}
		seen[State{q.p, q.keys, 0}] = struct{}{}

		// if we found a key, add it to our ring; if we've found all keys, we're
		// done. (There might be other routes that collect all keys, but since
		// this is a BFS, no future iteration will have a lower dist than this.)
		if k := grid[q.p.Y][q.p.X]; isKey(k) {
			q.keys = q.keys.add(k)
			if q.keys == fullRing {
				return q.dist
			}
		}
		// try next moves that aren't walls or un-openable doors
		for _, move := range m.ValidMoves(q.p) {
			if d := grid[move.Y][move.X]; isDoor(d) && !q.keys.opens(d) {
				continue
			}
			queue = append(queue, State{p: move, keys: q.keys, dist: q.dist + 1})
		}
	}
	panic("unreachable")
}

func main() {
	// part 1
	var fullRing keyring
	for y := range grid {
		for _, c := range grid[y] {
			if 'a' <= c && c <= 'z' {
				fullRing = fullRing.add(c)
			}
		}
	}
	m := utils.MakeSimpleMaze(grid, '#')
	// reduce search space by filling in dead ends (including dead ends
	// containing just a door; we don't need to go through all doors, just
	// collect all keys)
	m.IsWall = m.Minimize(func(p utils.Pos) bool {
		return grid[p.Y][p.X] == '.' || isDoor(grid[p.Y][p.X])
	})

	center := utils.Locate(grid, '@')
	utils.Println(shortestRoute(m, grid, State{p: center}, fullRing))

	// part 2
	copy(grid[center.Y-1][center.X-1:], "@#@")
	copy(grid[center.Y+0][center.X-1:], "###")
	copy(grid[center.Y+1][center.X-1:], "@#@")

	// BFS each quadrant, simply ignoring doors whose keys are in a different quadrant
	starts := [4]utils.Pos{
		center.Add(-1, -1),
		center.Add(+1, -1),
		center.Add(-1, +1),
		center.Add(+1, +1),
	}
	var cheatRings [4]keyring
	quadrant := func(x, y int) int {
		switch {
		case x <= starts[0].X && y <= starts[0].Y:
			return 0
		case x >= starts[1].X && y <= starts[1].Y:
			return 1
		case x <= starts[2].X && y >= starts[2].Y:
			return 2
		case x >= starts[3].X && y >= starts[3].Y:
			return 3
		default:
			panic("unreachable")
		}
	}
	for y := range grid {
		for x, c := range grid[y] {
			if 'a' <= c && c <= 'z' {
				// add key to all other quadrants' keyrings
				q := quadrant(x, y)
				for i := range cheatRings {
					if i != q {
						cheatRings[i] = cheatRings[i].add(c)
					}
				}
			}
		}
	}

	utils.Println(utils.Sum(len(starts), func(i int) int {
		return shortestRoute(m, grid, State{p: starts[i], keys: cheatRings[i]}, fullRing)
	}))
}
