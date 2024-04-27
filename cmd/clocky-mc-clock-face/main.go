package main

import (
	"fmt"
)

func main() {
	result := WhatTimeIsIt(50)
	fmt.Println(result)
}

func WhatTimeIsIt(angle int) string {
	hours := angle / 30
	if hours == 0 {
		hours = 12
	}
	minutes := angle % 30 * 2
	return fmt.Sprintf("%02d:%02d", hours, minutes)
}
