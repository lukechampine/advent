package main

import (
	"sort"
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2022, 16)

type valve struct {
	name  string
	rate  int
	conns []string
}

type entry struct {
	cur      string
	ele      string
	t        int
	rate     int
	pressure int
	open     uint16
}

func bfs(valves map[string]valve, queue []entry, expand func(entry, []entry) []entry) int {
	seen := make(map[entry]bool)
	var best entry
	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]
		if seen[q] {
			continue
		}
		seen[q] = true
		if q.t == 30 {
			if q.pressure > best.pressure {
				best = q
			}
			continue
		}
		queue = expand(q, queue)
		// only keep the most promising 50,000
		if len(queue) > 100000 {
			sort.Slice(queue, func(i, j int) bool {
				pi := queue[i].pressure + queue[i].rate*(30-queue[i].t)
				pj := queue[j].pressure + queue[j].rate*(30-queue[j].t)
				return pi > pj
			})
			queue = queue[:len(queue)/2]
		}
	}
	return best.pressure
}

func main() {
	valves := make(map[string]valve)
	for _, line := range utils.Lines(input) {
		var name string
		var rate int
		utils.Sscanf(line, "Valve %s has flow rate=%d", &name, &rate)
		conns := strings.Split(line, ", ")
		conns[0] = conns[0][len(conns[0])-2:]
		valves[name] = valve{name, rate, conns}
	}

	valveMask := make(map[string]int)
	for _, v := range valves {
		if v.rate > 0 {
			valveMask[v.name] = len(valveMask)
		}
	}
	canOpen := func(u uint16, v string) bool {
		i, ok := valveMask[v]
		return ok && u&(1<<i) == 0
	}
	open := func(u uint16, v string) uint16 {
		return u | (1 << valveMask[v])
	}

	expand1 := func(e entry, queue []entry) []entry {
		if canOpen(e.open, e.cur) {
			queue = append(queue, entry{
				cur:      e.cur,
				t:        e.t + 1,
				rate:     e.rate + valves[e.cur].rate,
				pressure: e.pressure + e.rate,
				open:     open(e.open, e.cur),
			})
		}
		for _, c := range valves[e.cur].conns {
			queue = append(queue, entry{
				cur:      c,
				t:        e.t + 1,
				rate:     e.rate,
				pressure: e.pressure + e.rate,
				open:     e.open,
			})
		}
		return queue
	}
	utils.Println(bfs(valves, []entry{{cur: "AA"}}, expand1))

	// now, with an elephant...
	expand2 := func(e entry, queue []entry) []entry {
		for _, o := range expand1(e, nil) {
			if canOpen(o.open, e.ele) {
				queue = append(queue, entry{
					cur:      o.cur,
					rate:     o.rate + valves[e.ele].rate,
					open:     open(o.open, e.ele),
					ele:      e.ele,
					t:        e.t + 1,
					pressure: e.pressure + e.rate,
				})
			}
			for _, c := range valves[e.ele].conns {
				queue = append(queue, entry{
					cur:      o.cur,
					rate:     o.rate,
					open:     o.open,
					ele:      c,
					t:        e.t + 1,
					pressure: e.pressure + e.rate,
				})
			}
		}
		return queue
	}
	utils.Println(bfs(valves, []entry{{cur: "AA", ele: "AA", t: 4}}, expand2))
}
