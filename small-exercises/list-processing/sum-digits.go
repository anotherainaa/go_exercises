package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sum(num int) int{
	digits := strings.Split(strconv.Itoa(num), "")
	sum := 0

	for _, digit := range digits {
		number , _ := strconv.Atoi(digit)
		sum += number
	}

	return sum
}

func main() {
	fmt.Println(sum(23))
	// fmt.Println(sum(496))
	// fmt.Println(sum(123456789))
}