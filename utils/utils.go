// Package utils provides helper types and functions for solving Advent of Code
// challenges.
package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// Input returns the input for the specified year and day as a string,
// downloading it if it does not already exist on disk.
func Input(year, day int) string {
	filename := fmt.Sprintf("day%v_input.txt", day)
	if _, err := os.Stat(filename); err != nil {
		est, err := time.LoadLocation("EST")
		if err != nil {
			panic(err)
		}
		if t := time.Date(year, time.December, day, 0, 0, 0, 0, est); time.Until(t) > 0 {
			fmt.Printf("Puzzle not unlocked yet! Sleeping for %v...\n", time.Until(t))
			time.Sleep(time.Until(t) + 3*time.Second) // don't want to fire too early
		}
		fmt.Println("Downloading input...")
		req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("https://adventofcode.com/%v/day/%v/input", year, day), nil)
		req.AddCookie(&http.Cookie{
			Name:  "session",
			Value: os.Getenv("AOC_USER_ID"),
		})
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		if err := ioutil.WriteFile(filename, data, 0660); err != nil {
			panic(err)
		}
	}
	return ReadInput(filename)
}

// ReadInput returns the contents of filename as a string.
func ReadInput(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(data)
}

// Window returns a slice of slices consisting of the original slice, split into
// groups of n elements. The final slice may contain fewer than n elements.
func Window(slice interface{}, n int) interface{} {
	v := reflect.ValueOf(slice)
	groups := v.Len() / n
	if v.Type().Kind() == reflect.String {
		w := make([]string, 0, groups+1)
		for i := 0; i < v.Len(); i += n {
			s := Min(n, v.Len()-i)
			w = append(w, v.Slice(i, i+s).String())
		}
		return w
	}
	w := reflect.MakeSlice(reflect.SliceOf(v.Type()), 0, groups+1)
	for i := 0; i < v.Len(); i += n {
		s := Min(n, v.Len()-i)
		w = reflect.Append(w, v.Slice(i, i+s))
	}
	return w.Interface()
}

// SlidingWindow returns a slice of slices consisting of the original slice,
// split into windows of n elements.
func SlidingWindow(slice interface{}, n int) interface{} {
	v := reflect.ValueOf(slice)
	groups := v.Len() - n + 1
	if v.Type().Kind() == reflect.String {
		w := make([]string, groups)
		for i := range w {
			w[i] = v.Slice(i, i+n).String()
		}
		return w
	}
	w := reflect.MakeSlice(reflect.SliceOf(v.Type()), groups, groups)
	for i := 0; i < w.Len(); i++ {
		w.Index(i).Set(v.Slice(i, i+n))
	}
	return w.Interface()
}

// Stride returns a slice of values from the original slice, consisting of
// every nth "step".
func Stride(slice interface{}, n int) interface{} {
	v := reflect.ValueOf(slice)
	if v.Type().Kind() == reflect.String {
		var s string
		for i := 0; i < v.Len(); i += n {
			s += v.Slice(i, i+1).String()
		}
		return s
	}
	w := reflect.MakeSlice(v.Type(), 0, v.Len()/n+1)
	for i := 0; i < v.Len(); i += n {
		w = reflect.Append(w, v.Index(i))
	}
	return w.Interface()
}

func Transpose(slice interface{}) interface{} {
	v := reflect.ValueOf(slice)
	if v.Type().Kind() != reflect.Slice {
		panic("not a slice")
	}
	switch v.Type().Elem().Kind() {
	case reflect.String:
		t := make([]string, len(v.Index(0).String()))
		for i := 0; i < v.Len(); i++ {
			for j, c := range v.Index(i).String() {
				t[j] += string(c)
			}
		}
		return t
	case reflect.Slice:
		t := reflect.MakeSlice(v.Type(), v.Index(0).Len(), v.Index(0).Len())
		for i := 0; i < v.Len(); i++ {
			s := v.Index(i)
			for j := 0; j < s.Len(); j++ {
				t.Index(j).Set(reflect.Append(t.Index(j), s.Index(j)))
			}
		}
		return t.Interface()
	default:
		panic("not a slice of slices or strings")
	}
}

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

