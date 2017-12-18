package main

import (
	"strconv"
	"strings"
	"time"

	"github.com/lukechampine/advent/utils"
)

const input = `set i 31
set a 1
mul p 17
jgz p p
mul a 2
add i -1
jgz i -2
add a -1
set i 127
set p 680
mul p 8505
mod p a
mul p 129749
add p 12345
mod p a
set b p
mod b 10000
snd b
add i -1
jgz i -9
jgz a 3
rcv b
jgz b -1
set f 0
set i 126
rcv a
rcv b
set p a
mul p -1
add p b
jgz p 4
snd a
set a b
jgz 1 3
snd b
set f 1
add i -1
jgz i -11
snd a
jgz f -16
jgz a -19`

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
	sound := 0
loop:
	for i := 0; i < len(insts); i++ {
		inst := strings.Fields(insts[i])
		switch inst[0] {
		case "snd":
			sound = regOrVal(inst[1])
		case "set":
			regs[inst[1]] = regOrVal(inst[2])
		case "add":
			regs[inst[1]] += regOrVal(inst[2])
		case "mul":
			regs[inst[1]] *= regOrVal(inst[2])
		case "mod":
			regs[inst[1]] %= regOrVal(inst[2])
		case "rcv":
			if regs[inst[1]] != 0 {
				utils.Println(sound)
				break loop
			}
		case "jgz":
			if regs[inst[1]] > 0 {
				i += regOrVal(inst[2]) - 1 // account for i++
			}
		}
	}

	// part 2
	prog := func(id int, out, in chan int) (sent int) {
		regs := make(map[string]int)
		regs["p"] = id

		regOrVal := func(s string) int {
			if i, err := strconv.Atoi(s); err == nil {
				return i
			}
			return regs[s]
		}

		insts := utils.Lines(input)
		for i := 0; i < len(insts); i++ {
			inst := strings.Fields(insts[i])
			switch inst[0] {
			case "set":
				regs[inst[1]] = regOrVal(inst[2])
			case "add":
				regs[inst[1]] += regOrVal(inst[2])
			case "mul":
				regs[inst[1]] *= regOrVal(inst[2])
			case "mod":
				regs[inst[1]] %= regOrVal(inst[2])
			case "snd":
				out <- regOrVal(inst[1])
				sent++
			case "rcv":
				select {
				case regs[inst[1]] = <-in:
				case <-time.After(time.Second):
					return
				}
			case "jgz":
				if regOrVal(inst[1]) > 0 {
					i += regOrVal(inst[2]) - 1
				}
			}
		}
		return
	}

	ch0, ch1 := make(chan int, 127), make(chan int, 127)
	go prog(0, ch0, ch1)
	utils.Println(prog(1, ch1, ch0))
}
