package main

import (
	"fmt"
	// "slices"
	"sort"
)

// func main() {
// 	result := MergeArrays([]int{1, 1, 2, 30, 4}, []int{5, 6, 7, 8})
// 	fmt.Println(result)
// }

// func MergeArrays(arr1, arr2 []int) []int {
// 	mergedArr := append(arr1, arr2...)
// 	slices.Sort([]int(mergedArr))
// 	return slices.Compact[[]int, int](mergedArr)
// }

// Codewars go version to low

func main() {
	result := MergeArrays([]int{1, 1, 2, 30, 4}, []int{5, 6, 7, 8})
	fmt.Println(result)
}

func MergeArrays(arr1, arr2 []int) []int {
	mergedArr := append(arr1, arr2...)
	sort.Ints(mergedArr)
	return removeDuplicateInt(mergedArr)
}

func removeDuplicateInt(intSlice []int) []int {
	allKeys := make(map[int]bool)
	list := []int{}
	for _, item := range intSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
