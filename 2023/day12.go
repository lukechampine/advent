package main

import (
	"fmt"
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2023, 12)

var memo = make(map[string]int)

func arrangements(springs string, runs []int) (res int) {
	id := fmt.Sprintf("%v|%v", springs, runs)
	if n, ok := memo[id]; ok {
		return n
	}
	defer func() { memo[id] = res }()

	if len(runs) == 0 {
		return utils.BoolToInt(!strings.Contains(springs, "#"))
	} else if len(springs) < runs[0] {
		return 0
	} else if len(springs) == runs[0] {
		return utils.BoolToInt(!strings.Contains(springs, ".") && len(runs) == 1)
	}
	n := 0
	if !strings.Contains(springs[:runs[0]], ".") && springs[runs[0]] != '#' {
		// taking is possible
		n += arrangements(springs[runs[0]+1:], runs[1:])
	}
	if springs[0] != '#' {
		// not taking is also ok
		n += arrangements(springs[1:], runs)
	}
	return n
}

func nfa(spring string, runs []int) int {
	isBlocks := []bool{false}
	for _, r := range runs {
		for i := 0; i < r; i++ {
			isBlocks = append(isBlocks, true)
		}
		isBlocks = append(isBlocks, false)
	}
	nfa := make([]int, len(isBlocks))
	nfa[0] = 1
	nfa[1] = 1
	for _, c := range spring {
		next := make([]int, len(isBlocks))
		for i, b := range isBlocks {
			switch {
			case b && c != '.':
				next[i+1] += nfa[i]
			case !b && c != '#':
				next[i] += nfa[i]
				if i+1 < len(nfa) {
					next[i+1] += nfa[i]
				}
			}
		}
		nfa = next
	}
	return nfa[len(nfa)-1]
}

func nfaArray(springs string, runs []int) int {
	isb := []int{0}
	for _, r := range runs {
		for i := 0; i < r; i++ {
			isb = append(isb, 1)
		}
		isb = append(isb, 0)
	}

	f := make([]int, len(isb))
	t := make([]int, len(isb))
	for i := range t {
		t[i] = 1
	}
	inhs := make([][]int, len(springs))
	inds := make([][]int, len(springs))
	for i, c := range springs {
		if c != '#' {
			inhs[i] = t
		} else {
			inhs[i] = f
		}
		if c != '.' {
			inds[i] = t
		} else {
			inds[i] = f
		}
	}

	nfa := make([]int, len(isb))
	nfa[0] = 1
	nfa[1] = 1
	for i := range springs {
		next := make([]int, len(nfa))
		ind, inh := inds[i], inhs[i]
		alts := make([]int, len(nfa))
		for i, n := range nfa {
			alts[i] = n * (isb[i]*ind[i] | (isb[i]^1)*inh[i])
		}
		copy(alts[1:], alts)
		alts[0] = 0
		for i, n := range nfa {
			next[i] = n*(isb[i]^1)*inh[i] + alts[i]
		}
		nfa = next
	}
	return nfa[len(nfa)-1]
}

func main() {
	if true {
		nfa(".??..??...?##.", []int{1, 1, 3})
		return
	}
	var sum int
	for _, line := range utils.Lines(input) {
		fields := strings.Fields(line)
		sum += nfa(fields[0], utils.ExtractInts(fields[1]))
	}
	utils.Println(sum)

	sum = 0
	for _, line := range utils.Lines(input) {
		fields := strings.Fields(line)
		fields[0] = fmt.Sprintf("%[1]s?%[1]s?%[1]s?%[1]s?%[1]s", fields[0])
		fields[1] = fmt.Sprintf("%[1]s,%[1]s,%[1]s,%[1]s,%[1]s", fields[1])
		//sum += arrangements(fields[0], utils.ExtractInts(fields[1]))
		sum += nfaArray(fields[0], utils.ExtractInts(fields[1]))
	}
	utils.Println(sum)
}
