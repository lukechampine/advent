package main

import (
	"strings"

	"github.com/lukechampine/advent/utils"
)

const input = `..##.##.######...#.######
##...#...###....##.#.#.##
###.#.#.#..#.##.####.#.#.
..##.##...#..#.##.....##.
##.##...#.....#.#..#.####
.###...#.........###.####
#..##....###...#######..#
###..#.####.###.#.#......
.#....##..##...###..###.#
###.#..#.##.###.#..###...
####.#..##.#.#.#.#.#...##
##.#####.#......#.#.#.#.#
..##..####...#..#.#.####.
.####.####.####...##.#.##
#####....#...#.####.#..#.
.#..###..........#..#.#..
.#.##.#.#.##.##.#..#.#...
..##...#..#.....##.####..
..#.#...######..##..##.#.
.####.###....##...####.#.
.#####..#####....####.#..
###..#..##.#......##.###.
.########...#.#...###....
...##.#.##.#####.###.####
.....##.#.#....#..#....#.`

func parse(s string) [][]int {
	s = strings.Repeat(".........................\n", 500) + s + strings.Repeat(".........................\n", 500)
	var grid [][]int
	for _, line := range utils.Lines(s) {
		line = strings.Repeat(".........................", 500) + line + strings.Repeat(".........................", 500)
		row := make([]int, len(line))
		for i := range row {
			if line[i] == '#' {
				row[i] = infected
			}
		}
		grid = append(grid, row)
	}
	return grid
}

const (
	clean = iota
	weakened
	infected
	flagged
)

type virus struct {
	utils.Agent
	infected int
}

func (v *virus) burst(grid [][]int) {
	if grid[v.Y][v.X] == infected {
		grid[v.Y][v.X] = clean
		v.TurnRight()
	} else {
		grid[v.Y][v.X] = infected
		v.TurnLeft()
		v.infected++
	}
	v.MoveForwardArray(1)
}

func (v *virus) burst2(grid [][]int) {
	switch grid[v.Y][v.X] {
	case clean:
		grid[v.Y][v.X] = weakened
		v.TurnLeft()
	case weakened:
		grid[v.Y][v.X] = infected
		v.infected++
	case infected:
		grid[v.Y][v.X] = flagged
		v.TurnRight()
	case flagged:
		grid[v.Y][v.X] = clean
		v.TurnAround()
	}
	v.MoveForwardArray(1)
}

func main() {
	// part 1
	grid := parse(input)
	v := virus{
		Agent: utils.NewAgent(len(grid[0])/2, len(grid)/2, utils.Up),
	}
	for i := 0; i < 10000; i++ {
		v.burst(grid)
	}
	utils.Println(v.infected)

	// part 2
	grid = parse(input)
	v = virus{
		Agent: utils.NewAgent(len(grid[0])/2, len(grid)/2, utils.Up),
	}
	for i := 0; i < 10000000; i++ {
		v.burst2(grid)
	}
	utils.Println(v.infected)
}
