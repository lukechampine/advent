package main

import (
	"fmt"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2021, 4)
var inputLines = utils.Lines(input)
var inputInts = utils.ExtractInts(input)

func isWinner(board []int) bool {
	for i := 0; i < 5; i++ {
		row := board[i*5:][:5]
		col := utils.Stride(board[i:], 5).([]int)
		if utils.IntSum(row) == -5 || utils.IntSum(col) == -5 {
			return true
		}
	}
	return false
}

func sumUnmarked(board []int) int {
	return utils.Sum(len(board), func(i int) int {
		if board[i] == -1 {
			return 0
		}
		return board[i]
	})
}

func markBoards(boards [][]int, n int) [][]int {
	var winners [][]int
	for _, board := range boards {
		if isWinner(board) {
			continue
		}
		for i, v := range board {
			if v == n {
				board[i] = -1
			}
		}
		if isWinner(board) {
			winners = append(winners, board)
		}
	}
	return winners
}

func main() {
	draws := utils.ExtractInts(inputLines[0])
	boards := utils.Window(inputInts[len(draws):], 25).([][]int)
	rem := len(boards)
	for _, n := range draws {
		winners := markBoards(boards, n)
		if len(winners) == 0 {
			continue
		}
		// part 1
		if rem == len(boards) {
			fmt.Println(winners[0])
			fmt.Println(sumUnmarked(winners[0]) * n)
		}
		// part 2
		rem -= len(winners)
		if rem == 0 {
			fmt.Println(sumUnmarked(winners[0]) * n)
		}
	}
}
