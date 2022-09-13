package main

import "fmt"

func asciiValue(str string) int{
	var sum int
	for _, r := range str {
		sum += int(r)
	}
	return sum
}

func main() {
	fmt.Println(asciiValue("Four score"))
	fmt.Println(asciiValue("Launch School"))
	fmt.Println(asciiValue("a"))
	fmt.Println(asciiValue(""))
}