package main

import (
	"sort"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2023, 17)

func bfs(grid [][]byte, minSteps, maxSteps int) int {
	g := utils.Grid{X: len(grid[0]) - 1, Y: len(grid) - 1}
	lossAt := func(p utils.Pos) int {
		return int(grid[p.Y][p.X] - '0')
	}

	end := utils.Pos{X: g.X, Y: g.Y}
	type entry struct {
		pos   utils.Pos
		dir   utils.Dir
		steps int
		loss  int
	}
	bestLoss := make(map[[4]uint8]int)
	var ends []int
	queue := []entry{{pos: utils.Pos{X: 0, Y: 0}, dir: utils.Right}}
	for len(queue) > 0 {
		e := queue[0]
		queue = queue[1:]
		hash := [4]uint8{uint8(e.pos.X), uint8(e.pos.Y), uint8(e.dir), uint8(e.steps)}
		if best, ok := bestLoss[hash]; ok && best <= e.loss {
			continue
		}
		bestLoss[hash] = e.loss
		if e.pos == end {
			if e.steps >= minSteps {
				ends = append(ends, e.loss)
			}
			continue
		}
		addQueue := func(dir utils.Dir, steps int) {
			next := entry{dir: dir, pos: e.pos.MoveArray(dir, 1), steps: steps}
			if g.Valid(next.pos) && next.steps <= maxSteps {
				next.loss = e.loss + lossAt(next.pos)
				queue = append(queue, next)
			}
		}
		addQueue(e.dir, e.steps+1)
		if e.steps >= minSteps {
			addQueue(e.dir.TurnLeft(), 1)
			addQueue(e.dir.TurnRight(), 1)
		}
		if len(queue) > 20000 {
			sort.Slice(queue, func(i, j int) bool {
				return queue[i].loss < queue[j].loss
			})
			queue = queue[:len(queue)-len(queue)/10]
		}
	}
	return utils.Minimum(len(ends), func(i int) int { return ends[i] })
}

func main() {
	grid := utils.ToByteGrid(utils.Lines(input))
	utils.Println(bfs(grid, 0, 3))
	utils.Println(bfs(grid, 4, 10))
}
