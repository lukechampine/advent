package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2021, 18)

type snailNumber struct {
	v           int
	left, right *snailNumber
}

func (sn snailNumber) isRegular() bool {
	return sn.left == nil
}

func (sn snailNumber) clone() snailNumber {
	if sn.isRegular() {
		return sn
	}
	l, r := sn.left.clone(), sn.right.clone()
	return snailNumber{left: &l, right: &r}
}

func (sn *snailNumber) tryExplode() bool {
	var rec func(sn *snailNumber, depth int, left, right *snailNumber) bool
	rec = func(sn *snailNumber, depth int, left, right *snailNumber) bool {
		if sn.isRegular() {
			return false
		} else if depth == 4 {
			if left != nil {
				left.v += sn.left.v
			}
			if right != nil {
				right.v += sn.right.v
			}
			*sn = snailNumber{v: 0}
			return true
		}
		leftOfRight := sn.right
		for !leftOfRight.isRegular() {
			leftOfRight = leftOfRight.left
		}
		rightOfLeft := sn.left
		for !rightOfLeft.isRegular() {
			rightOfLeft = rightOfLeft.right
		}
		return rec(sn.left, depth+1, left, leftOfRight) || rec(sn.right, depth+1, rightOfLeft, right)
	}

	return rec(sn, 0, nil, nil)
}

func (sn *snailNumber) trySplit() bool {
	if sn.isRegular() {
		if sn.v >= 10 {
			sn.left = &snailNumber{v: sn.v / 2}
			sn.right = &snailNumber{v: sn.v - sn.v/2}
			sn.v = 0
			return true
		}
		return false
	}
	return sn.left.trySplit() || sn.right.trySplit()
}

func (sn snailNumber) magnitude() int {
	if sn.isRegular() {
		return sn.v
	}
	return 3*sn.left.magnitude() + 2*sn.right.magnitude()
}

func add(x, y snailNumber) snailNumber {
	x = x.clone()
	y = y.clone()
	sum := snailNumber{left: &x, right: &y}
	for {
		if !sum.tryExplode() && !sum.trySplit() {
			return sum
		}
	}
}

func parse(line string) (snailNumber, string) {
	if line[0] == '[' {
		l, line := parse(line[1:])
		r, line := parse(line[1:])
		return snailNumber{left: &l, right: &r}, line[1:]
	}
	return snailNumber{v: utils.Atoi(line[:1])}, line[1:]
}

func main() {
	var sns []snailNumber
	for _, line := range utils.Lines(input) {
		sn, _ := parse(line)
		sns = append(sns, sn)
	}
	sum := sns[0]
	for _, sn := range sns[1:] {
		sum = add(sum, sn)
	}
	utils.Println(sum.magnitude())

	maxMag := 0
	for x := range sns {
		for y := range sns {
			if x == y {
				continue
			}
			sum := add(sns[x], sns[y])
			maxMag = utils.Max(maxMag, sum.magnitude())
		}
	}
	utils.Println(maxMag)
}
