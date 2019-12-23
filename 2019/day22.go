package main

import (
	"math/big"
	"strings"

	"github.com/lukechampine/advent/utils"
)

var input = utils.ReadInput("day22_input.txt")
var ils = utils.Lines(input)

func main() {
	// part 1
	deckSize := 10007
	pos := 2019
	for _, line := range ils {
		switch {
		case strings.Contains(line, "cut"):
			cut := utils.ExtractInts(line)[0] % deckSize
			pos -= cut

		case strings.Contains(line, "deal with"):
			inc := utils.ExtractInts(line)[0]
			pos *= inc

		case strings.Contains(line, "deal into new stack"):
			pos = (deckSize - pos - 1)
		}
		pos %= deckSize
	}
	utils.Println(pos)

	// part 2
	size := big.NewInt(119315717514047)
	one := big.NewInt(1)
	a, b := big.NewInt(1), big.NewInt(0) // y = ax + b
	for i := range ils {
		line := ils[len(ils)-i-1]
		switch {
		case strings.Contains(line, "cut"):
			// b += cut % size
			cut := big.NewInt(int64(utils.ExtractInts(line)[0]))
			cut.Mod(cut, size)
			b.Add(b, cut)

		case strings.Contains(line, "deal with"):
			// a *= invmod(inc, size)
			// b *= invmod(inc, size)
			inc := big.NewInt(int64(utils.ExtractInts(line)[0]))
			inv := inc.ModInverse(inc, size)
			a.Mul(a, inv)
			b.Mul(b, inv)

		case strings.Contains(line, "deal into new stack"):
			// a *= -1
			// b = size - b - 1
			a.Sub(big.NewInt(0), a)
			b.Sub(size, b)
			b.Sub(b, one)
		}
	}

	iters := big.NewInt(101741582076661)
	w := new(big.Int).Exp(a, iters, size)
	x := new(big.Int).Mul(w, big.NewInt(2020))        // (a^iters mod size) * 2020
	y := new(big.Int).Sub(w, one)                     // (a^iters mod size) - 1
	z := new(big.Int).ModInverse(a.Sub(a, one), size) // modinv(a-1, size)

	// (x + byz) mod size
	ans := y
	ans.Mul(ans, b)
	ans.Mul(ans, z)
	ans.Add(ans, x)
	ans.Mod(ans, size)
	utils.Println(ans)
}
