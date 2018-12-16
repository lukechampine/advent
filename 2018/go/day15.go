package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/lukechampine/advent/utils"
)

const input = `################################
##########.###.###..############
##########..##......############
#########...##....##############
######.....###..G..G############
##########..........############
##########.............#########
#######G..#.G...#......#########
#..G##....##..#.G#....#...######
##......###..##..####.#..#######
#G.G..#..#....#.###...G..#######
#.....GG##................######
#....G........#####....E.E.#####
#####G...#...#######........####
####.E#.G...#########.......####
#...G.....#.#########......#####
#.##........#########.......####
######......#########........###
######......#########..E.#....##
#######..E.G.#######..........##
#######E......#####............#
#######...G............E.......#
####............##............##
####..G.........##..........E.##
####.G.G#.....####E...##....#.##
#######.......####...####..#####
########....E....########..#####
##########.......#########...###
##########.......#########..####
##########....############..####
###########...##################
################################`

type Unit struct {
	utils.Pos
	isElf bool
	hp    int
	power int
}

func (u Unit) String() string {
	if u.isElf {
		return fmt.Sprintf("Elf at %v with %v hp", u.Pos, u.hp)
	}
	return fmt.Sprintf("Goblin at %v with %v hp", u.Pos, u.hp)
}

func (u Unit) tryAttack(units []Unit, board [][]byte) bool {
	moves := u.Moves()
	var targets []*Unit
	for i, t := range units {
		for _, m := range moves {
			if m == t.Pos && t.isElf != u.isElf && t.hp > 0 {
				targets = append(targets, &units[i])
				break
			}
		}
	}
	if len(targets) == 0 {
		return false
	}
	// select target with least hp, or by readability
	sel := targets[0]
	for _, t := range targets[1:] {
		if (t.hp < sel.hp) ||
			(t.hp == sel.hp && t.Y < sel.Y) ||
			(t.hp == sel.hp && t.Y == sel.Y && t.X < sel.X) {
			sel = t
		}
	}
	sel.hp -= u.power
	if sel.hp <= 0 {
		board[sel.Y][sel.X] = '.'
	}
	return true
}

type unitsByReadingOrder []Unit

func (us unitsByReadingOrder) Len() int      { return len(us) }
func (us unitsByReadingOrder) Swap(i, j int) { us[i], us[j] = us[j], us[i] }
func (us unitsByReadingOrder) Less(i, j int) bool {
	if us[i].Y != us[j].Y {
		return us[i].Y < us[j].Y
	}
	return us[i].X < us[j].X
}

func selectClosest(grid utils.Maze, options []utils.Pos, to utils.Pos) (utils.Pos, bool) {
	closest := grid.SelectClosest(options, to)
	if len(closest) == 0 {
		return utils.Pos{}, false
	}
	// pick first by "readability"
	best := closest[0]
	for _, p := range closest[1:] {
		if p.Y < best.Y || (p.Y == best.Y && p.X < best.X) {
			best = p
		}
	}
	return best, true
}

func run(input string, elfPower int) (outcome int, units []Unit) {
	lines := utils.Lines(input)
	for y := range lines {
		for x, c := range lines[y] {
			switch c {
			case 'E':
				units = append(units, Unit{
					Pos:   utils.Pos{X: x, Y: y},
					isElf: true,
					hp:    200,
					power: elfPower,
				})
			case 'G':
				units = append(units, Unit{
					Pos:   utils.Pos{X: x, Y: y},
					isElf: false,
					hp:    200,
					power: 3,
				})
			}
		}
	}
	board := make([][]byte, len(lines))
	for i := range board {
		board[i] = []byte(lines[i])
	}

	grid := utils.Maze{
		Grid: utils.Grid{
			X: len(lines[0]),
			Y: len(lines),
		},
		IsWall: func(p utils.Pos) bool {
			return board[p.Y][p.X] != '.'
		},
	}

	var rounds int
outer:
	for ; ; rounds++ {
		sort.Sort(unitsByReadingOrder(units))
		for i, u := range units {
			if u.hp <= 0 {
				continue // you are already dead
			}
			var haveTargets bool
			for _, t := range units {
				if t.hp > 0 && t.isElf != u.isElf {
					haveTargets = true
				}
			}
			if !haveTargets {
				// combat ends
				break outer
			}

			// attack if possible
			if u.tryAttack(units, board) {
				continue
			}

			// identify all squares "in range" of targets
			var inRange []utils.Pos
			for _, t := range units {
				if t.isElf == u.isElf || t.hp <= 0 {
					continue
				}
				inRange = append(inRange, grid.ValidMoves(t.Pos)...)
			}

			// find closest reachable target
			closestTarget, ok := selectClosest(grid, inRange, u.Pos)
			if !ok {
				continue
			}
			// find move that moves us closest to the closest target
			move, ok := selectClosest(grid, grid.ValidMoves(u.Pos), closestTarget)
			if !ok {
				continue
			}
			// swap board elements and update pos
			board[u.Pos.Y][u.Pos.X], board[move.Y][move.X] = board[move.Y][move.X], board[u.Pos.Y][u.Pos.X]
			u.Pos = move
			units[i] = u

			// attack if possible
			u.tryAttack(units, board)
		}
	}
	var finalUnits []Unit
	var hpSum int
	for _, u := range units {
		if u.hp > 0 {
			finalUnits = append(finalUnits, u)
			hpSum += u.hp
		}
	}
	return rounds * hpSum, finalUnits
}

func main() {
	// part 1
	outcome, _ := run(input, 3)
	utils.Println(outcome)

	// part 2
	origElves := strings.Count(input, "E")
	minPower := sort.Search(100, func(elfPower int) bool {
		_, units := run(input, elfPower)
		return units[0].isElf && len(units) == origElves
	})
	outcome, _ = run(input, minPower)
	utils.Println(outcome)

}
