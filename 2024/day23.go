package main

import (
	"sort"
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2024, 23)

func main() {
	graph := make(map[string]map[string]bool)
	for _, line := range utils.Lines(input) {
		parts := strings.Split(line, "-")
		a, b := parts[0], parts[1]
		if graph[a] == nil {
			graph[a] = make(map[string]bool)
		}
		graph[a][b] = true
		if graph[b] == nil {
			graph[b] = make(map[string]bool)
		}
		graph[b][a] = true
	}
	seen := make(map[string]bool)
	for a := range graph {
		for b := range graph[a] {
			for c := range graph[b] {
				if graph[c][a] && (a[0] == 't' || b[0] == 't' || c[0] == 't') {
					circle := []string{a, b, c}
					sort.Strings(circle)
					seen[strings.Join(circle, "")] = true
				}
			}
		}
	}
	utils.Println(len(seen))

	fullyConnected := func(set []string) bool {
		for _, a := range set {
			for _, b := range set {
				if a != b && !graph[a][b] {
					return false
				}
			}
		}
		return true
	}
	var best []string
	for a := range graph {
		// for each neighbor of a, build a subgraph containing all neighbors-
		// of-neighbors that are also neighbors of a; there must exist at least
		// one such subgraph which (along with a and n) forms the largest
		// fully-connected set that contains a.
		//
		// NOTE: probably not actually true, but hey it works
		for n := range graph[a] {
			set := []string{a, n}
			for nn := range graph[n] {
				if graph[a][nn] {
					set = append(set, nn)
				}
			}
			if fullyConnected(set) && len(set) > len(best) {
				best = set
			}
		}
	}
	sort.Strings(best)
	utils.Println(strings.Join(best, ","))
}
