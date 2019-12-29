package main

import (
	"math/rand"

	"github.com/lukechampine/advent/2019/intcode"
	"github.com/lukechampine/advent/utils"
)

var input = utils.ReadInput("day15_input.txt")
var prog = utils.ExtractInts(input)

const gm = 50

var origin = utils.Pos{
	X: gm / 2,
	Y: gm / 2,
}

type robot struct {
	m      *intcode.Machine
	p      utils.Pos
	isWall map[utils.Pos]bool
}

func newRobot() *robot {
	return &robot{
		m:      intcode.New(prog),
		p:      origin,
		isWall: make(map[utils.Pos]bool),
	}
}

func (r *robot) findOxy() utils.Pos {
	for {
		d := rand.Intn(4) + 1
		np := r.p.Moves()[d-1]
		switch out := r.m.Run(d); out[0] {
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

		switch out := r.m.Run(d); out[0] {
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
