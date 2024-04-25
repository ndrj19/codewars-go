package main

import (
	"fmt"
)

func main() {
	result := CleanString("abc#d##c")
	fmt.Println(result)
}

func CleanString(s string) string {
	var result []rune
	for _, char := range s {
		if char != '#' {
			result = append(result, char)
		} else {
			if len(result)-1 >= 0 {
				result = result[:len(result)-1]
			}
		}
	}

	return string(result)
}
