package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2023, 3)

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func main() {
	grid := utils.ToByteGrid(utils.Lines(input))
	g := utils.Grid{len(grid[0]) - 1, len(grid) - 1}
	parts := make(map[utils.Pos][]int)
	for y := range grid {
		for x := 0; x < len(grid[y]); x++ {
			var num string
			for x < len(grid[y]) && isDigit(grid[y][x]) {
				num += string(grid[y][x])
				x++
			}
			if num == "" {
				continue
			}
		outer:
			for p := (utils.Pos{x - len(num), y}); p.X < x; p.X++ {
				for _, adj := range p.ValidNumpad(g) {
					if c := grid[adj.Y][adj.X]; !isDigit(c) && c != '.' {
						parts[adj] = append(parts[adj], utils.Atoi(num))
						break outer
					}
				}
			}
		}
	}
	var part1, part2 int
	for p, nums := range parts {
		part1 += utils.IntSum(nums)
		if grid[p.Y][p.X] == '*' && len(nums) == 2 {
			part2 += nums[0] * nums[1]
		}
	}
	utils.Println(part1)
	utils.Println(part2)
}
