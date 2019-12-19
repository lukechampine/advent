package main

import (
	"math/rand"

	"github.com/lukechampine/advent/utils"
)

var input = utils.ReadInput("day15_input.txt")
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
		prog: append(prog[:len(prog):len(prog)], make([]int, 30000)...),
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

const gm = 50

var origin = utils.Pos{
	X: gm / 2,
	Y: gm / 2,
}

type robot struct {
	m      *machine
	p      utils.Pos
	isWall map[utils.Pos]bool
}

func newRobot() *robot {
	return &robot{
		m:      newMachine(prog),
		p:      origin,
		isWall: make(map[utils.Pos]bool),
	}
}

func (r *robot) findOxy() utils.Pos {
	for {
		d := rand.Intn(4) + 1
		np := r.p.Moves()[d-1]
		switch r.m.run(d) {
		case 0:
		case 1:
			r.p = np
		case 2:
			return np
		}
	}
}

func (r *robot) exploreFully(avoid utils.Pos) map[utils.Pos]bool {
	seen := make(map[utils.Pos]bool)
	toVisit := make(map[utils.Pos]struct{})
	toVisit[r.p] = struct{}{}
	for {
		seen[r.p] = true
		delete(toVisit, r.p)
		moves := r.p.Moves()
		for _, m := range r.p.Moves() {
			if !seen[m] && !r.isWall[m] {
				toVisit[m] = struct{}{}
			}
		}
	retry:
		d := rand.Intn(4) + 1
		np := moves[d-1]
		if np == avoid {
			if len(toVisit) == 1 {
				return seen
			}
			goto retry
		}

		switch r.m.run(d) {
		case 0:
			r.isWall[np] = true
			delete(toVisit, np)
		case 1:
			r.p = np // ha
		case 2:
			panic("should avoid this")
		}
	}
}

func main() {
	// part 1
	oxy := newRobot().findOxy()

	r := newRobot()
	toFill := r.exploreFully(oxy)

	maze := utils.Maze{
		Grid:   utils.Grid{X: 50, Y: 50},
		IsWall: func(p utils.Pos) bool { return r.isWall[p] },
	}
	utils.Println(maze.DistancesFrom(origin)[oxy])

	// part 2
	floods := []utils.Pos{oxy}
	minutes := 0
	for {
		for _, f := range floods {
			delete(toFill, f)
		}
		var newFloods []utils.Pos
		for _, f := range floods {
			for _, m := range maze.ValidMoves(f) {
				if toFill[m] {
					newFloods = append(newFloods, m)
				}
			}
		}
		floods = newFloods
		if len(floods) == 0 {
			break
		}
		minutes++
	}
	utils.Println(minutes)
}
