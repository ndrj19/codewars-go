package main

import (
	"fmt"
)

func main() {
	result := NumberOfPairs([]string{"gray", "black", "purple", "purple", "gray", "black", "black"})
	fmt.Println(result)
}

func NumberOfPairs(gloves []string) (numPairs int) {
	m := make(map[string]int)
	for _, glove := range gloves {
		m[glove]++
		if m[glove]%2 == 0 {
			numPairs++
		}
	}

	return
}
