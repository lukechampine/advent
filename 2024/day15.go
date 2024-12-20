package main

import (
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2024, 15)

func main() {
	groups := utils.Split(input, "\n\n")
	grid := utils.ToByteGrid(utils.Lines(groups[0]))
	pos := utils.Locate(grid, '@')
	grid[pos.Y][pos.X] = '.'
	a := utils.NewAgent(pos.X, pos.Y, utils.Up)
	moves := strings.TrimSpace(strings.Join(utils.Lines(groups[1]), ""))
	for i := range moves {
		a.Dir = utils.DirFromArrow(moves[i])
		a.MoveForwardArray(1)

		if grid[a.Y][a.X] == '#' {
			a.MoveForwardArray(-1)
			continue
		} else if grid[a.Y][a.X] == 'O' {
			// attempt to push
			p := a.Pos
			for grid[a.Y][a.X] == 'O' {
				a.MoveForwardArray(1)
			}
			if grid[a.Y][a.X] == '#' {
				// push failed
				a.Pos = p
				a.MoveForwardArray(-1)
				continue
			} else {
				// push succeeded
				grid[p.Y][p.X] = '.'
				grid[a.Y][a.X] = 'O'
				a.Pos = p
			}
		}
	}

	var sum int
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == 'O' {
				sum += x + 100*y
			}
		}
	}
	utils.Println(sum)

	grid = nil
	for _, row := range utils.ToByteGrid(utils.Lines(groups[0])) {
		var row2 []byte
		for _, c := range row {
			switch c {
			case '@':
				row2 = append(row2, '@', '.')
			case 'O':
				row2 = append(row2, '[', ']')
			default:
				row2 = append(row2, c, c)
			}
		}
		grid = append(grid, row2)
	}
	pos = utils.Locate(grid, '@')
	grid[pos.Y][pos.X] = '.'
	a = utils.NewAgent(pos.X, pos.Y, utils.Up)
	pushed := func(p utils.Pos, d utils.Dir) (map[utils.Pos]byte, bool) {
		ps := make(map[utils.Pos]byte)
		queue := []utils.Pos{p}
		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]
			if _, ok := ps[p]; ok {
				continue
			}
			switch grid[p.Y][p.X] {
			case '#':
				return nil, false
			case '.':
				continue
			case '[':
				ps[p] = '['
				ps[p.Add(1, 0)] = ']'
				queue = append(queue, p.MoveArray(d, 1), p.Add(1, 0).MoveArray(d, 1))
			case ']':
				ps[p.Add(-1, 0)] = '['
				ps[p] = ']'
				queue = append(queue, p.MoveArray(d, 1), p.Add(-1, 0).MoveArray(d, 1))
			}
		}
		return ps, true
	}

	for i := range moves {
		a.Dir = utils.DirFromArrow(moves[i])
		a.MoveForwardArray(1)

		if grid[a.Y][a.X] == '#' {
			a.MoveForwardArray(-1)
		} else if grid[a.Y][a.X] == '[' || grid[a.Y][a.X] == ']' {
			if ps, ok := pushed(a.Pos, a.Dir); ok {
				for p := range ps {
					grid[p.Y][p.X] = '.'
				}
				for p, c := range ps {
					p = p.MoveArray(a.Dir, 1)
					grid[p.Y][p.X] = c
				}
			} else {
				a.MoveForwardArray(-1)
			}
		}
	}

	sum = 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == '[' {
				sum += x + 100*y
			}
		}
	}
	utils.Println(sum)
}
