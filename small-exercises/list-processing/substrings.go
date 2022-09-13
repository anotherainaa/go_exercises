package main

import "fmt"

func substrings(str string) []string{
	substrings := []string{}

	for i := 0; i < len(str); i++ {
		for j := i + 1; j <= len(str); j++ {
			substrings = append(substrings, str[i:j])
		}
	}
	return substrings
}

func main() {
	fmt.Println(substrings("abc"))
	fmt.Println(substrings("a"))
	fmt.Println(substrings("xyzzy"))
	fmt.Println(substrings("abcde"))
}