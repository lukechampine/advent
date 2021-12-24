package main

import (
	"sort"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2021, 19)

type Vec3 struct {
	X, Y, Z int
}

func (v Vec3) Add(o Vec3) Vec3 {
	return Vec3{v.X + o.X, v.Y + o.Y, v.Z + o.Z}
}

func (v Vec3) Sub(o Vec3) Vec3 {
	return Vec3{v.X - o.X, v.Y - o.Y, v.Z - o.Z}
}

func (v Vec3) Dist(o Vec3) int {
	d := v.Sub(o)
	return utils.Abs(d.X) + utils.Abs(d.Y) + utils.Abs(d.Z)
}

func (v Vec3) Less(o Vec3) bool {
	if v.X != o.X {
		return v.X < o.X
	}
	if v.Y != o.Y {
		return v.Y < o.Y
	}
	return v.Z < o.Z
}

func (v Vec3) Rotate(r string) Vec3 {
	for _, c := range r {
		switch c {
		case 'X':
			v = Vec3{v.X, -v.Z, v.Y}
		case 'Y':
			v = Vec3{v.Z, v.Y, -v.X}
		case 'Z':
			v = Vec3{-v.Y, v.X, v.Z}
		}
	}
	return v
}

type scanner struct {
	loc       Vec3
	beacons   []Vec3
	rotations [24][]Vec3
}

func newScanner(beacons []Vec3) *scanner {
	fns := [...]string{
		"", "X", "Y", "Z",

		"XX", "XY", "XZ",
		"YY", "YZ",
		"ZX", "ZZ",

		"XXX", "XXY", "XXZ",
		"YYX", "YYY", "YYZ",
		"ZZX", "ZZY", "ZZZ",

		"XXYX", "XYXX",
		"ZXXY", "ZYXX",
	}
	var s scanner
	for i := range s.rotations {
		r := make([]Vec3, len(beacons))
		for j, v := range beacons {
			r[j] = v.Rotate(fns[i])
		}
		sort.Slice(r, func(i, j int) bool { return r[i].Less(r[j]) })
		s.rotations[i] = r
	}
	s.beacons = s.rotations[0]
	return &s
}

func countOverlap(x, y []Vec3, offset Vec3) (n int) {
	i, j := 0, 0
	for i < len(x) && j < len(y) {
		xv := x[i]
		yv := y[j].Add(offset)
		if xv == yv {
			n++
			i++
			j++
		} else if xv.Less(yv) {
			i++
		} else {
			j++
		}
	}
	return
}

func (t *scanner) tryAlign(s *scanner) bool {
	for _, r := range t.rotations {
		for _, p1 := range s.beacons {
			for _, p2 := range r {
				offset := p1.Sub(p2)
				if n := countOverlap(s.beacons, r, offset); n == 12 {
					t.loc = offset
					for i, v := range r {
						t.beacons[i] = v.Add(offset)
					}
					return true
				}
			}
		}
	}
	return false
}

func main() {
	var scanners []*scanner
	for _, group := range utils.Split(input, "\n\n") {
		beacons := utils.Lines(group)[1:]
		bs := make([]Vec3, len(beacons))
		for i, line := range beacons {
			utils.Parse(&bs[i], "%d,%d,%d", line)
		}
		scanners = append(scanners, newScanner(bs))
	}

	seen := make(map[*scanner]bool)
	queue := []*scanner{scanners[0]}
	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]
		if seen[s] {
			continue
		}
		for _, t := range scanners {
			if t.loc == (Vec3{}) && t.tryAlign(s) {
				queue = append(queue, t)
			}
		}
		seen[s] = true
	}

	field := make(map[Vec3]bool)
	for _, s := range scanners {
		for _, v := range s.beacons {
			field[v] = true
		}
	}
	utils.Println(len(field))

	var max int
	for _, s1 := range scanners {
		for _, s2 := range scanners {
			max = utils.Max(max, s1.loc.Dist(s2.loc))
		}
	}
	utils.Println(max)
}
