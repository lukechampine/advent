package main

import (
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2020, 7)

type bag struct {
	color    string
	contains map[string]int
}

func countBags(color string, bags []bag) int {
	var b bag
	for _, b = range bags {
		if b.color == color {
			break
		}
	}
	n := 1 // this bag
	for c, per := range b.contains {
		n += per * countBags(c, bags)
	}
	return n
}

func main() {
	var bags []bag
	for _, l := range utils.Lines(input) {
		fs := strings.Split(l, " bags contain ")
		b := bag{
			color:    fs[0],
			contains: make(map[string]int),
		}
		if !strings.Contains(l, "no other bags") {
			for _, c := range strings.Split(fs[1], ", ") {
				fs := strings.Fields(c)
				b.contains[strings.Join(fs[1:3], " ")] = utils.Atoi(fs[0])
			}
		}
		bags = append(bags, b)
	}

	queue := []string{"shiny gold"}
	seen := make(map[string]bool)
	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]
		for _, b := range bags {
			if !seen[b.color] && b.contains[q] > 0 {
				queue = append(queue, b.color)
				seen[b.color] = true
			}
		}
	}
	utils.Println(len(seen))

	utils.Println(countBags("shiny gold", bags) - 1) // not counting shiny gold bag itself
}
