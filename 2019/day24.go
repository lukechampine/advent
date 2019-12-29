package main

import (
	"bytes"

	"github.com/lukechampine/advent/utils"
)

var input = utils.Lines(utils.ReadInput("day24_input.txt"))

func rules(x, y int, grid [][]byte) byte {
	g := utils.Grid{X: len(grid[0]) - 1, Y: len(grid) - 1}
	p := utils.Pos{X: x, Y: y}
	moves := p.ValidMoves(g)
	bugs := utils.Count(len(moves), func(i int) bool {
		return grid[moves[i].Y][moves[i].X] == '#'
	})
	if grid[y][x] == '#' && bugs != 1 {
		return '.'
	} else if grid[y][x] == '.' && (bugs == 1 || bugs == 2) {
		return '#'
	}
	return grid[y][x]
}

func step(grid [][]byte) [][]byte {
	next := make([][]byte, len(grid))
	for i := range next {
		next[i] = make([]byte, len(grid[i]))
	}
	for y := range next {
		for x := range next[y] {
			next[y][x] = rules(x, y, grid)
		}
	}
	return next
}

func biodiversity(grid [][]byte) (diversity int) {
	pow := 1
	for y := range grid {
		for _, c := range grid[y] {
			if c == '#' {
				diversity += pow
			}
			pow *= 2
		}
	}
	return
}

func copyGrid(grid [][]byte) [][]byte {
	g := make([][]byte, len(grid))
	for i := range g {
		g[i] = append([]byte(nil), grid[i]...)
	}
	return g
}

type nestedGrid struct {
	grid  [][]byte
	above *nestedGrid
	below *nestedGrid
	level int
}

func (ng *nestedGrid) rulesAbove(x, y int, above [][]byte) byte {
	g := utils.Grid{X: 4, Y: 4}
	p := utils.Pos{X: x, Y: y}
	moves := p.ValidMoves(g)
	bugs := utils.Count(len(moves), func(i int) bool {
		return ng.grid[moves[i].Y][moves[i].X] == '#'
	})
	if y == 0 && above[1][2] == '#' {
		bugs++
	}
	if y == 4 && above[3][2] == '#' {
		bugs++
	}
	if x == 0 && above[2][1] == '#' {
		bugs++
	}
	if x == 4 && above[2][3] == '#' {
		bugs++
	}
	if ng.grid[y][x] == '#' && bugs != 1 {
		return '.'
	} else if ng.grid[y][x] == '.' && (bugs == 1 || bugs == 2) {
		return '#'
	}
	return ng.grid[y][x]
}

func (ng *nestedGrid) rulesBelow(x, y int, below [][]byte) byte {
	p := utils.Pos{X: x, Y: y}
	moves := p.Moves()
	bugs := utils.Count(len(moves), func(i int) bool {
		if moves[i] == (utils.Pos{X: 2, Y: 2}) {
			return false
		}
		return ng.grid[moves[i].Y][moves[i].X] == '#'
	})
	if y == 2 && x == 1 {
		bugs += utils.Count(5, func(y int) bool {
			return below[y][0] == '#'
		})
	}
	if y == 2 && x == 3 {
		bugs += utils.Count(5, func(y int) bool {
			return below[y][4] == '#'
		})
	}
	if x == 2 && y == 1 {
		bugs += utils.Count(5, func(x int) bool {
			return below[0][x] == '#'
		})
	}
	if x == 2 && y == 3 {
		bugs += utils.Count(5, func(x int) bool {
			return below[4][x] == '#'
		})
	}
	if ng.grid[y][x] == '#' && bugs != 1 {
		return '.'
	} else if ng.grid[y][x] == '.' && (bugs == 1 || bugs == 2) {
		return '#'
	}
	return ng.grid[y][x]
}

func (ng *nestedGrid) step() {
	var above, below [][]byte

	if ng.above == nil {
		above = utils.ByteGrid(len(ng.grid[0]), len(ng.grid), '.')
		if utils.CountGrid(ng.grid, '#') > 0 {
			ng.above = &nestedGrid{
				grid:  above,
				below: ng,
				level: ng.level - 1,
			}
			ng.above.step()
			return
		}
	} else {
		above = ng.above.grid
	}

	if ng.below == nil {
		below = utils.ByteGrid(len(ng.grid[0]), len(ng.grid), '.')
		if utils.CountGrid(ng.grid, '#') > 0 {
			ng.below = &nestedGrid{
				grid:  below,
				above: ng,
				level: ng.level + 1,
			}
			ng.below.step()
		}
	} else {
		below = copyGrid(ng.below.grid)
		ng.below.step()
	}

	next := copyGrid(ng.grid)
	for y := range next {
		for x := range next[y] {
			if x == 2 && y == 2 {
				continue
			}
			if len((utils.Pos{X: x, Y: y}).ValidMoves(utils.Grid{X: 4, Y: 4})) < 4 {
				next[y][x] = ng.rulesAbove(x, y, above)
			} else if (utils.Pos{X: x, Y: y}).Dist(utils.Pos{X: 2, Y: 2}) == 1 {
				next[y][x] = ng.rulesBelow(x, y, below)
			} else {
				next[y][x] = rules(x, y, ng.grid)
			}
		}
	}
	ng.grid = next

}

func main() {
	// part 1
	grid := utils.ToByteGrid(input)
	states := make(map[string]struct{})
	for {
		grid = step(grid)
		s := string(bytes.Join(grid, nil))
		if _, ok := states[s]; ok {
			break
		}
		states[s] = struct{}{}
	}
	utils.Println(biodiversity(grid))

	// part 2
	ngrid := &nestedGrid{
		grid: utils.ToByteGrid(input),
	}
	for minutes := 0; minutes < 200; minutes++ {
		ngrid.step()
		// zoom out
		for ngrid.above != nil {
			ngrid = ngrid.above
		}
	}
	var bugs int
	for p := ngrid; p != nil; p = p.below {
		bugs += utils.CountGrid(p.grid, '#')
	}
	utils.Println(bugs)
}
