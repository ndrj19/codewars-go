package main

import (
	"fmt"
	"slices"
)

func main() {
	result := solve("abababababfapeifapefijaefaepfjavnefjnfbhwyfnjsifjapnes")
	fmt.Println(result)
}

// use for < 1.21 versions
// func contains(s []rune, letter rune) bool {
// 	for _, v := range s {
// 		if v == letter {
// 			return true
// 		}
// 	}

// 	return false
// }

func solve(str string) int {
	var maxValue int
	var currentValue int
	for _, letter := range str {
		vowels := []rune{'a', 'e', 'i', 'o', 'u'}
		if slices.Contains(vowels, letter) {
			if currentValue > maxValue {
				maxValue = currentValue
			}
			currentValue = 0
		} else {
			currentValue += int(letter) - 96
		}
	}

	if currentValue > maxValue {
		maxValue = currentValue
	}
	return maxValue
}
