package main

import (
	"math/big"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2023, 24)

type Vec3 struct {
	X, Y, Z int
}

type hailstone struct {
	pos, vel Vec3
}

func collides(a, b hailstone, min, max float64) bool {
	a0x, a1x := float64(a.pos.X), float64(a.pos.X+a.vel.X)
	a0y, a1y := float64(a.pos.Y), float64(a.pos.Y+a.vel.Y)
	am := (a1y - a0y) / (a1x - a0x)
	b0x, b1x := float64(b.pos.X), float64(b.pos.X+b.vel.X)
	b0y, b1y := float64(b.pos.Y), float64(b.pos.Y+b.vel.Y)
	bm := (b1y - b0y) / (b1x - b0x)
	x := (b0y - a0y - bm*b0x + am*a0x) / (am - bm)
	y := am*(x-a0x) + a0y
	if a0x == a1x || b0x == b1x || am == bm {
		return false
	}
	if utils.Signum(int(x-a0x)) != utils.Signum(a.vel.X) ||
		utils.Signum(int(x-b0x)) != utils.Signum(b.vel.X) {
		return false // collision is in the past
	}
	return min <= x && x <= max && min <= y && y <= max
}

func solveMod(a1, n1, a2, n2 *big.Int) (*big.Int, *big.Int) {
	bezout := func(a, b *big.Int) *big.Int {
		s0, s1 := big.NewInt(1), big.NewInt(0)
		for b.Sign() != 0 {
			r := new(big.Int).Mul(a, s1)
			r.Sub(s0, r.Div(r, b))
			s0, s1 = s1, r
			a, b = b, new(big.Int).Mod(a, b)
		}
		return s0
	}
	m := bezout(n1, n2)
	r := new(big.Int).Sub(a2, a1)
	r.Mul(r, m)
	r.Mod(r, n2)
	r.Mul(r, n1)
	r.Add(r, a1)
	lcm := new(big.Int).Mul(n1, n2)
	return r.Mod(r, lcm), lcm
}

func crt2(mods []*big.Int, rem []*big.Int) (int, bool) {
	var a, n *big.Int
	for i, ni := range mods {
		if i == 0 {
			a, n = big.NewInt(int64(-i)), ni
		} else {
			a, n = solveMod(a, n, big.NewInt(int64(-i)), ni)
		}
	}
	return int(a.Int64()), a.IsInt64()
}

func check(stones []hailstone, i int) (_ int, ok bool) {
	var mods []*big.Int
	var rems []*big.Int
	for _, s := range stones[:10] {
		mod := big.NewInt(int64(s.vel.X))
		mod.Add(mod, big.NewInt(int64(s.vel.Y)))
		mod.Add(mod, big.NewInt(int64(s.vel.Z)))
		mod.Sub(mod, big.NewInt(int64(i)))
		mod.Abs(mod)
		if mod.Sign() == 0 {
			return 0, false
		}
		mods = append(mods, mod)
		rem := big.NewInt(int64(s.pos.X))
		rem.Add(rem, big.NewInt(int64(s.pos.Y)))
		rem.Add(rem, big.NewInt(int64(s.pos.Z)))
		rem.Mod(rem, mod)
		rems = append(rems, rem)
	}
	return crt2(mods, rems)
}

func main() {
	var stones []hailstone
	for _, line := range utils.Lines(input) {
		ints := utils.ExtractInts(line)
		stones = append(stones, hailstone{
			pos: Vec3{ints[0], ints[1], ints[2]},
			vel: Vec3{ints[3], ints[4], ints[5]},
		})
	}

	n := 0
	for i := range stones {
		for j := i + 1; j < len(stones); j++ {
			if collides(stones[i], stones[j], 200000000000000, 400000000000000) {
				n++
			}
		}
	}
	utils.Println(n)

	for i := -1000; i <= 1000; i++ {
		if sol, ok := check(stones, i); ok {
			utils.Println(i, sol)
			break
		}
	}
	utils.Println("no solutions")
	return
}
