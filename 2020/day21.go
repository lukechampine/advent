package main

import (
	"sort"
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2020, 21)

func main() {
	lines := utils.Lines(input)

	allergens := make(map[string]*utils.StringSet)
	for _, l := range lines {
		parts := strings.Split(l, "(contains ")
		ingredients := utils.NewStringSet(strings.Fields(parts[0])...)
		contains := strings.Split(strings.TrimSuffix(parts[1], ")"), ", ")
		for _, name := range contains {
			if set, ok := allergens[name]; !ok {
				allergens[name] = ingredients
			} else {
				allergens[name] = set.Intersection(ingredients)
			}
		}
	}
	suspect := utils.NewStringSet()
	for _, ingredients := range allergens {
		suspect = suspect.Union(ingredients)
	}

	n := 0
	for _, l := range lines {
		parts := strings.Split(l, "(contains ")
		for _, i := range strings.Fields(parts[0]) {
			n += utils.BoolToInt(!suspect.Contains(i))
		}
	}
	utils.Println(n)

	var allergenic []string
	for len(allergens) > 0 {
		for name, ingredients := range allergens {
			if ingredients.Len() == 1 {
				ingredient := ingredients.Elems()[0]
				allergenic = append(allergenic, name+":"+ingredient)
				for _, other := range allergens {
					other.Delete(ingredient)
				}
				delete(allergens, name)
			}
		}
	}
	sort.Strings(allergenic)
	for i, a := range allergenic {
		allergenic[i] = strings.Split(a, ":")[1]
	}
	utils.Println(strings.Join(allergenic, ","))
}
