package main

import (
	"bytes"
	"strings"

	"github.com/lukechampine/advent/utils"
)

var input = utils.ReadInput("day8_input.txt")

func countLayer(layer []string, b byte) int {
	return utils.Sum(len(layer), func(i int) int {
		return strings.Count(layer[i], string(b))
	})
}

func main() {
	// part 1
	const maxY = 6
	const maxX = 25

	rows := utils.WrapString(input, maxX)
	img := make([][]string, len(rows)/maxY)
	for i := range img {
		img[i] = rows[i*maxY:][:maxY]
	}

	minLayer := img[utils.MinimumIndex(len(img), func(i int) int {
		return countLayer(img[i], '0')
	})]
	utils.Println(countLayer(minLayer, '1') * countLayer(minLayer, '2'))

	// part 2
	final := utils.ByteGrid(maxX, maxY, '2')
	for _, l := range img {
		for y := range l {
			for x := range l[y] {
				if final[y][x] == '2' {
					final[y][x] = l[y][x]
				}
			}
		}
	}
	flat := string(bytes.Join(final, []byte("\n")))
	utils.Println(utils.Replace(flat, "0", " ", "1", "█"))
}
