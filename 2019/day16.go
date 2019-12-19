package main

import (
	"github.com/lukechampine/advent/utils"
)

var input = utils.ReadInput("day16_input.txt")

func phase(vec []int) []int {
	prod := make([]int, len(vec))
	for stride := 1; stride <= len(vec); stride++ {
		var sum int
		for j := stride - 1; j < len(vec); {
			for k := 0; k < stride && j < len(vec); k++ {
				sum += vec[j]
				j++
			}
			j += stride
			for k := 0; k < stride && j < len(vec); k++ {
				sum -= vec[j]
				j++
			}
			j += stride
		}
		prod[stride-1] = utils.Abs(sum) % 10
	}
	return prod
}

func addPhase(off int, vec []int) []int {
	for i := len(vec) - 2; i >= off-1; i-- {
		vec[i] += vec[i+1]
		vec[i+1] = utils.Abs(vec[i+1]) % 10
	}
	return vec
}

func main() {
	// part 1
	ds := utils.Digits(input)
	for p := 0; p < 100; p++ {
		ds = phase(ds)
	}
	for _, d := range ds[:8] {
		utils.Print(d)
	}
	utils.Println()

	// part 2
	ds = make([]int, len(ds)*10000)

	offset := utils.Atoi(input[:7])
	digs := utils.Digits(input)
	for i := range ds {
		ds[i] = digs[i%len(digs)]
	}
	for p := 0; p < 100; p++ {
		ds = addPhase(offset, ds)
	}
	for _, d := range ds[offset:][:8] {
		utils.Print(d)
	}
	utils.Println()
}
