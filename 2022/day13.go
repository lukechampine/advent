package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2022, 13)

func parse(line string) []interface{} {
	var p []interface{}
	json.Unmarshal([]byte(line), &p)
	return p
}

func cmp(left, right interface{}) int {
	li, lint := left.(float64)
	ri, rint := right.(float64)
	if lint && rint {
		return int(li - ri)
	}
	if lint {
		left = []interface{}{left}
	}
	if rint {
		right = []interface{}{right}
	}
	la := left.([]interface{})
	ra := right.([]interface{})
	for i := 0; i < len(la) && i < len(ra); i++ {
		c := cmp(la[i], ra[i])
		if c != 0 {
			return c
		}
	}
	return len(la) - len(ra)
}

func main() {
	n := 0
	var packets [][]interface{}
	for i, g := range strings.Split(input, "\n\n") {
		lines := utils.Lines(g)
		if cmp(parse(lines[0]), parse(lines[1])) < 0 {
			n += i + 1
		}
		packets = append(packets, parse(lines[0]), parse(lines[1]))
	}
	utils.Println(n)

	packets = append(packets, parse("[[2]]"), parse("[[6]]"))
	sort.Slice(packets, func(i, j int) bool {
		return cmp(packets[i], packets[j]) < 0
	})
	prod := 1
	for i, p := range packets {
		if s := fmt.Sprint(p); s == "[[2]]" || s == "[[6]]" {
			prod *= i + 1
		}
	}
	utils.Println(prod)
}
