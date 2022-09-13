package main

import "fmt"

// iterate from 1 to int
// if the number is a multiple of 3 or 5, add it to an array
// return the sum of the array

func multisum(num int) int {
	var sum int

	for i := 1; i <= num; i++ {
		if i % 3 == 0 || i % 5 == 0 {
			sum += i
		}
	}

	return sum
}

func main() {
	fmt.Println(multisum(3))
	fmt.Println(multisum(5))
	fmt.Println(multisum(10))
	fmt.Println(multisum(1000))
}