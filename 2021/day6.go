package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2021, 6)
var inputInts = utils.ExtractInts(input)

func step(fish *[9]int) {
	r := fish[0]
	copy(fish[:], fish[1:])
	fish[8] = r
	fish[6] += r
}

func main() {
	var fish [9]int
	for _, i := range inputInts {
		fish[i]++
	}
	for i := 0; i < 256; i++ {
		if i == 80 {
			utils.Println(utils.IntSum(fish[:]))
		}
		step(&fish)
	}
	utils.Println(utils.IntSum(fish[:]))
}
