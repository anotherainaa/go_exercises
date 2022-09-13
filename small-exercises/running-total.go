package main

import "fmt"

func runningTotal(numbers []int) []int {
	var sum int
	var result []int

	for _, el := range numbers {
		sum += el
		result = append(result, sum)
	}
	return result
}

func main() {
	fmt.Println(runningTotal([]int{2, 5, 13}))
	fmt.Println(runningTotal([]int{14, 11, 7, 15, 20}))
	fmt.Println(runningTotal([]int{3}))
	fmt.Println(runningTotal([]int{}))
}