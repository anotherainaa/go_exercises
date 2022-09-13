package main

import (
	"fmt"
	"strconv"
)

// calculate the century
// if 1901 - 2000 - 20th
// if  

func contains(numbers []int, num int) bool {
	for _, v := range numbers {
		if v == num {
			return true
		}
	}
	return false
}

func centurySuffix(century int) string {
	var r string
	y := []int{11, 12, 13}

	if contains(y, century % 100) == true {
		return "th"
	}

	lastDigit := century % 10

	switch lastDigit {
	case 1:
		r = "st"
	case 2:
		r = "nd"
	case 3: 
		r = "rd"
	default:
		r = "th"
	}
	
	return r
}

func century(year int) string {
	century := year / 100 + 1
	if year % 100 == 0 {
		century -= 1
	}

	r := strconv.Itoa(century) + centurySuffix(century)
	return r
}

func main() {
	fmt.Println(century(2000))
	fmt.Println(century(2001))
	fmt.Println(century(1965))
	fmt.Println(century(256))
	fmt.Println(century(5))
	fmt.Println(century(10103))
	fmt.Println(century(1052))
	fmt.Println(century(1127))
	fmt.Println(century(11201))
}