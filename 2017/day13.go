package main

import (
	"github.com/lukechampine/advent/utils"
)

const input = `0: 3
1: 2
2: 6
4: 4
6: 4
8: 8
10: 9
12: 8
14: 5
16: 6
18: 8
20: 6
22: 12
24: 6
26: 12
28: 8
30: 8
32: 10
34: 12
36: 12
38: 8
40: 12
42: 12
44: 14
46: 12
48: 14
50: 12
52: 12
54: 12
56: 10
58: 14
60: 14
62: 14
64: 14
66: 17
68: 14
72: 14
76: 14
80: 14
82: 14
88: 18
92: 14
98: 18`

func parse(s string) map[int]int {
	m := make(map[int]int)
	for _, line := range utils.Lines(s) {
		var x, y int
		utils.Sscanf(line, "%d: %d", &x, &y)
		m[x] = y
	}
	return m
}

func main() {
	// part 1
	scanners := parse(input)
	numLayers := 0
	for l := range scanners {
		numLayers = utils.Max(numLayers, l+1)
	}
	layers := make([]int, numLayers)
	// intialize scanners
	for i := range layers {
		if _, ok := scanners[i]; !ok {
			layers[i] = -1 // no scanner
		}
	}

	severity := 0
	for i := range layers {
		if layers[i] == 0 {
			// caught
			severity += scanners[i] * i
		}
		// advance all scanners
		for l, r := range scanners {
			if i%(2*(r-1)) < r-1 {
				layers[l]++
			} else {
				layers[l]--
			}
		}
	}
	i := len(layers) - 1
	if layers[i] == 0 {
		severity += scanners[i] * i
	}
	utils.Println(severity)

	// part 2
	layers = make([]int, numLayers)
	for i := range layers {
		if _, ok := scanners[i]; !ok {
			layers[i] = -1 // no scanner
		}
	}
	picoLayers := make([][]int, 5e6)
	for i := range picoLayers {
		picoLayers[i] = append([]int(nil), layers...)
		for l, r := range scanners {
			if i%(2*(r-1)) < r-1 {
				layers[l]++
			} else {
				layers[l]--
			}
		}
	}

outer:
	for pico := range picoLayers {
		for i := 0; i < numLayers; i++ {
			if picoLayers[pico+i][i] == 0 {
				// caught
				continue outer
			}
		}
		utils.Println(pico)
		return
	}
}
