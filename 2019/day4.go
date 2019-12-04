package main

import (
	"strconv"

	"github.com/lukechampine/advent/utils"
)

var input = utils.ReadInput("day4_input.txt")
var iis = utils.ExtractInts(input)

func twoAdjSame(i int) bool {
	digs := utils.Digits(strconv.Itoa(i))
	for i := 0; i < len(digs)-1; i++ {
		if digs[i] == digs[i+1] {
			return true
		}
	}
	return false
}

func monotonic(i int) bool {
	digs := utils.Digits(strconv.Itoa(i))
	for i := 0; i < len(digs)-1; i++ {
		if digs[i] > digs[i+1] {
			return false
		}
	}
	return true
}

func twoAdjSameAlone(i int) bool {
	digs := utils.Digits(strconv.Itoa(i))
	for i := 0; i < len(digs)-1; i++ {
		count := 1
		for i < len(digs)-1 && digs[i] == digs[i+1] {
			count++
			i++
		}
		if count == 2 {
			return true
		}
	}
	return false
}

func main() {
	// part 1
	var num int
	min, max := iis[0], iis[1]
	for i := min; i < max; i++ {
		if twoAdjSame(i) && monotonic(i) {
			num++
		}
	}
	utils.Println(num)

	// part 2
	num = 0
	for i := min; i < max; i++ {
		if twoAdjSameAlone(i) && monotonic(i) {
			num++
		}
	}
	utils.Println(num)
}
