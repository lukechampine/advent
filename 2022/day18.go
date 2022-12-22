package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2022, 18)

type cube struct {
	x, y, z int
}

func (c cube) adj() []cube {
	return []cube{
		{c.x - 1, c.y, c.z},
		{c.x + 1, c.y, c.z},
		{c.x, c.y - 1, c.z},
		{c.x, c.y + 1, c.z},
		{c.x, c.y, c.z - 1},
		{c.x, c.y, c.z + 1},
	}
}

func main() {
	lava := make(map[cube]bool)
	for _, line := range utils.Lines(input) {
		var c cube
		utils.Sscanf(line, "%d,%d,%d", &c.x, &c.y, &c.z)
		lava[c] = true
	}

	surfaceArea := 0
	for c := range lava {
		for _, n := range c.adj() {
			surfaceArea += utils.BoolToInt(!lava[n])
		}
	}
	utils.Println(surfaceArea)

	var min, max cube
	for c := range lava {
		max.x = utils.Max(max.x, c.x)
		max.y = utils.Max(max.y, c.y)
		max.z = utils.Max(max.z, c.z)
		min.x = utils.Min(min.x, c.x)
		min.y = utils.Min(min.y, c.y)
		min.z = utils.Min(min.z, c.z)
	}
	min.x--
	min.y--
	min.z--
	max.x++
	max.y++
	max.z++

	// flood fill air, within bounds
	queue := []cube{max}
	air := make(map[cube]bool)
	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]
		if air[q] || q.x < min.x || q.y < min.y || q.z < min.z || q.x > max.x || q.y > max.y || q.z > max.z {
			continue
		}
		air[q] = true
		for _, adj := range q.adj() {
			if !lava[adj] {
				queue = append(queue, adj)
			}
		}
	}
	// delete non-adjacent
outer:
	for c := range air {
		for _, n := range c.adj() {
			if lava[n] {
				continue outer
			}
		}
		delete(air, c)
	}
	utils.Println("air:", len(air))

	exteriorArea := 0
	for c := range lava {
		for _, n := range c.adj() {
			exteriorArea += utils.BoolToInt(air[n])
		}
	}
	utils.Println(exteriorArea)
}
