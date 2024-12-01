package main

import (
	"sort"
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2023, 7)

const (
	highCard = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

func handRank(cards string) int {
	counts := make(map[int]int)
	for _, c := range "AKQJT98765432" {
		counts[strings.Count(cards, string(c))]++
	}
	switch {
	case counts[5] == 1:
		return fiveOfAKind
	case counts[4] == 1:
		return fourOfAKind
	case counts[3] == 1 && counts[2] == 1:
		return fullHouse
	case counts[3] == 1:
		return threeOfAKind
	case counts[2] == 2:
		return twoPair
	case counts[2] == 1:
		return onePair
	default:
		return highCard
	}
}

func beats(x, y string) bool {
	xt, yt := handRank(x), handRank(y)
	if xt != yt {
		return xt > yt
	}
	for i := range x {
		const order = "AKQJT98765432"
		xo, yo := strings.IndexByte(order, x[i]), strings.IndexByte(order, y[i])
		if xo != yo {
			return xo < yo
		}
	}
	return false
}

func handRank2(cards string) int {
	most := 'A'
	for _, c := range "AKQT98765432" {
		if strings.Count(cards, string(c)) > strings.Count(cards, string(most)) {
			most = c
		}
	}
	return handRank(strings.ReplaceAll(cards, "J", string(most)))
}

func beats2(x, y string) bool {
	xt, yt := handRank2(x), handRank2(y)
	if xt != yt {
		return xt > yt
	}
	for i := range x {
		const order = "AKQT98765432J"
		xo, yo := strings.IndexByte(order, x[i]), strings.IndexByte(order, y[i])
		if xo != yo {
			return xo < yo
		}
	}
	return false
}

func main() {
	hands := utils.Lines(input)
	sort.Slice(hands, func(i, j int) bool {
		icards := strings.Fields(hands[i])[0]
		jcards := strings.Fields(hands[j])[0]
		return !beats(icards, jcards)
	})
	utils.Println(utils.Sum(len(hands), func(i int) int {
		return (i + 1) * utils.Atoi(strings.Fields(hands[i])[1])
	}))

	sort.Slice(hands, func(i, j int) bool {
		icards := strings.Fields(hands[i])[0]
		jcards := strings.Fields(hands[j])[0]
		return !beats2(icards, jcards)
	})
	utils.Println(utils.Sum(len(hands), func(i int) int {
		return (i + 1) * utils.Atoi(strings.Fields(hands[i])[1])
	}))
}
