package main

import (
	"fmt"
)

func main() {
	result := FoldArray([]int{1, 2, 3, 4, 5}, 3)
	fmt.Println(result)
}

func FoldArray(arr []int, runs int) []int {
	var result []int
	for i, j := 0, len(arr)-1; i <= j; i, j = i+1, j-1 {
		if i == j {
			result = append(result, arr[i])
		} else {
			result = append(result, arr[i]+arr[j])
		}
	}

	if runs > 1 {
		return FoldArray(result, runs-1)
	}
	return result
}
