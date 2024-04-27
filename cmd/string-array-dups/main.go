package main

import (
	"fmt"
	"strings"
)

func main() {
	result := Dup([]string{"abracadabra", "allottee", "assessee"})
	fmt.Println(result)
}

func Dup(arr []string) []string {
	var result []string
	for _, word := range arr {
		var sb strings.Builder
		for i, letter := range word {
			if i == 0 {
				sb.WriteRune(letter)
			}
			if i > 0 {
				if letter != rune(word[i-1]) {
					sb.WriteRune(letter)
				}
			}
		}
		result = append(result, sb.String())
	}
	return result
}
