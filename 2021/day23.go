package main

import (
	"container/heap"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2021, 23)

var (
	dests1 = []utils.Pos{
		{1, 1}, {2, 1}, {4, 1}, {6, 1}, {8, 1}, {10, 1}, {11, 1},
		{3, 2}, {5, 2}, {7, 2}, {9, 2},
		{3, 3}, {5, 3}, {7, 3}, {9, 3},
	}
	dests2 = append(dests1,
		utils.Pos{3, 4}, utils.Pos{5, 4}, utils.Pos{7, 4}, utils.Pos{9, 4},
		utils.Pos{3, 5}, utils.Pos{5, 5}, utils.Pos{7, 5}, utils.Pos{9, 5},
	)

	destCol  = [...]int{'A': 3, 'B': 5, 'C': 7, 'D': 9}
	moveCost = [...]int{'A': 1, 'B': 10, 'C': 100, 'D': 1000}
)

type move struct {
	grid [][]byte
	cost int
}

func (m move) dests() []utils.Pos {
	if len(m.grid) == 5 {
		return dests1
	}
	return dests2
}

// for memoization
func (m move) String() string {
	s := make([]byte, len(m.dests()))
	for i, p := range m.dests() {
		s[i] = m.grid[p.Y][p.X]
	}
	return string(s)
}

func (m move) distance(a, b utils.Pos) (int, bool) {
	lineIsClear := func(a, b utils.Pos) bool {
		for a != b {
			a.X += utils.Signum(b.X - a.X)
			a.Y += utils.Signum(b.Y - a.Y)
			if m.grid[a.Y][a.X] != '.' {
				return false
			}
		}
		return true
	}

	switch {
	case a.Y > b.Y:
		// move from room to hallway
		door := utils.Pos{a.X, b.Y}
		return a.Dist(b), lineIsClear(a, door) && lineIsClear(door, b)
	case a.Y < b.Y:
		// move from hallway to room
		door := utils.Pos{b.X, a.Y}
		return a.Dist(b), lineIsClear(a, door) && lineIsClear(door, b)
	case a.Y == b.Y:
		// move from room to room, via hallway
		hall := utils.Pos{a.X + utils.Signum(b.X-a.X), 1}
		d1, ok1 := m.distance(a, hall)
		d2, ok2 := m.distance(hall, b)
		return d1 + d2, ok1 && ok2
	default:
		panic("unreachable")
	}
}

func (m move) canMove(a, b utils.Pos) (int, bool) {
	c := m.grid[a.Y][a.X]
	switch {
	case false,
		a == b,                         // can't move to same spot
		m.grid[b.Y][b.X] != '.',        // or occupied spot
		(b.X == a.X && b.Y != a.Y),     // or purely vertically
		(b.Y > 1 && b.X != destCol[c]), // or a room not meant for us
		(a.Y == 1 && b.Y == 1):         // or from the hallway to the hallway
		return 0, false
	}
	if b.X == destCol[c] {
		// can't move to room if there are other types there
		for y := 2; y < len(m.grid)-1; y++ {
			if m.grid[y][b.X] != c && m.grid[y][b.X] != '.' {
				return 0, false
			}
		}
		// must move to the lowest open slot in the room
		for y := b.Y + 1; y < len(m.grid)-1; y++ {
			if m.grid[y][b.X] != c {
				return 0, false
			}
		}
	}
	return m.distance(a, b)
}

func (m move) validMoves() []*move {
	inFinalPosition := func(p utils.Pos) bool {
		c := m.grid[p.Y][p.X]
		if p.X != destCol[c] {
			return false
		}
		for y := p.Y; y < len(m.grid)-1; y++ {
			if m.grid[y][p.X] != c {
				return false
			}
		}
		return true
	}

	var moves []*move
	for _, p := range m.dests() {
		cost := moveCost[m.grid[p.Y][p.X]]
		if cost == 0 || inFinalPosition(p) {
			continue
		}
		for _, d := range m.dests() {
			dist, ok := m.canMove(p, d)
			if !ok {
				continue
			}
			newGrid := utils.CloneGrid(m.grid)
			newGrid[p.Y][p.X] = m.grid[d.Y][d.X]
			newGrid[d.Y][d.X] = m.grid[p.Y][p.X]
			moves = append(moves, &move{newGrid, m.cost + dist*cost})
		}
	}
	return moves
}

// priority queue
type moves []*move

func (ms moves) Len() int           { return len(ms) }
func (ms moves) Less(i, j int) bool { return ms[i].cost < ms[j].cost }
func (ms moves) Swap(i, j int)      { ms[i], ms[j] = ms[j], ms[i] }

func (ms *moves) Push(x interface{}) {
	*ms = append(*ms, x.(*move))
}

func (ms *moves) Pop() interface{} {
	old := *ms
	*ms = old[:len(old)-1]
	return old[len(old)-1]
}

func finished(grid [][]byte) bool {
	for y := 2; y < len(grid)-1; y++ {
		if grid[y][3] != 'A' ||
			grid[y][5] != 'B' ||
			grid[y][7] != 'C' ||
			grid[y][9] != 'D' {
			return false
		}
	}
	return true
}

func djikstra(grid [][]byte) int {
	var queue moves
	for _, m := range (&move{grid: grid}).validMoves() {
		heap.Push(&queue, m)
	}
	seen := make(map[string]bool)
	for len(queue) > 0 {
		m := heap.Pop(&queue).(*move)
		if seen[m.String()] {
			continue
		} else if finished(m.grid) {
			return m.cost
		}
		for _, next := range m.validMoves() {
			if !seen[next.String()] {
				heap.Push(&queue, next)
			}
		}
		seen[m.String()] = true
	}
	panic("unreachable")
}

func main() {
	grid := utils.ToByteGrid(utils.Lines(input))
	utils.Println(djikstra(grid))

	grid = append(grid, grid[3], grid[4])
	grid[3] = []byte("  #D#C#B#A#  ")
	grid[4] = []byte("  #D#B#A#C#  ")
	utils.Println(djikstra(grid))
}
