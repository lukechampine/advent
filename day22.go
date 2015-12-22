package main

import (
	"github.com/lukechampine/advent/utils"
	"math/rand"
	"time"
)

type actor struct {
	hp, mp        int
	damage, armor int
	effects       []effect
}

func (a *actor) attack(b *actor) {
	b.hp -= utils.Max(a.damage-b.armor, 1)
}

func (a *actor) runEffects() {
	for i := 0; i < len(a.effects); i++ {
		a.effects[i].fn(a)
		a.effects[i].turns--
		if a.effects[i].turns <= 0 {
			// absolutely_disgusting.jpg
			if a.effects[i].name == "Shield" {
				a.armor = 0
			}
			a.effects = append(a.effects[:i], a.effects[i+1:]...)
			i--
		}
	}
}

func (a actor) fight(b actor, spells []*spell, hardmode bool) (int, bool) {
	var cost int
	for _, spell := range spells {
		if hardmode {
			if a.hp--; a.hp <= 0 {
				return cost, false
			}
		}

		// player turn
		a.runEffects()
		b.runEffects()
		if spell.cost > a.mp {
			// more efficient if we assume this loses immediately
			return cost, false
		}
		a.mp -= spell.cost
		cost += spell.cost
		valid := spell.cast(&a, &b)
		if !valid {
		}
		if b.hp <= 0 {
			return cost, true
		}

		// boss turn
		a.runEffects()
		b.runEffects()
		if b.hp <= 0 {
			return cost, true
		}
		b.attack(&a)
		if a.hp <= 0 {
			return cost, false
		}
	}
	// not enough spells
	return cost, false
}

type spell struct {
	name         string
	cost         int
	playerEffect effect
	bossEffect   effect
}

type effect struct {
	name    string
	fn      func(a *actor)
	instant bool
	turns   int
}

func (s spell) cast(a, b *actor) (valid bool) {
	// check for duplicates
	for _, e := range a.effects {
		if e.name == s.name {
			return false
		}
	}
	for _, e := range b.effects {
		if e.name == s.name {
			return false
		}
	}

	if s.playerEffect.instant {
		s.playerEffect.fn(a)
	} else if s.playerEffect.fn != nil {
		a.effects = append(a.effects, s.playerEffect)
	}
	if s.bossEffect.instant {
		s.bossEffect.fn(b)
	} else if s.bossEffect.fn != nil {
		b.effects = append(b.effects, s.bossEffect)
	}
	return true
}

// input
var (
	spells = []spell{
		{name: "Magic Missile", cost: 53, bossEffect: effect{
			name:    "Magic Missile",
			fn:      func(a *actor) { a.hp -= 4 },
			instant: true,
		}},
		{name: "Drain", cost: 73, playerEffect: effect{
			name:    "Drain",
			fn:      func(a *actor) { a.hp += 2 },
			instant: true,
		}, bossEffect: effect{
			name:    "Drain",
			fn:      func(a *actor) { a.hp -= 2 },
			instant: true,
		}},
		{name: "Shield", cost: 113, playerEffect: effect{
			name:  "Shield",
			fn:    func(a *actor) { a.armor = 7 }, // hack
			turns: 6,
		}},
		{name: "Poison", cost: 173, bossEffect: effect{
			name:  "Poison",
			fn:    func(a *actor) { a.hp -= 3 },
			turns: 6,
		}},
		{name: "Recharge", cost: 229, playerEffect: effect{
			name:  "Recharge",
			fn:    func(a *actor) { a.mp += 101 },
			turns: 5,
		}},
	}

	boss   = actor{hp: 58, damage: 9}
	player = actor{hp: 50, mp: 500}
)

// generate a bunch of random spell sequences
func spellSeqs() [][]*spell {
	var seqs [][]*spell
	for i := 0; i < 3000000; i++ {
		var seq []*spell
		for j := 0; j < 10; j++ { // totally arbitrary
			seq = append(seq, &spells[rand.Intn(len(spells))])
		}
		seqs = append(seqs, seq)
	}
	return seqs
}

func main() {
	rand.Seed(time.Now().Unix())
	seqs := spellSeqs()

	// part 1
	println(utils.Minimum(len(seqs), func(i int) int {
		if cost, won := player.fight(boss, seqs[i], false); won {
			return cost
		}
		return 999999
	}))

	// part 2
	println(utils.Minimum(len(seqs), func(i int) int {
		if cost, won := player.fight(boss, seqs[i], true); won {
			return cost
		}
		return 999999
	}))
}
