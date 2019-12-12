package main

import (
	"fmt"

	"github.com/lukechampine/advent/utils"
)

var input = utils.ReadInput("day11_input.txt")
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
			imm := &p[m.i+n]
			if len(flags) < n || flags[len(flags)-n] == 0 {
				return &p[*imm]
			} else if flags[len(flags)-n] == 1 {
				return imm
			} else if flags[len(flags)-n] == 2 {
				return &p[m.rel+*imm]
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

const (
	Black = 0
	White = 1
)

func main() {
	// part 1
	prog = append(prog, make([]int, 30000)...)
	m := newMachine(prog)

	robot := utils.NewAgent(0, 0, utils.Up)
	panels := make(map[utils.Pos]int)
	getInput := func() int {
		if panels[robot.Pos] == White {
			return 1
		}
		return 0
	}
	paintPanel := func(color int) {
		switch color {
		case 0:
			panels[robot.Pos] = Black
		case 1:
			panels[robot.Pos] = White
		default:
			panic(fmt.Sprint(color, m.state))
		}
	}
	turn := func(dir int) {
		switch dir {
		case 0:
			robot.TurnLeft()
		case 1:
			robot.TurnRight()
		default:
			panic(fmt.Sprint(dir, m.state))
		}
		robot.MoveForward(1)
	}
	for {
		in := getInput()
		color, dir := m.run(in), m.run(in)
		if m.state == stateHalted {
			break
		}
		paintPanel(color)
		turn(dir)
	}
	utils.Println(len(panels))

	// part 2
	m = newMachine(prog)
	robot = utils.NewAgent(0, 0, utils.Up)
	panels = make(map[utils.Pos]int)
	panels[utils.Pos{0, 0}] = White
	for {
		in := getInput()
		color, dir := m.run(in), m.run(in)
		if m.state == stateHalted {
			break
		}
		paintPanel(color)
		turn(dir)
	}
	var painted []utils.Pos
	for p, color := range panels {
		if color == White {
			painted = append(painted, p)
		}
	}
	utils.PrintPositions(painted, ' ', '█')
}
