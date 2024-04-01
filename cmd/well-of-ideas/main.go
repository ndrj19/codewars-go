package main

import (
	"fmt"
)

func main() {
	result := Well([]string{"good", "bad", "good", "good", "bad", "good", "bad", "bad", "good", "bad", "bad"})
	fmt.Println(result)
}

func Well(x []string) string {
	var goodIdeas int
	for _, idea := range x {
		if idea == "good" {
			goodIdeas++
		}
	}

	if goodIdeas > 2 {
		return "I smell a series!"
	} else if goodIdeas > 0 {
		return "Publish!"
	} else {
		return "Fail!"
	}
}
