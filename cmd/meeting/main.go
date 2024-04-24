package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	result := Meeting("Fred:Corwill;Wilfred:Corwill;Barney:Tornbull;Betty:Tornbull;Bjon:Tornbull;Raphael:Corwill;Alfred:Corwill")
	fmt.Println(result)
}

func Meeting(s string) string {
	var result []string
	friends := strings.Split(s, ";")
	for _, friend := range friends {
		lastFirst := strings.Split(friend, ":")
		lastFirst[0], lastFirst[1] = lastFirst[1], lastFirst[0]
		lastFirstStr := "(" + strings.ToUpper(strings.Join(lastFirst, ", ")) + ")"
		result = append(result, lastFirstStr)
	}
	sort.Strings(result)
	return strings.Join(result, "")
}
