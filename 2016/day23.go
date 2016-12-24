package main

import (
	"strconv"
	"strings"

	"github.com/lukechampine/advent/utils"
)

const input = `cpy a b
dec b
cpy a d
cpy 0 a
cpy b c
inc a
dec c
jnz c -2
dec d
jnz d -5
dec b
cpy b c
cpy c d
dec d
inc c
jnz d -2
tgl c
cpy -16 c
jnz 1 c
cpy 89 c
jnz 77 d
inc a
inc d
jnz d -2
inc c
jnz c -5`

const input2 = `cpy a b
dec b
mul a b
jnz 1 1
jnz 1 1
jnz 1 1
jnz 1 1
jnz 1 1
jnz 1 1
jnz 1 1
dec b
cpy b c
cpy c d
dec d
inc c
jnz d -2
tgl c
cpy -16 c
jnz 1 c
cpy 89 c
jnz 77 d
inc a
inc d
jnz d -2
inc c
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
			off := i + eval(in.y)
			if off >= len(insts) || off < 0 {
				continue
			}
			if eval(in.x) != 0 {
				i = off - 1 // offset i++
			}

		case "inc":
			regs[in.x]++

		case "dec":
			regs[in.x]--

		case "mul":
			regs[in.x] = eval(in.x) * eval(in.y)

		case "tgl":
			off := i + eval(in.x)
			if off >= len(insts) {
				continue
			}
			insts[off].op = map[string]string{
				// one arg
				"inc": "dec",
				"dec": "inc",
				"tgl": "inc",
				// two args
				"cpy": "jnz",
				"jnz": "cpy",
			}[insts[off].op]
		}
	}
}

func main() {
	// part 1
	insts := parse(input)
	regs := map[string]int{"a": 7}
	run(regs, insts)
	println(regs["a"])

	// part 2
	insts = parse(input2)
	regs = map[string]int{"a": 12}
	run(regs, insts)
	println(regs["a"])
}
