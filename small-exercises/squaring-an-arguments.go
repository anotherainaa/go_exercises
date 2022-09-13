package main

import "fmt"

func square(num int) int{
	return num * num
}

func main() {
	fmt.Println(square(5))
	fmt.Println(square(-8))
}