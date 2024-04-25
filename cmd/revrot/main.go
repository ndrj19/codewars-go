package main

import (
	"fmt"
	"strings"
)

func main() {
	result := Revrot("123456987653", 6)
	fmt.Println(result)
}

func Revrot(s string, n int) string {
	if s == "" || n <= 0 || n > len(s) {
		return ""
	}

	var chunks []string

	for i := 0; i <= len(s)-n; i += n {
		currentChunk := s[i : i+n]
		if stringSum(currentChunk)%2 == 0 {
			chunks = append(chunks, reverse(currentChunk))
		} else {
			chunks = append(chunks, rotateLeft(currentChunk))
		}
	}

	result := strings.Join(chunks, "")
	return result
}

func stringSum(s string) int {
	var result int
	for _, num := range s {
		result += int(num) - '0'
	}
	return result
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func rotateLeft(s string) string {
	first := string(s[0])
	remainder := s[1:]
	return remainder + first
}
