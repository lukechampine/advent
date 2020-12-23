package main

import (
	"container/ring"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2020, 23)

func locate(m []*ring.Ring, label int, pick *ring.Ring) *ring.Ring {
	picked := make([]int, 3)
	for i := range picked {
		picked[i] = pick.Value.(int)
		pick = pick.Next()
	}

again:
	if label == 0 {
		label = len(m) - 1
	}
	if label == picked[0] || label == picked[1] || label == picked[2] {
		label--
		goto again
	}
	return m[label]
}

func simulate(nums []int, turns int) *ring.Ring {
	m := make([]*ring.Ring, len(nums)+1) // for O(1) lookups
	r := ring.New(len(nums))
	for _, n := range nums {
		r.Value = n
		m[n] = r
		r = r.Next()
	}
	for turn := 0; turn < turns; turn++ {
		pick := r.Unlink(3)
		dest := locate(m, r.Value.(int)-1, pick)
		pick.Prev().Link(dest.Next())
		dest.Link(pick)
		r = r.Next()
	}
	return m[1]
}

func main() {
	nums := utils.Digits(input)
	r := simulate(nums, 100)
	for r = r.Next(); r.Value.(int) != 1; r = r.Next() {
		utils.Print(r.Value)
	}
	utils.Println()

	for i := len(nums) + 1; i <= 1e6; i++ {
		nums = append(nums, i)
	}
	r = simulate(nums, 10e6)
	utils.Println(r.Next().Value.(int) * r.Next().Next().Value.(int))
}
