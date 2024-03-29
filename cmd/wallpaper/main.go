package main

import (
	"fmt"
	"math"
)

func main()  {
  result := wallPaper(6.3, 4.5, 3.29)
  fmt.Println(result)
}

func wallPaper(l, w, h float64) string {
  if (l == 0.0 || w == 0.0) {
    return "zero"
  }
  numbers := [21] string {"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve","thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen", "twenty",}
  area := 2 * l * h + 2 * w * h
  areaPlus15 := area + area * 0.15
  rolls := int(math.Ceil(areaPlus15 / 5.2))
  
  return numbers[rolls]
}
