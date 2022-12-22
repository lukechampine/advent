package main

import (
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2022, 21)

type node struct {
	name string
	op   string
	args []string
}

func compute(name string, nodes map[string]node) int {
	n := nodes[name]
	switch n.op {
	case "set":
		return utils.Atoi(n.args[0])
	case "+":
		return compute(n.args[0], nodes) + compute(n.args[1], nodes)
	case "-":
		return compute(n.args[0], nodes) - compute(n.args[1], nodes)
	case "*":
		return compute(n.args[0], nodes) * compute(n.args[1], nodes)
	case "/":
		return compute(n.args[0], nodes) / compute(n.args[1], nodes)
	}
	panic("unreachable")
}

func derive(name string, inTermsOf string, nodes map[string]node) (int, func(int) int) {
	n := nodes[name]
	if name == inTermsOf {
		return 0, func(x int) int { return x }
	}
	if n.op == "set" {
		return utils.Atoi(n.args[0]), nil
	}

	l, lfn := derive(n.args[0], inTermsOf, nodes)
	r, rfn := derive(n.args[1], inTermsOf, nodes)
	if lfn == nil && rfn == nil {
		switch n.op {
		case "+":
			return l + r, nil
		case "-":
			return l - r, nil
		case "*":
			return l * r, nil
		case "/":
			return l / r, nil
		}
	}
	if lfn != nil {
		switch n.op {
		case "+":
			return 0, func(x int) int { return lfn(x - r) }
		case "-":
			return 0, func(x int) int { return lfn(x + r) }
		case "*":
			return 0, func(x int) int { return lfn(x / r) }
		case "/":
			return 0, func(x int) int { return lfn(x * r) }
		}
	} else if rfn != nil {
		switch n.op {
		case "+":
			return 0, func(x int) int { return rfn(x - l) }
		case "-":
			return 0, func(x int) int { return rfn(l - x) }
		case "*":
			return 0, func(x int) int { return rfn(x / l) }
		case "/":
			return 0, func(x int) int { return rfn(r / x) }
		}
	}
	panic("unreachable")
}

func main() {
	nodes := make(map[string]node)
	for _, line := range utils.Lines(input) {
		parts := strings.Fields(line)
		n := node{name: parts[0][:len(parts[0])-1]}
		if len(parts) == 2 {
			n.op = "set"
			n.args = parts[1:]
		} else {
			parts[1], parts[2] = parts[2], parts[1]
			n.op = parts[1]
			n.args = parts[2:]
		}
		nodes[n.name] = n
	}
	utils.Println(compute("root", nodes))

	root := nodes["root"]
	l, lfn := derive(root.args[0], "humn", nodes)
	r, rfn := derive(root.args[1], "humn", nodes)
	if lfn != nil {
		utils.Println(lfn(r))
	} else {
		utils.Println(rfn(l))
	}
}
