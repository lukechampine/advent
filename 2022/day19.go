package main

import (
	"sort"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2022, 19)

func maxGeodes(oreCost, clayCost, obsidianOreCost, obsidianClayCost, geodeOreCost, geodeObsidianCost, timeLimit int) int {
	type entry struct {
		ore, clay, obsidian, geodes,
		oreBots, clayBots, obsidianBots, geodeBots,
		t int
	}
	heuristic := func(q entry) int {
		obsidianCost := obsidianOreCost + obsidianClayCost*clayCost
		geodeCost := geodeOreCost + geodeObsidianCost*obsidianCost
		if q.oreBots > geodeCost || q.clayBots > obsidianClayCost || q.obsidianBots > geodeObsidianCost {
			return -1 // excessive
		}
		return q.oreBots + q.clayBots*clayCost + q.obsidianBots*obsidianCost + q.geodeBots*geodeCost
	}
	queue := []entry{{oreBots: 1}}
	seen := make(map[entry]bool)
	var best int
	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]
		if seen[q] {
			continue
		}
		seen[q] = true
		if q.t == timeLimit {
			if q.geodes > best {
				best = q.geodes
			}
			continue
		}
		canBuildOreBot := q.ore >= oreCost
		canBuildClayBot := q.ore >= clayCost
		canBuildObsidianBot := q.ore >= obsidianOreCost && q.clay >= obsidianClayCost
		canBuildGeodeBot := q.ore >= geodeOreCost && q.obsidian >= geodeObsidianCost
		q.ore += q.oreBots
		q.clay += q.clayBots
		q.obsidian += q.obsidianBots
		q.geodes += q.geodeBots
		q.t++
		queue = append(queue, q)
		if canBuildOreBot {
			e := q
			e.oreBots++
			e.ore -= oreCost
			queue = append(queue, e)
		}
		if canBuildClayBot {
			e := q
			e.clayBots++
			e.ore -= clayCost
			queue = append(queue, e)
		}
		if canBuildObsidianBot {
			e := q
			e.obsidianBots++
			e.ore -= obsidianOreCost
			e.clay -= obsidianClayCost
			queue = append(queue, e)
		}
		if canBuildGeodeBot {
			e := q
			e.geodeBots++
			e.ore -= geodeOreCost
			e.obsidian -= geodeObsidianCost
			queue = append(queue, e)
		}
		if len(queue) > 200000 {
			sort.Slice(queue, func(i, j int) bool {
				return heuristic(queue[i]) > heuristic(queue[j])
			})
			queue = queue[:20000]
		}
	}
	return best
}

func main() {
	sum := 0
	for _, line := range utils.Lines(input) {
		ints := utils.ExtractInts(line)
		sum += ints[0] * maxGeodes(ints[1], ints[2], ints[3], ints[4], ints[5], ints[6], 24)
	}
	utils.Println(sum)

	prod := 1
	for _, line := range utils.Lines(input)[:3] {
		ints := utils.ExtractInts(line)
		prod *= maxGeodes(ints[1], ints[2], ints[3], ints[4], ints[5], ints[6], 32)
	}
	utils.Println(prod)
}
