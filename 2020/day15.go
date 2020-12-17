package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2020, 15)

func main() {
	m := make(map[int][]int)
	var i, n int
	for i, n = range utils.ExtractInts(input) {
		m[n] = []int{i + 1}
	}
	for i++; i < 2020; i++ {
		if m[n] = append(m[n], i); len(m[n]) >= 2 {
			n, m[n] = m[n][1]-m[n][0], m[n][1:]
		} else {
			n = 0
		}
	}
	utils.Println(n)

	for ; i < 30000000; i++ {
		if m[n] = append(m[n], i); len(m[n]) >= 2 {
			n, m[n] = m[n][1]-m[n][0], m[n][1:]
		} else {
			n = 0
		}
	}
	utils.Println(n)
}
