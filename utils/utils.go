package utils

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Min returns the greater of x and y.
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Min returns the lesser of x and y.
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// Maximum returns the maximum value of fn mapped over the range [0,n).
func Maximum(n int, fn func(int) int) int {
	max := ^int(^uint(0) >> 1) // smallest possible integer
	for i := 0; i < n; i++ {
		max = Max(max, fn(i))
	}
	return max
}

// Minimum returns the minimum value of fn mapped over the range [0,n).
func Minimum(n int, fn func(int) int) int {
	min := int(^uint(0) >> 1) // largest possible integer
	for i := 0; i < n; i++ {
		min = Min(min, fn(i))
	}
	return min
}

// MinimumIndex returns the integer in [0,n) that produces the smallest value
// of fn.
func MinimumIndex(n int, fn func(int) int) int {
	min := int(^uint(0) >> 1) // largest possible integer
	mini := -1
	for i := 0; i < n; i++ {
		if f := fn(i); f < min {
			min, mini = f, i
		}
	}
	return mini
}

// And returns true if all of its arguments are true.
func And(preds ...bool) bool {
	for _, pred := range preds {
		if !pred {
			return false
		}
	}
	return true
}

// And returns true if any of its arguments are true.
func Or(preds ...bool) bool {
	for _, pred := range preds {
		if pred {
			return true
		}
	}
	return false
}

// Lines splits a string by newlines.
func Lines(input string) []string {
	return strings.Split(input, "\n")
}

// Itoa is a passthrough for strconv.Itoa
func Itoa(i int) string {
	return strconv.Itoa(i)
}

// Atoi is a passthrough for strconv.Atoi that panics upon failure.
func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

// Digits parses a string into its consituent digits.
func Digits(s string) []int {
	digits := make([]int, len(s))
	for i, c := range s {
		digits[i] = Atoi(string(c))
	}
	return digits
}

// IntList parses a list of ints.
func IntList(s string) []int {
	fs := strings.Fields(s)
	ints := make([]int, len(fs))
	for i, n := range fs {
		ints[i] = Atoi(n)
	}
	return ints
}

// IntSum returns the sum of a list of ints.
func IntSum(xs []int) int {
	sum := 0
	for _, x := range xs {
		sum += x
	}
	return sum
}

// CharCounts returns the count of each character in s.
func CharCounts(s string) map[rune]int {
	chars := make(map[rune]int)
	for _, c := range s {
		chars[c]++
	}
	return chars
}

// SortString sorts a string lexicographically.
func SortString(s string) string {
	b := []byte(s)
	sort.Slice(b, func(i int, j int) bool {
		return b[i] < b[j]
	})
	return string(b)
}

// DiffIndex returns the index at which s1 and s2 diverge.
func DiffIndex(s1, s2 string) int {
	i := 0
	for i < len(s1) && i < len(s2) {
		if s1[i] != s2[i] {
			return i
		}
		i++
	}
	return i
}

// Split is a passthrough for strings.Split.
func Split(s string, sep string) []string {
	return strings.Split(s, sep)
}

// Sscanf is a passthrough for fmt.Sscanf that panics upon failure.
func Sscanf(str, format string, args ...interface{}) {
	_, err := fmt.Sscanf(str, format, args...)
	if err != nil {
		panic(err)
	}
}

// Println is a passthrough for fmt.Println.
func Println(args ...interface{}) {
	fmt.Println(args...)
}

// Perms returns all possible permutations of the numbers [0,n).
func Perms(n int) [][]int {
	if n == 1 {
		return [][]int{{0}}
	}
	perms := Perms(n - 1)

	// interleave
	leaved := make([][]int, 0, len(perms)*n)
	for _, perm := range perms {
		for i := 0; i <= len(perm); i++ {
			withN := make([]int, len(perm)+1)
			copy(withN[:i], perm[:i])
			withN[i] = n - 1
			copy(withN[i+1:], perm[i:])
			leaved = append(leaved, withN)
		}
	}
	return leaved
}

