package main

import (
	"sort"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2022, 24)

func main() {
	g := utils.ToByteGrid(utils.Lines(input))
	lcm := utils.LCM((len(g) - 2), (len(g[0]) - 2))

	blizzards := make(map[utils.Pos]byte)
	for y := range g {
		for x, c := range g[y] {
			if c != '#' && c != '.' {
				blizzards[utils.Pos{x, y}] = c
			}
		}
	}
	blizzardAt := func(p utils.Pos, dir byte, t int) utils.Pos {
		t %= lcm
		switch dir {
		case '^':
			p.Y = (p.Y+(lcm-t)-1)%(len(g)-2) + 1
		case 'v':
			p.Y = (p.Y+t-1)%(len(g)-2) + 1
		case '>':
			p.X = (p.X+t-1)%(len(g[0])-2) + 1
		case '<':
			p.X = (p.X+(lcm-t)-1)%(len(g[0])-2) + 1
		}
		return p
	}

	for y := range g {
		for x, c := range g[y] {
			if c != '#' {
				g[y][x] = '.'
			}
		}
	}

	boards := make([][][]byte, lcm)
	for t := range boards {
		boards[t] = utils.CloneGrid(g)
		for b, dir := range blizzards {
			b = blizzardAt(b, dir, t)
			boards[t][b.Y][b.X] = dir
		}
	}
	canMove := func(p utils.Pos, t int) bool {
		return boards[t%len(boards)][p.Y][p.X] == '.'
	}

	start := utils.Pos{1, 0}
	end := utils.Pos{len(g[0]) - 2, len(g) - 1}

	type entry struct {
		p   utils.Pos
		t   int
		leg int
	}
	queue := []entry{{start, 0, 0}}
	seen := make(map[entry]bool)
	var best1, best2 int
	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]
		if seen[q] {
			continue
		}
		if q.t > 2000 {
			continue
		}
		seen[q] = true
		if q.p == start && q.leg == 1 {
			q.leg = 2
		}
		if q.p == end {
			if q.leg == 0 {
				if best1 == 0 || q.t < best1 {
					best1 = q.t
				}
				q.leg = 1
			} else if q.leg == 2 {
				if best2 == 0 || q.t < best2 {
					best2 = q.t
				}
				continue
			}
		}

		// move
		for _, d := range q.p.ValidMoves(utils.Grid{len(g[0]) - 1, len(g) - 1}) {
			if canMove(d, q.t+1) {
				queue = append(queue, entry{d, q.t + 1, q.leg})
			}
		}
		// wait
		if canMove(q.p, q.t+1) {
			queue = append(queue, entry{q.p, q.t + 1, q.leg})
		}

		// only keep the most promising 5,000
		if len(queue) > 5000 {
			sort.Slice(queue, func(i, j int) bool {
				pi := queue[i].t + queue[i].p.Dist(end)
				pj := queue[j].t + queue[j].p.Dist(end)
				return pi < pj
			})
			queue = queue[:len(queue)/2]
		}
	}
	utils.Println(best1)
	utils.Println(best2)
}
