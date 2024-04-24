package main

import (
	"fmt"
)

func main() {
	result := Factorial(0)
	fmt.Println(result)
}

func Factorial(n int) int {
	x := 1
	for i := 2; i <= n; i++ {
		x *= i
	}
	return x
}
