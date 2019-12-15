package main

import (
	"github.com/lukechampine/advent/utils"
)

var input = utils.ReadInput("day13_input.txt")
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
	rel   int
	state machineState
}

func newMachine(prog []int) *machine {
	return &machine{
		prog: append([]int(nil), prog...),
	}
}

func (m *machine) run(input int) (output int) {
	m.state = stateRunning
	p := m.prog
	for m.i < len(p) {
		op, flags := p[m.i]%100, utils.Digits(utils.Itoa(p[m.i]/100))
		getArg := func(n int) *int {
			imm := p[m.i+n]
			if len(flags) < n || flags[len(flags)-n] == 0 {
				return &p[imm]
			} else if flags[len(flags)-n] == 1 {
				return &imm
			} else if flags[len(flags)-n] == 2 {
				return &p[m.rel+imm]
			}
			panic("unreachable")
		}
		switch op {
		case 1:
			*getArg(3) = *getArg(1) + *getArg(2)
			m.i += 4
		case 2:
			*getArg(3) = *getArg(1) * *getArg(2)
			m.i += 4
		case 3:
			*getArg(1) = input
			m.i += 2
		case 4:
			x := *getArg(1)
			m.i += 2
			return x
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
		case 9:
			m.rel += *getArg(1)
			m.i += 2
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
	prog = append(prog, make([]int, 30000)...)
	m := newMachine(prog)

	tiles := make(map[utils.Pos]int)
	for m.state != stateHalted {
		x, y, id := m.run(0), m.run(0), m.run(0)
		tiles[utils.Pos{X: x, Y: y}] = id
	}

	var n int
	for _, id := range tiles {
		if id == 2 {
			n++
		}
	}
	utils.Println(n)

	// part 2
	var tilt int
	var ballX int
	var paddleX int
	prog[0] = 2
	m = newMachine(prog)
	for m.state != stateHalted {
		x, y, id := m.run(tilt), m.run(tilt), m.run(tilt)
		tiles[utils.Pos{X: x, Y: y}] = id
		if id == 4 {
			ballX = x
		} else if id == 3 {
			paddleX = x
		}
		if ballX == paddleX {
			tilt = 0
		} else if ballX > paddleX {
			tilt = 1
		} else {
			tilt = -1
		}
	}
	utils.Println(tiles[utils.Pos{X: -1, Y: 0}])
}