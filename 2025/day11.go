package main

import (
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2025, 11)

func main() {
	graph := make(map[string][]string)
	for _, line := range utils.Lines(input) {
		from, tos, _ := strings.Cut(line, ": ")
		graph[from] = strings.Fields(tos)
	}

	var part1 int
	var dfs func(node string)
	dfs = func(node string) {
		part1 += utils.BoolToInt(node == "out")
		for _, next := range graph[node] {
			dfs(next)
		}
	}
	dfs("you")
	utils.Println(part1)

	var countPaths func(start, end string, memo map[string]int) int
	countPaths = func(start, end string, memo map[string]int) int {
		if start == end {
			return 1
		} else if v, ok := memo[start]; ok {
			return v
		}
		var count int
		for _, next := range graph[start] {
			count += countPaths(next, end, memo)
		}
		memo[start] = count
		return count
	}
	svr2dac := countPaths("svr", "dac", make(map[string]int))
	svr2fft := countPaths("svr", "fft", make(map[string]int))
	dac2fft := countPaths("dac", "fft", make(map[string]int))
	fft2dac := countPaths("fft", "dac", make(map[string]int))
	dac2out := countPaths("dac", "out", make(map[string]int))
	fft2out := countPaths("fft", "out", make(map[string]int))
	utils.Println((svr2dac * dac2fft * fft2out) + (svr2fft * fft2dac * dac2out))
}
