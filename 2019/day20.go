package main

import (
	"bytes"
	"io/ioutil"

	"github.com/lukechampine/advent/utils"
)

func main() {
	// part 1
	data, err := ioutil.ReadFile("day20_input.txt")
	if err != nil {
		panic(err)
	}
	grid := bytes.Split(bytes.TrimSuffix(data, []byte("\n")), []byte("\n"))

	m := utils.MakeSimpleMaze(grid, '#')
	m.IsWall = func(p utils.Pos) bool { return grid[p.Y][p.X] != '.' }

	teleports := make(map[string][2]utils.Pos)
	for y := range grid {
		for x := range grid[y] {
			c := grid[y][x]
			if !utils.IsLetter(c) {
				continue
			}
			var dot utils.Pos
			var name string
			for _, m := range (utils.Pos{X: x, Y: y}).Moves() {
				if m.X < 0 || m.X >= len(grid[y]) || m.Y < 0 || m.Y >= len(grid) {
					continue
				}
				if d := grid[m.Y][m.X]; utils.IsLetter(d) {
					name = utils.SortString(string(c) + string(d))
				} else if d == '.' {
					dot = m
				}
			}
			if name != "" && dot != (utils.Pos{}) {
				a := teleports[name]
				if a[0] == (utils.Pos{}) {
					a[0] = dot
				} else {
					a[1] = dot
				}
				teleports[name] = a
			}
		}
	}
	revTeleports := make(map[utils.Pos]utils.Pos)
	for _, v := range teleports {
		if v[1] != (utils.Pos{}) {
			revTeleports[v[0]] = v[1]
			revTeleports[v[1]] = v[0]
		}
	}

	type State struct {
		p     utils.Pos
		dist  int
		level int
	}
	start := State{p: teleports["AA"][0]}
	end := teleports["ZZ"][0]
	seen := make(map[State]struct{})
	queue := []State{start}
	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]
		if q.p == end {
			utils.Println(q.dist)
			break
		}
		if _, ok := seen[State{p: q.p}]; ok {
			continue
		}
		seen[State{p: q.p}] = struct{}{}

		moves := m.ValidMoves(q.p)
		if t, ok := revTeleports[q.p]; ok {
			moves = append(moves, t)
		}

		for _, move := range moves {
			queue = append(queue, State{move, q.dist + 1, 0})
		}
	}

	// part 2
	isOuter := func(p utils.Pos) bool {
		return p.X == 2 || p.X == len(grid[0])-3 || p.Y == 2 || p.Y == len(grid)-3
	}
	seen = make(map[State]struct{})
	queue = []State{start}
	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]
		if q.p == end && q.level == 0 {
			utils.Println(q.dist)
			break
		}
		if _, ok := seen[State{p: q.p, level: q.level}]; ok {
			continue
		}
		seen[State{p: q.p, level: q.level}] = struct{}{}

		for _, move := range m.ValidMoves(q.p) {
			queue = append(queue, State{move, q.dist + 1, q.level})
		}
		if t, ok := revTeleports[q.p]; ok {
			level := q.level
			if isOuter(q.p) {
				level--
			} else {
				level++
			}
			if level < 0 {
				continue
			}
			queue = append(queue, State{t, q.dist + 1, level})
		}
	}
}
