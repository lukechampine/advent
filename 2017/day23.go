package main

import (
	"math/big"
	"strconv"
	"strings"

	"github.com/lukechampine/advent/utils"
)

const input = `set b 93
set c b
jnz a 2
jnz 1 5
mul b 100
sub b -100000
set c b
sub c -17000
set f 1
set d 2
set e 2
set g d
mul g e
sub g b
jnz g 2
set f 0
sub e -1
set g e
sub g b
jnz g -8
sub d -1
set g d
sub g b
jnz g -13
jnz f 2
sub h -1
set g b
sub g c
jnz g 2
jnz 1 3
sub b -17
jnz 1 -23`

func main() {
	// part 1
	insts := utils.Lines(input)
	regs := make(map[string]int)
	regOrVal := func(s string) int {
		if i, err := strconv.Atoi(s); err == nil {
			return i
		}
		return regs[s]
	}
	muls := 0
	for i := 0; i < len(insts); i++ {
		inst := strings.Fields(insts[i])
		switch inst[0] {
		case "set":
			regs[inst[1]] = regOrVal(inst[2])
		case "sub":
			regs[inst[1]] -= regOrVal(inst[2])
		case "mul":
			regs[inst[1]] *= regOrVal(inst[2])
			muls++
		case "jnz":
			if regOrVal(inst[1]) != 0 {
				i += regOrVal(inst[2]) - 1 // account for i++
			}
		}
	}
	utils.Println(muls)

	// part 2

	/*
		The decompiled program is:

			h := 0
			for b := 109300; b <= 126300; b += 17 {
				f := true
				for d := 2; d != b; d++ {
					for e := 2; e != b; e++ {
						if d*e == b {
							f = false
						}
					}
				}
				if !f {
					h++
				}
			}
			utils.Println(n)

		In other words, it counts the composite numbers between 109300 and 126300,
		incrementing by 17 each time. It runs slowly because of the prime-check
		routine, which brute-forces every combination of factors.
	*/

	isPrime := func(x int) bool {
		return new(big.Int).SetInt64(int64(x)).ProbablyPrime(0)
	}

	n := 0
	for b := 109300; b <= 126300; b += 17 {
		if !isPrime(b) {
			n++
		}
	}
	utils.Println(n)
}
