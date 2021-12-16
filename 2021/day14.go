package main

import (
	"math"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2021, 14)

func applyRules(template map[string]int, rules map[string]string) map[string]int {
	t2 := make(map[string]int)
	for p, n := range template {
		if r, ok := rules[p]; ok {
			t2[p[:1]+r] += n
			t2[r+p[1:]] += n
		} else {
			t2[p] = n
		}
	}
	return t2
}

func delta(template map[string]int) int {
	cc := make(map[byte]int)
	for p, n := range template {
		cc[p[0]] += n
		cc[p[1]] += n
	}
	min, max := math.MaxInt64, math.MinInt64
	for _, n := range cc {
		max = utils.Max(max, n)
		min = utils.Min(min, n)
	}
	return (max-min)/2 + 1
}

func main() {
	parts := utils.Split(input, "\n\n")
	template := make(map[string]int)
	for i := 1; i < len(parts[0]); i++ {
		template[parts[0][i-1:i+1]]++
	}
	rules := make(map[string]string)
	for _, line := range utils.Lines(parts[1]) {
		parts := utils.Split(line, " -> ")
		rules[parts[0]] = parts[1]
	}
	for step := 1; step <= 40; step++ {
		template = applyRules(template, rules)
		if step == 10 || step == 40 {
			utils.Println(delta(template))
		}
	}
}
