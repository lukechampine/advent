package main

import (
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2020, 13)

// We have:
//
//   x ≡ a1 (mod n1)
//   x ≡ a2 (mod n2)
//
// First, solve the system with a1 = 0 and a2 = 1 using the extended Euclidean
// algorithm, producing m. Then compute (m*(a2-a1) mod n2)*n1 + a1 to get x.
//
// Example:
//
//   x ≡ 1 (mod 3)
//   x ≡ 5 (mod 7)
//
// For x ≡ 0 (mod 3) ≡ 1 (mod 7), we compute m = 5. (5*(5-1) mod 7)*3 + 1 = 19.
//
// But why does this work? Well, consider the system:
//
//   x ≡ 0 (mod 3)
//   x ≡ n (mod 7)
//
// For n = 0..6. We can find solutions manually by looking for the first
// multiple of 3 in the sequence 7b + n. Visually:
//
//   3a     =  0  3  6  9 12 15 18 21 24 27 30 33
//   7b + 0 = >0<     7     14    >21<    28
//   7b + 1 =   1      8    >15<    22     29
//   7b + 2 =    2     >9<    16     23    >30<
//   7b + 3 =    >3<    10     17    >24<    31
//   7b + 4 =      4     11    >18<    25     32
//   7b + 5 =       5    >12<    19     26    >33<
//   7b + 6 =       >6<    13     20    >27<    34
//
// Now consider the *index* of these solutions in each sequence. For the
// sequence 7b + n, the solutions are at index 0,2,1,0,2,1, which is pretty
// clearly the sequence -b mod 3. For the sequence 3a, the indices are
// 0,5,3,1,6,4,2. Here, the pattern is a bit harder to see; it's 5a mod 7.
//
// This is great, because if we know the index in the sequence 3a, we can just
// multiply by 3 to get the actual value; and this value is x. For example, if n
// is 4, then the index is 5*4 mod 7 = 6, so x = 6*3 = 18. (And in case it
// wasn't obvious, this method generalizes to any n1 and n2, not just 3 and 7.)
//
// The other thing to note is that if we can solve a system with 0 and n, we can
// solve it for any a1 and a2. All we need to do is subtract a1 from both
// equations to get a system with 0 and n; then we solve that, and add a1 to the
// result. This works because "shifting" both equations by a fixed offset
// doesn't change the relationship between their sequences.
//
// Now there's only one missing piece of the puzzle: how do we figure out the
// pattern of indices in the sequences? We know that the pattern will be of the
// form m*a mod n2; we just need to solve for m. Notice that m is the index
// where n = 1; thus, we can restate our problem as solving for x in the system:
//
//   x ≡ 0 (mod n1)
//   x ≡ 1 (mod n2)
//
// Fortunately, there is a well-known algorithm for this: the extended Euclidean
// algorithm. Whereas the standard Euclidean algorithm only produces the
// greatest common denominator, the *extended* algorithm also produces the
// so-called Bezout coefficients s and t, satisfying n1*s + n2*t = gcd(n1, n2).
// Since our n1 and n2 are always coprime, their gcd is 1; furthermore, since n1
// and n2 are greater than 1, s and t must have opposite signs. Thus we can
// rewrite the equation as n1*s = n2*t + 1, which makes it clear that x = n1*s.
//
// Hmm? How does the extended Euclidean algorithm work? No idea, sorry.
//
// Anyway, putting it all together: we use the EEA to get the solution for the
// 0,1 system; this solution gives us the coefficient m in our "index equation"
// m*a mod n2; we use this equation to get the solution for the 0,a system,
// where a is (a2-a1); and finally, we add a1 to solve our original system.
//
// NOTE: apparently you can also solve the system with the equation:
//
//   x = a1*t*n2 + a2*s*n1
//
// where t is the other Bezout coefficient. However, this equation overflows our
// native int representation, whereas my equation does not --- probably because
// my largest term is a2-a1, rather than max(a1, a2).
func solveMod(a1, n1, a2, n2 int) (int, int) {
	bezout := func(a, b int) int {
		s0, s1 := 1, 0
		for b != 0 {
			s0, s1 = s1, s0-(a/b)*s1
			a, b = b, a%b
		}
		return s0
	}
	m := bezout(n1, n2)
	r := (m*(a2-a1)%n2)*n1 + a1
	lcm := n1 * n2
	return r % lcm, lcm
}

func main() {
	lines := utils.Lines(input)
	start := utils.Atoi(lines[0])
	buses := utils.ExtractInts(lines[1])
	b := buses[utils.MinimumIndex(len(buses), func(i int) int {
		return buses[i] - start%buses[i]
	})]
	utils.Println(b * (b - start%b))

	// We have a system of congruence equations:
	//
	//   x ≡ a1 (mod n1)
	//   x ≡ a2 (mod n2)
	//    ...
	//   x ≡ ak (mod nk)
	//
	// To solve the system, we repeatedly reduce pairs of two equations to a
	// single equation x ≡ r (mod ni*nj) until only one equation remains.
	var a, n int
	for i, bus := range strings.Split(lines[1], ",") {
		if bus == "x" {
			continue
		}
		ni := utils.Atoi(bus)
		if i == 0 {
			a, n = -i, ni
		} else {
			a, n = solveMod(a, n, -i, ni)
		}
	}
	// ensure a is positive
	if a < 0 {
		a += n
	}
	utils.Println(a)
}
