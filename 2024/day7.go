package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2024, 7)

func concat(x, y int) int {
	return utils.Atoi(utils.Itoa(x) + utils.Itoa(y))
}

func canMake(r int, acc int, xs []int) bool {
	if len(xs) == 0 {
		return r == acc
	}
	return canMake(r, acc+xs[0], xs[1:]) ||
		canMake(r, acc*xs[0], xs[1:])
}

func canMakeConcat(r int, acc int, xs []int) bool {
	if len(xs) == 0 {
		return r == acc
	}
	return canMakeConcat(r, acc+xs[0], xs[1:]) ||
		canMakeConcat(r, acc*xs[0], xs[1:]) ||
		canMakeConcat(r, concat(acc, xs[0]), xs[1:])
}

func main() {
	lines := utils.Lines(input)
	utils.Println(utils.Sum(len(lines), func(i int) int {
		ints := utils.ExtractInts(lines[i])
		return ints[0] * utils.BoolToInt(canMake(ints[0], ints[1], ints[2:]))
	}))
	utils.Println(utils.Sum(len(lines), func(i int) int {
		ints := utils.ExtractInts(lines[i])
		return ints[0] * utils.BoolToInt(canMakeConcat(ints[0], ints[1], ints[2:]))
	}))
}
