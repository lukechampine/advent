package main

import (
	"github.com/lukechampine/advent/utils"
)

func main() {
	// part 1
	a, b := uint64(634), uint64(301)
	matches := 0
	for i := 0; i < 40e6; i++ {
		a = (a * 16807) % 2147483647
		b = (b * 48271) % 2147483647
		if uint16(a) == uint16(b) {
			matches++
		}
	}
	utils.Println(matches)

	a, b = uint64(634), uint64(301)
	matches = 0
	for i := 0; i < 5e6; i++ {
		a = (a * 16807) % 2147483647
		for a%4 != 0 {
			a = (a * 16807) % 2147483647
		}
		b = (b * 48271) % 2147483647
		for b%8 != 0 {
			b = (b * 48271) % 2147483647
		}
		if uint16(a) == uint16(b) {
			matches++
		}
	}
	utils.Println(matches)
}
