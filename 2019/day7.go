package main

import (
	"github.com/lukechampine/advent/utils"
)

var input = utils.ReadInput("day7_input.txt")
var prog = utils.ExtractInts(input)

type machineState byte

const (
	stateRunning machineState = iota
	stateAwaitingInput
	stateEmittedOutput
	stateHalted
)

type machine struct {
	prog  []int
	i     int
	state machineState
}

func newMachine(prog []int) machine {
	return machine{
		prog: append([]int(nil), prog...),
	}
}

func (m *machine) run(input int) (output int) {
	m.state = stateRunning
	inputUsed := false
	p := m.prog
	for m.i < len(p) {
		op, flags := p[m.i]%100, utils.Digits(utils.Itoa(p[m.i]/100))
		getArg := func(n int) *int {
			imm := p[m.i+n]
			if len(flags) < n || flags[len(flags)-n] == 0 {
				return &p[imm]
			}
			return &imm
		}
		switch op {
		case 1:
			*getArg(3) = *getArg(1) + *getArg(2)
			m.i += 4
		case 2:
			*getArg(3) = *getArg(1) * *getArg(2)
			m.i += 4
		case 3:
			if inputUsed {
				m.state = stateAwaitingInput
				return
			}
			*getArg(1) = input
			inputUsed = true
			m.i += 2
		case 4:
			x := *getArg(1)
			m.i += 2
			if x != 0 {
				m.state = stateEmittedOutput
				return x
			}
		case 5:
			if *getArg(1) != 0 {
				m.i = *getArg(2)
			} else {
				m.i += 3
			}
		case 6:
			if *getArg(1) == 0 {
				m.i = *getArg(2)
			} else {
				m.i += 3
			}
		case 7:
			*getArg(3) = utils.BoolToInt(*getArg(1) < *getArg(2))
			m.i += 4
		case 8:
			*getArg(3) = utils.BoolToInt(*getArg(1) == *getArg(2))
			m.i += 4
		case 99:
			m.state = stateHalted
			return -1
		default:
			m.i++
		}
	}
	panic("unreachable")
}

func main() {
	// part 1
	seqs := utils.Perms(5)
	utils.Println(utils.Maximum(len(seqs), func(i int) int {
		out := 0
		for _, s := range seqs[i] {
			m := newMachine(prog)
			m.run(s)
			out = m.run(out)
		}
		return out
	}))

	// part 2
	utils.Println(utils.Maximum(len(seqs), func(i int) int {
		seq := seqs[i]
		ms := make([]machine, len(seq))
		for i := range ms {
			ms[i] = newMachine(prog)
			ms[i].run(seq[i] + 5)
		}
		out := 0
		for {
			for i := range ms {
				next := ms[i].run(out)
				if ms[i].state == stateHalted {
					return out
				}
				out = next
			}
		}
	}))
}
