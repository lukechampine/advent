package main

import (
	"sort"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2020, 10)
var ads = utils.ExtractInts(input)

func countPaths(cur int, ads []int, memo map[int]int) (total int) {
	if len(ads) == 0 {
		return 1 // complete path
	}
	for _, next := range ads {
		ads = ads[1:]
		if next-cur > 3 {
			break
		}
		if _, ok := memo[next]; !ok {
			memo[next] = countPaths(next, ads, memo)
		}
		total += memo[next]
	}
	return total
}

func main() {
	sort.Ints(ads)
	dist := make(map[int]int)
	for i := range ads[1:] {
		dist[ads[i+1]-ads[i]]++
	}
	dist[1]++
	dist[3]++
	utils.Println(dist[1] * dist[3])

	utils.Println(countPaths(0, ads, make(map[int]int)))
}
