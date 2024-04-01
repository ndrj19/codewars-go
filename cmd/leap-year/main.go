package main

import (
	"fmt"
)

func main() {
	result := IsLeapYear(2000)
	fmt.Println(result)
}

func IsLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	}
	if year%100 == 0 {
		return false
	}
	if year%4 == 0 {
		return true
	}
	return false
}
