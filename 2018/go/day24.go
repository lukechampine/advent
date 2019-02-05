package foo

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/lukechampine/advent/utils"
)

const test = `Immune System:
17 units each with 5390 hit points (weak to radiation, bludgeoning) with an attack that does 4507 fire damage at initiative 2
989 units each with 1274 hit points (immune to fire; weak to bludgeoning, slashing) with an attack that does 25 slashing damage at initiative 3

Infection:
801 units each with 4706 hit points (weak to radiation) with an attack that does 116 bludgeoning damage at initiative 1
4485 units each with 2961 hit points (immune to radiation; weak to fire, cold) with an attack that does 12 slashing damage at initiative 4`

const input = `Immune System:
2991 units each with 8084 hit points (weak to fire) with an attack that does 19 radiation damage at initiative 11
4513 units each with 3901 hit points (weak to slashing; immune to bludgeoning, radiation) with an attack that does 7 bludgeoning damage at initiative 12
5007 units each with 9502 hit points (immune to bludgeoning; weak to fire) with an attack that does 16 fire damage at initiative 2
2007 units each with 5188 hit points (weak to radiation) with an attack that does 23 cold damage at initiative 9
1680 units each with 1873 hit points (immune to bludgeoning; weak to radiation) with an attack that does 10 bludgeoning damage at initiative 10
1344 units each with 9093 hit points (immune to bludgeoning, cold; weak to radiation) with an attack that does 63 cold damage at initiative 16
498 units each with 2425 hit points (immune to fire, bludgeoning, cold) with an attack that does 44 slashing damage at initiative 3
1166 units each with 7295 hit points with an attack that does 56 bludgeoning damage at initiative 8
613 units each with 13254 hit points (immune to radiation, cold, fire) with an attack that does 162 radiation damage at initiative 15
1431 units each with 2848 hit points (weak to radiation) with an attack that does 19 cold damage at initiative 1

Infection:
700 units each with 47055 hit points (weak to fire; immune to slashing) with an attack that does 116 fire damage at initiative 14
2654 units each with 13093 hit points (weak to radiation) with an attack that does 8 radiation damage at initiative 19
5513 units each with 18026 hit points (immune to radiation; weak to slashing) with an attack that does 6 slashing damage at initiative 20
89 units each with 48412 hit points (weak to cold) with an attack that does 815 radiation damage at initiative 17
2995 units each with 51205 hit points (weak to cold) with an attack that does 28 slashing damage at initiative 7
495 units each with 21912 hit points with an attack that does 82 cold damage at initiative 13
2911 units each with 13547 hit points with an attack that does 7 slashing damage at initiative 18
1017 units each with 28427 hit points (immune to fire) with an attack that does 52 fire damage at initiative 4
2048 units each with 29191 hit points (weak to bludgeoning) with an attack that does 22 bludgeoning damage at initiative 6
1718 units each with 15725 hit points (immune to cold) with an attack that does 18 slashing damage at initiative 5`

type group struct {
	units        int
	unitHP       int
	weaknesses   []string
	immunities   []string
	attackDamage int
	attackType   string
	initiative   int

	id         int
	infection  bool
	target     *group
	isTargeted bool
}

func (g *group) effectivePower() int {
	return g.units * g.attackDamage
}

func (g *group) weakTo(typ string) bool {
	return utils.Any(len(g.weaknesses), func(i int) bool {
		return g.weaknesses[i] == typ
	})
}

func (g *group) immuneTo(typ string) bool {
	return utils.Any(len(g.immunities), func(i int) bool {
		return g.immunities[i] == typ
	})
}

func (g *group) damageDealtTo(target *group) int {
	damage := g.effectivePower()
	if target.weakTo(g.attackType) {
		damage *= 2
	}
	if target.immuneTo(g.attackType) {
		damage = 0
	}
	return damage
}

func (g group) selectTarget(targets []*group) *group {
	var best *group
	for _, t := range targets {
		if t.infection == g.infection || t.units == 0 || t.isTargeted || g.damageDealtTo(t) == 0 {
			continue
		}

		//fmt.Printf("%v group %v would deal defending group %v %v damage\n", g.teamName(), g.id, t.id, g.damageDealtTo(t))
		if best == nil ||
			g.damageDealtTo(t) > g.damageDealtTo(best) ||
			(g.damageDealtTo(t) == g.damageDealtTo(best) && t.effectivePower() > best.effectivePower()) ||
			(g.damageDealtTo(t) == g.damageDealtTo(best) && t.effectivePower() == best.effectivePower() && t.initiative > best.initiative) {
			best = t
		}
	}
	return best
}

