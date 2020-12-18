package main

import (
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2020, 18)

func eval(expr string, precedences map[string]int) int {
	toks := strings.Fields(utils.Replace(expr, "(", " ( ", ")", " ) "))
	_, n := evalExpr(toks, 0, precedences)
	return n
}

func evalExpr(expr []string, prec int, precedences map[string]int) (rest []string, n int) {
	if expr[0] == "(" {
		rest, n = evalExpr(expr[1:], 0, precedences)
		rest = rest[1:]
	} else {
		n = utils.Atoi(expr[0])
		rest = expr[1:]
	}
	for {
		if len(rest) == 0 || rest[0] == ")" {
			return rest, n
		}
		op := rest[0]
		if op != "+" && op != "*" {
			panic(op)
		}
		opPrec := precedences[op]
		if opPrec <= prec {
			return rest, n
		}
		var m int
		rest, m = evalExpr(rest[1:], opPrec, precedences)
		if op == "+" {
			n += m
		} else if op == "*" {
			n *= m
		}
	}
}

func main() {
	exprs := utils.Lines(input)
	utils.Println(utils.Sum(len(exprs), func(i int) int {
		return eval(exprs[i], map[string]int{"+": 1, "*": 1})
	}))
	utils.Println(utils.Sum(len(exprs), func(i int) int {
		return eval(exprs[i], map[string]int{"+": 2, "*": 1})
	}))
}
