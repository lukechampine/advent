package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2024, 13)

type machine struct {
	a, b  utils.Pos
	prize utils.Pos
}

func (m machine) fracs() (n0 int, d0 int, n1 int, d1 int) {
	n := m.a.X*m.prize.Y - m.a.Y*m.prize.X
	d := m.a.X*m.b.Y - m.a.Y*m.b.X
	return (m.prize.Y - m.b.Y*n/d), m.a.Y, n, d
}

func (m machine) winnable() bool {
	n0, d0, n1, d1 := m.fracs()
	return n0%d0 == 0 && n1%d1 == 0
}

func (m machine) winCost() int {
	n0, d0, n1, d1 := m.fracs()
	return 3*n0/d0 + n1/d1
}

func main() {
	var machines []machine
	for _, group := range utils.Split(input, "\n\n") {
		ints := utils.ExtractInts(group)
		machines = append(machines, machine{
			a:     utils.Pos{X: ints[0], Y: ints[1]},
			b:     utils.Pos{X: ints[2], Y: ints[3]},
			prize: utils.Pos{X: ints[4], Y: ints[5]},
		})
	}
	var sum int
	for _, m := range machines {
		if m.winnable() {
			sum += m.winCost()
		}
	}
	utils.Println(sum)

	sum = 0
	for _, m := range machines {
		m.prize.X += 10000000000000
		m.prize.Y += 10000000000000
		if m.winnable() {
			sum += m.winCost()
		}
	}
	utils.Println(sum)
}
