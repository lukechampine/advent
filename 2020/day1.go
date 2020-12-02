package main

import (
	"fmt"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2020, 1)
var inputInts = utils.ExtractInts(input)

func main() {
outer:
	for _, i := range inputInts {
		for _, j := range inputInts {
			if i+j == 2020 {
				fmt.Println(i * j)
				break outer
			}
		}
	}
	for _, i := range inputInts {
		for _, j := range inputInts {
			for _, k := range inputInts {
				if i+j+k == 2020 {
					fmt.Println(i * j * k)
					return
				}
			}
		}
	}
}
