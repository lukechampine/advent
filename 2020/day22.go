package main

import (
	"fmt"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2020, 22)

func play(p1, p2 []int) ([]int, []int) {
	for len(p1) > 0 && len(p2) > 0 {
		d1, d2 := p1[0], p2[0]
		p1, p2 = p1[1:], p2[1:]
		if d1 > d2 {
			p1 = append(p1, d1, d2)
		} else {
			p2 = append(p2, d2, d1)
		}
	}
	return p1, p2
}

func playrec(p1, p2 []int) ([]int, []int) {
	seen := make(map[string]bool)
	for len(p1) > 0 && len(p2) > 0 {
		id := fmt.Sprint(p1) + fmt.Sprint(p2)
		if seen[id] {
			return p1, nil
		}
		seen[id] = true

		d1, d2 := p1[0], p2[0]
		p1, p2 = p1[1:], p2[1:]

		var p1wins bool
		if len(p1) >= d1 && len(p2) >= d2 {
			rp1, _ := playrec(p1[:d1:d1], p2[:d2:d2])
			p1wins = len(rp1) > 0
		} else {
			p1wins = d1 > d2
		}
		if p1wins {
			p1 = append(p1, d1, d2)
		} else {
			p2 = append(p2, d2, d1)
		}
	}
	return p1, p2
}

func score(p1, p2 []int) (n int) {
	for i, c := range append(p1, p2...) { // weird, but it works
		n += c * (len(p1) - i)
	}
	return
}

func main() {
	groups := utils.Split(input, "\n\n")
	p1 := utils.ExtractInts(groups[0])[1:]
	p2 := utils.ExtractInts(groups[1])[1:]

	utils.Println(score(play(p1, p2)))
	utils.Println(score(playrec(p1, p2)))
}
