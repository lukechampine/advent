package main

import (
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2021, 12)
var inputLines = utils.Lines(input)

func dfs(n string, edges map[string]map[string]struct{}, path []string, canDouble bool) (paths int) {
	if n == "end" {
		return 1
	}
	timesVisited := func(e string) (n int) {
		for _, v := range path {
			if v == e {
				n++
			}
		}
		return n
	}
	for e := range edges[n] {
		if e == "start" {
			continue
		}
		if t := timesVisited(e); utils.IsUpper(e) || t == 0 || (t == 1 && canDouble) {
			paths += dfs(e, edges, append(path, n), canDouble && (utils.IsUpper(e) || t == 0))
		}
	}
	return
}

func main() {
	edges := make(map[string]map[string]struct{})
	for _, line := range inputLines {
		parts := strings.Split(line, "-")
		if _, ok := edges[parts[0]]; !ok {
			edges[parts[0]] = make(map[string]struct{})
		}
		if _, ok := edges[parts[1]]; !ok {
			edges[parts[1]] = make(map[string]struct{})
		}
		edges[parts[0]][parts[1]] = struct{}{}
		edges[parts[1]][parts[0]] = struct{}{}
	}
	utils.Println(dfs("start", edges, nil, false))
	utils.Println(dfs("start", edges, nil, true))
}
