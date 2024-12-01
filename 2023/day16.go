package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2023, 16)

type entry struct {
	pos utils.Pos
	dir utils.Dir
}

func numEnergized(grid [][]byte, start entry) int {
	g := utils.Grid{X: len(grid[0]) - 1, Y: len(grid) - 1}

	seen := make(map[entry]bool)
	queue := []entry{start}
	for len(queue) > 0 {
		e := queue[0]
		queue = queue[1:]
		if seen[e] {
			continue
		}
		seen[e] = true

		e.pos = e.pos.MoveArray(e.dir, 1)
		if !g.Valid(e.pos) {
			continue
		}
		switch grid[e.pos.Y][e.pos.X] {
		case '.':
			queue = append(queue, entry{e.pos, e.dir})
		case '/':
			if e.dir == utils.Up || e.dir == utils.Down {
				e.dir = e.dir.TurnRight()
			} else {
				e.dir = e.dir.TurnLeft()
			}
			queue = append(queue, entry{e.pos, e.dir})
		case '\\':
			if e.dir == utils.Up || e.dir == utils.Down {
				e.dir = e.dir.TurnLeft()
			} else {
				e.dir = e.dir.TurnRight()
			}
			queue = append(queue, entry{e.pos, e.dir})
		case '|':
			if e.dir == utils.Up || e.dir == utils.Down {
				queue = append(queue, entry{e.pos, e.dir})
			} else {
				// split
				queue = append(queue, entry{e.pos, utils.Up})
				queue = append(queue, entry{e.pos, utils.Down})
			}
		case '-':
			if e.dir == utils.Up || e.dir == utils.Down {
				// split
				queue = append(queue, entry{e.pos, utils.Left})
				queue = append(queue, entry{e.pos, utils.Right})
			} else {
				queue = append(queue, entry{e.pos, e.dir})
			}
		}
	}
	delete(seen, start)
	for e := range seen {
		delete(seen, e)
		e.dir = 0
		seen[e] = true
	}
	return len(seen)
}

func main() {
	if false {
		input = `.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`
	}
	grid := utils.ToByteGrid(utils.Lines(input))

	utils.Println(numEnergized(grid, entry{utils.Pos{X: -1, Y: 0}, utils.Right}))

	var max int
	for i := range grid {
		max = utils.Max(max, numEnergized(grid, entry{utils.Pos{X: -1, Y: i}, utils.Right}))
		max = utils.Max(max, numEnergized(grid, entry{utils.Pos{X: len(grid[0]), Y: i}, utils.Left}))
		max = utils.Max(max, numEnergized(grid, entry{utils.Pos{X: i, Y: -1}, utils.Down}))
		max = utils.Max(max, numEnergized(grid, entry{utils.Pos{X: i, Y: len(grid)}, utils.Up}))
	}
	utils.Println(max)
}
