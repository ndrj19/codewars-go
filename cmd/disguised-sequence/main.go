package main

import (
	"fmt"
)

func main() {
	result := Fct(17)
	fmt.Println(result)
}

func Fct(n uint) uint {
	return 1 << n
}
