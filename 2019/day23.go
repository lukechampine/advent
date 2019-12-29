package main

import (
	"github.com/lukechampine/advent/2019/intcode"
	"github.com/lukechampine/advent/utils"
)

var input = utils.ReadInput("day23_input.txt")
var prog = utils.ExtractInts(input)

func runMachine(m *intcode.Machine, q, nat chan [2]int, queues []chan [2]int) {
	var dst int = -2
	select {
	case packet := <-q:
		dst = m.Run(packet[0], packet[1])
	default:
		dst = m.Run(-1)
	}
	for !m.AwaitingInput {
		x := m.Run()
		y := m.Run()
		if dst == 255 {
			nat <- [2]int{x, y}
		} else if dst < len(queues) {
			queues[dst] <- [2]int{x, y}

		}
		dst = m.Run()
	}
}

func main() {
	// part 1
	queues := make([]chan [2]int, 50)
	for i := range queues {
		queues[i] = make(chan [2]int, 5000)
	}
	nat := make(chan [2]int, 5000)
	machines := make([]*intcode.Machine, 50)
	for i := range machines {
		m := intcode.New(prog)
		dst := m.Run(i)

		for !m.AwaitingInput {
			x := m.Run()
			y := m.Run()
			if dst == 255 {
				nat <- [2]int{x, y}
			} else if dst < len(queues) {
				queues[dst] <- [2]int{x, y}
			}
			dst = m.Run()
		}
		machines[i] = m
	}
	ys := make(map[int]struct{})
	for {
		for i, m := range machines {
			runMachine(m, queues[i], nat, queues)
		}
		allHalted := utils.All(len(machines), func(i int) bool { return machines[i].Halted })
		if allHalted {
			break
		}
		allQueuesEmpty := utils.All(len(machines), func(i int) bool { return len(queues[i]) == 0 })
		allAwaitingInput := utils.All(len(machines), func(i int) bool { return machines[i].AwaitingInput })
		if allQueuesEmpty && allAwaitingInput {
			var natVal [2]int
			for len(nat) > 0 {
				natVal = <-nat
			}
			if _, ok := ys[natVal[1]]; ok {
				utils.Println(natVal[1])
				return
			}
			if len(ys) == 0 {
				utils.Println(natVal[1]) // first value sent
			}
			ys[natVal[1]] = struct{}{}
			queues[0] <- natVal
		}
	}
}
