package main

import (
	"strings"

	"github.com/lukechampine/advent/utils"
)

const input = `../.. => ..#/#../.#.
#./.. => #../#../...
##/.. => ###/#.#/#..
.#/#. => ###/##./.#.
##/#. => .../.#./..#
##/## => ##./#.#/###
.../.../... => ##../.#../#.#./....
#../.../... => ..../##.#/...#/##.#
.#./.../... => ###./####/#.../#..#
##./.../... => ###./.##./...#/..##
#.#/.../... => .###/.##./#.../#.##
###/.../... => ##.#/#..#/#.#./#.##
.#./#../... => #.#./.###/#.../#.##
##./#../... => #.../####/#.##/....
..#/#../... => #.##/..#./...#/...#
#.#/#../... => #.##/####/.#.#/#.#.
.##/#../... => #.../##../##.#/.##.
###/#../... => ..../#.#./.###/#...
.../.#./... => .#.#/#..#/##../#.##
#../.#./... => ###./.###/.#.#/..#.
.#./.#./... => ..##/.##./..##/.#.#
##./.#./... => ..#./##../###./...#
#.#/.#./... => ..##/.##./.###/###.
###/.#./... => ..#./.###/###./#.##
.#./##./... => ###./..../.#../#...
##./##./... => .#.#/##../##.#/...#
..#/##./... => ##.#/.##./.###/..##
#.#/##./... => .###/..#./#.##/####
.##/##./... => ##.#/..#./..##/###.
###/##./... => ..../.#.#/.#../#...
.../#.#/... => ###./.#.#/.#../#.##
#../#.#/... => ####/#..#/..../....
.#./#.#/... => #.../..##/#.##/#.#.
##./#.#/... => #.#./###./##../#.#.
#.#/#.#/... => ...#/.##./.##./.#..
###/#.#/... => ..../.##./####/#.#.
.../###/... => .###/.#../.###/#.##
#../###/... => ..##/..##/.##./##..
.#./###/... => .#.#/..#./..##/##.#
##./###/... => ...#/#.##/#.#./##.#
#.#/###/... => #.##/.##./...#/###.
###/###/... => ##../...#/..##/####
..#/.../#.. => #.##/#.../.#../#.#.
#.#/.../#.. => .##./.##./.#.#/.##.
.##/.../#.. => .#.#/#.##/...#/##.#
###/.../#.. => ##../..#./...#/##..
.##/#../#.. => ##../..##/#..#/#..#
###/#../#.. => ##../..#./#.#./....
..#/.#./#.. => .##./##.#/##../####
#.#/.#./#.. => ####/...#/.#.#/..#.
.##/.#./#.. => .#.#/..#./##.#/.#..
###/.#./#.. => #.../#.##/..../##.#
.##/##./#.. => #.#./#.#./#.##/#.#.
###/##./#.. => ...#/###./.##./.#.#
#../..#/#.. => ####/####/..../.##.
.#./..#/#.. => #.##/...#/..#./####
##./..#/#.. => ..#./#.../..##/####
#.#/..#/#.. => #.../#.##/#.##/..##
.##/..#/#.. => ####/..../##../####
###/..#/#.. => ..../##.#/.##./####
#../#.#/#.. => ...#/..##/###./#..#
.#./#.#/#.. => #..#/..#./.###/##.#
##./#.#/#.. => ###./####/#.##/..#.
..#/#.#/#.. => ##../##.#/..##/.##.
#.#/#.#/#.. => .#.#/.##./#.../##.#
.##/#.#/#.. => .#.#/#..#/.##./..#.
###/#.#/#.. => ...#/.#../.##./##.#
#../.##/#.. => ###./##../#.#./####
.#./.##/#.. => .#../##../#.#./.#.#
##./.##/#.. => ##.#/.#../.#.#/####
#.#/.##/#.. => ####/.#.#/..../....
.##/.##/#.. => ####/##../#..#/####
###/.##/#.. => .###/##.#/.#../#.##
#../###/#.. => #..#/###./####/.#.#
.#./###/#.. => ..##/##../##.#/.#.#
##./###/#.. => #..#/.#../####/...#
..#/###/#.. => ##../##.#/...#/#..#
#.#/###/#.. => ..#./.##./#..#/....
.##/###/#.. => #..#/#.../..../.#..
###/###/#.. => ..#./#.##/.##./#...
.#./#.#/.#. => .#.#/.##./##.#/.##.
##./#.#/.#. => #..#/.###/.#.#/.##.
#.#/#.#/.#. => #.../##../#.../.###
###/#.#/.#. => ###./.###/###./....
.#./###/.#. => .#../####/...#/##..
##./###/.#. => ####/###./..../....
#.#/###/.#. => ...#/.###/..../####
###/###/.#. => ..../#.../..#./.###
#.#/..#/##. => #.#./#.../####/#.##
###/..#/##. => .#.#/#..#/.###/#...
.##/#.#/##. => ..##/..#./..../##..
###/#.#/##. => #.#./##.#/####/#..#
#.#/.##/##. => ..../.#../#.#./##.#
###/.##/##. => ..../..../.#../##.#
.##/###/##. => #.#./.###/#.#./#.##
###/###/##. => ##.#/##.#/.###/..#.
#.#/.../#.# => #..#/.#../#.../...#
###/.../#.# => ##../.#../##.#/..#.
###/#../#.# => ..##/#.#./####/.#..
#.#/.#./#.# => ...#/...#/#..#/#.#.
###/.#./#.# => ..../####/.##./.#.#
###/##./#.# => #..#/.#.#/..##/####
#.#/#.#/#.# => #.#./..#./...#/.#..
###/#.#/#.# => ...#/##.#/.###/.#..
#.#/###/#.# => .#.#/###./.#../.##.
###/###/#.# => ...#/.###/.#.#/###.
###/#.#/### => #.##/.#.#/...#/.#..
###/###/### => ..##/.#../#.#./.#..`

