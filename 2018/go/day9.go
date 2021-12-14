package main

import (
	"container/ring"

	"lukechampine.com/advent/utils"
)

const numPlayers = 479
const lastMarblePoints = 71035

func winningScore(numPlayers, numMarbles int) int {
	pos := &ring.Ring{Value: 0}
	players := make([]int, numPlayers)
	curPlayer := 0
	for m := 0; m <= numMarbles; m++ {
		curPlayer = (curPlayer + 1) % len(players)
		if m%23 != 0 {
			pos = pos.Next()
			pos.Link(&ring.Ring{Value: m})
		} else {
			pos = pos.Move(-8)
			players[curPlayer] += m + pos.Unlink(1).Value.(int)
		}
		pos = pos.Next()
	}
	return utils.Maximum(len(players), func(i int) int {
		return players[i]
	})
}

func main() {
	// part 1
	utils.Println(winningScore(numPlayers, lastMarblePoints))

	// part 2
	utils.Println(winningScore(numPlayers, lastMarblePoints*100))
}
