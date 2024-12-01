package main

import (
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2023, 19)

type rules []string

type part map[string]int

func satisfies(p part, rule string) (string, bool) {
	if cond, dst, ok := strings.Cut(rule, ":"); !ok {
		return rule, true
	} else if rating, limit, ok := strings.Cut(cond, "<"); ok {
		return dst, p[rating] < utils.Atoi(limit)
	} else if rating, limit, ok := strings.Cut(cond, ">"); ok {
		return dst, p[rating] > utils.Atoi(limit)
	}
	panic("unreachable")
}

func accepted(p part, ws map[string]rules) (ok bool) {
	label := "in"
next:
	for _, r := range ws[label] {
		if label, ok = satisfies(p, r); ok {
			goto next
		}
	}
	return label == "A"
}

func combinations(label string, bounds [4][2]int, ws map[string]rules) (n int) {
	if label == "A" {
		n = 1
		for _, b := range bounds {
			n *= b[1] - b[0] - 1
		}
		return n
	} else if label == "R" {
		return 0
	}

	for _, rule := range ws[label] {
		if cond, dst, ok := strings.Cut(rule, ":"); !ok {
			n += combinations(rule, bounds, ws)
		} else if rating, limit, ok := strings.Cut(cond, "<"); ok {
			i := strings.Index("xmas", rating)
			subbounds := bounds
			subbounds[i][1] = utils.Min(subbounds[i][1], utils.Atoi(limit))
			n += combinations(dst, subbounds, ws)
			bounds[i][0] = utils.Max(bounds[i][0], utils.Atoi(limit)-1)
		} else if rating, limit, ok := strings.Cut(cond, ">"); ok {
			i := strings.Index("xmas", rating)
			subbounds := bounds
			subbounds[i][0] = utils.Max(subbounds[i][0], utils.Atoi(limit))
			n += combinations(dst, subbounds, ws)
			bounds[i][1] = utils.Min(bounds[i][1], utils.Atoi(limit)+1)
		}
	}
	return
}

func main() {
	workflows, ratings, _ := strings.Cut(input, "\n\n")
	ws := make(map[string]rules)
	for _, line := range utils.Lines(workflows) {
		name, rest, _ := strings.Cut(line, "{")
		rules := strings.Split(strings.Trim(rest, "}"), ",")
		ws[name] = rules
	}
	var sum int
	for _, line := range utils.Lines(ratings) {
		ints := utils.ExtractInts(line)
		p := part{"x": ints[0], "m": ints[1], "a": ints[2], "s": ints[3]}
		if accepted(p, ws) {
			sum += utils.IntSum(ints)
		}
	}
	utils.Println(sum)

	var bounds [4][2]int
	for i := range bounds {
		bounds[i] = [2]int{0, 4001}
	}
	utils.Println(combinations("in", bounds, ws))
}
