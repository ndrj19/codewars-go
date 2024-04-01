package main

import (
	"fmt"
	"unicode"
)

func main() {
	result := ReverseLetters("krish21an")
	fmt.Println(result)
}

func ReverseLetters(s string) (result string) {
	for _, v := range s {
		if unicode.IsLetter(v) {
			result = string(v) + result
		}
	}
	return
}
