package main

import (
	"fmt"
)

func main() {
	result := Angle(3)
	fmt.Println(result)
}

func Angle(n int) int {
	return (n - 2) * 180
}
