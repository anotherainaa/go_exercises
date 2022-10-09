package main 

import (
	"fmt"
	"strconv"
)

/*
Given a string of integers, return the number of odd-numbered substrings that can be formed.

For example, in the case of "1341", they are 1, 1, 3, 13, 41, 341, 1341, a total of 7 numbers.

we have to create all substrings and then count the number of odd numbers? 
Is there an efficient way to do this question? 

*/

func getSubstrings(str string) []string {
	substrings := []string{}

	for i := 0; i < len(str); i++ {
		for j := i+1; j < len(str)+1; j++ {
			substrings = append(substrings, string(str[i:j]))
		}
	}
	return substrings
}

func nonEvenSubstrings(str string) int {
	substrings := getSubstrings(str)
	count := 0

	for _, v := range substrings {
		num, _ := strconv.Atoi(v)
		if num % 2 == 1 {
			count += 1
		}
	}
	return count
}

func main() {
	fmt.Println(nonEvenSubstrings("1341"))
	fmt.Println(nonEvenSubstrings("1357"))
	fmt.Println(nonEvenSubstrings("13471"))
	fmt.Println(nonEvenSubstrings("134721"))
	fmt.Println(nonEvenSubstrings("1347231"))
	fmt.Println(nonEvenSubstrings("13472315"))
}