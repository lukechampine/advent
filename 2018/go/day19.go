package main

import (
	"strings"

	"github.com/lukechampine/advent/utils"
)

const input = `#ip 4
addi 4 16 4
seti 1 2 3
seti 1 6 1
mulr 3 1 2
eqrr 2 5 2
addr 2 4 4
addi 4 1 4
addr 3 0 0
addi 1 1 1
gtrr 1 5 2
addr 4 2 4
seti 2 8 4
addi 3 1 3
gtrr 3 5 2
addr 2 4 4
seti 1 4 4
mulr 4 4 4
addi 5 2 5
mulr 5 5 5
mulr 4 5 5
muli 5 11 5
addi 2 5 2
mulr 2 4 2
addi 2 18 2
addr 5 2 5
addr 4 0 4
seti 0 6 4
setr 4 8 2
mulr 2 4 2
addr 4 2 2
mulr 4 2 2
muli 2 14 2
mulr 2 4 2
addr 5 2 5
seti 0 1 0
seti 0 5 4`

var opcodes = map[string]func(regs [6]int, a, b, c int) [6]int{
	"addr": func(regs [6]int, a, b, c int) [6]int {
		regs[c] = regs[a] + regs[b]
		return regs
	},
	"addi": func(regs [6]int, a, b, c int) [6]int {
		regs[c] = regs[a] + b
		return regs
	},
	"mulr": func(regs [6]int, a, b, c int) [6]int {
		regs[c] = regs[a] * regs[b]
		return regs
	},
	"muli": func(regs [6]int, a, b, c int) [6]int {
		regs[c] = regs[a] * b
		return regs
	},
	"banr": func(regs [6]int, a, b, c int) [6]int {
		regs[c] = regs[a] & regs[b]
		return regs
	},
	"bani": func(regs [6]int, a, b, c int) [6]int {
		regs[c] = regs[a] & b
		return regs
	},
	"borr": func(regs [6]int, a, b, c int) [6]int {
		regs[c] = regs[a] | regs[b]
		return regs
	},
	"bori": func(regs [6]int, a, b, c int) [6]int {
		regs[c] = regs[a] | b
		return regs
	},
	"setr": func(regs [6]int, a, b, c int) [6]int {
		regs[c] = regs[a]
		return regs
	},
	"seti": func(regs [6]int, a, b, c int) [6]int {
		regs[c] = a
		return regs
	},
	"gtir": func(regs [6]int, a, b, c int) [6]int {
		if a > regs[b] {
			regs[c] = 1
		} else {
			regs[c] = 0
		}
		return regs
	},
	"gtri": func(regs [6]int, a, b, c int) [6]int {
		if regs[a] > b {
			regs[c] = 1
		} else {
			regs[c] = 0
		}
		return regs
	},
	"gtrr": func(regs [6]int, a, b, c int) [6]int {
		if regs[a] > regs[b] {
			regs[c] = 1
		} else {
			regs[c] = 0
		}
		return regs
	},
	"eqir": func(regs [6]int, a, b, c int) [6]int {
		if a == regs[b] {
			regs[c] = 1
		} else {
			regs[c] = 0
		}
		return regs
	},
	"eqri": func(regs [6]int, a, b, c int) [6]int {
		if regs[a] == b {
			regs[c] = 1
		} else {
			regs[c] = 0
		}
		return regs
	},
	"eqrr": func(regs [6]int, a, b, c int) [6]int {
		if regs[a] == regs[b] {
			regs[c] = 1
		} else {
			regs[c] = 0
		}
		return regs
	},
}

type inst struct {
	op      string
	a, b, c int
}

func main() {
	// part 1
	lines := utils.Lines(input)
	ipreg := utils.Atoi(strings.TrimPrefix(lines[0], "#ip "))
	var insts []inst
	for _, line := range lines[1:] {
		var i inst
		utils.Sscanf(line, "%s %d %d %d", &i.op, &i.a, &i.b, &i.c)
		insts = append(insts, i)
	}

	var regs [6]int
	for ip := 0; ip < len(lines); ip++ {
		in := insts[ip]
		regs[ipreg] = ip
		regs = opcodes[in.op](regs, in.a, in.b, in.c)
		ip = regs[ipreg]
	}
	utils.Println(regs[0])

	// part 2
	regs = [6]int{1, 0, 0, 0, 0, 0}
	for ip := 0; ip < len(lines); ip++ {
		// jet
		if ip == 3 {
			if regs[5]%regs[3] == 0 {
				regs[0] += regs[3]
			}
			ip = 11
			continue
		}

		in := insts[ip]
		regs[ipreg] = ip
		regs = opcodes[in.op](regs, in.a, in.b, in.c)
		ip = regs[ipreg]
	}
	utils.Println(regs[0])
}
