package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	result := solve("aaaaaaaaBBBBBBBBB")
	fmt.Println(result)
}

func solve(str string) string {
	var lower int
	for _, char := range str {
		if char == unicode.ToLower(char) {
			lower++
		}
	}

	upper := len(str) - lower
	if upper > lower {
		return strings.ToUpper(str)
	} else {
		return strings.ToLower(str)
	}
}
