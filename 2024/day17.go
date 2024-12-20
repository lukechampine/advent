package main

import (
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2024, 17)

func runFast(a, b, c int) (out []int) {
loop:
	out = append(out, (a^(a>>((a%8)^1))^4)%8)
	a /= 8
	if a != 0 {
		goto loop
	}
	return
}

func run(a, b, c int, prog []int) (out []int) {
	combo := func(arg int) int {
		switch arg {
		case 0, 1, 2, 3:
			return arg
		case 4:
			return a
		case 5:
			return b
		case 6:
			return c
		default:
			panic(arg)
		}
	}
	for i := 0; i+1 < len(prog); i += 2 {
		switch op, arg := prog[i], prog[i+1]; op {
		case 0:
			a /= (1 << combo(arg))
		case 1:
			b ^= arg
		case 2:
			b = combo(arg) % 8
		case 3:
			if a != 0 {
				i = arg - 2
			}
		case 4:
			b ^= c
		case 5:
			out = append(out, combo(arg)%8)
		case 6:
			b = a / (1 << combo(arg))
		case 7:
			c = a / (1 << combo(arg))
		}
	}
	return
}

func main() {
	groups := utils.Split(input, "\n\n")
	regs := utils.ExtractInts(groups[0])
	a, b, c := regs[0], regs[1], regs[2]
	prog := utils.ExtractInts(groups[1])
	var out []string
	for _, n := range run(a, b, c, prog) {
		out = append(out, utils.Itoa(n))
	}
	utils.Println(strings.Join(out, ","))

	i := 0
	for len(run(1<<i, b, c, prog)) != len(prog) {
		i++
	}
	a = 0
	for shift := i &^ 15; shift >= 0; shift -= 16 {
		d := utils.MaximumIndex(1<<16, func(d int) int {
			out := run(a|(d<<shift), b, c, prog)
			for j := range out {
				if out[len(out)-j-1] != prog[len(prog)-j-1] {
					return j
				}
			}
			return len(prog)
		})
		a |= d << shift
	}
	utils.Println(a)
}
