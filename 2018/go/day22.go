package main

import (
	"container/heap"

	"github.com/lukechampine/advent/utils"
)

const depth = 11817

var target = utils.Pos{X: 9, Y: 751}

const rocky = 0
const wet = 1
const narrow = 2

const neither = 0
const torch = 1
const gear = 2

func validTool(typ int, tool int) bool {
	return (typ == 0 && tool != neither) ||
		(typ == 1 && tool != torch) ||
		(typ == 2 && tool != gear)
}

type region struct {
	index   int
	erosion int
	typ     int
}

// for Djikstra
type vertex struct {
	utils.Pos
	tool int
}

// heap implementation for Djikstra
type distHeap struct {
	v       []vertex
	indices map[vertex]int // for efficient heap.Fix calls later
	dists   map[vertex]int
}

func (h distHeap) Len() int           { return len(h.v) }
func (h distHeap) Less(i, j int) bool { return h.dists[h.v[i]] < h.dists[h.v[j]] }
func (h distHeap) Swap(i, j int) {
	h.v[i], h.v[j] = h.v[j], h.v[i]
	h.indices[h.v[i]] = i
	h.indices[h.v[j]] = j
}

func (h *distHeap) Push(x interface{}) {
	h.v = append(h.v, x.(vertex))
	h.indices[x.(vertex)] = len(h.v) - 1
}

func (h *distHeap) Pop() interface{} {
	u := h.v[len(h.v)-1]
	h.v = h.v[:len(h.v)-1]
	delete(h.indices, u)
	return u
}

func main() {
	// part 1
	grid := make([][]region, target.Y+20)
	for i := range grid {
		grid[i] = make([]region, target.X+20)
	}
	for y := range grid {
		for x := range grid[y] {
			var index int
			switch {
			case x == 0 && y == 0:
				index = 0
			case x == target.X && y == target.Y:
				index = 0
			case y == 0:
				index = x * 16807
			case x == 0:
				index = y * 48271
			default:
				index = grid[y][x-1].erosion * grid[y-1][x].erosion
			}
			erosion := (index + depth) % 20183
			grid[y][x] = region{
				index:   index,
				erosion: erosion,
				typ:     erosion % 3,
			}
		}
	}
	var risk int
	for y := range grid[:target.Y+1] {
		for _, r := range grid[y][:target.X+1] {
			risk += r.typ
		}
	}
	utils.Println(risk)

	// part 2
	// create vertices
	var vertices []vertex
	for y := range grid {
		for x, r := range grid[y] {
			if (x == 0 && y == 0) || (x == target.X && y == target.Y) {
				vertices = append(vertices, vertex{utils.Pos{X: x, Y: y}, torch})
			} else {
				for tool := neither; tool <= gear; tool++ {
					if validTool(r.typ, tool) {
						vertices = append(vertices, vertex{utils.Pos{X: x, Y: y}, tool})
					}
				}
			}
		}
	}

	// create edges
	edges := make(map[vertex][]vertex)
	g := utils.Grid{X: len(grid[0]) - 1, Y: len(grid) - 1}
	for _, v := range vertices {
		// add an edge for changing tools
		for otherTool := neither; otherTool <= gear; otherTool++ {
			if otherTool != v.tool && validTool(grid[v.Y][v.X].typ, otherTool) {
				edges[v] = append(edges[v], vertex{v.Pos, otherTool})
			}
		}
		// add edges for moving to a new region
		for _, m := range v.Pos.ValidMoves(g) {
			if validTool(grid[m.Y][m.X].typ, v.tool) {
				edges[v] = append(edges[v], vertex{m, v.tool})
			}
		}
	}

	// initialize all distances
	dists := make(map[vertex]int)
	for _, v := range vertices {
		dists[v] = 1e9
	}
	start := vertex{utils.Pos{X: 0, Y: 0}, torch}
	dists[start] = 0

	// initialize min-heap
	indices := make(map[vertex]int, len(vertices))
	for i, v := range vertices {
		indices[v] = i
	}
	h := &distHeap{
		v:       vertices,
		indices: indices,
		dists:   dists,
	}
	heap.Init(h)

	// helper function for Djikstra
	edgeWeight := func(u, v vertex) int {
		if u.tool == v.tool {
			return 1
		}
		return 7
	}

	// Djikstra
	for h.Len() > 0 {
		u := heap.Pop(h).(vertex)
		for _, v := range edges[u] {
			d := dists[u] + edgeWeight(u, v)
			if d < dists[v] {
				dists[v] = d
				heap.Fix(h, h.indices[v])
			}
		}
	}
	targetVertex := vertex{target, torch}
	utils.Println(dists[targetVertex])
}
