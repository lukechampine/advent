package utils

import (
	"bytes"
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

// Max returns the greater of x and y.
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

// Or returns true if any of its arguments are true.
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
	return strings.Split(strings.TrimSpace(input), "\n")
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

// Print is a passthrough for fmt.Print.
func Print(args ...interface{}) {
	fmt.Print(args...)
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

type LightBoard struct {
	board [][]bool
}

func (l *LightBoard) Set(x, y int, lit bool) {
	l.board[y][x] = lit
}

func (l *LightBoard) Get(x, y int) bool {
	return l.board[y][x]
}

func (l *LightBoard) Print() {
	for y := range l.board {
		for x := range l.board[y] {
			if l.board[y][x] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func NewLightBoard(x, y int) *LightBoard {
	board := make([][]bool, y)
	for i := range board {
		board[i] = make([]bool, x)
	}
	return &LightBoard{board}
}

type Maze struct {
	Grid
	IsWall func(Pos) bool
}

func (m Maze) Valid(p Pos) bool {
	return m.Grid.Valid(p) && !m.IsWall(p)
}

func (m Maze) ValidMoves(p Pos) []Pos {
	valid := make([]Pos, 0, 4)
	for _, move := range p.Moves() {
		if m.Valid(move) {
			valid = append(valid, move)
		}
	}
	return valid
}

func (p Pos) ValidMoves(g Grid) []Pos {
	valid := make([]Pos, 0, 4)
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
	dist := make(map[Pos]int, m.X*m.Y)
	m.recdistances(dist, 0, p)
	return dist
}

func (m Maze) recdistances(distances map[Pos]int, dist int, cur Pos) {
	distances[cur] = dist

	if p := (Pos{cur.X, cur.Y - 1}); m.Valid(p) {
		if d, ok := distances[p]; !ok || d > dist+1 {
			m.recdistances(distances, dist+1, p)
		}
	}
	if p := (Pos{cur.X, cur.Y + 1}); m.Valid(p) {
		if d, ok := distances[p]; !ok || d > dist+1 {
			m.recdistances(distances, dist+1, p)
		}
	}
	if p := (Pos{cur.X - 1, cur.Y}); m.Valid(p) {
		if d, ok := distances[p]; !ok || d > dist+1 {
			m.recdistances(distances, dist+1, p)
		}
	}
	if p := (Pos{cur.X + 1, cur.Y}); m.Valid(p) {
		if d, ok := distances[p]; !ok || d > dist+1 {
			m.recdistances(distances, dist+1, p)
		}
	}
}

// SelectClosest returns the positions closest to 'to' among the supplied
// options. If multiple options have the same distance, they are returned in
// unspecified order.
func (m Maze) SelectClosest(options []Pos, to Pos) []Pos {
	// first check that we're not already on an option
	for _, o := range options {
		if o == to {
			return []Pos{to}
		}
	}

	// BFS using a queue
	type entry struct {
		Pos
		dist int
	}
	queue := make([]entry, 0, 100)
	if p := (Pos{to.X, to.Y - 1}); m.Valid(p) {
		queue = append(queue, entry{p, 1})
	}
	if p := (Pos{to.X, to.Y + 1}); m.Valid(p) {
		queue = append(queue, entry{p, 1})
	}
	if p := (Pos{to.X - 1, to.Y}); m.Valid(p) {
		queue = append(queue, entry{p, 1})
	}
	if p := (Pos{to.X + 1, to.Y}); m.Valid(p) {
		queue = append(queue, entry{p, 1})
	}
	seen := make(map[Pos]struct{}, 1000)
	closest := make([]Pos, 0, 10)
	var limit int = 1e9
outer:
	for len(queue) > 0 {
		e := queue[0]
		copy(queue, queue[1:])
		queue = queue[:len(queue)-1]
		if _, ok := seen[e.Pos]; ok {
			continue
		}
		seen[e.Pos] = struct{}{}
		if e.dist > limit {
			continue
		}
		for _, o := range options {
			if o == e.Pos {
				closest = append(closest, e.Pos)
				limit = e.dist
				continue outer
			}
		}

		if p := (Pos{e.X, e.Y - 1}); m.Valid(p) {
			queue = append(queue, entry{p, e.dist + 1})
		}
		if p := (Pos{e.X, e.Y + 1}); m.Valid(p) {
			queue = append(queue, entry{p, e.dist + 1})
		}
		if p := (Pos{e.X - 1, e.Y}); m.Valid(p) {
			queue = append(queue, entry{p, e.dist + 1})
		}
		if p := (Pos{e.X + 1, e.Y}); m.Valid(p) {
			queue = append(queue, entry{p, e.dist + 1})
		}
	}
	return closest
}

type Dir int

const (
	Up Dir = iota
	Right
	Down
	Left
)

func (d Dir) String() string {
	return map[Dir]string{
		Up:    "Up",
		Down:  "Down",
		Left:  "Left",
		Right: "Right",
	}[d]
}

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

func (a Agent) String() string {
	return fmt.Sprintf("(%v, %v, %v)", a.X, a.Y, a.Dir)
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

func PrintGrid(grid [][]byte) {
	grid2 := make([][]byte, len(grid))
	for i := range grid2 {
		grid2[i] = append([]byte(nil), grid[i]...)
	}
	fmt.Println(string(bytes.Join(grid2, []byte("\n"))))
}

func CountGrid(grid [][]byte, c byte) int {
	sep := []byte{c}
	var sum int
	for y := range grid {
		sum += bytes.Count(grid[y], sep)
	}
	return sum
}
