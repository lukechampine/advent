package main

import (
	"container/heap"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2021, 15)

type entry struct {
	utils.Pos
	cost int
}

type entries []entry

func (es entries) Len() int           { return len(es) }
func (es entries) Less(i, j int) bool { return es[i].cost < es[j].cost }
func (es entries) Swap(i, j int)      { es[i], es[j] = es[j], es[i] }

func (es *entries) Push(x interface{}) {
	*es = append(*es, x.(entry))
}

func (es *entries) Pop() interface{} {
	old := *es
	*es = old[:len(old)-1]
	return old[len(old)-1]
}

func dijkstra(grid [][]byte) int {
	end := utils.Pos{X: len(grid[0]) - 1, Y: len(grid) - 1}

	queue := entries{{utils.Pos{0, 0}, 0}}
	seen := make(map[utils.Pos]bool)
	for len(queue) > 0 {
		p := heap.Pop(&queue).(entry)
		if seen[p.Pos] {
			continue
		} else if p.Pos == end {
			return p.cost
		}
		for _, m := range p.ValidMoves(utils.Grid{X: end.X, Y: end.Y}) {
			if !seen[m] {
				heap.Push(&queue, entry{m, p.cost + int(grid[m.Y][m.X]-'0')})
			}
		}
		seen[p.Pos] = true
	}

	panic("unreachable")
}

func inc(b []byte) []byte {
	b2 := make([]byte, len(b))
	for i := range b {
		b2[i] = b[i] + 1
		if b2[i] > '9' {
			b2[i] = '1'
		}
	}
	return b2
}

func main() {
	grid := utils.ToByteGrid(utils.Lines(input))
	utils.Println(dijkstra(grid))

	fullMaze := make([][]byte, len(grid)*5)
	for i := range grid {
		row := make([]byte, 0, len(grid[i])*5)
		row = append(row, grid[i]...)
		for j := 0; j < 4; j++ {
			row = append(row, inc(row[j*len(grid[i]):][:len(grid[i])])...)
		}
		fullMaze[i] = row
	}
	for i := len(grid); i < len(fullMaze); i++ {
		fullMaze[i] = inc(fullMaze[i-len(grid)])
	}

	utils.Println(dijkstra(fullMaze))
}
