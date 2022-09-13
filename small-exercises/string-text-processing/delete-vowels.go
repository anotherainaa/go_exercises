package main

import (
	"fmt"
	"regexp"
)

func removeVowels(words []string) []string {
	reg, _ := regexp.Compile("[aeiouAEIOU]+")
	var result []string

	for _, word := range words {
		w := reg.ReplaceAllString(word, "")
		result = append(result, w)
	}
	return result
}

func main() {
	fmt.Println(removeVowels([]string{"abcdefghijklmnopqrstuvwxyz"}))
	fmt.Println(removeVowels([]string{"green", "YELLOW", "black", "white"}))
	fmt.Println(removeVowels([]string{"ABC", "AEIOU", "XYZ"}))
}