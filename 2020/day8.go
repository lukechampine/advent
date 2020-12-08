package main

import (
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2020, 8)
var prog = utils.Lines(input)

func runProg(prog []string) (int, bool) {
	var pc, acc int
	seen := make(map[int]bool)
	for {
		l := prog[pc]
		fs := strings.Fields(l)
		arg := utils.Atoi(fs[1])
		switch fs[0] {
		case "acc":
			acc += arg
		case "jmp":
			pc += arg - 1
		case "nop":
		}
		pc++
		if pc >= len(prog) {
			return acc, true
		} else if seen[pc] {
			return acc, false
		}
		seen[pc] = true
	}
}

func main() {
	// part 1
	prog := utils.Lines(input)
	acc, _ := runProg(prog)
	utils.Println(acc)

	// part 2
	for i, l := range prog {
		prog[i] = utils.Replace(l, "jmp", "nop", "nop", "jmp")
		acc, halted := runProg(prog)
		if halted {
			utils.Println(acc)
			return
		}
		prog[i] = l
	}
}
