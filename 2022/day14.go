package main

import (
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2022, 14)

func drop(grid [][]byte, p utils.Pos) utils.Pos {
	for p.Y+1 < len(grid) {
		if grid[p.Y+1][p.X] == '.' {
			p.Y++
		} else if grid[p.Y+1][p.X-1] == '.' {
			p.Y++
			p.X--
		} else if grid[p.Y+1][p.X+1] == '.' {
			p.Y++
			p.X++
		} else {
			break
		}
	}
	return p
}

func main() {
	const offsetX = 100
	var veins [][]utils.Pos
	var maxX, maxY int
	for _, line := range utils.Lines(input) {
		var points []utils.Pos
		for _, p := range strings.Split(line, " -> ") {
			var x, y int
			utils.Sscanf(p, "%d,%d", &x, &y)
			x += offsetX
			points = append(points, utils.Pos{x, y})
			if x > maxX {
				maxX = x
			}
			if y > maxY {
				maxY = y
			}
		}
		veins = append(veins, points)
	}
	maxX += offsetX * 2
	maxY += 1
	grid := utils.ByteGrid(maxX+1, maxY+1, '.')
	for _, vein := range veins {
		p := vein[0]
		grid[p.Y][p.X] = '#'
		for _, q := range vein[1:] {
			for p != q {
				p = p.StrideTowards(q)
				grid[p.Y][p.X] = '#'
			}
			p = q
		}
	}

	for {
		sand := drop(grid, utils.Pos{500 + offsetX, 0})
		if sand.Y == maxY {
			break
		}
		grid[sand.Y][sand.X] = 'o'
	}
	utils.Println(utils.CountGrid(grid, 'o'))

	for {
		sand := drop(grid, utils.Pos{500 + offsetX, 0})
		grid[sand.Y][sand.X] = 'o'
		if sand.Y == 0 {
			break
		}
	}
	utils.Println(utils.CountGrid(grid, 'o'))
}
