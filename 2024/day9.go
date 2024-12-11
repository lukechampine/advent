package main

import (
	"fmt"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2024, 9)

func rep(i int, n int) []int {
	is := make([]int, n)
	for j := range is {
		is[j] = i
	}
	return is
}

func main() {
	type run struct {
		id, len int
	}
	var fs []int
	var runs []run
	var id int
	for i := range input {
		if i%2 == 0 {
			fs = append(fs, rep(id, int(input[i]-'0'))...)
			runs = append(runs, run{id, int(input[i] - '0')})
			id++
		} else {
			fs = append(fs, rep(-1, int(input[i]-'0'))...)
			runs = append(runs, run{-1, int(input[i] - '0')})
		}
	}

	i, j := 0, len(fs)-1
	for i < j {
		if fs[i] != -1 {
			i++
		} else if fs[j] == -1 {
			j--
		} else {
			fs[i], fs[j] = fs[j], fs[i]
		}
	}
	var sum int
	for i := range fs {
		if fs[i] != -1 {
			sum += i * fs[i]
		}
	}
	fmt.Println(sum)

	for j := len(runs) - 1; j > 0; j-- {
		if runs[j].id == -1 {
			continue
		}
		for i := 0; i < j; i++ {
			if runs[i].id == -1 && runs[i].len >= runs[j].len {
				gap := run{-1, runs[i].len - runs[j].len}
				runs[i] = runs[j]
				runs[j].id = -1
				runs = append(append(runs[:i+1:i+1], gap), runs[i+1:]...)
				break
			}
		}
	}
	sum = 0
	i = 0
	for _, r := range runs {
		for j := 0; j < r.len; j++ {
			if r.id != -1 {
				sum += i * r.id
			}
			i++
		}
	}
	fmt.Println(sum)
}
