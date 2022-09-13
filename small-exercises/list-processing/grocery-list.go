package main

import (
	"fmt"
	"strconv"
)

func buyFruit(fruits [][]string) []string{
	result := []string{}

	for _, fruit := range fruits {
		num, _ := strconv.Atoi(fruit[1])
		for i := 1; i <= num; i++ {
			result = append(result, fruit[0])
		}
	}

	return result
}

func main() {
	a := [][]string{{"apple", "3"}, {"orange", "1"}, {"banana", "2"}}
	fmt.Println(buyFruit(a))
}