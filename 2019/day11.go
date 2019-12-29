package main

import (
	"github.com/lukechampine/advent/2019/intcode"
	"github.com/lukechampine/advent/utils"
)

var input = utils.ReadInput("day11_input.txt")
var prog = utils.ExtractInts(input)

const (
	Black = 0
	White = 1
)

func main() {
	// part 1
	m := intcode.New(prog)

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
			panic("unreachable")
		}
	}
	turn := func(dir int) {
		switch dir {
		case 0:
			robot.TurnLeft()
		case 1:
			robot.TurnRight()
		default:
			panic("unreachable")
		}
		robot.MoveForward(1)
	}
	out := m.Run(getInput())
	for !m.Halted {
		paintPanel(out[0])
		turn(out[1])
		out = m.Run(getInput())
	}
	utils.Println(len(panels))

	// part 2
	m = intcode.New(prog)
	robot = utils.NewAgent(0, 0, utils.Up)
	panels = make(map[utils.Pos]int)
	panels[utils.Pos{0, 0}] = White

	out = m.Run(getInput())
	for !m.Halted {
		paintPanel(out[0])
		turn(out[1])
		out = m.Run(getInput())
	}
	var painted []utils.Pos
	for p, color := range panels {
		if color == White {
			painted = append(painted, p)
		}
	}
	utils.PrintPositions(painted, ' ', '█')
}
