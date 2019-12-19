package main

import (
	"github.com/lukechampine/advent/utils"
)

var input = utils.ReadInput("day17_input.txt")
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

	input []int
}

func newMachine(prog []int) *machine {
	return &machine{
		prog: append(prog[:len(prog):len(prog)], make([]int, 30000)...),
	}
}

func (m *machine) run() (output int) {
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
			*getArg(1) = m.input[0]
			m.input = m.input[1:]
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
	m := newMachine(prog)
	var grid [][]byte
	var row []byte
	for m.state != stateHalted {
		c := m.run()
		if c == '\n' {
			grid = append(grid, row)
			row = nil
		} else {
			row = append(row, byte(c))
		}
	}

	var sum int
	for y := range grid {
		for x := range grid[y] {
			p := utils.Pos{X: x, Y: y}
			if grid[y][x] != '#' {
				continue
			}
			intersection := true
			for _, move := range p.Moves() {
				if move.Y < 0 || move.X < 0 || move.Y >= len(grid) || move.X >= len(grid[move.Y]) {
					intersection = false
					break
				}
				intersection = intersection && grid[move.Y][move.X] == '#'
			}
			if intersection {
				sum += x * y
			}
		}
	}
	utils.Println(sum)

	prog[0] = 2
	m = newMachine(prog)
	for _, c := range []byte("A,B,A,C,B,C,A,B,A,C\nR,6,L,10,R,8,R,8\nR,12,L,8,L,10\nR,12,L,10,R,6,L,10\nn\n") {
		m.input = append(m.input, int(c))
	}
	for m.state != stateHalted {
		out := m.run()
		if out > 128 {
			utils.Println(out)
			break
		}
	}
}
