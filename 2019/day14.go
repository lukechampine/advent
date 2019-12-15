package main

import (
	"sort"

	"github.com/lukechampine/advent/utils"
)

var input = utils.ReadInput("day14_input.txt")
var ils = utils.Lines(input)

type chemical struct {
	name   string
	amount int
}

type reaction struct {
	inputs []chemical
	output chemical
}

func produce(c chemical, excess map[string]int, rs map[string]reaction) int {
	if e := utils.Min(c.amount, excess[c.name]); e > 0 {
		excess[c.name] -= e
		c.amount -= e
	}
	if c.amount == 0 || c.name == "ORE" {
		return c.amount
	}
	// lookup inputs and reaction multiplier
	r := rs[c.name]
	m := c.amount / r.output.amount
	if c.amount%r.output.amount != 0 {
		m++
	}
	// produce each input, summing ore required
	var ore int
	for _, in := range r.inputs {
		in.amount *= m
		ore += produce(in, excess, rs)
	}
	excess[c.name] += (m * r.output.amount) - c.amount
	return ore
}

func main() {
	// part 1
	rs := make(map[string]reaction)
	for _, line := range ils {
		chems := utils.Split(line, " => ")
		ins, outs := chems[0], chems[1]
		var r reaction
		utils.Sscanf(outs, "%d %s", &r.output.amount, &r.output.name)
		for _, in := range utils.Split(ins, ", ") {
			var i chemical
			utils.Sscanf(in, "%d %s", &i.amount, &i.name)
			r.inputs = append(r.inputs, i)
		}
		rs[r.output.name] = r
	}
	excess := make(map[string]int)
	utils.Println(produce(chemical{"FUEL", 1}, excess, rs))

	// part 2
	fuel := sort.Search(100000000, func(i int) bool {
		return produce(chemical{"FUEL", i}, map[string]int{"ORE": 1e12}, rs) > 0
	}) - 1
	utils.Println(fuel)
}
