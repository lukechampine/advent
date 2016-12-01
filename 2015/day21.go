package main

import (
	"github.com/lukechampine/advent/utils"
)

type actor struct {
	hp, damage, armor int
}

func (a *actor) attack(b *actor) {
	b.hp -= utils.Max(a.damage-b.armor, 1)
}

func (a actor) beats(b actor) bool {
	for {
		a.attack(&b)
		if b.hp <= 0 {
			return true
		}
		b.attack(&a)
		if a.hp <= 0 {
			return false
		}
	}
}

type item struct {
	name                string
	cost, damage, armor int
}

// input
var (
	weapons = []item{
		{"Dagger", 8, 4, 0},
		{"Shortsword", 10, 5, 0},
		{"Warhammer", 25, 6, 0},
		{"Longsword", 40, 7, 0},
		{"Greataxe", 74, 8, 0},
	}

	armor = []item{
		{"Leather", 13, 0, 1},
		{"Chainmail", 31, 0, 2},
		{"Splintmail", 53, 0, 3},
		{"Bandedmail", 75, 0, 4},
		{"Platemail", 102, 0, 5},
		{"No armor", 0, 0, 0}, // armor is optional
	}

	rings = []item{
		{"Damage +1", 25, 1, 0},
		{"Damage +2", 50, 2, 0},
		{"Damage +3", 100, 3, 0},
		{"Defense +1", 20, 0, 1},
		{"Defense +2", 40, 0, 2},
		{"Defense +3", 80, 0, 3},
		{"No ring 1", 0, 0, 0}, // rings are optional
		{"No ring 2", 0, 0, 0}, // rings are optional
	}

	boss = actor{100, 8, 2}
)

type loadout struct {
	cost, damage, armor int
}

func loadouts() []loadout {
	var ls []loadout
	for _, w := range weapons {
		for _, a := range armor {
			for _, r1 := range rings {
				for _, r2 := range rings {
					if r2 == r1 {
						continue
					}
					ls = append(ls, loadout{
						cost:   w.cost + a.cost + r1.cost + r2.cost,
						damage: w.damage + r1.damage + r2.damage,
						armor:  a.armor + r1.armor + r2.armor,
					})
				}
			}
		}
	}
	return ls
}

func main() {
	// part 1
	var mincost int = 999
	for _, l := range loadouts() {
		player := actor{100, l.damage, l.armor}
		if player.beats(boss) {
			mincost = utils.Min(mincost, l.cost)
		}
	}
	println(mincost)

	// part 2
	var maxcost int
	for _, l := range loadouts() {
		player := actor{100, l.damage, l.armor}
		if !player.beats(boss) {
			maxcost = utils.Max(maxcost, l.cost)
		}
	}
	println(maxcost)
}
