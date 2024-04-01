package main

import (
	"fmt"
	"sort"
)

func main() {
	result := TwoOldestAges([]int{6, 5, 83, 5, 3, 18})
	fmt.Println(result)
}

func TwoOldestAges(ages []int) [2]int {
	sort.Ints(ages)
	return [2]int{ages[len(ages)-2], ages[len(ages)-1]}
}
