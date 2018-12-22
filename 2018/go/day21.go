package main

import (
	"fmt"
	"strings"

	"github.com/lukechampine/advent/utils"
)

const input = `#ip 3
seti 123 0 1
bani 1 456 1
eqri 1 72 1
addr 1 3 3
seti 0 0 3
seti 0 0 1
bori 1 65536 2
seti 10605201 9 1
bani 2 255 5
addr 1 5 1
bani 1 16777215 1
muli 1 65899 1
bani 1 16777215 1
gtir 256 2 5
addr 5 3 3
addi 3 1 3
seti 27 3 3
seti 0 3 5
addi 5 1 4
muli 4 256 4
gtrr 4 2 4
addr 4 3 3
addi 3 1 3
seti 25 3 3
addi 5 1 5
seti 17 5 3
setr 5 5 2
seti 7 6 3
eqrr 1 0 5
addr 5 3 3
seti 5 8 3`

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
		if lines[ip+1] == "eqrr 1 0 5" {
			break
		}
	}
	utils.Println(regs[1])

	// part 2
	counts := make(map[int]struct{}, 1e6)
	var lastAdded int
	regs = [6]int{}
	for i := 0; i < 1e6; i++ {
		regs[2] = regs[1] | 65536
		regs[1] = 10605201
		for {
			regs[1] = ((((regs[1] + ((regs[2] & 255) & 0xFF)) & 16777215) * 65899) & 16777215)
			if regs[2] < 256 {
				break
			}
			regs[2] /= 256
		}
		if _, ok := counts[regs[1]]; !ok {
			counts[regs[1]] = struct{}{}
			lastAdded = regs[1]
		}
	}
	fmt.Println(lastAdded)
	return
}
