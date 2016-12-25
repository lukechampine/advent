package main

import (
	"strconv"
	"strings"

	"github.com/lukechampine/advent/utils"
)

const input = `cpy a d
cpy 14 c
cpy 182 b
inc d
dec b
jnz b -2
dec c
jnz c -5
cpy d a
jnz 0 0
cpy a b
cpy 0 a
cpy 2 c
jnz b 2
jnz 1 6
dec b
dec c
jnz c -4
inc a
jnz 1 -7
cpy 2 b
jnz c 2
jnz 1 4
dec b
dec c
jnz 1 -4
jnz 0 0
out b
jnz a -19
jnz 1 -21`

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

func run(regs map[string]int, insts []inst) bool {
	eval := func(s string) int {
		if n, err := strconv.Atoi(s); err == nil {
			return n
		}
		return regs[s]
	}

	n := 0
	for i := 0; i < len(insts) && n < 20; i++ {
		switch in := insts[i]; in.op {
		case "cpy":
			regs[in.y] = eval(in.x)

		case "jnz":
			if eval(in.x) != 0 {
				i += eval(in.y)
				i-- // offset i++
			}

		case "inc":
			regs[in.x]++

		case "dec":
			regs[in.x]--

		case "out":
			if regs[in.x] != n%2 {
				return false
			}
			n++
		}
	}
	return true
}

func main() {
	// part 1
	insts := parse(input)
	for a := 0; ; a++ {
		regs := map[string]int{"a": a}
		if run(regs, insts) {
			println(a)
			return
		}
	}
}
