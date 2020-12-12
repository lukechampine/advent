package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2020, 12)

func main() {
	lines := utils.Lines(input)

	ship := utils.NewAgent(0, 0, utils.Right)
	for _, l := range lines {
		n := utils.Atoi(l[1:])
		switch l[0] {
		case 'R':
			ship.SpinRight(n / 90)
		case 'L':
			ship.SpinLeft(n / 90)
		case 'F':
			ship.MoveForward(n)
		case 'N', 'E', 'W', 'S':
			ship.Move(utils.DirFromNEWS(l[0]), n)
		}
	}
	utils.Println(ship.Dist(utils.Pos{0, 0}))

	ship.Pos = utils.Pos{0, 0}
	waypoint := utils.Pos{10, 1}
	for _, l := range lines {
		n := utils.Atoi(l[1:])
		switch l[0] {
		case 'R':
			waypoint = waypoint.RotateClockwiseAround(ship.Pos, float64(n))
		case 'L':
			waypoint = waypoint.RotateCounterClockwiseAround(ship.Pos, float64(n))
		case 'F':
			r := waypoint.Rel(ship.Pos)
			ship.Pos = ship.Pos.Add(r.X*n, r.Y*n)
			waypoint = waypoint.Add(r.X*n, r.Y*n)
		case 'N', 'E', 'W', 'S':
			waypoint = waypoint.Move(utils.DirFromNEWS(l[0]), n)
		}
	}
	utils.Println(ship.Dist(utils.Pos{0, 0}))
}
