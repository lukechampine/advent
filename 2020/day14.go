package main

import (
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2020, 14)

func applyPart1(index, val int, mask string, mem map[int]int) {
	for i := range mask {
		j := len(mask) - i - 1
		if mask[j] == '1' {
			val |= 1 << i
		} else if mask[j] == '0' {
			val &^= 1 << i
		}
	}
	mem[index] = val
}

func applyPart2(index, val int, mask string, mem map[int]int) {
	var floating []int
	for i := range mask {
		j := len(mask) - i - 1
		if mask[j] == '1' {
			index |= 1 << i
		} else if mask[j] == 'X' {
			floating = append(floating, i)
		}
	}

	x := 1 << len(floating)
	for i := 0; i < x; i++ {
		idx := index
		for b, l := range floating {
			if i&(1<<b) != 0 {
				idx ^= 1 << l
			}
		}
		mem[idx] = val
	}
}

func runProg(lines []string, apply func(int, int, string, map[int]int)) int {
	mem := make(map[int]int)
	var mask string
	for _, l := range lines {
		switch {
		case strings.HasPrefix(l, "mask"):
			mask = strings.Fields(l)[2]
		case strings.HasPrefix(l, "mem"):
			ns := utils.ExtractInts(l)
			apply(ns[0], ns[1], mask, mem)
		}
	}
	var sum int
	for _, i := range mem {
		sum += i
	}
	return sum
}

func main() {
	lines := utils.Lines(input)
	utils.Println(runProg(lines, applyPart1))
	utils.Println(runProg(lines, applyPart2))
}
