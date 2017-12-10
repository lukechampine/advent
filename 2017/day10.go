package main

import (
	"encoding/hex"
	"strings"

	"github.com/lukechampine/advent/utils"
)

const input = `34,88,2,222,254,93,150,0,199,255,39,32,137,136,1,167`

func parse(s string) []int {
	return utils.IntList(strings.Replace(s, ",", " ", -1))
}

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

func main() {
	// part 1
	pos := 0
	skip := 0
	list := make([]int, 256)
	for i := range list {
		list[i] = i
	}
	lengths := parse(input)
	round(list, lengths, pos, skip)
	utils.Println(list[0] * list[1])

	// part 2
	pos = 0
	skip = 0
	for i := range list {
		list[i] = i
	}
	lengths = append(make([]int, len(input)), 17, 31, 73, 47, 23)
	for i, c := range input {
		lengths[i] = int(c)
	}
	for i := 0; i < 64; i++ {
		pos, skip = round(list, lengths, pos, skip)
	}
	utils.Println(hex.EncodeToString(dense(list)))
}
