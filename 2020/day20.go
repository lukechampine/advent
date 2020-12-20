package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2020, 20)

func rotateGrid(grid [][]byte) [][]byte {
	rot := make([][]byte, len(grid[0]))
	for i := range rot {
		rot[i] = make([]byte, len(grid))
	}
	for y, row := range rot {
		for x := range row {
			rot[y][x] = grid[len(grid)-x-1][y]
		}
	}
	return rot
}

func flipGrid(grid [][]byte) [][]byte {
	flipped := make([][]byte, len(grid))
	for i := range flipped {
		flipped[i] = []byte(utils.ReverseString(string(grid[i])))
	}
	return flipped
}

type tile struct {
	id   int
	grid [][]byte
	adj  [4]*tile // top, right, bottom, left
}

func (t *tile) side(i int) string {
	var b []byte
	switch i {
	case 0:
		b = t.grid[0]
	case 1:
		for _, l := range t.grid {
			b = append(b, l[len(l)-1])
		}
	case 2:
		b = t.grid[len(t.grid)-1]
	case 3:
		for _, l := range t.grid {
			b = append(b, l[0])
		}
	}
	return string(b)
}

func (t *tile) hasSide(side string) bool {
	return utils.Any(4, func(i int) bool {
		s := t.side(i)
		return side == s || side == utils.ReverseString(s)
	})
}

func (t *tile) flip() {
	t.grid = flipGrid(t.grid)
	t.adj[1], t.adj[3] = t.adj[3], t.adj[1]
}

func (t *tile) rotate() {
	t.grid = rotateGrid(t.grid)
	t.adj = [4]*tile{t.adj[3], t.adj[0], t.adj[1], t.adj[2]}
}

func (t *tile) alignedTo(q *tile) *tile {
	if t == nil {
		return nil
	}
	matches := func(i int) bool { return t.side(i) == q.side((i+2)%4) }
	for i := 0; !utils.Any(4, matches); i++ {
		t.rotate()
		if i == 3 {
			t.flip()
		}
	}
	return t
}

func main() {
	tiles := make(map[int]*tile)
	for _, t := range utils.Split(input, "\n\n") {
		lines := utils.Lines(t)
		tl := &tile{
			id:   utils.ExtractInts(lines[0])[0],
			grid: utils.ToByteGrid(lines[1:]),
		}
		tiles[tl.id] = tl
	}
	// find adj set for each tile (not correctly oriented yet)
	for _, t := range tiles {
		for i := range t.adj {
			side := t.side(i)
			for _, q := range tiles {
				if q.id != t.id && q.hasSide(side) {
					t.adj[i] = q
					break
				}
			}
		}
	}
	// corners will have 2 nil adj
	var corners []int
	for _, t := range tiles {
		n := utils.Count(4, func(i int) bool { return t.adj[i] == nil })
		if n == 2 {
			corners = append(corners, t.id)
		}
	}
	utils.Println(corners[0] * corners[1] * corners[2] * corners[3])

	// rotate+flip corners[0] to make it the "top left"
	start := tiles[corners[0]]
	for i := 0; !(start.adj[0] == nil && start.adj[3] == nil); i++ {
		start.rotate()
		if i == 3 {
			start.flip()
		}
	}

	// build image one row at a time, aligning tiles as we go
	var img [][]byte
	cur := start
	for cur != nil {
		rows := make([][]byte, len(cur.grid)-2)
		down := cur.adj[2].alignedTo(cur)
		for cur != nil {
			right := cur.adj[1].alignedTo(cur)
			for i, row := range cur.grid[1 : len(cur.grid)-1] {
				rows[i] = append(rows[i], row[1:len(row)-1]...)
			}
			cur = right
		}
		img = append(img, rows...)
		cur = down
	}

	// hunt for sea monsters
	for i := 0; ; i++ {
		if n := countMonsters(img); n > 0 {
			utils.Println(utils.CountGrid(img, '#') - n*utils.CountGrid(monster, '#'))
			return
		}
		img = rotateGrid(img)
		if i == 3 {
			img = flipGrid(img)
		}
	}
}

var monster = [][]byte{
	[]byte("                  # "),
	[]byte("#    ##    ##    ###"),
	[]byte(" #  #  #  #  #  #   "),
}

func countMonsters(img [][]byte) int {
	lineMatch := func(l, m []byte) bool {
		if len(l) < len(m) {
			return false
		}
		for i := range m {
			if m[i] == '#' && l[i] != '#' {
				return false
			}
		}
		return true
	}
	var n int
	for i := 0; i < len(img)-2; i++ {
		for x := range img[i] {
			n += utils.BoolToInt(utils.All(3, func(r int) bool {
				return lineMatch(img[i+r][x:], monster[r])
			}))
		}
	}
	return n
}