type Pos struct {
	X, Y int
}

func (p Pos) Dist(q Pos) int {
	return Abs(p.X-q.X) + Abs(p.Y-q.Y)
}

func (p Pos) Moves() []Pos {
	return []Pos{
		{p.X, p.Y - 1},
		{p.X, p.Y + 1},
		{p.X - 1, p.Y},
		{p.X + 1, p.Y},
	}
}

func (p Pos) Numpad() []Pos {
	return []Pos{
		{p.X - 1, p.Y - 1},
		{p.X - 1, p.Y + 0},
		{p.X - 1, p.Y + 1},
		{p.X + 0, p.Y - 1},
		{p.X + 0, p.Y + 1},
		{p.X + 1, p.Y - 1},
		{p.X + 1, p.Y + 0},
		{p.X + 1, p.Y + 1},
	}
}

type Grid struct {
	X, Y int
}

func (g Grid) Valid(p Pos) bool {
	return 0 <= p.X && p.X <= g.X && 0 <= p.Y && p.Y <= g.Y
}

type Maze struct {
	Grid
	IsWall func(Pos) bool
}

func (m Maze) Valid(p Pos) bool {
	return m.Grid.Valid(p) && !m.IsWall(p)
}

func (m Maze) ValidMoves(p Pos) []Pos {
	var valid []Pos
	for _, move := range p.Moves() {
		if m.Valid(move) {
			valid = append(valid, move)
		}
	}
	return valid
}

func (p Pos) ValidMoves(g Grid) []Pos {
	var valid []Pos
	for _, m := range p.Moves() {
		if g.Valid(m) {
			valid = append(valid, m)
		}
	}
	return valid
}

func (p Pos) ValidNumpad(g Grid) []Pos {
	var valid []Pos
	for _, m := range p.Numpad() {
		if g.Valid(m) {
			valid = append(valid, m)
		}
	}
	return valid
}

func (m Maze) DistancesFrom(p Pos) map[Pos]int {
	dist := make(map[Pos]int)
	m.recdistances(dist, 0, p)
	return dist
}

func (m Maze) recdistances(distances map[Pos]int, dist int, cur Pos) {
	distances[cur] = dist

	for _, p := range m.ValidMoves(cur) {
		if d, ok := distances[p]; ok && d <= dist+1 {
			// already saw this position, and took fewer steps to reach it
			continue
		}
		m.recdistances(distances, dist+1, p)
	}
}

type Dir int

const (
	Up Dir = iota
	Right
	Down
	Left
)

func (d Dir) SpinRight(n int) Dir {
	return (((d + Dir(n)) % 4) + 4) % 4
}

func (d Dir) TurnRight() Dir  { return d.SpinRight(1) }
func (d Dir) TurnAround() Dir { return d.SpinRight(2) }
func (d Dir) TurnLeft() Dir   { return d.SpinRight(-1) }

type Agent struct {
	Pos
	Dir
}

func (a *Agent) MoveForward(n int) {
	switch a.Dir {
	case Up:
		a.Y += n
	case Right:
		a.X += n
	case Down:
		a.Y -= n
	case Left:
		a.X -= n
	}
}

func (a *Agent) MoveForwardArray(n int) {
	switch a.Dir {
	case Up:
		a.Y -= n
	case Right:
		a.X += n
	case Down:
		a.Y += n
	case Left:
		a.X -= n
	}
}

func (a *Agent) TurnRight()  { a.Dir = a.Dir.SpinRight(1) }
func (a *Agent) TurnAround() { a.Dir = a.Dir.SpinRight(2) }
func (a *Agent) TurnLeft()   { a.Dir = a.Dir.SpinRight(-1) }

func NewAgent(x, y int, d Dir) Agent {
	return Agent{
		Pos: Pos{x, y},
		Dir: d,
	}
}
