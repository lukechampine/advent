package main

import (
	"fmt"

	"github.com/lukechampine/advent/utils"
)

const input = `4	1	15	12	0	9	9	5	5	8	7	3	14	5	12	3`

func parse(s string) []int {
	return utils.IntList(s)
}

func redistribute(blocks []int) {
	// find highest block
	var maxIndex int
	for i := range blocks {
		if blocks[i] > blocks[maxIndex] {
			maxIndex = i
		}
	}
	distrib := blocks[maxIndex]
	blocks[maxIndex] = 0
	for j := (maxIndex + 1) % len(blocks); distrib > 0; j = (j + 1) % len(blocks) {
		blocks[j]++
		distrib--
	}
}

func main() {
	// part 1
	blocks := parse(input)
	seen := make(map[string]int)
	cycles := 0
	for {
		key := fmt.Sprint(blocks)
		if _, ok := seen[key]; ok {
			break
		}
		seen[key] = cycles
		redistribute(blocks)
		cycles++
	}
	utils.Println(cycles)

	// part 2
	utils.Println(cycles - seen[fmt.Sprint(blocks)])
}
