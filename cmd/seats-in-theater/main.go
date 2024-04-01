package main

import (
	"fmt"
)

func main() {
	result := SeatsInTheater(16, 11, 5, 3)
	fmt.Println(result)
}

func SeatsInTheater(nCols int, nRows int, col int, row int) int {
	return (nCols - col + 1) * (nRows - row)
}
