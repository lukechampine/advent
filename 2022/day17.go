package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2022, 17)

func canPlace(p []string, pos utils.Pos, field [][]byte) bool {
	for y := range p {
		for x, b := range p[y] {
			if b != '.' {
				px := pos.X + x
				py := len(field) - 1 - pos.Y + y
				if py >= len(field) || px >= len(field[0]) || py < 0 || px < 0 || field[py][px] != '.' {
					return false
				}
			}
		}
	}
	return true
}

func place(p []string, pos utils.Pos, field [][]byte) {
	for y := range p {
		for x, b := range p[y] {
			if b != '.' {
				px := pos.X + x
				py := len(field) - 1 - pos.Y + y
				field[py][px] = byte(b)
			}
		}
	}
}

func main() {
	pieces := [][]string{
		{
			"1111",
		},
		{
			".2.",
			"222",
			".2.",
		},
		{
			"..3",
			"..3",
			"333",
		},
		{
			"4",
			"4",
			"4",
			"4",
		},
		{
			"55",
			"55",
		},
	}

	// determined manually by observing output
	const cycleBlocks = 1720
	const cycleHeight = 2704
	cycles := 1000000000000 / cycleBlocks
	cycleRemBlocks := 1000000000000 % cycleBlocks
	skippedHeight := cycles * cycleHeight

	field := make([][]byte, 10000)
	for i := range field {
		field[i] = []byte(".......")
	}
	height := 0
	g := 0
	blockHeights := make([]int, 5000)
	for i := range blockHeights {
		blockHeights[i] = height
		p := pieces[i%len(pieces)]
		pos := utils.Pos{X: 2, Y: height + len(p) + 2}
		for {
			// blow
			m := pos
			switch input[g%len(input)] {
			case '<':
				m.X--
			case '>':
				m.X++
			}
			if canPlace(p, m, field) {
				pos = m
			}
			g++
			// drop
			m = pos
			m.Y--
			if canPlace(p, m, field) {
				pos = m
			} else {
				break
			}
		}
		place(p, pos, field)
		if pos.Y+1 > height {
			height = pos.Y + 1
		}
	}
	utils.Println(blockHeights[2022])
	utils.Println(blockHeights[cycleRemBlocks] + skippedHeight)
}
