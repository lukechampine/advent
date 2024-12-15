package main

import (
	"bytes"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2024, 14)

const maxX = 101
const maxY = 103

type robot struct {
	pos utils.Pos
	vel utils.Pos
}

func updateRobots(robots []robot) {
	for i := range robots {
		r := &robots[i]
		r.pos = r.pos.Add(r.vel.X, r.vel.Y)
		r.pos.X = (r.pos.X + maxX) % maxX
		r.pos.Y = (r.pos.Y + maxY) % maxY
	}
}

func renderRobots(robots []robot) [][]byte {
	grid := utils.ByteGrid(maxX, maxY, '.')
	for _, r := range robots {
		grid[r.pos.Y][r.pos.X] = '#'
	}
	return grid
}

func gridEntropy(grid [][]byte) (runs int) {
	flat := bytes.Join(grid, nil)
	for i := range flat[1:] {
		runs += utils.BoolToInt(flat[i] != flat[i+1])
	}
	return
}

func main() {
	var robots []robot
	for _, line := range utils.Lines(input) {
		ints := utils.ExtractInts(line)
		robots = append(robots, robot{
			pos: utils.Pos{X: ints[0], Y: ints[1]},
			vel: utils.Pos{X: ints[2], Y: ints[3]},
		})
	}

	for i := 0; i < 100; i++ {
		updateRobots(robots)
	}
	var quads [4]int
	for i, min := range []utils.Pos{
		{X: 0, Y: 0},
		{X: maxX/2 + 1, Y: 0},
		{X: 0, Y: maxY/2 + 1},
		{X: maxX/2 + 1, Y: maxY/2 + 1},
	} {
		quad := utils.Rect{Min: min, Max: min.Add(maxX/2-1, maxY/2-1)}
		for _, r := range robots {
			quads[i] += utils.BoolToInt(quad.Contains(r.pos))
		}
	}
	utils.Println(quads[0] * quads[1] * quads[2] * quads[3])

	// minimize entropy
	minE := gridEntropy(renderRobots(robots))
	longevity := 0
	for seconds := 100; ; seconds++ {
		if e := gridEntropy(renderRobots(robots)); e < minE {
			minE = e
			longevity = 0
		} else if longevity++; longevity == 10000 {
			utils.Println(seconds - 10000)
			break
		}
		updateRobots(robots)
	}
}
