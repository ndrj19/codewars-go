package main

import (
	"fmt"
)

func main() {
	result := multipleOfIndex([]int{22, -6, 32, 82, 9, 25})
	fmt.Println(result)
}

func multipleOfIndex(ints []int) []int {
	var result []int
	for i := 1; i < len(ints); i++ {
		if ints[i]%i == 0 {
			result = append(result, ints[i])
		}
	}
	return result
}
