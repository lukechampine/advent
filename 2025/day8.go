package main

import (
	"math"
	"sort"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2025, 8)

func dist(a, b [3]int) float64 {
	d := 0.0
	for i := range a {
		d += float64((a[i] - b[i]) * (a[i] - b[i]))
	}
	return math.Sqrt(d)
}

func main() {
	var boxes [][3]int
	for _, line := range utils.Lines(input) {
		var box [3]int
		copy(box[:], utils.ExtractInts(line))
		boxes = append(boxes, box)
	}
	var pairs [][2][3]int
	for i := range boxes {
		for j := i + 1; j < len(boxes); j++ {
			pairs = append(pairs, [2][3]int{boxes[i], boxes[j]})
		}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return dist(pairs[i][0], pairs[i][1]) < dist(pairs[j][0], pairs[j][1])
	})

	var circuits [][][3]int
	index := func(box [3]int) int {
		for i, circuit := range circuits {
			for _, b := range circuit {
				if b == box {
					return i
				}
			}
		}
		return -1
	}
	for i, p := range pairs {
		a, b := p[0], p[1]
		ia, ib := index(a), index(b)
		switch {
		case ia == -1 && ib == -1:
			circuits = append(circuits, [][3]int{a, b})
		case ia != -1 && ib != -1 && ia != ib:
			// merge b into a
			circuits[ia] = append(circuits[ia], circuits[ib]...)
			circuits = append(circuits[:ib], circuits[ib+1:]...)
		case ia != -1 && ib == -1:
			circuits[ia] = append(circuits[ia], b)
		case ia == -1 && ib != -1:
			circuits[ib] = append(circuits[ib], a)
		}
		if i == 999 {
			sizes := make([]int, len(circuits))
			for i := range circuits {
				sizes[i] = len(circuits[i])
			}
			sort.Sort(sort.Reverse(sort.IntSlice(sizes)))
			utils.Println(sizes[0] * sizes[1] * sizes[2])
		}

		if len(circuits) == 1 && len(circuits[0]) == len(boxes) {
			utils.Println(p[0][0] * p[1][0])
			break
		}
	}
}
