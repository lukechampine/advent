package main

import (
	"fmt"
	"math/rand"
	"strings"
)

const input = `Tristram to AlphaCentauri = 34
Tristram to Snowdin = 100
Tristram to Tambi = 63
Tristram to Faerun = 108
Tristram to Norrath = 111
Tristram to Straylight = 89
Tristram to Arbre = 132
AlphaCentauri to Snowdin = 4
AlphaCentauri to Tambi = 79
AlphaCentauri to Faerun = 44
AlphaCentauri to Norrath = 147
AlphaCentauri to Straylight = 133
AlphaCentauri to Arbre = 74
Snowdin to Tambi = 105
Snowdin to Faerun = 95
Snowdin to Norrath = 48
Snowdin to Straylight = 88
Snowdin to Arbre = 7
Tambi to Faerun = 68
Tambi to Norrath = 134
Tambi to Straylight = 107
Tambi to Arbre = 40
Faerun to Norrath = 11
Faerun to Straylight = 66
Faerun to Arbre = 144
Norrath to Straylight = 115
Norrath to Arbre = 135
Straylight to Arbre = 127`

type route struct {
	from, to string
	dist     int
}

func parse(str string) route {
	var r route
	fmt.Sscanf(str, "%s to %s = %d", &r.from, &r.to, &r.dist)
	return r
}

func main() {
	// part 1
	cities := make(map[string]map[string]int)
	for _, str := range strings.Split(input, "\n") {
		r := parse(str)
		if _, ok := cities[r.from]; !ok {
			cities[r.from] = make(map[string]int)
		}
		if _, ok := cities[r.to]; !ok {
			cities[r.to] = make(map[string]int)
		}
		cities[r.from][r.to] = r.dist
		cities[r.to][r.from] = r.dist
	}
	var names []string
	for name := range cities {
		names = append(names, name)
	}

	var shortest int = 999999999
	for n := 0; n < 50000; n++ {
		var dist int
		order := rand.Perm(len(names))
		for i := 0; i < len(order)-1; i++ {
			from, to := names[order[i]], names[order[i+1]]
			dist += cities[from][to]
		}
		if dist < shortest {
			shortest = dist
		}
	}
	fmt.Println(shortest)

	// part 2
	var longest int
	for n := 0; n < 50000; n++ {
		var dist int
		order := rand.Perm(len(names))
		for i := 0; i < len(order)-1; i++ {
			from, to := names[order[i]], names[order[i+1]]
			dist += cities[from][to]
		}
		if dist > longest {
			longest = dist
		}
	}
	fmt.Println(longest)
}
