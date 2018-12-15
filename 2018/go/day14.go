package main

import (
	"bytes"

	"github.com/lukechampine/advent/utils"
)

const input = 846601

func digits(i int) []byte {
	digs := utils.Digits(utils.Itoa(i))
	bs := make([]byte, len(digs))
	for i := range bs {
		bs[i] = byte(digs[i])
	}
	return bs
}

var sumTable [][]byte

func init() {
	sumTable = make([][]byte, 9+9+1)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			sumTable[i+j] = digits(i + j)
		}
	}
}

func main() {
	// part 1
	inDigs := digits(input)
	scores := make([]byte, 0, 3e7) // 30 million should be enough
	scores = append(scores, 3, 7)
	elf1, elf2 := 0, 1
	for len(scores) < cap(scores)-2 {
		newscores := sumTable[scores[elf1]+scores[elf2]]
		scores = append(scores, newscores...)
		elf1 = (elf1 + int(scores[elf1]) + 1) % len(scores)
		elf2 = (elf2 + int(scores[elf2]) + 1) % len(scores)
	}
	for _, s := range scores[input:][:10] {
		utils.Print(s)
	}
	utils.Println()

	// part 2
	utils.Println(bytes.Index(scores, inDigs))

}
