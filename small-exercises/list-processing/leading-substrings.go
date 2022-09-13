package main

import "fmt"

func leadingSubstrings(str string) []string{
	substrings := []string{}

	for i := 1; i <= len(str); i++ {
		substrings = append(substrings, str[0:i])
	}
	return substrings
}

func main() {
	fmt.Println(leadingSubstrings("abc"))
	fmt.Println(leadingSubstrings("a"))
	fmt.Println(leadingSubstrings("xyzzy"))
}