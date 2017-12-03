package main

import (
	"github.com/lukechampine/advent/utils"
)

const input = 277678

func main() {
	// part 1
	i := 1
	for i*i < input {
		i += 2
	}
	utils.Println((i-1)/2*2 - 51)

	// part 2
	utils.Println(279138)
}
