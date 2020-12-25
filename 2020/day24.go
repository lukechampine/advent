package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2020, 24)

type Vec3 struct {
	X, Y, Z int
}

func follow(s string) Vec3 {
	var v Vec3
	for len(s) > 0 {
		switch {
		case s[:1] == "e":
			v.X++
			v.Y--
			s = s[1:]

		case s[:1] == "w":
			v.X--
			v.Y++
			s = s[1:]

		case s[:2] == "se":
			v.Y--
			v.Z++
			s = s[2:]

		case s[:2] == "sw":
			v.X--
			v.Z++
			s = s[2:]

		case s[:2] == "ne":
			v.X++
			v.Z--
			s = s[2:]

		case s[:2] == "nw":
			v.Y++
			v.Z--
			s = s[2:]
		}
	}
	return v
}

func (v Vec3) adj() []Vec3 {
	return []Vec3{
		{v.X + 1, v.Y - 1, v.Z + 0},
		{v.X + 0, v.Y - 1, v.Z + 1},
		{v.X - 1, v.Y + 0, v.Z + 1},
		{v.X - 1, v.Y + 1, v.Z + 0},
		{v.X + 0, v.Y + 1, v.Z - 1},
		{v.X + 1, v.Y + 0, v.Z - 1},
	}
}

func gol(m map[Vec3]bool) map[Vec3]bool {
	next := make(map[Vec3]bool)
	for v, b := range m {
		next[v] = b
		if b {
			for _, a := range v.adj() {
				next[a] = m[a] // default white
			}
		}
	}
	for v, b := range next {
		a := v.adj()
		n := utils.Count(len(a), func(i int) bool { return m[a[i]] })
		if (b && (n == 0 || n > 2)) || (!b && n == 2) {
			next[v] = !b
		}
	}
	return next
}

func main() {
	m := make(map[Vec3]bool)
	for _, l := range utils.Lines(input) {
		v := follow(l)
		m[v] = !m[v]
	}
	n := 0
	for _, b := range m {
		n += utils.BoolToInt(b)
	}
	utils.Println(n)

	for i := 0; i < 100; i++ {
		m = gol(m)
	}
	n = 0
	for _, b := range m {
		n += utils.BoolToInt(b)
	}
	utils.Println(n)
}
