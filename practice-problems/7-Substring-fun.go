package main

import "fmt"

/*
Complete the function that takes an array of words.

You must concatenate the nth letter from each word to construct a new word which should be returned as a string, where n is the position of the word in the list.

For example:

["yoda", "best", "has"]  -->  "yes"

*/

func nth_char(words []string) string {
	result := ""

	for i, word := range words {
		result += string(word[i])
	}

	return result
}

func main() {
	fmt.Println(nth_char([]string{"yoda", "best", "has"}))
	fmt.Println(nth_char([]string{}) == "")
	fmt.Println(nth_char([]string{"X-ray"}))
	fmt.Println(nth_char([]string{"No", "No'"}))
	fmt.Println(nth_char([]string{"Chad", "Morocco", "India", "Algeria", "Botswana", "Bahamas", "Ecuador", "Micronesia"}))
}