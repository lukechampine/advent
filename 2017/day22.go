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
	up = iota
	right
	down
	left
)

const (
	clean = iota
	weakened
	infected
	flagged
)

type virus struct {
	x, y     int
	dir      int
	infected int
}

func (v *virus) burst(grid [][]int) {
	if grid[v.y][v.x] == infected {
		v.dir = (v.dir + 1) % 4 // turn right
		grid[v.y][v.x] = clean
	} else {
		v.dir = (v.dir + 3) % 4 // turn left
		grid[v.y][v.x] = infected
		v.infected++
	}

	switch v.dir {
	case up:
		v.y--
	case right:
		v.x++
	case down:
		v.y++
	case left:
		v.x--
	}
}

func (v *virus) burst2(grid [][]int) {
	switch grid[v.y][v.x] {
	case clean:
		grid[v.y][v.x] = weakened
		v.dir = (v.dir + 3) % 4 // turn left
	case weakened:
		grid[v.y][v.x] = infected
		v.infected++
	case infected:
		grid[v.y][v.x] = flagged
		v.dir = (v.dir + 1) % 4 // turn right
	case flagged:
		grid[v.y][v.x] = clean
		v.dir = (v.dir + 2) % 4 // turn around
	}

	switch v.dir {
	case up:
		v.y--
	case right:
		v.x++
	case down:
		v.y++
	case left:
		v.x--
	}
}

func main() {
	// part 1
	grid := parse(input)
	v := virus{
		x:   len(grid[0]) / 2,
		y:   len(grid) / 2,
		dir: up,
	}

	for i := 0; i < 10000; i++ {
		v.burst(grid)
	}
	utils.Println(v.infected)

	// part 2
	grid = parse(input)
	v = virus{
		x:   len(grid[0]) / 2,
		y:   len(grid) / 2,
		dir: up,
	}

	for i := 0; i < 10000000; i++ {
		v.burst2(grid)
	}
	utils.Println(v.infected)
}