func (g *group) attack(target *group) {
	unitsKilled := utils.Min(target.units, g.damageDealtTo(target)/target.unitHP)
	target.units -= unitsKilled
}

func (g group) teamName() string {
	if g.infection {
		return "Infection"
	}
	return "Immune System"
}

func parseAttributes(s string) (ws, is []string) {
	sp := strings.Split(s, "; ")
	for _, ats := range sp {
		if strings.HasPrefix(ats, "weak") {
			ats = strings.TrimPrefix(ats, "weak to ")
			for _, at := range strings.Split(ats, ", ") {
				ws = append(ws, at)
			}
		} else {
			ats = strings.TrimPrefix(ats, "immune to ")
			for _, at := range strings.Split(ats, ", ") {
				is = append(is, at)
			}
		}
	}
	return
}

func fight(groups []*group) {
	// select targets
	sort.Slice(groups, func(i, j int) bool {
		if groups[i].effectivePower() == groups[j].effectivePower() {
			return groups[i].initiative > groups[j].initiative
		}
		return groups[i].effectivePower() > groups[j].effectivePower()
	})
	for _, g := range groups {
		g.isTargeted = false
		g.target = nil
	}
	for _, g := range groups {
		if g.units == 0 {
			continue
		}
		g.target = g.selectTarget(groups)
		if g.target != nil {
			g.target.isTargeted = true
		}
	}

	// attack
	sort.Slice(groups, func(i, j int) bool {
		return groups[i].initiative > groups[j].initiative
	})
	for _, g := range groups {
		if g.units == 0 || g.target == nil {
			continue
		}
		g.attack(g.target)
	}
}

func remainingUnits(groups []*group) (immune, infection int) {
	var immuneUnits, infectionUnits int
	for _, g := range groups {
		if g.infection {
			infectionUnits += g.units
		} else {
			immuneUnits += g.units
		}
	}
	return immuneUnits, infectionUnits
}

func fightOver(groups []*group) bool {
	immune, infection := remainingUnits(groups)
	return immune == 0 || infection == 0
}

func parseGroups(s string) []*group {
	var groups []*group
	var isInfection bool
	var id int
	for _, line := range utils.Lines(s) {
		if line == "" || strings.Contains(line, ":") {
			isInfection = strings.Contains(line, "Infection")
			id = 0
			continue
		}
		id++
		var g group
		g.id = id
		g.infection = isInfection
		if strings.Contains(line, "(") {
			openParen := strings.IndexByte(line, '(')
			closeParen := strings.IndexByte(line, ')') + 1
			utils.Sscanf(line[:openParen], "%d units each with %d hit points ", &g.units, &g.unitHP)
			utils.Sscanf(line[closeParen:], " with an attack that does %d %s damage at initiative %v", &g.attackDamage, &g.attackType, &g.initiative)
			g.weaknesses, g.immunities = parseAttributes(line[openParen+1 : closeParen-1])
		} else {
			utils.Sscanf(line, "%d units each with %d hit points with an attack that does %d %s damage at initiative %v", &g.units, &g.unitHP, &g.attackDamage, &g.attackType, &g.initiative)
		}
		groups = append(groups, &g)
	}
	return groups
}

func TestFoo(t *testing.T) {
	// part 1
	groups := parseGroups(input)
	for !fightOver(groups) {
		fight(groups)
	}
	var finalUnits int
	for _, g := range groups {
		finalUnits += g.units
	}
	utils.Println(finalUnits)

	// part 2
	var remUnits int
	for boost := 0; ; boost++ {
		groups = parseGroups(input)
		for _, g := range groups {
			if !g.infection {
				g.attackDamage += boost
			}
		}
		for !fightOver(groups) {
			immuneBefore, infectionBefore := remainingUnits(groups)
			fight(groups)
			immuneAfter, infectionAfter := remainingUnits(groups)
			if immuneBefore == immuneAfter && infectionBefore == infectionAfter {
				fmt.Println("deadlock!")
				break
			}
		}
		immune, infection := remainingUnits(groups)
		fmt.Println(boost, immune, infection)
		if infection == 0 {
			remUnits = immune
			break
		}
	}
	utils.Println(remUnits)
}
