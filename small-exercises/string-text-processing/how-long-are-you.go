package main

import (
	"fmt"
	"strings"
	"strconv"
)

func wordLengths(words string) []string{
	splitWords := strings.Split(words, " ")
	var result []string

	for _, word := range splitWords {
		currentWord := string(word)
		wordAndCount := currentWord + " " + strconv.Itoa(len(currentWord))
		result = append(result, wordAndCount)
	}

	return result
}

func main() {
	fmt.Println(wordLengths("cow sheep chicken"))
	fmt.Println(wordLengths("baseball hot dogs and apple pie"))
	fmt.Println(wordLengths("It ain't easy, is it?"))
	fmt.Println(wordLengths("Supercalifragilisticexpialidocious"))
	fmt.Println(wordLengths(""))
}