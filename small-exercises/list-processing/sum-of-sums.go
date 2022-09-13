package main

import "fmt"

func sumOfSums(numbers []int) int{
	newNums := [][]int{}

	for i, _ := range numbers {
		newNums = append(newNums, numbers[0:i + 1])
	}

	sum := 0

	for _, slice := range newNums {
		innerSum := 0
		for _, is := range slice {
			innerSum += is
		}
		sum += innerSum
	}
	return sum
}

func main() {
	fmt.Println(sumOfSums([]int{3, 5, 2}))
	fmt.Println(sumOfSums([]int{1, 5, 7, 3}))
	fmt.Println(sumOfSums([]int{4}))
	fmt.Println(sumOfSums([]int{1, 2, 3, 4, 5}))
}