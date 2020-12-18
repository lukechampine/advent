package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2020, 17)

type Vec3 struct {
	X, Y, Z int
}

func (v Vec3) numActiveNeighbors(active map[Vec3]struct{}) (n int) {
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				if dx == 0 && dy == 0 && dz == 0 {
					continue
				}
				nv := Vec3{v.X + dx, v.Y + dy, v.Z + dz}
				if _, ok := active[nv]; ok {
					n++
				}
			}
		}
	}
	return
}

func (v Vec3) inactiveNeighbors(active map[Vec3]struct{}) (ns []Vec3) {
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				if dx == 0 && dy == 0 && dz == 0 {
					continue
				}
				nv := Vec3{v.X + dx, v.Y + dy, v.Z + dz}
				if _, ok := active[nv]; !ok {
					ns = append(ns, nv)
				}
			}
		}
	}
	return
}

type Vec4 struct {
	X, Y, Z, T int
}

func (v Vec4) numActiveNeighbors(active map[Vec4]struct{}) (n int) {
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				for dt := -1; dt <= 1; dt++ {
					if dx == 0 && dy == 0 && dz == 0 && dt == 0 {
						continue
					}
					nv := Vec4{v.X + dx, v.Y + dy, v.Z + dz, v.T + dt}
					if _, ok := active[nv]; ok {
						n++
					}
				}
			}
		}
	}
	return
}

func (v Vec4) inactiveNeighbors(active map[Vec4]struct{}) (ns []Vec4) {
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				for dt := -1; dt <= 1; dt++ {
					if dx == 0 && dy == 0 && dz == 0 && dt == 0 {
						continue
					}
					nv := Vec4{v.X + dx, v.Y + dy, v.Z + dz, v.T + dt}
					if _, ok := active[nv]; !ok {
						ns = append(ns, nv)
					}
				}
			}
		}
	}
	return
}

func main() {
	grid := utils.ToByteGrid(utils.Lines(input))
	active3 := make(map[Vec3]struct{})
	for y := range grid {
		for x, c := range grid[y] {
			if c == '#' {
				active3[Vec3{x, y, 0}] = struct{}{}
			}
		}
	}

	for step := 0; step < 6; step++ {
		next := make(map[Vec3]struct{})
		for v := range active3 {
			if n := v.numActiveNeighbors(active3); n == 2 || n == 3 {
				next[v] = struct{}{}
			}
			for _, in := range v.inactiveNeighbors(active3) {
				if n := in.numActiveNeighbors(active3); n == 3 {
					next[in] = struct{}{}
				}
			}
		}
		active3 = next
	}
	utils.Println(len(active3))

	active4 := make(map[Vec4]struct{})
	for y := range grid {
		for x, c := range grid[y] {
			if c == '#' {
				active4[Vec4{x, y, 0, 0}] = struct{}{}
			}
		}
	}

	for step := 0; step < 6; step++ {
		next := make(map[Vec4]struct{})
		for v := range active4 {
			if n := v.numActiveNeighbors(active4); n == 2 || n == 3 {
				next[v] = struct{}{}
			}
			for _, in := range v.inactiveNeighbors(active4) {
				if n := in.numActiveNeighbors(active4); n == 3 {
					next[in] = struct{}{}
				}
			}
		}
		active4 = next
	}
	utils.Println(len(active4))
}