func parse2x2(s string) [2][2]bool {
	var sq [2][2]bool
	for i, line := range strings.Split(s, "/") {
		for j, c := range line {
			sq[i][j] = c == '#'
		}
	}
	return sq
}

func parse3x3(s string) [3][3]bool {
	var sq [3][3]bool
	for i, line := range strings.Split(s, "/") {
		for j, c := range line {
			sq[i][j] = c == '#'
		}
	}
	return sq
}

func parse4x4(s string) [4][4]bool {
	var sq [4][4]bool
	for i, line := range strings.Split(s, "/") {
		for j, c := range line {
			sq[i][j] = c == '#'
		}
	}
	return sq
}

func rotateAndFlip2x2(sq [2][2]bool) [][2][2]bool {
	return [][2][2]bool{
		// 0 degrees
		{{sq[0][0], sq[0][1]}, {sq[1][0], sq[1][1]}},
		{{sq[0][1], sq[0][0]}, {sq[1][1], sq[1][0]}}, // flip horizontal
		{{sq[1][0], sq[1][1]}, {sq[0][0], sq[0][1]}}, // flip vertical
		// 90 degrees
		{{sq[1][0], sq[0][0]}, {sq[1][1], sq[0][1]}},
		{{sq[0][0], sq[1][0]}, {sq[0][1], sq[1][1]}}, // flip horizontal
		{{sq[1][1], sq[0][1]}, {sq[1][0], sq[0][0]}}, // flip vertical
		// 180 degrees
		{{sq[1][1], sq[1][0]}, {sq[0][1], sq[0][0]}},
		{{sq[1][0], sq[1][1]}, {sq[0][0], sq[0][1]}}, // flip horizontal
		{{sq[0][1], sq[0][0]}, {sq[1][1], sq[1][0]}}, // flip vertical
		// 270 degrees + flip
		{{sq[0][1], sq[1][1]}, {sq[0][0], sq[1][0]}},
		{{sq[1][1], sq[0][1]}, {sq[1][0], sq[0][0]}}, // flip horizontal
		{{sq[0][0], sq[1][0]}, {sq[0][1], sq[1][1]}}, // flip vertical
	}
}

func rotateAndFlip3x3(sq [3][3]bool) [][3][3]bool {
	return [][3][3]bool{
		// 0 degrees
		{{sq[0][0], sq[0][1], sq[0][2]}, {sq[1][0], sq[1][1], sq[1][2]}, {sq[2][0], sq[2][1], sq[2][2]}},
		{{sq[0][2], sq[0][1], sq[0][0]}, {sq[1][2], sq[1][1], sq[1][0]}, {sq[2][2], sq[2][1], sq[2][0]}},
		{{sq[2][0], sq[2][1], sq[2][2]}, {sq[1][0], sq[1][1], sq[1][2]}, {sq[0][0], sq[0][1], sq[0][2]}},
		// 90 degrees
		{{sq[2][0], sq[1][0], sq[0][0]}, {sq[2][1], sq[1][1], sq[0][1]}, {sq[2][2], sq[1][2], sq[0][2]}},
		{{sq[0][0], sq[1][0], sq[2][0]}, {sq[0][1], sq[1][1], sq[2][1]}, {sq[0][2], sq[1][2], sq[2][2]}},
		{{sq[2][2], sq[1][2], sq[0][2]}, {sq[2][1], sq[1][1], sq[0][1]}, {sq[2][0], sq[1][0], sq[0][0]}},
		// 180 degrees
		{{sq[2][2], sq[2][1], sq[2][0]}, {sq[1][2], sq[1][1], sq[1][0]}, {sq[0][2], sq[0][1], sq[0][0]}},
		{{sq[2][0], sq[2][1], sq[2][2]}, {sq[1][0], sq[1][1], sq[1][2]}, {sq[0][0], sq[0][1], sq[0][2]}},
		{{sq[0][2], sq[0][1], sq[0][0]}, {sq[1][2], sq[1][1], sq[1][0]}, {sq[2][2], sq[2][1], sq[2][0]}},
		// 270 degrees
		{{sq[0][2], sq[1][2], sq[2][2]}, {sq[0][1], sq[1][1], sq[2][1]}, {sq[0][0], sq[1][0], sq[2][0]}},
		{{sq[2][2], sq[1][2], sq[0][2]}, {sq[2][1], sq[1][1], sq[0][1]}, {sq[2][0], sq[1][0], sq[0][0]}},
		{{sq[0][0], sq[1][0], sq[2][0]}, {sq[0][1], sq[1][1], sq[2][1]}, {sq[0][2], sq[1][2], sq[2][2]}},
	}
}

