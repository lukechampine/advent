package main

import (
	"github.com/lukechampine/advent/utils"
)

var input = utils.ReadInput("day6_input.txt")
var ils = utils.Lines(input)

func recOrbits(o string, orbits map[string]string) int {
	if _, ok := orbits[o]; !ok {
		return 0
	}
	return 1 + recOrbits(orbits[o], orbits)
}

func orbitPath(o string, orbits map[string]string) []string {
	if _, ok := orbits[o]; !ok {
		return nil
	}
	return append(orbitPath(orbits[o], orbits), o)
}

func main() {
	// part 1
	orbits := make(map[string]string)
	for _, line := range ils {
		xs := utils.Split(line, ")")
		orbits[xs[1]] = xs[0]
	}
	var n int
	for o := range orbits {
		n += recOrbits(o, orbits)
	}
	utils.Println(n)

	// part 2
	youPath := orbitPath("YOU", orbits)
	sanPath := orbitPath("SAN", orbits)

	for i := range youPath {
		if youPath[i] != sanPath[i] {
			utils.Println(len(youPath[i+1:]) + len(sanPath[i+1:]))
			return
		}
	}
}
