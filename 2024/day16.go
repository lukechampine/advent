package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2024, 16)

func main() {
	grid := utils.ToByteGrid(utils.Lines(input))
	start := utils.Locate(grid, 'S')
	end := utils.Locate(grid, 'E')
	m := utils.MakeSimpleMaze(grid, '#')

	type step struct {
		pos utils.Pos
		dir utils.Dir
	}
	type entry struct {
		path []step
		cost int
	}
	queue := []entry{{[]step{{start, utils.Right}}, 0}}
	best := entry{nil, 1<<31 - 1}
	var bests []entry
	costs := make(map[step]int)
	for len(queue) > 0 {
		e := queue[0]
		queue = queue[1:]
		s := e.path[len(e.path)-1]
		if s.pos == end {
			if e.cost < best.cost {
				best = e
				bests = append(bests[:0], e)
			} else if e.cost == best.cost {
				bests = append(bests, e)
			}
			continue
		}
		for _, ns := range []step{
			{s.pos.MoveArray(s.dir, 1), s.dir},
			{s.pos, s.dir.SpinLeft(1)},
			{s.pos, s.dir.SpinRight(1)},
		} {
			cost := e.cost + 1 + 999*utils.BoolToInt(ns.dir != s.dir)
			if c, ok := costs[ns]; (ok && cost > c) || !m.Valid(ns.pos) {
				continue
			}
			costs[ns] = cost
			queue = append(queue, entry{append(append([]step(nil), e.path...), ns), cost})
		}
	}
	utils.Println(best.cost)
	seats := make(map[utils.Pos]struct{})
	for _, e := range bests {
		for _, s := range e.path {
			seats[s.pos] = struct{}{}
		}
	}
	utils.Println(len(seats))
}
