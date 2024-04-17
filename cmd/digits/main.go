package main

import (
	"fmt"
	"strconv"
)

func main() {
	result := Digits(128685)
	fmt.Println(result)
}

func Digits(n uint64) int {
	strDig := strconv.FormatUint(n, 10)
	return len(strDig)
}
