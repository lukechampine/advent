package main

const input = `Disc #1 has 17 positions; at time=0, it is at position 15.
Disc #2 has 3 positions; at time=0, it is at position 2.
Disc #3 has 19 positions; at time=0, it is at position 4.
Disc #4 has 13 positions; at time=0, it is at position 2.
Disc #5 has 7 positions; at time=0, it is at position 2.
Disc #6 has 5 positions; at time=0, it is at position 0.
`

type disc struct {
	n   int
	pos int
}

func (d disc) rotate(t int) int {
	return (d.pos + t) % d.n
}

func falls(discs []disc, t0 int) bool {
	for i, d := range discs {
		if d.rotate(t0+i+1) != 0 {
			return false
		}
	}
	return true
}

func findFirst(discs []disc) int {
	i := 0
	for !falls(discs, i) {
		i++
	}
	return i
}

func main() {
	// part 1
	discs := []disc{
		{17, 15},
		{3, 2},
		{19, 4},
		{13, 2},
		{7, 2},
		{5, 0},
	}
	println(findFirst(discs))

	// part 2
	discs = append(discs, disc{11, 0})
	println(findFirst(discs))
}
