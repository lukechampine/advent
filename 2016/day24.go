package main

import (
	"strconv"

	"github.com/lukechampine/advent/utils"
)

const input = `#########################################################################################################################################################################################
#.#.......#.#.#.....#.#.......#.................#.......#.#.....#.....#...#...#.......#...#...........#.#.....#.............#.........#.............#.........#.....#.#.............#...#
#.#####.#.#.#.#.#.#.#.#.#.###.#.#.###.#.#.###.###.#.#.#.#.#.#####.#.#.###.#.#.###.#.#.###.###.#.###.###.###.#.#.###.#.#.#.#.#.#.###.#.###.#.#.#.#.#.#.#.#.#.#.#.#.#.#.#.###.#####.#.###.#
#.....#.............#...#....6#.....#.....#.#...#.....#...#.........#.......#...#.#.....#.....#.#...#...#.....#.#.......#.........#...#...#.#.#.......#.........#.....#...#.#.#.#.....#.#
#####.#.###.###.#.#.#.#####.#.#####.#.#####.#.#######.#.#.#.#.#.#####.#####.#.#########.#.###.#.#.#.#.#.#.#.#######.#.#.#####.###.###.###.#.#.###.#.#.#.#.#.#.#.#.###########.#.#.#####.#
#.....#.......#.#.....#.......#.#...#.#.#.#...........#.....#.#.#...#.#...#.#.#.....#...#.....#.....#...............#.......#.#.........#.#...........#.....#...#.......#.#...#.#.....#.#
###.#.#.#.#####.#####.#.#.#.#.#.#.#.#.#.#.#########.#.#########.#.#.###.###.###.###.#.#.#.#.#.#.#.#######.#.#.###.#.###.#.#.#.#.#.#.#######.#.#.#.###.#.#.#.###.###.#.#.#.#.#.###.#.#####
#...#...#.#.....#.............#...#...#...#.#.......#.#.#...#...#.#.#...#.....#.#.#.....#.#...#.#.......#.#.#...#.......#.....#.....#.......#...#...#.#...#.....#.#.......#.#.#...#...#.#
#.#.#####.#.###.###.#.#########.#.#.#.###.#.#.#.#.#.#.#.#.#.#.#.#.#.###.#.#.#.###.#.#.#.#.#######.#.#.###.#.###.###.#.#.#.#####.#####.#####.#.#.#.#.#.###.#.#.###.#.#####.#.#####.#.###.#
#...#.#...#.....#.#.........#.#.....#...#.....#...#...#...#...#...#.........#.......#...#.#.....#...#.#...#.#.......#.........#...#.#...#...#2..#.#.#.#...#.......#...#.#...........#...#
#.###.#####.#####.#.#.###.###.#.###.#.#.#.#.###.#.#.###.#.#.#.###.#.#.#####.#.#####.#.###.#.#.#.#.###.#.#.#.###.#.#.#.#####.#.#.#.#.#.###.#####.#.#.#.#.#.#####.###.#.#.#.#.#.#.#####.#.#
#...#.#.#4#.#.........#.......#...........#.#...#...#...#.........#.......#.....#...#...#.#...#...#...#.#...#.........#.....#...#.....#.........#...#...............#.#.#...#.........#.#
#######.#.#.###.###.###.###.#.#####.#####.#.#.###.#.#.#########.#####.#.#####.#.#.#.###.#.#.###.#.#.#.#.#.#.#.#####.#.#####.#.#.#.#.#.#.#.#.#.###.#.#######.#####.#.#.#.#.#.#.###.#.#####
#.........#.#.....#.....#...................#.....#...#.........#.....#.#.....#.#...........#...#...#.........#...#...#.....#...#.#.#.#.....#.......#.#3....#.......#.....#.#.#.....#...#
#.#####.###.#.###.#####.#.###.###.###.#.#####.#.#.###.###.#.###.#.#.#.###.#.###.#.###.#.#####.#.###.###.#######.#.#.#.#.###.#.#.#.#########.#.###.###.###.#.#.###.#.#######.###.#.#.#.#.#
#...#...#.........#...#.......#...#...#...#.........#.....#.....#.....#...#.......#...#...#...#.#.#.#...#.................#.............................#...#.....#.....#.#.....#.....#.#
#.#.#.#.#####.###.#.#.#.#.###.###.###.#.###.#.#.#####.###.#.###.#.#.###.#.#.#####.###.#.#.###.#.#.#.###.#.#.###.#####.#.###.###.#.#.#.#.#.#.###.#.#.#.#.#.###.###.#.#.#.#.#####.#####.#.#
#...#.......#...#.....#...#.#.#.#.#...#.#...#.....#.#.#.....#.....#.......#.#.......#.........#.....#...#...#...#.#.........#.#.#.....#.#...#...#.....#...#.......#.........#.....#...#.#
###.#.#.#.#.#####.###.#####.#.#.#.#.###.#####.###.#.#.#.#####.#.#.#.#.#.#.###.#.#.###.#.#.###.#.#.###.#.#.###.###.#.#.#.#.#.#.#.#.#####.#.#.#.###.#.#.#.###.###.###.###.###.#.#.#####.#.#
#7#.......#...#...#...#...#.......#.#.........#.#.......#.....#...#.....#.#...#...#...#.#.#...#.#.......#.....#.#.#.....#...#...#...#.....#.....#.#.#.#.#...#...#...#.#.......#.....#...#
#.#.#.#.#.#.#######.#.#.#.#####.#.###.###.#.###.#.###.#.###.#.#######.###.#.#.#.#.#####.###.###.#####.#.#####.#.###.#####.#.#.#.###.#.#######.#.#.###.#.###.#####.#.#.#.#.###.###.###.#.#
#.......#...#.#.....#.#.....#.#...#.......#...#.......#.#.....#.........#.#.#.....#.........#.#.#...............#...#.....#...#.#...#.......#.#...........#...#.#.#.#.#.#.....#...#.....#
#.#.#.#######.#.###.#.#.###.#.#.#.#.#####.#.#.#.#.#.#.#.#.#.#.#.#######.#####.###.#.#####.###.#####.#.#.###.#.###.#.#.###.#.###.#.#.#.#.#.###.###.#.###.#.#.#.#.#.#.###.#.#.#.#.#.#.#.#.#
#.....#.#...#.#.....#...........#...#.#.....#.#.......#.....#.#...#...........#.....#.........#...#.#.#.#.#.#.....#...#.#.............#.....#.#.#.#.#.#.#.#.#.....#.#.#.#...........#...#
#####.#.###.#.#.#.#.#.###.#.#.#.###.#.#.#######.#####.###.#.#.#.#.#.#.#.#.#.###.#.#####.#.#######.#.#####.#.#.#########.#########.#.#.###.###.#.#####.#.###.###.#.#.#.#.#.#.#.#######.#.#
#.....#...#.#.......#...........#...#.#.....#.#.......#.#...#.#...#.......#...#.......#...#.........#...#.........................#...#...#...#.#.#...#.#.........#...#...#...#...#...#.#
#.#.#.#.#.#.#####.#.#.#########.#.#####.#.#.#.#.#.###.#.#.#.#.###.#.#.#.#.#.#.###.#.#######.###.#.#.#.#######.#.###.#####.###.#.#.#####.#.#.#####.#######.#.#.###.###.#.#.#.#.#.###.###.#
#.#...#...........#.#.......#...#.#.............#.......#...#.....#.#.....#...#.....#.......#...#.#.#...#...#.#...#.#.....#...#...#...#.....#.....#...#.....#.......#.....#.#.#.#0#.....#
###.#.#.#.#.#.#######.#.###.###.#.#.###.#.#######.#.###.#########.#.#.#.#.###.#.###.###########.#.###.###.#.#.###.#.#.###.#.#.###.###.#.###.#.###########.###.#.###.#####.#.#.#.#.#.#.#.#
#.#...#.#.#...#.........#...#.......#.#.................#.#...........#.....#.......#.....#.#.#.#.#...#.....#...#...#.#.........#.#.....#.#...#.....#.....#.#...#.#...#.......#.#.....#.#
#.###.#.#.#.#.#.#.#.#.#.#.#.#.###.#.#.#.#.#####.#.#####.#.###.#.#.###.#.###.###.###.#.#####.#.#.#######.#.#####.###.###.#######.###.#####.#.#.#.#.#.#.#.#.#.###.#.#.#.#.#.###.#.#.#.#.#.#
#.#...#.........#.#...#...........#.#.....#5........#...#.....#.......#.#...#.......#.........#.......#.........#.#.#.......#...#.#...#.#.#.....#...#.#.#.#.......#.#.....#...#.#.....#.#
###.###.#.#####.###.###.###.#.#.###.#.#.#####.###.#.#.#.#.###.#####.#.#.#.#.#.#.#.###.#.###.#.###.#.#.#.#.###.#.#.#.#######.#.###.#.#.#.#.###.###.#.#.#.#.#.#######.#.#.#.#####.###.#.#.#
#...#...........#...#...#.#.......#.....#.#.#.#...............#.....#.....#...#.#.#...#.#...#.#.....#...#.#...#...#.#.#.......#...#.#...#.......#...#.......#.......#.#...#.#.....#.....#
#.#.#.###.#.#.#######.###.#####.#.#######.#.###.#######.#.###.###.#.#.#####.#.###.#.#.#.#.###.###.#.#.#.###.#.#.###.#.###.#########.###.#.#.#####.#.#####.###.#.#.#.###.#.#.#####.#######
#...#.................#.....#.....#...#...#.#...#...#.....#.......#.#...#...#.....#.....#.........#...#.....#.........#...#.....#...#.......#1..#.#.#.#...#...#.#.......#.#...#.......#.#
#.#.#.#####.#.#.###.###.###.#####.#.#.#####.#.#.#.#####.#.###.###.#.###.#.#.#.###.#.#.#.#.#.#.#.#.#.#.#.###.#######.#.#.#####.###.#######.#.#.#.#.#.#.#.###.#.#.#.#####.#.#######.#.#.#.#
#.....#.....#.#...............#.#...#.....#.................#...#.#...#.#...#.....#.#...#.........#...........#.#.......#.....#.#.....#.....#.......#...........#...#.....#.#.......#.#.#
#.#.#.#.#.#.#.#######.###.###.#.#.#.#.#.#.###.###.#.#.###.#.#.#.#.#.#.#####.#.#.#.#.###.#.#.#######.#.#.#.###.#.#.#.#.###.#.#.#.#.#.#.#.#####.#.#.#.#.#####.#####.###.###.#.#.#.#.#.#.#.#
#...#.#...#.#.......#.......#.....#...#.......#...#...#...#.....#.#.....#...#.#...#.....#.#.........#.#.#.....#.....#...#.........#.#.#.......#.........#...#.....#.#...#.#...#.....#...#
#########################################################################################################################################################################################`

