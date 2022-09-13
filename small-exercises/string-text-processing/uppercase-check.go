package main

import (
	"fmt"
	"strings"
)

func isUppercase(str string) bool {
	return strings.ToUpper(str) == str
}

func main() {
	fmt.Println(isUppercase("t"))
	fmt.Println(isUppercase("T"))
	fmt.Println(isUppercase("Four Score"))
	fmt.Println(isUppercase("FOUR SCORE"))
	fmt.Println(isUppercase("4SCORE"))
	fmt.Println(isUppercase(""))
}