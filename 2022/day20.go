package main

import (
	"container/ring"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2022, 20)

func mix(ps []*ring.Ring) {
	for _, p := range ps {
		p = p.Prev()
		v := p.Unlink(1)
		p.Move(v.Value.(int) % (len(ps) - 1)).Link(v)
	}
}

func grove(r *ring.Ring) (g int) {
	for r.Value.(int) != 0 {
		r = r.Next()
	}
	for i := 0; i < 3; i++ {
		r = r.Move(1000)
		g += r.Value.(int)
	}
	return
}

func main() {
	ns := utils.ExtractInts(input)
	r := ring.New(len(ns))
	ps := make([]*ring.Ring, r.Len())
	for i, n := range ns {
		r.Value = n
		ps[i] = r
		r = r.Next()
	}
	mix(ps)
	utils.Println(grove(r))

	for i, n := range ns {
		r.Value = n * 811589153
		ps[i] = r
		r = r.Next()
	}
	for i := 0; i < 10; i++ {
		mix(ps)
	}
	utils.Println(grove(r))
}
