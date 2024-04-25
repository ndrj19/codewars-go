package main

import (
	"fmt"
	"strconv"
)

func main() {
	result := Thirt(321)
	fmt.Println(result)
}

func Thirt(n int) int {
	currentCalc := Calc(n)
	if n != currentCalc {
		return Thirt(currentCalc)
	}
	return currentCalc
}

func Calc(n int) int {
	numStr := strconv.Itoa(n)
	numArr := []int{}

	for _, digit := range numStr {
		numArr = append(numArr, int(digit)-'0')
	}

	multiSequence := []int{1, 10, 9, 12, 3, 4}
	var multi int
	var position int

	for i := len(numArr) - 1; i >= 0; i-- {
		multi += numArr[i] * multiSequence[position]
		if position == 5 {
			position = 0
		} else {
			position++
		}
	}
	return multi
}
