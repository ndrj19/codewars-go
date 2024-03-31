package main

import (
	"fmt"
	"math"
)

func main() {
	result := SquareOrSquareRoot([]int{4, 3, 9, 7, 2, 1})
	fmt.Println(result)
}

func SquareOrSquareRoot(arr []int) []int {
	result := []int{}
	for i := 0; i < len(arr); i++ {
		currentSqrt := math.Sqrt(float64(arr[i]))
		currentSqrtFloored := math.Floor(math.Sqrt(float64(arr[i])))
		if currentSqrt == currentSqrtFloored {
			result = append(result, int(currentSqrt))
		} else {
			result = append(result, int(math.Pow(float64(arr[i]), 2)))
		}
	}
	return result
}
