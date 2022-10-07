package main

import (
	"fmt"
	"strings"
)

/*
Alphabet symmetry

approach
- iterate through each string in the array
- check if given character occupies it's characters in the 

*/

func generateMap() map[string]int {
	position := 1
	alphabets := make(map[string]int)
	for ch:= 'A'; ch <= 'Z'; ch++ {
		alphabets[string(ch)] = position
		position += 1;
	}
	return alphabets
}

func solve(words []string) []int {
	alphabets := generateMap()
	results := []int{}
	for _, word := range words {
		count := 0

		for i, ch := range word {
			char := strings.ToUpper(string(ch))

			if (alphabets[char] == (i+1)) {
				count += 1
			}
		}
		results = append(results, count)
	}
	return results
}

func main() {
	fmt.Println(solve([]string{"abode","ABc","xyzD"}))
}