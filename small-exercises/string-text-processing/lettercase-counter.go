package main 

import (
	"fmt"
	"regexp"
)

func letterCaseCount(word string) map[string]int{
	count := map[string]int{"lowercase": 0, "uppercase": 0, "neither": 0}

	for _, r := range word {
		lowerMatch , _ := regexp.MatchString("[a-z]", string(r))
		upperMatch , _ := regexp.MatchString("[A-Z]", string(r))
		// neitherMatch , _ := regexp.MatchString("[^a-zA-Z]", string(r))

		if lowerMatch == true {
			count["lowercase"] += 1
		} else if upperMatch == true {
			count["uppercase"] += 1
		} else{
			count["neither"] += 1
		}
	}

	return count
}

func main() {
	fmt.Println(letterCaseCount("abCdef 123"))
	fmt.Println(letterCaseCount("AbCd +Ef"))
	fmt.Println(letterCaseCount("123"))
	fmt.Println(letterCaseCount(""))
}