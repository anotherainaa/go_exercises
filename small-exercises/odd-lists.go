package main

import "fmt"

func oddities(slice []int) []int{
  var sliceResult []int
	
	for i, el := range slice {
		if i % 2 == 0 {
			sliceResult = append(sliceResult, el)
		}
	}
	return sliceResult
}

func main() {
	v := oddities([]int{2, 3, 4, 5, 6})
	fmt.Println(v)
}