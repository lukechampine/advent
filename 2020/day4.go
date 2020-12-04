package main

import (
	"encoding/hex"
	"math/big"
	"strconv"
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2020, 4)

func isNumInRange(s string, l, h int) bool {
	n, err := strconv.Atoi(s)
	return err == nil && l <= n && n <= h
}

var validators = map[string]func(s string) bool{
	"byr": func(s string) bool { return len(s) == 4 && isNumInRange(s, 1920, 2002) },
	"iyr": func(s string) bool { return len(s) == 4 && isNumInRange(s, 2010, 2020) },
	"eyr": func(s string) bool { return len(s) == 4 && isNumInRange(s, 2020, 2030) },
	"hgt": func(s string) bool {
		if len(s) < 2 {
			return false
		}
		val, unit := s[:len(s)-2], s[len(s)-2:]
		return (unit == "cm" && isNumInRange(val, 150, 193)) || (unit == "in" && isNumInRange(val, 59, 76))
	},
	"hcl": func(s string) bool {
		_, err := hex.DecodeString(strings.TrimPrefix(s, "#"))
		return err == nil && len(s) == 7 && s[0] == '#'
	},
	"ecl": func(s string) bool {
		switch s {
		case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
			return true
		default:
			return false
		}
	},
	"pid": func(s string) bool {
		_, ok := new(big.Int).SetString(s, 10)
		return ok && len(s) == 9
	},
}

func main() {
	var validFields, validValues int
	for _, l := range strings.Split(input, "\n\n") {
		var vf, vv int
		for _, f := range strings.Fields(l) {
			kv := strings.Split(f, ":")
			if valid, ok := validators[kv[0]]; ok {
				vf++
				vv += utils.BoolToInt(valid(kv[1]))
			}
		}
		validFields += utils.BoolToInt(vf >= 7)
		validValues += utils.BoolToInt(vv >= 7)
	}
	utils.Println(validFields)
	utils.Println(validValues)
}