type maze struct {
	utils.Maze
	locs []utils.Pos
}

func parse(str string) maze {
	lines := utils.Lines(str)
	var m maze
	m.X = len(lines[0])
	m.Y = len(lines)
	m.IsWall = func(p utils.Pos) bool {
		return lines[p.Y][p.X] == '#'
	}
	// search for each location
	locs := make(map[int]utils.Pos)
	for y, line := range lines {
		for x, c := range line {
			if n, err := strconv.Atoi(string(c)); err == nil {
				locs[n] = utils.Pos{x, y}
			}
		}
	}
	for i := 0; i < len(locs); i++ {
		m.locs = append(m.locs, locs[i])
	}
	return m
}

func locDists(m maze) [][]int {
	dists := make([][]int, len(m.locs))
	for i := range dists {
		dists[i] = make([]int, len(m.locs))
	}
	for i, a := range m.locs {
		dist := m.DistancesFrom(a)
		for j, b := range m.locs {
			dists[i][j] = dist[b]
		}
	}
	return dists
}

func shortestPath(m maze, dists [][]int) int {
	perms := utils.Perms(len(m.locs))
	var min int = 100000000
	for _, perm := range perms {
		if perm[0] != 0 {
			continue
		}
		dist := 0
		for i := 0; i < len(perm)-1; i++ {
			dist += dists[perm[i]][perm[i+1]]
		}
		min = utils.Min(min, dist)
	}
	return min
}

func shortestReturnPath(m maze, dists [][]int) int {
	perms := utils.Perms(len(m.locs))
	var min int = 100000000
	for _, perm := range perms {
		if perm[0] != 0 {
			continue
		}
		perm = append(perm, 0) // return to 0
		dist := 0
		for i := 0; i < len(perm)-1; i++ {
			dist += dists[perm[i]][perm[i+1]]
		}
		min = utils.Min(min, dist)
	}
	return min
}

func main() {
	// part 1
	maze := parse(input)
	dists := locDists(maze)
	println(shortestPath(maze, dists))

	// part 2
	println(shortestReturnPath(maze, dists))
}
