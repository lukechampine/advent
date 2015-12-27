package main

const inputRow = 2981
const inputCol = 3075

func makeCodes() [][]int {
	codes := make([][]int, 8000)
	for i := range codes {
		codes[i] = make([]int, 8000)
	}
	lastcode := 20151125
	for i := 0; i < 8000; i++ {
		for j := 0; j <= i; j++ {
			codes[i-j][j] = lastcode
			lastcode = (lastcode * 252533) % 33554393
		}
	}
	return codes
}

func main() {
	codes := makeCodes()
	println(codes[inputRow-1][inputCol-1]) // 1-indexed
}