func parse(s string) (map[[2][2]bool][3][3]bool, map[[3][3]bool][4][4]bool) {
	enhance2x2 := make(map[[2][2]bool][3][3]bool)
	enhance3x3 := make(map[[3][3]bool][4][4]bool)
	for _, line := range utils.Lines(s) {
		fs := strings.Fields(line)
		left, right := fs[0], fs[2]
		if len(line) == 20 {
			ruleIn, ruleOut := parse2x2(left), parse3x3(right)
			enhance2x2[ruleIn] = ruleOut
			for _, variant := range rotateAndFlip2x2(ruleIn) {
				enhance2x2[variant] = ruleOut
			}
		} else {
			ruleIn, ruleOut := parse3x3(left), parse4x4(right)
			enhance3x3[ruleIn] = ruleOut
			for _, variant := range rotateAndFlip3x3(ruleIn) {
				enhance3x3[variant] = ruleOut
			}
		}
	}

	return enhance2x2, enhance3x3
}

func enhance2x2(grid [][]bool, enhance map[[2][2]bool][3][3]bool) [][]bool {
	newgrid := make([][]bool, (len(grid)/2)*3)
	for i := range newgrid {
		newgrid[i] = make([]bool, (len(grid)/2)*3)
	}
	for i := 0; i < len(grid); i += 2 {
		for j := 0; j < len(grid); j += 2 {
			sq := [2][2]bool{
				{grid[i+0][j+0], grid[i+0][j+1]},
				{grid[i+1][j+0], grid[i+1][j+1]},
			}
			ni, nj := (i/2)*3, (j/2)*3
			for k, row := range enhance[sq] {
				copy(newgrid[ni+k][nj:], row[:])
			}
		}
	}
	return newgrid
}

func enhance3x3(grid [][]bool, enhance map[[3][3]bool][4][4]bool) [][]bool {
	newgrid := make([][]bool, (len(grid)/3)*4)
	for i := range newgrid {
		newgrid[i] = make([]bool, (len(grid)/3)*4)
	}
	for i := 0; i < len(grid); i += 3 {
		for j := 0; j < len(grid); j += 3 {
			sq := [3][3]bool{
				{grid[i+0][j+0], grid[i+0][j+1], grid[i+0][j+2]},
				{grid[i+1][j+0], grid[i+1][j+1], grid[i+1][j+2]},
				{grid[i+2][j+0], grid[i+2][j+1], grid[i+2][j+2]},
			}
			ni, nj := (i/3)*4, (j/3)*4
			for k, row := range enhance[sq] {
				copy(newgrid[ni+k][nj:], row[:])
			}
		}
	}
	return newgrid
}

func main() {
	// part 1
	e2, e3 := parse(input)

	grid := [][]bool{
		{false, true, false},
		{false, false, true},
		{true, true, true},
	}
	for i := 0; i < 5; i++ {
		if len(grid)%2 == 0 {
			grid = enhance2x2(grid, e2)
		} else if len(grid)%3 == 0 {
			grid = enhance3x3(grid, e3)
		} else {
			panic(len(grid))
		}
	}

	on := 0
	for _, row := range grid {
		for _, p := range row {
			if p {
				on++
			}
		}
	}
	utils.Println(on)

	// part 2
	grid = [][]bool{
		{false, true, false},
		{false, false, true},
		{true, true, true},
	}
	for i := 0; i < 18; i++ {
		if len(grid)%2 == 0 {
			grid = enhance2x2(grid, e2)
		} else if len(grid)%3 == 0 {
			grid = enhance3x3(grid, e3)
		} else {
			panic(len(grid))
		}
	}

	on = 0
	for _, row := range grid {
		for _, p := range row {
			if p {
				on++
			}
		}
	}
	utils.Println(on)
}
