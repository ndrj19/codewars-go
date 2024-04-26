package main

import (
	"fmt"
)

func main() {
	result := NbrOfLaps(4, 6)
	fmt.Println(result)
}

func NbrOfLaps(x, y uint16) [2]uint16 {
	gcd := GCD(x, y)
	return [2]uint16{y / gcd, x / gcd}
}

func GCD(a, b uint16) uint16 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
