package main

import (
	"container/ring"

	"github.com/lukechampine/advent/utils"
)

func main() {
	// part 1
	spin := &ring.Ring{Value: 0}
	steps := 382
	for i := 1; i <= 2017; i++ {
		spin = spin.Move(steps)
		next := &ring.Ring{Value: i}
		next.Link(spin.Next())
		spin.Link(next)
		spin = next
	}
	utils.Println(spin.Next().Value.(int))

	// part 2
	root := &ring.Ring{Value: 0}
	spin = root
	for i := 1; i <= 50e6; i++ {
		spin = spin.Move(steps)
		next := &ring.Ring{Value: i}
		next.Link(spin.Next())
		spin.Link(next)
		spin = next
	}
	utils.Println(root.Next().Value.(int))
}
