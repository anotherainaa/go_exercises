package main 

import "fmt"

/*
Return substring instance count
Complete the solution so that it returns the number of times the search_text is found within the full_text.

Usage example:

solution('aa_bb_cc_dd_bb_e', 'bb') # should return 2 since bb shows up twice
solution('aaabbbcccc', 'bbb') # should return 1

approach

brute force approach
- generate all the substrings
- find all substrings that match the given string

the optimised approach 
- slide window? 
- 

subsrtings
- iterate through the first array 
	- iterate through second array
		- create substring
		- save substring
- return all substring


abc
a
ab
abc
b

*/

func generateSubstring(str string) []string {
	substrings := []string{}

	for i:= 0; i < len(str); i++ {
		for j := (i+1); j < len(str)+1; j++ {
			substr := str[i:j]
			substrings = append(substrings, substr)
		}
	}

	return substrings
}



func solution(str string, match string) int{
	count := 0
	
	substrings := generateSubstring(str)

	for _, v := range substrings {
		if v == match {
			count += 1
		}
	}
	return count
}


func main() {
	// s := "hello"
	// fmt.Println(s[:3])
	fmt.Println(generateSubstring("abc"))
	fmt.Println(solution("abcdeb","b"))
	fmt.Println(solution("abcdeb", "a"))
	fmt.Println(solution("aa_bb_cc_dd_bb_e", "bb"))
}