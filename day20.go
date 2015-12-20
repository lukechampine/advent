package main

const input = 29000000

func main() {
	// part 1
	houses := make([]int, input/10) // answer won't be higher than input/10
	for elf := 1; elf < len(houses); elf++ {
		for h := elf; h < len(houses); h += elf {
			houses[h] += elf * 10
		}
	}
	for h, presents := range houses {
		if presents > input {
			println(h)
			break
		}
	}

	// // part 2
	houses = make([]int, input/10)
	for elf := 1; elf < len(houses); elf++ {
		count := 0
		for h := elf; h < len(houses) && count < 50; h += elf {
			houses[h] += elf * 11
			count++
		}
	}
	for h, presents := range houses {
		if presents > input {
			println(h)
			break
		}
	}
}
