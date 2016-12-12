package main

import (
	"strconv"
	"strings"

	"github.com/lukechampine/advent/utils"
)

const input = `cpy 1 a
cpy 1 b
cpy 26 d
jnz c 2
jnz 1 5
cpy 7 c
inc d
dec c
jnz c -2
cpy a c
inc a
dec b
jnz b -2
cpy c b
dec d
jnz d -6
cpy 13 c
cpy 14 d
inc a
dec d
jnz d -2
dec c
jnz c -5`

type inst struct {
	op   string
	x, y string
}

func parse(s string) []inst {
	var insts []inst
	for _, line := range utils.Lines(s) {
		fs := append(strings.Fields(line), "")
		insts = append(insts, inst{
			op: fs[0],
			x:  fs[1],
			y:  fs[2],
		})
	}
	return insts
}

func run(regs map[string]int, insts []inst) {
	eval := func(s string) int {
		if n, err := strconv.Atoi(s); err == nil {
			return n
		}
		return regs[s]
	}
	for i := 0; i < len(insts); i++ {
		in := insts[i]
		switch in.op {
		case "cpy":
			regs[in.y] = eval(in.x)

		case "jnz":
			if eval(in.x) != 0 {
				i += eval(in.y)
				i-- // offset ++
			}

		case "inc":
			regs[in.x]++

		case "dec":
			regs[in.x]--
		}
	}
}

func main() {
	// part 1
	insts := parse(input)
	regs := make(map[string]int)
	run(regs, insts)
	println(regs["a"])

	// part 2
	regs = map[string]int{
		"c": 1,
	}
	run(regs, insts)
	println(regs["a"])
}
