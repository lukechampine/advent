package main

import "github.com/lukechampine/advent/utils"

const input = 1362

func isWall(p utils.Pos) bool {
	n := uint64(p.X*p.X + 3*p.X + 2*p.X*p.Y + p.Y + p.Y*p.Y + input)
	n -= (n >> 1) & 0x5555555555555555
	n = (n>>2)&0x3333333333333333 + n&0x3333333333333333
	n += n >> 4
	n &= 0x0f0f0f0f0f0f0f0f
	n *= 0x0101010101010101
	return byte(n>>56)%2 != 0
}

func main() {
	// part 1
	m := utils.Maze{
		Grid:   utils.Grid{50, 50},
		IsWall: isWall,
	}
	dist := m.DistancesFrom(utils.Pos{1, 1})
	println(dist[utils.Pos{31, 39}])

	// part 2
	var total int
	for y := 0; y < 30; y++ {
		for x := 0; x < 30; x++ {
			if d, ok := dist[utils.Pos{x, y}]; ok && d <= 50 {
				total++
			}
		}
	}
	println(total)
}
