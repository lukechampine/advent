package main

import (
	"sort"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2022, 15)

func excludedRange(s, b utils.Pos, y int) [2]int {
	d := s.Dist(b) - s.Dist(utils.Pos{X: s.X, Y: y})
	if d < 0 {
		return [2]int{0, 0}
	}
	return [2]int{s.X - d, s.X + d}
}

func mergeRanges(rs [][2]int) [][2]int {
	sort.Slice(rs, func(i, j int) bool {
		return rs[i][0] < rs[j][0]
	})
	merged := [][2]int{rs[0]}
	for _, r := range rs[1:] {
		last := &merged[len(merged)-1]
		if last[1] < r[0] {
			merged = append(merged, r)
		} else {
			last[1] = utils.Max(last[1], r[1])
		}
	}
	return merged
}

func excludedRanges(sensors map[utils.Pos]utils.Pos, y int) [][2]int {
	var rs [][2]int
	for s, b := range sensors {
		rs = append(rs, excludedRange(s, b, y))
	}
	return mergeRanges(rs)
}

func main() {
	sensors := make(map[utils.Pos]utils.Pos)
	for _, line := range utils.Lines(input) {
		var s, b utils.Pos
		utils.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &s.X, &s.Y, &b.X, &b.Y)
		sensors[s] = b
	}

	ex := excludedRanges(sensors, 2000000)
	utils.Println(ex[0][1] - ex[0][0])

	for y := 0; y < 4000000; y++ {
		ex := excludedRanges(sensors, y)
		if len(ex) == 2 {
			utils.Println((ex[0][1]+1)*4000000 + y)
			break
		}
	}
}
