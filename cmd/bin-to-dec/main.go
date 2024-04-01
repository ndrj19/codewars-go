package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	result := BinToDec("101010")
	fmt.Println(result)
}

func BinToDec(bin string) int {
	var dec int
	arr := strings.Split(bin, "")
	for i := 0; i < len(arr); i++ {
		if arr[i] == "1" {
			dec += int(math.Pow(2, float64(len(arr)-1-i)))
		}
	}
	return dec
}
