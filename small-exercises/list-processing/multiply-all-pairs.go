package main

import (
	"fmt"
	"sort"
)

func multiplyAllPairs(slice1, slice2 []int) []int{
	products := []int{}

	for _, num1 := range slice1 {
		for _, num2 := range slice2 {
			products = append(products, num1 * num2)
		}
	}

	sort.Ints(products)
	return products
}

func main() {
	fmt.Println(multiplyAllPairs([]int{2, 4}, []int{4, 3, 1, 2}))
}