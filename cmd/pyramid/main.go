package main

import (
	"fmt"
)

func main() {
	result := Pyramid(0)
	fmt.Println(result)
}

func Pyramid(n int) [][]int {
	result := [][]int{}
	for i := 0; i < n; i++ {
		var subArr []int
		for j := 0; j <= i; j++ {
			subArr = append(subArr, 1)
		}
		result = append(result, subArr)
	}
	return result
}
