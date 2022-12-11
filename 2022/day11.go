package main

import (
	"sort"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2022, 11)

type monkey struct {
	items []int
	op    func(int) int
	test  func(int) int
}

func main() {
	var monkeys []monkey
	modulus := 1
	for _, group := range utils.Split(input, "\n\n") {
		lines := utils.Lines(group)
		op := lines[2][23]
		operand := lines[2][25:]
		testN := utils.ExtractInts(lines[3])[0]
		modulus *= testN
		testTrue := utils.ExtractInts(lines[4])[0]
		testFalse := utils.ExtractInts(lines[5])[0]
		m := monkey{
			items: utils.ExtractInts(lines[1]),
			op: func(old int) int {
				switch op {
				case '+':
					if operand == "old" {
						return old + old
					}
					return old + utils.Atoi(operand)
				case '*':
					if operand == "old" {
						return old * old
					}
					return old * utils.Atoi(operand)
				default:
					panic(string(op))
				}
			},
			test: func(item int) int {
				if item%testN == 0 {
					return testTrue
				}
				return testFalse
			},
		}
		monkeys = append(monkeys, m)
	}
	saved := append([]monkey(nil), monkeys...)

	counts := make([]int, len(monkeys))
	for round := 0; round < 20; round++ {
		for i, m := range monkeys {
			for _, item := range m.items {
				item = m.op(item) / 3
				dst := m.test(item)
				monkeys[dst].items = append(monkeys[dst].items, item)
				counts[i]++
			}
			monkeys[i].items = nil
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))
	utils.Println(counts[0] * counts[1])

	monkeys = saved
	counts = make([]int, len(monkeys))
	for round := 0; round < 10000; round++ {
		for i, m := range monkeys {
			for _, item := range m.items {
				item = m.op(item) % modulus
				dst := m.test(item)
				monkeys[dst].items = append(monkeys[dst].items, item)
				counts[i]++
			}
			monkeys[i].items = nil
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))
	utils.Println(counts[0] * counts[1])
}
