package main

import (
	"github.com/lukechampine/advent/utils"
)

var input = utils.ReadInput("day12_input.txt")

type Vec3 struct {
	X, Y, Z int
}

func (v Vec3) energy() int {
	return utils.Abs(v.X) + utils.Abs(v.Y) + utils.Abs(v.Z)
}

type Moon struct {
	Pos Vec3
	Vel Vec3
}

func (m *Moon) energy() int {
	return m.Pos.energy() * m.Vel.energy()
}

func (m *Moon) gravitate(n Moon) {
	adjust := func(a, b int) int {
		if a == b {
			return 0
		} else if a < b {
			return 1
		} else {
			return -1
		}
	}
	m.Vel.X += adjust(m.Pos.X, n.Pos.X)
	m.Vel.Y += adjust(m.Pos.Y, n.Pos.Y)
	m.Vel.Z += adjust(m.Pos.Z, n.Pos.Z)
}

func (m *Moon) move() {
	m.Pos.X += m.Vel.X
	m.Pos.Y += m.Vel.Y
	m.Pos.Z += m.Vel.Z
}

func main() {
	// part 1
	var positions []Vec3
	utils.Parse(&positions, "<x=%d, y=%d, z=%d>", input)
	moons := make([]Moon, len(positions))
	for i := range moons {
		moons[i].Pos = positions[i]
	}

	for step := 0; step < 1000; step++ {
		for i := range moons {
			for j := range moons {
				if i == j {
					continue
				}
				moons[i].gravitate(moons[j])
			}
		}
		for i := range moons {
			moons[i].move()
		}
	}

	utils.Println(utils.Sum(len(moons), func(i int) int {
		return moons[i].energy()
	}))

	// part 2
	for i := range moons {
		moons[i] = Moon{
			Pos: positions[i],
		}
	}
	var hist [3]map[[8]int]int
	for i := range hist {
		hist[i] = make(map[[8]int]int)
	}
	var cycles Vec3
	for step := 0; cycles.X*cycles.Y*cycles.Z == 0; step++ {
		mxs := [8]int{
			moons[0].Pos.X, moons[1].Pos.X, moons[2].Pos.X, moons[3].Pos.X,
			moons[0].Vel.X, moons[1].Vel.X, moons[2].Vel.X, moons[3].Vel.X,
		}
		mys := [8]int{
			moons[0].Pos.Y, moons[1].Pos.Y, moons[2].Pos.Y, moons[3].Pos.Y,
			moons[0].Vel.Y, moons[1].Vel.Y, moons[2].Vel.Y, moons[3].Vel.Y,
		}
		mzs := [8]int{
			moons[0].Pos.Z, moons[1].Pos.Z, moons[2].Pos.Z, moons[3].Pos.Z,
			moons[0].Vel.Z, moons[1].Vel.Z, moons[2].Vel.Z, moons[3].Vel.Z,
		}
		if cycles.X == 0 {
			if oldstep, ok := hist[0][mxs]; ok {
				cycles.X = step - oldstep
			}
			hist[0][mxs] = step
		}
		if cycles.Y == 0 {
			if oldstep, ok := hist[1][mys]; ok {
				cycles.Y = step - oldstep
			}
			hist[1][mys] = step
		}
		if cycles.Z == 0 {
			if oldstep, ok := hist[2][mzs]; ok {
				cycles.Z = step - oldstep
			}
			hist[2][mzs] = step
		}

		for i := range moons {
			for j := range moons {
				if i == j {
					continue
				}
				moons[i].gravitate(moons[j])
			}
		}
		for i := range moons {
			moons[i].move()
		}
	}
	utils.Println(utils.LCM(cycles.X, cycles.Y, cycles.Z))
}
