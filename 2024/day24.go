package main

import (
	"fmt"
	"sort"
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2024, 24)

type op struct {
	l, r string
	op   string
}

func main() {
	groups := utils.Split(input, "\n\n")
	wires := make(map[string]op)
	vals := make(map[string]int)
	for _, line := range utils.Lines(groups[0]) {
		vr, vl, _ := strings.Cut(line, ": ")
		vals[vr] = utils.Atoi(vl)
	}
	for _, line := range utils.Lines(groups[1]) {
		parts := strings.Fields(line)
		if parts[0] > parts[2] {
			parts[0], parts[2] = parts[2], parts[0]
		}
		wires[parts[4]] = op{parts[0], parts[2], parts[1]}
	}

	var rec func(string) int
	rec = func(wire string) int {
		if v, ok := vals[wire]; ok {
			return v
		}
		op := wires[wire]
		vals[wire] = map[string]func(int, int) int{
			"AND": func(l, r int) int { return l & r },
			"OR":  func(l, r int) int { return l | r },
			"XOR": func(l, r int) int { return l ^ r },
		}[op.op](rec(op.l), rec(op.r))
		return vals[wire]
	}
	var z int
	for i := 0; i <= 45; i++ {
		z |= rec(fmt.Sprintf("z%02d", i)) << i
	}
	utils.Println(z)

	// hideous and bespoke
	swapped := make(map[string]bool)
	for i := 2; i < 45; i++ {
		wire := fmt.Sprintf("z%02d", i)
		w := wires[wire]
		if w.op != "XOR" {
			swapped[wire] = true
			continue
		}
		l, r := wires[w.l], wires[w.r]
		if l.op > r.op {
			l, r = r, l
		}
		if l.op != "OR" {
			swapped[w.l] = true
		} else {
			w := l
			l, r := wires[w.l], wires[w.r]
			if l.op != "AND" {
				swapped[w.l] = true
			}
			if r.op != "AND" {
				swapped[w.r] = true
			}
		}
		if l.op == "OR" && r.op != "XOR" {
			swapped[w.r] = true
		}
	}
	var sorted []string
	for k := range swapped {
		sorted = append(sorted, k)
	}
	sort.Strings(sorted)
	utils.Println(strings.Join(sorted, ","))
}
