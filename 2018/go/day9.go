package main

import (
	"container/ring"

	"github.com/lukechampine/advent/utils"
)

const numPlayers = 479
const lastMarblePoints = 71035

func winningScore(numPlayers, numMarbles int) int {
	marbles := make([]int, numMarbles)
	for i := range marbles {
		marbles[i] = i
	}
	pos := &ring.Ring{Value: 0}
	players := make([]int, numPlayers)
	curPlayer := 0
	for _, m := range marbles[1:] {
		curPlayer = (curPlayer + 1) % len(players)
		if m%23 == 0 {
			players[curPlayer] += m
			pos = pos.Move(-7)
			players[curPlayer] += pos.Value.(int)
			pos.Prev().Unlink(1)
		} else {
			pos = pos.Next()
			pos.Link(&ring.Ring{Value: m})
		}
		pos = pos.Next()
	}
	return utils.Maximum(len(players), func(i int) int {
		return players[i]
	})
}

func main() {
	// part 1
	utils.Println(winningScore(numPlayers, lastMarblePoints+1))

	// part 2
	utils.Println(winningScore(numPlayers, lastMarblePoints*100+1))
}
