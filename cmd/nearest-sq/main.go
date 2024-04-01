package main

import (
	"fmt"
	"math"
)

func main() {
	result := NearestSq(111)
	fmt.Println(result)
}

func NearestSq(n int) int {
	sqrt := math.Sqrt(float64(n))
	sqrtRounded := math.Round(sqrt)
	if sqrt == sqrtRounded {
		return n
	}
	return int(sqrtRounded * sqrtRounded)
}
