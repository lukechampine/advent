package main

import (
	"fmt"
)

const test = 7803

func calcPower(x, y int) int {
	rackID := x + 10
	power := rackID * y
	power += test
	power *= rackID
	hundreds := (power / 100) % 10
	return hundreds - 5
}

func gridPower(grid *[300][300]int, x, y, n int) int {
	var sum int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum += grid[y+j][x+i]
		}
	}
	return sum
}

func main() {
	// part 1
	var grid [300][300]int
	for y := range grid {
		for x := range grid[y] {
			grid[y][x] = calcPower(x+1, y+1)
		}
	}
	var maxPower int
	var maxX, maxY int
	for y := range grid[3:] {
		for x := range grid[y][3:] {
			if p := gridPower(&grid, x, y, 3); p > maxPower {
				maxPower = p
				maxX, maxY = x+1, y+1
			}
		}
	}
	fmt.Printf("%v,%v\n", maxX, maxY)

	// part 2
	maxPower, maxX, maxY = 0, 0, 0
	var maxN int
	// area probably isn't larger than 50
	for n := 1; n < 50; n++ {
		for y := range grid[n:] {
			for x := range grid[y][n:] {
				if p := gridPower(&grid, x, y, n); p > maxPower {
					maxPower = p
					maxX, maxY = x+1, y+1
					maxN = n
				}
			}
		}
	}
	fmt.Printf("%v,%v,%v\n", maxX, maxY, maxN)
}