// MinimumIndex returns the integer in [0,n) that produces the smallest value
// of fn.
func MaximumIndex(n int, fn func(int) int) int {
	max := ^int(^uint(0) >> 1) // smallest possible integer
	maxi := -1
	for i := 0; i < n; i++ {
		if f := fn(i); f > max {
			max, maxi = f, i
		}
	}
	return maxi
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

// Any returns true if fn returns true for any value in [0..n).
func Any(n int, fn func(i int) bool) bool {
	for i := 0; i < n; i++ {
		if fn(i) {
			return true
		}
	}
	return false
}

// All returns true if fn returns true for all values in [0..n).
func All(n int, fn func(i int) bool) bool {
	for i := 0; i < n; i++ {
		if !fn(i) {
			return false
		}
	}
	return true
}

// Count returns the number of values in [0..n) for which fn returns true.
func Count(n int, fn func(i int) bool) (c int) {
	for i := 0; i < n; i++ {
		if fn(i) {
			c++
		}
	}
	return
}

// Sum returns the sum of f(i) for i in [0..n).
func Sum(n int, fn func(i int) int) (c int) {
	for i := 0; i < n; i++ {
		c += fn(i)
	}
	return
}

// Product returns the product of f(i) for i in [0..n).
func Product(n int, fn func(i int) int) (c int) {
	c = 1
	for i := 0; i < n; i++ {
		c *= fn(i)
	}
	return
}

func GCD(a, b int) int {
	a, b = Abs(a), Abs(b)
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(nums ...int) int {
	result := nums[0] * nums[1] / GCD(nums[0], nums[1])
	for _, i := range nums[2:] {
		result = LCM(result, i)
	}
	return result
}

func Signum(x int) int {
	if x < 0 {
		return -1
	}
	if x > 0 {
		return 1
	}
	return 0
}

func IntToBool(i int) bool { return i != 0 }
func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func ReverseString(s string) string {
	b := []byte(s)
	for i := 0; i < len(b)/2; i++ {
		j := len(b) - i - 1
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func DeleteSliceIndex(sliceptr interface{}, i int) {
	sp := reflect.ValueOf(sliceptr)
	if sp.Kind() != reflect.Ptr || sp.Elem().Kind() != reflect.Slice {
		panic("not a pointer to a slice")
	}
	s := sp.Elem()
	if s.Len() <= i {
		return
	}
	s.Set(reflect.AppendSlice(s.Slice(0, i), s.Slice(i+1, s.Len())))
}

func DeleteSliceValue(sliceptr interface{}, v interface{}) {
	sp := reflect.ValueOf(sliceptr)
	if sp.Kind() != reflect.Ptr || sp.Elem().Kind() != reflect.Slice {
		panic("not a pointer to a slice")
	}
	s := sp.Elem()
	for i := 0; i < s.Len(); i++ {
		if reflect.DeepEqual(s.Index(i).Interface(), v) {
			s.Set(reflect.AppendSlice(s.Slice(0, i), s.Slice(i+1, s.Len())))
			return
		}
	}
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
	return Sum(len(xs), func(i int) int { return xs[i] })
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

func IsLetter(c byte) bool {
	return ('A' <= c && c <= 'Z') || ('a' <= c && c <= 'z')
}

func IsLower(s string) bool {
	return s == strings.ToLower(s)
}

func IsUpper(s string) bool {
	return s == strings.ToUpper(s)
}

func ContainsByte(s string, b byte) bool {
	return strings.IndexByte(s, b) >= 0
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
	_, err := fmt.Sscanf(strings.TrimSpace(str), format, args...)
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

var Origin = Pos{0, 0}

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

// StrideTowards "moves" p towards q, taking the smallest "step" that results in
// integer coordinates.
func (p Pos) StrideTowards(q Pos) Pos {
	if q == p {
		return q
	}
	dx, dy := q.X-p.X, q.Y-p.Y
	gcd := GCD(dx, dy)
	return Pos{p.X + dx/gcd, p.Y + dy/gcd}
}

func (p Pos) Add(x, y int) Pos {
	return Pos{p.X + x, p.Y + y}
}

func (p Pos) Rel(q Pos) Pos {
	return Pos{p.X - q.X, p.Y - q.Y}
}

func (p Pos) Polar() (rho float64, phi float64) {
	x, y := float64(p.X), float64(p.Y)
	return math.Hypot(x, y), math.Atan2(y, x)
}

func (p Pos) DirTo(q Pos) Dir {
	_, phi := q.Rel(p).Polar()
	return Dir(2*phi/math.Pi) + 1
}

func (p Pos) RotateClockwiseAround(q Pos, deg float64) Pos {
	rho, phi := p.Rel(q).Polar()
	phi -= (deg / 180) * math.Pi
	return FromPolar(rho, phi).Add(q.X, q.Y)
}

func (p Pos) RotateCounterClockwiseAround(q Pos, deg float64) Pos {
	rho, phi := p.Rel(q).Polar()
	phi += (deg / 180) * math.Pi
	return FromPolar(rho, phi).Add(q.X, q.Y)
}

func FromPolar(rho, phi float64) Pos {
	return Pos{int(math.Round(rho * math.Cos(phi))), int(math.Round(rho * math.Sin(phi)))}
}

func BoundingBox(ps []Pos) (min, max Pos) {
	min.X, min.Y = int(1e9), int(1e9)
	max.X, max.Y = int(-1e9), int(-1e9)
	for _, p := range ps {
		min.X = Min(min.X, p.X)
		min.Y = Min(min.Y, p.Y)
		max.X = Max(max.X, p.X)
		max.Y = Max(max.Y, p.Y)
	}
	return
}

func BoundingRect(ps []Pos) Rect {
	min, max := BoundingBox(ps)
	return Rect{min, max}
}

func PrintPositions(ps []Pos, empty, full rune) {
	min, max := BoundingBox(ps)
	for i := range ps {
		ps[i].X -= min.X
		ps[i].Y -= min.Y
	}
	g := make([][]rune, 1+max.Y-min.Y)
	for y := range g {
		g[y] = make([]rune, 1+max.X-min.X)
		for x := range g[y] {
			g[y][x] = empty
		}
	}
	for _, p := range ps {
		y := len(g) - p.Y - 1 // flip vertically
		g[y][p.X] = full
	}
	for _, row := range g {
		fmt.Println(string(row))
	}
}

type Rect struct {
	Min, Max Pos
}

func (r Rect) Width() int  { return r.Max.Rel(r.Min).X }
func (r Rect) Height() int { return r.Max.Rel(r.Min).Y }
func (r Rect) Area() int {
	p := r.Max.Rel(r.Min)
	return (p.X + 1) * (p.Y + 1)
}

func (r Rect) Corners() []Pos {
	return []Pos{
		{r.Min.X, r.Min.Y},
		{r.Max.X, r.Min.Y},
		{r.Max.X, r.Max.Y},
		{r.Min.X, r.Max.Y},
	}
}

func (r Rect) Contains(p Pos) bool {
	return r.Min.X <= p.X && p.X <= r.Max.X &&
		r.Min.Y <= p.Y && p.Y <= r.Max.Y
}

func (r Rect) ContainsRect(p Rect) bool {
	return r.Contains(p.Min) && r.Contains(p.Max)
}

func (r Rect) Overlaps(p Rect) bool {
	return r.Min.X <= p.Max.X && p.Min.X <= r.Max.X &&
		r.Min.Y <= p.Max.Y && p.Min.Y <= r.Max.Y
}

func (r Rect) Overlap(p Rect) Rect {
	return Rect{
		Min: Pos{Max(r.Min.X, p.Min.X), Max(r.Min.Y, p.Min.Y)},
		Max: Pos{Min(r.Max.X, p.Max.X), Min(r.Max.Y, p.Max.Y)},
	}
}

func (r Rect) ForEach(fn func(Pos)) {
	for y := r.Min.Y; y <= r.Max.Y; y++ {
		for x := r.Min.X; x <= r.Max.X; x++ {
			fn(Pos{x, y})
		}
	}
}

type Grid struct {
	X, Y int
}

func (g Grid) Valid(p Pos) bool {
	return 0 <= p.X && p.X <= g.X && 0 <= p.Y && p.Y <= g.Y
}

func (g Grid) ForEach(fn func(Pos)) {
	for y := 0; y <= g.Y; y++ {
		for x := 0; x <= g.X; x++ {
			fn(Pos{x, y})
		}
	}
}

type LightBoard struct {
	Board [][]bool
}

func (l *LightBoard) Set(x, y int, lit bool) {
	l.Board[y][x] = lit
}

func (l *LightBoard) Get(x, y int) bool {
	return l.Board[y][x]
}

func (l *LightBoard) ForEach(fn func(p Pos, b bool)) {
	for y := range l.Board {
		for x := range l.Board[y] {
			fn(Pos{x, y}, l.Board[y][x])
		}
	}
}

func (l *LightBoard) Count(set bool) (n int) {
	l.ForEach(func(p Pos, b bool) {
		if b == set {
			n++
		}
	})
	return
}

func (l *LightBoard) Print() {
	l.ForEach(func(p Pos, b bool) {
		if b {
			fmt.Print("█")
		} else {
			fmt.Print(" ")
		}
		if p.X == len(l.Board[0])-1 {
			fmt.Println()
		}
	})
}

func NewLightBoard(x, y int) *LightBoard {
	board := make([][]bool, y)
	for i := range board {
		board[i] = make([]bool, x)
	}
	return &LightBoard{board}
}

func Locate(grid [][]byte, c byte) Pos {
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == c {
				return Pos{x, y}
			}
		}
	}
	panic("not found")
}

type Maze struct {
	Grid
	IsWall  func(Pos) bool
	CanMove func(Pos, Pos) bool
}

func (m Maze) canMove(p, d Pos) bool {
	return m.Grid.Valid(d) && !m.IsWall(d) && (m.CanMove == nil || m.CanMove(p, d))
}

func (m Maze) Valid(p Pos) bool {
	return m.Grid.Valid(p) && !m.IsWall(p)
}

func (m Maze) ValidMoves(p Pos) []Pos {
	valid := make([]Pos, 0, 4)
	for _, move := range p.Moves() {
		if m.canMove(p, move) {
			valid = append(valid, move)
		}
	}
	return valid
}

func (m Maze) Minimize(canBacktrack func(Pos) bool) func(Pos) bool {
	walls := make(map[Pos]struct{})
	oldFn := m.IsWall
	m.IsWall = func(p Pos) bool {
		if _, ok := walls[p]; ok {
			return true
		}
		return oldFn(p)
	}

	// DFS; first recurse to all valid moves; then, if only one move is
	// possible, mark current position as a wall
	seen := make(map[Pos]struct{})
	var recMinimize func(Pos)
	recMinimize = func(p Pos) {
		if _, ok := seen[p]; ok {
			return
		}
		seen[p] = struct{}{}
		for _, move := range m.ValidMoves(p) {
			recMinimize(move)
		}
		if len(m.ValidMoves(p)) == 1 && canBacktrack(p) {
			walls[p] = struct{}{}
		}
	}

	// starting point doesn't matter; use first non-wall position
	var start Pos
outer:
	for start.X = 0; start.X < m.X; start.X++ {
		for start.Y = 0; start.Y < m.Y; start.Y++ {
			if !m.IsWall(start) {
				break outer
			}
		}
	}
	recMinimize(start)

	return m.IsWall
}

func MakeSimpleMaze(grid [][]byte, wall byte) Maze {
	return Maze{
		Grid:   Grid{X: len(grid[0]) - 1, Y: len(grid) - 1},
		IsWall: func(p Pos) bool { return grid[p.Y][p.X] == wall },
	}
}

func PrintMaze(m Maze, wall byte) {
	var p Pos
	for p.Y = 0; p.Y < m.Y; p.Y++ {
		for p.X = 0; p.X < m.X; p.X++ {
			if m.IsWall(p) {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
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

	if p := (Pos{cur.X, cur.Y - 1}); m.canMove(cur, p) {
		if d, ok := distances[p]; !ok || d > dist+1 {
			m.recdistances(distances, dist+1, p)
		}
	}
	if p := (Pos{cur.X, cur.Y + 1}); m.canMove(cur, p) {
		if d, ok := distances[p]; !ok || d > dist+1 {
			m.recdistances(distances, dist+1, p)
		}
	}
	if p := (Pos{cur.X - 1, cur.Y}); m.canMove(cur, p) {
		if d, ok := distances[p]; !ok || d > dist+1 {
			m.recdistances(distances, dist+1, p)
		}
	}
	if p := (Pos{cur.X + 1, cur.Y}); m.canMove(cur, p) {
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
	if p := (Pos{to.X, to.Y - 1}); m.canMove(p, to) {
		queue = append(queue, entry{p, 1})
	}
	if p := (Pos{to.X, to.Y + 1}); m.canMove(p, to) {
		queue = append(queue, entry{p, 1})
	}
	if p := (Pos{to.X - 1, to.Y}); m.canMove(p, to) {
		queue = append(queue, entry{p, 1})
	}
	if p := (Pos{to.X + 1, to.Y}); m.canMove(p, to) {
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

		if p := (Pos{e.X, e.Y - 1}); m.canMove(e.Pos, p) {
			queue = append(queue, entry{p, e.dist + 1})
		}
		if p := (Pos{e.X, e.Y + 1}); m.canMove(e.Pos, p) {
			queue = append(queue, entry{p, e.dist + 1})
		}
		if p := (Pos{e.X - 1, e.Y}); m.canMove(e.Pos, p) {
			queue = append(queue, entry{p, e.dist + 1})
		}
		if p := (Pos{e.X + 1, e.Y}); m.canMove(e.Pos, p) {
			queue = append(queue, entry{p, e.dist + 1})
		}
	}
	return closest
}

type MazeSolver struct {
	Pos
	Ch   byte
	Dist int
	Done bool
}

func recvisit(grid [][]byte, s MazeSolver, seen map[Pos]struct{}, fn func(*MazeSolver)) {
	seen[s.Pos] = struct{}{}
	fn(&s)
	if s.Done {
		return
	}
	s.Dist++
	for _, p := range MakeSimpleMaze(grid, '#').ValidMoves(s.Pos) {
		if _, ok := seen[p]; ok {
			continue
		}
		s.Pos = p
		s.Ch = grid[s.Y][s.X]
		recvisit(grid, s, seen, fn)
	}
}

func SolveMaze(grid [][]byte, start Pos, fn func(*MazeSolver)) {
	seen := make(map[Pos]struct{})
	s := MazeSolver{
		Pos: start,
		Ch:  grid[start.Y][start.X],
	}
	recvisit(grid, s, seen, fn)
}

func Flood(start Pos, visit func(Pos) []Pos) {
	seen := make(map[Pos]bool)
	queue := []Pos{start}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		if seen[p] {
			continue
		}
		seen[p] = true
		queue = append(queue, visit(p)...)
	}
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

func (d Dir) SpinLeft(n int) Dir {
	return d.SpinRight(-n)
}

func (d Dir) TurnRight() Dir  { return d.SpinRight(1) }
func (d Dir) TurnAround() Dir { return d.SpinRight(2) }
func (d Dir) TurnLeft() Dir   { return d.SpinLeft(1) }

func DirFromNEWS(c byte) Dir {
	switch c {
	case 'N':
		return Up
	case 'E':
		return Right
	case 'W':
		return Left
	case 'S':
		return Down
	}
	panic("invalid NEWS")
}

func DirFromUDLR(c byte) Dir {
	switch c {
	case 'U':
		return Up
	case 'D':
		return Down
	case 'L':
		return Left
	case 'R':
		return Right
	}
	panic("invalid UDLR")
}

func DirFromArrow(c byte) Dir {
	switch c {
	case '^':
		return Up
	case 'v':
		return Down
	case '<':
		return Left
	case '>':
		return Right
	}
	panic("invalid ^v<>")
}

func (p Pos) Move(d Dir, n int) Pos {
	switch d {
	case Up:
		p.Y += n
	case Right:
		p.X += n
	case Down:
		p.Y -= n
	case Left:
		p.X -= n
	}
	return p
}

func (p Pos) MoveArray(d Dir, n int) Pos {
	switch d {
	case Up:
		p.Y -= n
	case Right:
		p.X += n
	case Down:
		p.Y += n
	case Left:
		p.X -= n
	}
	return p
}

func (p Pos) Tread(d Dir, n int, fn func(Pos)) Pos {
	for i := 0; i < n; i++ {
		p = p.Move(d, 1)
		fn(p)
	}
	return p
}

func (p Pos) TreadArray(d Dir, n int, fn func(Pos)) Pos {
	for i := 0; i < n; i++ {
		p = p.MoveArray(d, 1)
		fn(p)
	}
	return p
}

type Agent struct {
	Pos
	Dir
}

func (a Agent) String() string {
	return fmt.Sprintf("(%v, %v, %v)", a.X, a.Y, a.Dir)
}

func (a *Agent) Move(d Dir, n int) {
	a.Pos = a.Pos.Move(d, n)
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

func (a *Agent) SpinRight(n int) { a.Dir = a.Dir.SpinRight(n) }
func (a *Agent) SpinLeft(n int)  { a.Dir = a.Dir.SpinLeft(n) }

func (a *Agent) TurnRight()  { a.SpinRight(1) }
func (a *Agent) TurnAround() { a.SpinRight(2) }
func (a *Agent) TurnLeft()   { a.SpinLeft(1) }

func (a *Agent) ValidMoves(g Grid) []Agent {
	as := make([]Agent, 0, 4)
	for i := range as {
		a := *a
		a.SpinLeft(i)
		a.MoveForward(1)
		if g.Valid(a.Pos) {
			as = append(as, a)
		}
	}
	return as
}

func NewAgent(x, y int, d Dir) Agent {
	return Agent{
		Pos: Pos{x, y},
		Dir: d,
	}
}

func PrintGrid(grid [][]byte) {
	fmt.Println(string(bytes.Join(CloneGrid(grid), []byte("\n"))))
}

func CountGrid(grid [][]byte, c byte) int {
	return Sum(len(grid), func(i int) int {
		return bytes.Count(grid[i], []byte{c})
	})
}

func ToByteGrid(grid []string) [][]byte {
	b := make([][]byte, len(grid))
	for i := range b {
		b[i] = []byte(grid[i])
	}
	return b
}

func CloneGrid(grid [][]byte) [][]byte {
	b := make([][]byte, len(grid))
	for i := range b {
		b[i] = append([]byte(nil), grid[i]...)
	}
	return b
}

// GameOfLife runs one step of a Game-of-Life-style cellular automaton grid.
func GameOfLife(grid [][]byte, update func(c byte, p Pos, adj []Pos) byte) [][]byte {
	next := CloneGrid(grid)
	for y := range grid {
		for x, c := range grid[y] {
			p := Pos{x, y}
			g := Grid{len(grid[y]) - 1, len(grid) - 1}
			next[y][x] = update(c, p, p.ValidMoves(g))
		}
	}
	return next
}

type UnionFinder struct {
	parent map[int]int
	rank   map[int]int
}

func NewUnionFinder() UnionFinder {
	return UnionFinder{
		parent: make(map[int]int),
		rank:   make(map[int]int),
	}
}

func (u UnionFinder) hasParent(p int) bool {
	_, ok := u.parent[p]
	return ok
}

func (u UnionFinder) Find(p int) int {
	root := p

	// Find the root of the element by following parent pointers until an element
	// without a parent is found.
	for u.hasParent(root) {
		root = u.parent[root]
	}

	// Compress the connections between the element and the located root by making
	// every element found on the way to root point directly to it.
	for p != root {
		p, u.parent[p] = u.parent[p], root
	}

	return root
}

func (u UnionFinder) Join(p int, q int) {
	pr := u.Find(p)
	qr := u.Find(q)

	if pr == qr {
		return
	}

	// Merge the lower-ranking component into the larger-ranking component.
	if u.rank[pr] < u.rank[qr] {
		u.parent[pr] = qr
	} else {
		u.parent[qr] = pr
	}

	// Increase the rank of the merged component if joining two components that
	// have the same rank.
	if u.rank[pr] == u.rank[qr] {
		u.rank[pr]++
	}
}

func (u UnionFinder) Connected(p int, q int) bool { return u.Find(p) == u.Find(q) }

func ExtractInts(s string) []int {
	fs := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsDigit(r) && r != '-'
	})
	ints := make([]int, 0, len(fs))
	for _, w := range fs {
		i, err := strconv.Atoi(w)
		if err == nil {
			ints = append(ints, i)
		}
	}
	return ints
}

func Parse(obj interface{}, format string, input string) {
	typ := reflect.TypeOf(obj)
	if typ.Kind() != reflect.Ptr {
		panic("not a pointer!")
	}
	switch typ.Elem().Kind() {
	case reflect.Slice:
		if typ.Elem().Elem().Kind() != reflect.Struct {
			panic("not a pointer to a slice of structs!")
		}
		parseStructSlice(reflect.ValueOf(obj).Elem(), format, input)

	case reflect.Struct:
		parseStruct(reflect.ValueOf(obj).Elem(), format, input)

	default:
		panic("not a pointer to a struct or slice of structs!")
	}
}

func parseStructSlice(obj reflect.Value, format, input string) {
	lines := Lines(input)
	obj.Set(reflect.MakeSlice(obj.Type(), len(lines), len(lines)))
	for i, line := range lines {
		parseStruct(obj.Index(i), format, line)
	}
}

func parseStruct(obj reflect.Value, format string, input string) {
	var args []interface{}
	for i := 0; i < obj.NumField(); i++ {
		args = append(args, obj.Field(i).Addr().Interface())
	}
	Sscanf(input, format, args...)
}

func Wrap(xs []int, n int) [][]int {
	if len(xs)%n != 0 {
		panic("untidy wrapping")
	}
	lines := make([][]int, len(xs)/n)
	for i := range lines {
		lines[i] = xs[i*n:][:n]
	}
	return lines
}

func WrapString(s string, n int) []string {
	if len(s)%n != 0 {
		panic("untidy wrapping")
	}
	lines := make([]string, len(s)/n)
	for i := range lines {
		lines[i] = s[i*n:][:n]
	}
	return lines
}

func Replace(s string, oldnew ...string) string {
	return strings.NewReplacer(oldnew...).Replace(s)
}

func ByteGrid(x, y int, init ...byte) [][]byte {
	g := make([][]byte, y)
	for i := range g {
		g[i] = make([]byte, x)
		if len(init) != 0 {
			for j := range g[i] {
				g[i][j] = init[0]
			}
		}
	}
	return g
}

func Surround(grid [][]byte, c byte) [][]byte {
	g := make([][]byte, len(grid)+2)
	for y := range g {
		g[y] = bytes.Repeat([]byte{c}, len(grid[0])+2)
		if 0 < y && y < len(g)-1 {
			copy(g[y][1:], grid[y-1])
		}
	}
	return g
}

type StringSet struct {
	m map[string]struct{}
}

func (set *StringSet) Add(s string) bool {
	added := !set.Contains(s)
	set.m[s] = struct{}{}
	return added
}

func (set *StringSet) Delete(s string) bool {
	deleted := set.Contains(s)
	delete(set.m, s)
	return deleted
}

func (set *StringSet) Contains(s string) bool {
	_, ok := set.m[s]
	return ok
}

func (set *StringSet) Len() int {
	return len(set.m)
}

func (set *StringSet) Elems() []string {
	ss := make([]string, 0, len(set.m))
	for s := range set.m {
		ss = append(ss, s)
	}
	return ss
}

func (set *StringSet) Union(other *StringSet) *StringSet {
	u := NewStringSet()
	for s := range set.m {
		u.m[s] = struct{}{}
	}
	for s := range other.m {
		u.m[s] = struct{}{}
	}
	return u
}

func (set *StringSet) Difference(other *StringSet) *StringSet {
	u := NewStringSet()
	for s := range set.m {
		if !other.Contains(s) {
			u.Add(s)
		}
	}
	return u
}

func (set *StringSet) Intersection(other *StringSet) *StringSet {
	u := NewStringSet()
	for s := range set.m {
		if other.Contains(s) {
			u.Add(s)
		}
	}
	return u
}

func NewStringSet(ss ...string) *StringSet {
	m := make(map[string]struct{}, len(ss))
	for _, s := range ss {
		m[s] = struct{}{}
	}
	return &StringSet{m: m}
}
