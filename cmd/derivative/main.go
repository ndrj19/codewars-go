package main

import (
	"fmt"
)

func main() {
	result := Derive(7, 8)
	fmt.Println(result)
}

func Derive(coefficient, exponent int) string {
	multi := coefficient * exponent
	return fmt.Sprintf("%dx^%d", multi, exponent-1)
}
