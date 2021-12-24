package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2021, 21)

func playPart1(p1, p2 int) int {
	die := 0
	roll := func() int {
		die = die%100 + 1
		return die
	}

	var s1, s2 int
	for rolls := 0; ; rolls += 6 {
		p1 = (p1 + roll() + roll() + roll()) % 10
		if s1 += p1 + 1; s1 >= 1000 {
			return s2 * (rolls + 3)
		}
		p2 = (p2 + roll() + roll() + roll()) % 10
		if s2 += p2 + 1; s2 >= 1000 {
			return s1 * (rolls + 6)
		}
	}
}

var memo = map[[4]int][2]int{}

func playPart2(p1, p2 int, s1, s2 int) (p1wins, p2wins int) {
	if w, ok := memo[[4]int{p1, p2, s1, s2}]; ok {
		return w[0], w[1]
	}

	rolls := []int{1, 2, 3}
	for _, roll1 := range rolls {
		for _, roll2 := range rolls {
			for _, roll3 := range rolls {
				p1 := (p1 + roll1 + roll2 + roll3) % 10
				s1 := s1 + p1 + 1
				if s1 >= 21 {
					p1wins++
					continue
				}
				for _, roll4 := range rolls {
					for _, roll5 := range rolls {
						for _, roll6 := range rolls {
							p2 := (p2 + roll4 + roll5 + roll6) % 10
							s2 := s2 + p2 + 1
							if s2 >= 21 {
								p2wins++
								continue
							}
							// still no winner; recurse to next round
							w1, w2 := playPart2(p1, p2, s1, s2)
							p1wins += w1
							p2wins += w2
						}
					}
				}
			}
		}
	}
	memo[[4]int{p1, p2, s1, s2}] = [2]int{p1wins, p2wins}
	return
}

func main() {
	if false {
		input = `Player 1 starting position: 4
Player 2 starting position: 8`
	}
	ps := utils.ExtractInts(input)
	p1, p2 := ps[1]-1, ps[3]-1

	utils.Println(playPart1(p1, p2))
	utils.Println(utils.Max(playPart2(p1, p2, 0, 0)))
}
