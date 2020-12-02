package main

import (
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2020, 2)
var inputInts = utils.ExtractInts(input)

type policy struct {
	L, U int
	Ch   byte
	Pw   string
}

func main() {
	var policies []policy
	utils.Parse(&policies, "%d-%d %c: %s", input)

	utils.Println(utils.Count(len(policies), func(i int) bool {
		p := policies[i]
		n := strings.Count(p.Pw, string(p.Ch))
		return p.L <= n && n <= p.U
	}))

	utils.Println(utils.Count(len(policies), func(i int) bool {
		p := policies[i]
		return (p.Pw[p.L-1] == p.Ch) != (p.Pw[p.U-1] == p.Ch)
	}))
}
