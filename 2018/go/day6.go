package main

import (
	"fmt"

	"github.com/lukechampine/advent/utils"
)

const input = `156, 193
81, 315
50, 197
84, 234
124, 162
339, 345
259, 146
240, 350
97, 310
202, 119
188, 331
199, 211
117, 348
350, 169
131, 355
71, 107
214, 232
312, 282
131, 108
224, 103
83, 122
352, 142
208, 203
319, 217
224, 207
327, 174
89, 332
254, 181
113, 117
120, 161
322, 43
115, 226
324, 222
151, 240
248, 184
207, 136
41, 169
63, 78
286, 43
84, 222
81, 167
128, 192
127, 346
213, 102
313, 319
207, 134
154, 253
50, 313
160, 330
332, 163`

type elem struct {
	dist  int
	coord int
}

func main() {
	// part 1
	var coords []utils.Pos
	for _, line := range utils.Lines(input) {
		var c utils.Pos
		utils.Sscanf(line, "%d, %d", &c.X, &c.Y)
		// avoid dealing with negative coordinates by offsetting
		c.X += 500
		c.Y += 500
		coords = append(coords, c)
	}
	var grid [1000][1000]elem
	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = elem{1e9, -1} // sentinel
		}
	}

	for id, c := range coords {
		for i := range grid {
			for j, e := range grid[i] {
				dist := c.Dist(utils.Pos{X: i, Y: j})
				if dist == e.dist {
					// sentinel for equal-dist points
					grid[i][j] = elem{dist, -1}
				} else if dist < e.dist {
					grid[i][j] = elem{dist, id}
				}
			}
		}
	}

	counts := make(map[int]int)
	for i := range grid {
		for _, e := range grid[i] {
			if e.coord != -1 {
				counts[e.coord]++
			}
		}
	}
	// delete coords that appear on an edge
	for i := range grid {
		delete(counts, grid[i][0].coord)
		delete(counts, grid[0][i].coord)
		delete(counts, grid[i][len(grid)-1].coord)
		delete(counts, grid[len(grid)-1][i].coord)
	}
	var max int
	for _, n := range counts {
		max = utils.Max(max, n)
	}
	fmt.Println(max)

	// part 2
	reg := 0
	for i := range grid {
		for j := range grid[i] {
			dist := 0
			for _, c := range coords {
				dist += c.Dist(utils.Pos{X: i, Y: j})
			}
			if dist < 10000 {
				reg++
			}
		}
	}

	fmt.Println(reg)
	return
}
