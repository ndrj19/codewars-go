package main

import (
	"fmt"
)

func main() {
	result := InAscOrder([]int{1, 6, 10, 18, 2, 4, 20})
	fmt.Println(result)
}

func InAscOrder(numbers []int) bool {
	for i, n := range numbers {
		if i == 0 {
		} else {
			if numbers[i-1] > n {
				return false
			}
		}
	}
	return true
}
