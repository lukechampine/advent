package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2023, 23)

type graphEdge struct {
	dest  utils.Pos
	steps int
}

func makeGraph(grid [][]byte, start, end utils.Pos) map[utils.Pos][]graphEdge {
	graph := make(map[utils.Pos][]graphEdge)
	m := utils.MakeSimpleMaze(grid, '#')
	m.Grid.ForEach(func(p utils.Pos) {
		if len(m.ValidMoves(p)) > 2 {
			graph[p] = nil
		}
	})
	graph[start] = nil
	graph[end] = nil

	findEdge := func(prev, cur utils.Pos) graphEdge {
		dist := 1
		for {
			moves := m.ValidMoves(cur)
			if len(moves) != 2 {
				return graphEdge{cur, dist}
			}
			next := moves[0]
			if next == prev {
				next = moves[1]
			}
			prev, cur = cur, next
			dist++
			if c := grid[cur.Y][cur.X]; utils.ContainsByte("^>v<", c) {
				prev, cur = cur, cur.MoveArray(utils.DirFromArrow(c), 1)
				dist++
			}
		}
	}

	for v := range graph {
		for _, p := range m.ValidMoves(v) {
			graph[v] = append(graph[v], findEdge(v, p))
		}
	}
	return graph
}

func maxDistance(graph map[utils.Pos][]graphEdge, cur, end utils.Pos, path []utils.Pos) int {
	if cur == end {
		return 0
	}
	return utils.Maximum(len(graph[cur]), func(i int) int {
		e := graph[cur][i]
		if utils.Any(len(path), func(i int) bool { return path[i] == e.dest }) {
			return -1e9
		}
		return e.steps + maxDistance(graph, e.dest, end, append(path, e.dest))
	})
}

func main() {
	grid := utils.ToByteGrid(utils.Lines(input))
	start, end := utils.Pos{X: 1, Y: 0}, utils.Pos{X: len(grid[0]) - 2, Y: len(grid) - 1}
	utils.Println(maxDistance(makeGraph(grid, start, end), start, end, nil))
	grid = utils.ToByteGrid(utils.Lines(utils.Replace(input, "^", ".", "<", ".", "v", ".", ">", ".")))
	utils.Println(maxDistance(makeGraph(grid, start, end), start, end, nil))
}
