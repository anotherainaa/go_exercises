package main

import (
	"fmt"
	"strings"
)

func wordCap(words string) string {
	wordsSlice := strings.Split(words, " ")
	var results []string

	for _, w := range wordsSlice {
		w = strings.ToUpper(string(w[0])) + w[1:]
		results = append(results, w)
	}
	return strings.Join(results, " ")
}

func main() {
	fmt.Println(wordCap("four score and seven"))
	fmt.Println(wordCap("the javaScript language"))
	fmt.Println(wordCap("this is a \"quoted\" word"))
}