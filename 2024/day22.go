package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2024, 22)

func next(x int) int {
	x ^= x << 6
	x &= (1 << 24) - 1
	x ^= x >> 5
	x &= (1 << 24) - 1
	x ^= x << 11
	x &= (1 << 24) - 1
	return x
}

func main() {
	secrets := utils.ExtractInts(input)
	utils.Println(utils.Sum(len(secrets), func(i int) int {
		n := secrets[i]
		for i := 0; i < 2000; i++ {
			n = next(n)
		}
		return n
	}))

	sales := make(map[[4]int8]int)
	for _, n := range secrets {
		var prevPrice int8
		var deltas [4]int8
		seen := make(map[[4]int8]bool, 2000)
		for i := 0; i <= 2000; i++ {
			price := int8(n % 10)
			n = next(n)
			copy(deltas[:], deltas[1:])
			deltas[3] = price - prevPrice
			prevPrice = price
			if i >= 4 && !seen[deltas] {
				sales[deltas] += int(price)
				seen[deltas] = true
			}
		}
	}
	var best int
	for _, v := range sales {
		best = utils.Max(best, v)
	}
	utils.Println(best)
}
