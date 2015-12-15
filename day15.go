package main

import (
	"github.com/lukechampine/advent/utils"
)

const input = `Sugar: capacity 3, durability 0, flavor 0, texture -3, calories 2
Sprinkles: capacity -3, durability 3, flavor 0, texture 0, calories 9
Candy: capacity -1, durability 0, flavor 4, texture 0, calories 1
Chocolate: capacity 0, durability 0, flavor -2, texture 2, calories 8`

type ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func score(is []ingredient, props []int) int {
	var capacity, durability, flavor, texture int
	for i := range is {
		ing, ts := is[i], props[i]
		capacity += ing.capacity * ts
		durability += ing.durability * ts
		flavor += ing.flavor * ts
		texture += ing.texture * ts
	}
	// set any negative values to 0
	zero := func(x int) int {
		if x < 0 {
			return 0
		}
		return x
	}
	return zero(capacity) * zero(durability) * zero(flavor) * zero(texture)
}

func calories(is []ingredient, props []int) int {
	var calories int
	for i := range is {
		calories += is[i].calories * props[i]
	}
	return calories
}

func proportions(ingredients, teaspoons int) [][]int {
	if ingredients == 1 {
		return [][]int{{teaspoons}}
	}
	var props [][]int
	for i := 0; i <= teaspoons; i++ {
		subProps := proportions(ingredients-1, teaspoons-i)
		for _, sp := range subProps {
			prop := make([]int, ingredients)
			prop[0] = i
			copy(prop[1:], sp)
			props = append(props, prop)
		}
	}
	return props
}

func parse(str string) ingredient {
	var i ingredient
	utils.Sscanf(str, "%s capacity %d, durability %d, flavor %d, texture %d, calories %d", &i.name, &i.capacity, &i.durability, &i.flavor, &i.texture, &i.calories)
	return i
}

func main() {
	// part 1
	var is []ingredient
	for _, str := range utils.Lines(input) {
		is = append(is, parse(str))
	}
	var optimal int
	for _, props := range proportions(len(is), 100) {
		optimal = utils.Max(optimal, score(is, props))
	}
	println(optimal)

	// part 2
	optimal = 0
	for _, props := range proportions(len(is), 100) {
		if calories(is, props) == 500 {
			optimal = utils.Max(optimal, score(is, props))
		}
	}
	println(optimal)
}
