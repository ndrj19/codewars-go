package main

import (
	"fmt"
)

func main() {
	result := HighestRank([]int{12, 10, 8, 8, 3, 3, 3, 3, 2, 4, 10, 12, 10})
	fmt.Println(result)
}

func HighestRank(nums []int) int {
	var highest int
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
		if m[num] == m[highest] && num > highest {
			highest = num
		}
		if m[num] > m[highest] {
			highest = num
		}
	}

	return highest
}
