package main

import (
	"strconv"
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2020, 19)

type rule struct {
	left, right []int
	txt         string
}

func satisfies(msg string, rules map[int]rule) bool {
	rem, ok := recsatisfies(msg, rules[0], rules)
	return ok && rem == ""
}

func recsatisfies(msg string, r rule, rules map[int]rule) (string, bool) {
	if r.txt != "" {
		return strings.TrimPrefix(msg, r.txt), strings.HasPrefix(msg, r.txt)
	}
	lrem := msg
	leftOk := utils.All(len(r.left), func(i int) (ok bool) {
		lrem, ok = recsatisfies(lrem, rules[r.left[i]], rules)
		return
	})
	if leftOk || len(r.right) == 0 {
		return lrem, leftOk
	}
	rrem := msg
	rightOk := utils.All(len(r.right), func(i int) (ok bool) {
		rrem, ok = recsatisfies(rrem, rules[r.right[i]], rules)
		return
	})
	return rrem, rightOk
}

func satisfies2(msg string, rules map[int]rule) bool {
	// The new rules are:
	//
	//  8: 42 | 8             "one or more 42s"
	// 11: 42 31 | 42 11 31   "one or more 42s followed by the same number of 31s"
	//
	// Putting these together, rule 0 is:
	//
	//  0: 8 11               "n 42s followed by m 31s, where n >= 2, m >= 1, m < n"

	num42 := 0
	for {
		rem, ok := recsatisfies(msg, rules[42], rules)
		if !ok {
			break
		}
		msg = rem
		num42++
	}
	num31 := 0
	for {
		rem, ok := recsatisfies(msg, rules[31], rules)
		if !ok {
			break
		}
		msg = rem
		num31++
	}
	return num42 >= 2 && num31 >= 1 && num31 < num42 && msg == ""
}

func main() {
	parts := utils.Split(input, "\n\n")
	msgs := utils.Lines(parts[1])

	rules := make(map[int]rule)
	for _, l := range utils.Lines(parts[0]) {
		fs := strings.Split(l, ":")
		ruleNum := utils.Atoi(fs[0])
		var r rule
		if body := strings.TrimSpace(fs[1]); strings.Contains(body, `"`) {
			r.txt, _ = strconv.Unquote(body)
		} else if parts := strings.Split(body, "|"); len(parts) > 1 {
			r.left = utils.ExtractInts(parts[0])
			r.right = utils.ExtractInts(parts[1])
		} else {
			r.left = utils.ExtractInts(body)
		}
		rules[ruleNum] = r
	}

	utils.Println(utils.Count(len(msgs), func(i int) bool {
		return satisfies(msgs[i], rules)
	}))

	utils.Println(utils.Count(len(msgs), func(i int) bool {
		return satisfies2(msgs[i], rules)
	}))
}
