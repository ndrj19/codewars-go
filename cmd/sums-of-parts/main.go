package main

import (
	"fmt"
)

func main() {
	result := PartsSums([]uint64{0, 1, 3, 6, 10})
	fmt.Println(result)
}

func PartsSums(ls []uint64) []uint64 {
	var listSum uint64
	for _, num := range ls {
		listSum += num
	}

	result := []uint64{listSum}
	var subtract uint64

	for i := 0; i < len(ls); i++ {
		subtract += ls[i]
		result = append(result, listSum-subtract)
	}

	return result
}
