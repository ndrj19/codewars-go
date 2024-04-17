package main

import (
	"fmt"
)

func main() {
	result := FindMissingLetter([]rune{'O', 'Q', 'R', 'S'})
	fmt.Println(result)
}

// []rune("café")
func FindMissingLetter(chars []rune) rune {
	var encodedSum int
	var encodedSumWithMissing int
	for i := 0; i <= len(chars); i++ {
		if i < len(chars) {
			encodedSum += int(rune(chars[i]))
		}
		encodedSumWithMissing += int(chars[0]) + i
	}

	encodedMissing := encodedSumWithMissing - encodedSum
	fmt.Println(encodedSum)
	fmt.Println(encodedSumWithMissing)
	fmt.Println(string(rune(encodedMissing)))
	return rune(encodedMissing)
}
