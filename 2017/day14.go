package main

import (
	"encoding/hex"

	"github.com/lukechampine/advent/utils"
)

const input = `oundnydw`

const test = `flqrgnkx`

func reverse(list []int, pos, length int) {
	double := append(list, list...)
	for i, j := pos, pos+length-1; i < j; i, j = i+1, j-1 {
		double[i], double[j] = double[j], double[i]
	}

	copy(list[pos:], double[pos:pos+length])
	if overlap := (pos + length) - len(list); overlap > 0 {
		copy(list[:overlap], double[len(list):])
	}
}

func round(list, lengths []int, pos, skip int) (int, int) {
	for _, length := range lengths {
		reverse(list, pos, length)
		pos = (pos + length + skip) % len(list)
		skip++
	}
	return pos, skip
}

func dense(list []int) []byte {
	dense := make([]byte, len(list)/16)
	for i := range dense {
		d := 0
		for _, l := range list[i*16:][:16] {
			d ^= l
		}
		dense[i] = byte(d)
	}
	return dense
}

func knotHash(s string) string {
	pos := 0
	skip := 0
	list := make([]int, 256)
	for i := range list {
		list[i] = i
	}
	lengths := append(make([]int, len(s)), 17, 31, 73, 47, 23)
	for i, c := range s {
		lengths[i] = int(c)
	}
	for i := 0; i < 64; i++ {
		pos, skip = round(list, lengths, pos, skip)
	}
	return hex.EncodeToString(dense(list))
}

func hexToBits(s string) (bits []bool) {
	for _, c := range []byte(s) {
		var x byte
		if c >= 'a' {
			x = 10 + c - 'a'
		} else {
			x = c - '0'
		}
		for j := 3; j >= 0; j-- {
			bits = append(bits, x&(1<<uint(j)) != 0)
		}
	}
	return bits
}

func visitNeighbors(visited map[utils.Pos]bool, m *utils.Maze, p utils.Pos) {
	visited[p] = true
	for _, move := range m.ValidMoves(p) {
		if !visited[move] {
			visitNeighbors(visited, m, move)
		}
	}
}

func main() {
	// part 1
	grid := make([][]bool, 128)
	for i := range grid {
		grid[i] = hexToBits(knotHash(input + "-" + utils.Itoa(i)))
	}
	squares := 0
	for i := range grid {
		for _, b := range grid[i] {
			if b {
				squares++
			}
		}
	}
	utils.Println(squares)

	// part 2
	m := utils.Maze{
		Grid: utils.Grid{127, 127},
		IsWall: func(p utils.Pos) bool {
			return !grid[p.Y][p.X]
		},
	}
	visited := make(map[utils.Pos]bool)
	regions := 0
	for y := range grid {
		for x, isSquare := range grid[y] {
			p := utils.Pos{x, y}
			if isSquare && !visited[p] {
				visitNeighbors(visited, &m, p)
				regions++
			}
		}
	}
	utils.Println(regions)
}
