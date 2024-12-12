package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2024, 11)

var memo = map[[2]int]int{}

func totalStones(stone int, steps int) int {
	if steps == 0 {
		return 1
	} else if n, ok := memo[[2]int{stone, steps}]; ok {
		return n
	}

	var n int
	if stone == 0 {
		n = totalStones(1, steps-1)
	} else if str := utils.Itoa(stone); len(str)%2 == 0 {
		n += totalStones(utils.Atoi(str[:len(str)/2]), steps-1)
		n += totalStones(utils.Atoi(str[len(str)/2:]), steps-1)
	} else {
		n = totalStones(stone*2024, steps-1)
	}
	memo[[2]int{stone, steps}] = n
	return n
}

func main() {
	stones := utils.ExtractInts(input)
	var sum int
	for _, s := range stones {
		sum += totalStones(s, 25)
	}
	utils.Println(sum)

	sum = 0
	for _, s := range stones {
		sum += totalStones(s, 75)
	}
	utils.Println(sum)
}
